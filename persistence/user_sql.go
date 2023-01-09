package persistence

import (
	"context"
	"github.com/maestre3d/coinlog/exception"
	"log"
	"strconv"

	"github.com/maestre3d/coinlog/ent"
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

func (u UserSQL) Get(ctx context.Context, v *entity.User) (found bool, err error) {
	if v == nil {
		return false, nil
	}
	usr, err := u.client.User.Get(ctx, v.ID)
	if err != nil && ent.IsNotFound(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	model.NewUserFromSQL(usr, v)
	return true, nil
}

func (u UserSQL) Search(ctx context.Context, c valueobject.Criteria) (items []entity.User, nextPage valueobject.PageToken, err error) {
	pageTokenDecode := valueobject.DecodePageToken(c.PageToken)
	pageOffset, _ := strconv.Atoi(pageTokenDecode)
	usersSQL, err := u.client.User.Query().
		Where(user.IsActive(true)).
		Limit(c.Limit).
		Offset(pageOffset).
		All(ctx)
	if err != nil {
		return nil, nil, err
	} else if len(usersSQL) == 0 {
		return nil, nil, exception.ResourceNotFound{Resource: "user"}
	}

	items = make([]entity.User, 0, len(usersSQL))
	for _, usrSQL := range usersSQL {
		usr := entity.User{}
		model.NewUserFromSQL(usrSQL, &usr)
		items = append(items, usr)
	}
	pageOffset += len(items)
	nextPage = valueobject.NewPageToken(strconv.Itoa(pageOffset))
	log.Print(nextPage.String())
	return
}

func (u UserSQL) Remove(ctx context.Context, v entity.User) error {
	return u.client.User.DeleteOneID(v.ID).Exec(ctx)
}
