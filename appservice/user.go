package appservice

import (
	"context"

	"github.com/maestre3d/coinlog/entity"
	"github.com/maestre3d/coinlog/repository"
)

type User struct {
	repository repository.User
}

func NewUser(r repository.User) User {
	return User{repository: r}
}

func (u User) CreateUser(ctx context.Context, id, name string) error {
	usr, err := entity.NewUser(entity.NewUserArgs{
		ID:          id,
		DisplayName: name,
	})
	if err != nil {
		return err
	}
	return u.repository.Save(ctx, usr)
}
