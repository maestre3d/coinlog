//go:build integration

package sql_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/maestre3d/coinlog/customtype"
	"github.com/maestre3d/coinlog/domain/user"
	"github.com/maestre3d/coinlog/ent"
	"github.com/maestre3d/coinlog/identifier"
	"github.com/maestre3d/coinlog/storage"
	"github.com/maestre3d/coinlog/storage/sql"
	"github.com/stretchr/testify/suite"
)

type userTestSuite struct {
	suite.Suite
	db          *ent.Client
	repo        sql.UserStorage
	idFactory   identifier.FactoryFunc
	flushBuffer sync.Map
}

func TestUserStorage(t *testing.T) {
	suite.Run(t, &userTestSuite{})
}

func (s *userTestSuite) SetupSuite() {
	cfg := sql.NewConfig()
	client, _, err := sql.NewEntClientWithAutoMigrate(cfg)
	s.Require().NoError(err)
	s.db = client
	s.repo = sql.NewUserStorage(client)
	s.idFactory = identifier.NewKSUID
}

func (s *userTestSuite) TearDownSuite() {
	s.flushBuffer.Range(func(key, value any) bool {
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second*15)
		defer cancel()
		keyStr := key.(string)
		if keyStr == "" {
			return true
		}
		s.T().Logf("removing user %s", keyStr)
		_ = s.repo.Remove(ctx, keyStr)
		return true
	})
	_ = s.db.Close()
}

