package persistence

import (
	"context"
	"log"

	"github.com/maestre3d/coinlog/entity"
	"github.com/maestre3d/coinlog/repository"
)

type UserSQL struct{}

var _ repository.User = UserSQL{}

func NewUserSQL() UserSQL {
	return UserSQL{}
}

func (u UserSQL) Save(ctx context.Context, v entity.User) error {
	log.Printf("at repository save, got %+v", v)
	return nil
}

func (u UserSQL) Get(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (u UserSQL) Search(ctx context.Context) (items []entity.User, nextPage string, err error) {
	//TODO implement me
	panic("implement me")
}

func (u UserSQL) Remove(ctx context.Context, v entity.User) error {
	//TODO implement me
	panic("implement me")
}
