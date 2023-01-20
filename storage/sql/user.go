package sql

import (
	"context"

	"github.com/maestre3d/coinlog/customtype"
	"github.com/maestre3d/coinlog/domain/user"
	"github.com/maestre3d/coinlog/ent"
	"github.com/maestre3d/coinlog/ent/predicate"
	entschema "github.com/maestre3d/coinlog/ent/user"
	"github.com/maestre3d/coinlog/parser"
	"github.com/maestre3d/coinlog/storage"
)

type UserStorage struct {
	db *ent.Client
}

var _ user.Repository = UserStorage{}

func NewUserStorage(e *ent.Client) UserStorage {
	return UserStorage{
		db: e,
	}
}

var _ parser.ParseFunc[*ent.User, user.User] = newUserFromEnt

func newUserFromEnt(src *ent.User) user.User {
	if src == nil {
		return user.User{}
	}
	return user.User{
		ID:          src.ID,
		DisplayName: src.DisplayName,
		Auditable: customtype.Auditable{
			IsActive:  src.IsActive,
			Version:   src.Version,
			CreatedAt: src.CreatedAt,
			UpdatedAt: src.UpdatedAt,
		},
	}
}

func (u UserStorage) Save(ctx context.Context, v user.User) error {
	return u.db.User.Create().
		SetID(v.ID).
		SetDisplayName(v.DisplayName).
		SetIsActive(v.Auditable.IsActive).
		SetVersion(v.Auditable.Version).
		SetCreatedAt(v.Auditable.CreatedAt).
		SetUpdatedAt(v.Auditable.UpdatedAt).
		OnConflictColumns(entschema.FieldID).
		UpdateNewValues().
		Exec(ctx)
}

func (u UserStorage) Get(ctx context.Context, id string) (*user.User, error) {
	usr, err := u.db.User.Get(ctx, id)
	if err != nil && ent.IsNotFound(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	usrCopy := newUserFromEnt(usr)
	return &usrCopy, nil
}

func (u UserStorage) buildQueryFunc(pred ...predicate.User) querySQLFunc[*ent.User] {
	return func(ctx context.Context, limit, offset int) ([]*ent.User, error) {
		return u.db.User.Query().
			Where(pred...).
			Limit(limit).
			Offset(offset).
			All(ctx)
	}
}

func (u UserStorage) find(ctx context.Context, cr storage.Criteria, pred ...predicate.User) ([]user.User,
	storage.PageToken, error) {
	return paginateSQLFunc(ctx, cr, newUserFromEnt, u.buildQueryFunc(pred...))
}

func (u UserStorage) Find(ctx context.Context, cr storage.Criteria) ([]user.User, storage.PageToken, error) {
	if cr.Limit == 0 {
		return nil, nil, nil
	}
	return u.find(ctx, cr, entschema.IsActive(true))
}

func (u UserStorage) Remove(ctx context.Context, id string) error {
	err := u.db.User.DeleteOneID(id).Exec(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil
	}

	return err
}
