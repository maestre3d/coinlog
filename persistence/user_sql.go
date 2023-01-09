package persistence

import (
	"context"

	"github.com/maestre3d/coinlog/ent"
	"github.com/maestre3d/coinlog/ent/predicate"
	"github.com/maestre3d/coinlog/ent/user"
	"github.com/maestre3d/coinlog/entity"
	"github.com/maestre3d/coinlog/model"
	"github.com/maestre3d/coinlog/repository"
	"github.com/maestre3d/coinlog/valueobject"
)

type UserSQL struct {
	client *ent.Client
}

var _ repository.User = UserSQL{}

func NewUserSQL(c *ent.Client) UserSQL {
	return UserSQL{
		client: c,
	}
}

func (u UserSQL) Save(ctx context.Context, v entity.User) error {
	return u.client.User.Create().
		SetID(v.ID).
		SetDisplayName(v.DisplayName).
		SetIsActive(v.Auditable.IsActive).
		SetVersion(v.Auditable.Version).
		SetCreatedAt(v.Auditable.CreatedAt).
		SetUpdatedAt(v.Auditable.UpdatedAt).
		OnConflictColumns(user.FieldID).
		UpdateNewValues().
		Exec(ctx)
}

func (u UserSQL) Get(ctx context.Context, key string) (entity.User, error) {
	out, err := u.client.User.Get(ctx, key)
	if err != nil && ent.IsNotFound(err) {
		return entity.User{}, entity.ErrUserNotFound
	} else if err != nil {
		return entity.User{}, err
	}
	return model.NewUserFromSQL(out), nil
}

func (u UserSQL) search(ctx context.Context, c valueobject.Criteria,
	predicates ...predicate.User) ([]entity.User, valueobject.PageToken, error) {
	return paginateSQLFunc[*ent.User, entity.User](ctx, c,
		model.NewUserFromSQL,
		func(ctx context.Context, limit, offset int) ([]*ent.User, error) {
			return u.client.User.Query().
				Where(predicates...).
				Limit(c.Limit).
				Offset(offset).
				All(ctx)
		})
}

func (u UserSQL) Search(ctx context.Context, c valueobject.Criteria) ([]entity.User,
	valueobject.PageToken, error) {
	return u.search(ctx, c, user.IsActive(true))
}

func (u UserSQL) Remove(ctx context.Context, k string) error {
	return u.client.User.DeleteOneID(k).Exec(ctx)
}
