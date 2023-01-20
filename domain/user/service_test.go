package user_test

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/maestre3d/coinlog/customtype"
	"github.com/maestre3d/coinlog/domain/user"
	"github.com/maestre3d/coinlog/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type serviceTestSuite struct {
	suite.Suite

	repo *mocks.UserRepository
	svc  user.Service
}

func TestService(t *testing.T) {
	suite.Run(t, &serviceTestSuite{})
}

func (s *serviceTestSuite) SetupSuite() {
	s.repo = mocks.NewUserRepository(s.T())
	s.svc = user.NewService(s.repo)
}

func (s *serviceTestSuite) SetupTest() {
	s.repo.Mock = mock.Mock{
		ExpectedCalls: nil,
		Calls:         nil,
	}
}

func (s *serviceTestSuite) TestService_Create() {
	ctx := context.TODO()

	s.repo.On("Save", ctx, mock.MatchedBy(func(u user.User) bool {
		return u.Version == 1 && u.IsActive &&
			!u.CreatedAt.IsZero() && !u.UpdatedAt.IsZero() &&
			u.ID == "123" && u.DisplayName == "Foo"
	})).Return(error(nil))

	err := s.svc.Create(ctx, user.CreateCommand{
		ID:          "",
		DisplayName: "",
	})
	assert.Error(s.T(), err)
	assert.True(s.T(), strings.Contains(err.Error(), "required"))

	err = s.svc.Create(ctx, user.CreateCommand{
		ID:          "123",
		DisplayName: "Foo",
	})
	assert.NoError(s.T(), err)
}

func (s *serviceTestSuite) TestService_Update() {
	ctx := context.TODO()

	createTime := time.Now().UTC()
	expUser := &user.User{
		ID:          "123",
		DisplayName: "Foo",
		Auditable: customtype.Auditable{
			IsActive:  true,
			Version:   1,
			CreatedAt: createTime,
			UpdatedAt: createTime,
		},
	}
	s.repo.On("Get", ctx, "456").Return(nil, error(nil))
	s.repo.On("Get", ctx, "789").Return(nil, errors.New("database error"))
	s.repo.On("Get", ctx, "123").Return(expUser, error(nil))
	s.repo.On("Save", ctx, mock.MatchedBy(func(u user.User) bool {
		return u.Version == 2 && u.IsActive &&
			u.CreatedAt.Equal(createTime) && createTime.Before(u.UpdatedAt) &&
			u.ID == "123" && u.DisplayName == "Bar"
	})).Return(error(nil))

	err := s.svc.Update(ctx, user.UpdateCommand{
		ID:          "",
		DisplayName: "",
	})
	assert.Error(s.T(), err)
	assert.True(s.T(), strings.Contains(err.Error(), "required"))

	err = s.svc.Update(ctx, user.UpdateCommand{
		ID:          "789",
		DisplayName: "Bar",
	})
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "database error", err.Error())

	err = s.svc.Update(ctx, user.UpdateCommand{
		ID:          "456",
		DisplayName: "Bar",
	})
	assert.Error(s.T(), err)
	assert.ErrorIs(s.T(), err, user.ErrNotFound)

	err = s.svc.Update(ctx, user.UpdateCommand{
		ID:          "123",
		DisplayName: "Bar",
	})
	assert.NoError(s.T(), err)
}

func (s *serviceTestSuite) TestService_Get() {
	ctx := context.TODO()

	createTime := time.Now().UTC()
	expUser := &user.User{
		ID:          "123",
		DisplayName: "Foo",
		Auditable: customtype.Auditable{
			IsActive:  true,
			Version:   1,
			CreatedAt: createTime,
			UpdatedAt: createTime,
		},
	}
	s.repo.On("Get", ctx, "456").Return(nil, error(nil))
	s.repo.On("Get", ctx, "789").Return(nil, errors.New("database error"))
	s.repo.On("Get", ctx, "123").Return(expUser, error(nil))

	usr, err := s.svc.GetByID(ctx, "789")
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "database error", err.Error())

	usr, err = s.svc.GetByID(ctx, "456")
	assert.Error(s.T(), err)
	assert.ErrorIs(s.T(), err, user.ErrNotFound)

	usr, err = s.svc.GetByID(ctx, "123")
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expUser.ID, usr.ID)
	assert.Equal(s.T(), expUser.DisplayName, usr.DisplayName)
	assert.Equal(s.T(), expUser.Version, usr.Version)
	assert.Equal(s.T(), expUser.IsActive, usr.IsActive)
	assert.True(s.T(), expUser.CreatedAt.Equal(usr.CreatedAt))
	assert.True(s.T(), expUser.UpdatedAt.Equal(usr.UpdatedAt))
}