func (s *userTestSuite) TestUserStorage_Create() {
	id, errID := s.idFactory()
	s.Require().NoError(errID)
	err := s.repo.Save(context.TODO(), user.User{
		ID:          id,
		DisplayName: "foo",
		Auditable: customtype.Auditable{
			IsActive:  true,
			Version:   1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	})
	defer func() {
		if err != nil {
			s.flushBuffer.Store(id, struct{}{})
		}
	}()
	s.Assert().NoError(err)
}

func (s *userTestSuite) insertRow(uu ...user.User) (ids []string) {
	ids = make([]string, 0, len(uu))

	for _, u := range uu {
		idCreate, errCreateID := s.idFactory()
		s.Require().NoError(errCreateID)
		u.ID = idCreate
		errCreate := s.repo.Save(context.TODO(), u)
		s.Require().NoError(errCreate)
		ids = append(ids, idCreate)
	}
	for _, id := range ids {
		s.flushBuffer.Store(id, struct{}{})
	}

	return
}

func (s *userTestSuite) TestUserStorage_Update() {
	idCreate := s.insertRow(user.User{
		DisplayName: "bar",
		Auditable: customtype.Auditable{
			IsActive:  true,
			Version:   1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	})

	tests := []struct {
		name       string
		updateFunc func(existingID string) (usedID string, err error)
		wantErr    error
	}{
		{
			name: "upsert ignore id",
			updateFunc: func(_ string) (string, error) {
				id, _ := s.idFactory()
				err := s.repo.Save(context.TODO(), user.User{
					ID:          id,
					DisplayName: "bar new",
					Auditable: customtype.Auditable{
						IsActive:  true,
						Version:   1,
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				})
				return id, err
			},
			wantErr: nil,
		},
		{
			name: "update",
			updateFunc: func(existingID string) (string, error) {
				err := s.repo.Save(context.TODO(), user.User{
					ID:          existingID,
					DisplayName: "bar v2",
					Auditable: customtype.Auditable{
						IsActive:  true,
						Version:   2,
						CreatedAt: time.Now(),
						UpdatedAt: time.Now().Add(time.Second * 10),
					},
				})
				return "", err
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			usedID, err := tt.updateFunc(idCreate[0])
			defer func() {
				if err == nil && usedID != idCreate[0] {
					s.flushBuffer.Store(usedID, struct{}{})
				}
			}()
			s.Assert().Equal(tt.wantErr, err)
		})
	}
}

func (s *userTestSuite) TestUserStorage_Get() {
	createTime := time.Now().UTC()
	idCreate := s.insertRow(user.User{
		DisplayName: "foo",
		Auditable: customtype.Auditable{
			IsActive:  true,
			Version:   1,
			CreatedAt: createTime,
			UpdatedAt: createTime,
		},
	})

	tests := []struct {
		name       string
		inLookupID string
		want       *user.User
		wantErr    error
	}{
		{
			name:       "empty",
			inLookupID: "",
			want:       nil,
			wantErr:    nil,
		},
		{
			name:       "missing",
			inLookupID: "this_is_a_fake_id",
			want:       nil,
			wantErr:    nil,
		},
		{
			name:       "hit",
			inLookupID: idCreate[0],
			want: &user.User{
				ID:          idCreate[0],
				DisplayName: "foo",
				Auditable: customtype.Auditable{
					IsActive:  true,
					Version:   1,
					CreatedAt: createTime,
					UpdatedAt: createTime,
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			usr, err := s.repo.Get(context.TODO(), tt.inLookupID)
			s.Assert().Equal(tt.wantErr, err)
			if tt.want != nil && usr != nil {
				s.Assert().Equal(tt.want.ID, usr.ID)
				s.Assert().Equal(tt.want.DisplayName, usr.DisplayName)
			}
		})
	}
}

func (s *userTestSuite) TestUserStorage_List() {
	createTime := time.Now().UTC()
	_ = s.insertRow(
		user.User{
			DisplayName: "foo",
			Auditable: customtype.Auditable{
				IsActive:  true,
				Version:   1,
				CreatedAt: createTime,
				UpdatedAt: createTime,
			},
		},
		user.User{
			DisplayName: "bar",
			Auditable: customtype.Auditable{
				IsActive:  true,
				Version:   1,
				CreatedAt: createTime,
				UpdatedAt: createTime,
			},
		})

	tests := []struct {
		name string
		in   storage.Criteria
		exp  int
		// Token is a hash of a database offset. Moreover, tests are running concurrently and
		// could cause these tests to deliver false positives and thus fail unexpectedly.
		//
		// Instead, we use a boolean indicating if a token is expected from the search exec.
		expToken bool
		wantErr  error
	}{
		{
			name:     "empty",
			in:       storage.Criteria{},
			exp:      0,
			expToken: false,
			wantErr:  nil,
		},
		{
			name: "exact",
			in: storage.Criteria{
				Limit:     2,
				PageToken: nil,
			},
			exp:      2,
			expToken: true,
			wantErr:  nil,
		},
		{
			name: "less",
			in: storage.Criteria{
				Limit:     1,
				PageToken: nil,
			},
			exp:      1,
			expToken: true,
			wantErr:  nil,
		},
		{
			name: "more",
			in: storage.Criteria{
				Limit:     100,
				PageToken: nil,
			},
			exp:      2,
			expToken: true,
			wantErr:  nil,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			usr, nextPage, err := s.repo.Find(context.TODO(), tt.in)
			s.Assert().Equal(tt.wantErr, err)
			s.Assert().Equal(tt.expToken, nextPage != nil)
			// We're running tests concurrently meaning other tests might insert rows at the time of this proc exec.
			// Thus, we use greater or equal assertion instead of equals or len assertions
			// to avoid tests delivering false positives and thus fail unexpectedly.
			s.Assert().GreaterOrEqual(len(usr), tt.exp)
		})
	}
}

func (s *userTestSuite) TestUserStorage_Delete() {
	createTime := time.Now().UTC()
	idCreate := s.insertRow(user.User{
		DisplayName: "foo",
		Auditable: customtype.Auditable{
			IsActive:  true,
			Version:   1,
			CreatedAt: createTime,
			UpdatedAt: createTime,
		},
	})

	tests := []struct {
		name       string
		inLookupID string
		wantErr    error
	}{
		{
			name:       "empty",
			inLookupID: "",
			wantErr:    nil,
		},
		{
			name:       "missing",
			inLookupID: "this_is_a_fake_id",
			wantErr:    nil,
		},
		{
			name:       "hit",
			inLookupID: idCreate[0],
			wantErr:    nil,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			err := s.repo.Remove(context.TODO(), tt.inLookupID)
			s.Assert().Equal(tt.wantErr, err)
		})
	}
}
