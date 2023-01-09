package appservice

import (
	"context"
	"github.com/maestre3d/coinlog/exception"
	"github.com/maestre3d/coinlog/valueobject"
	"github.com/maestre3d/coinlog/view"
	"github.com/rs/zerolog/log"

	"github.com/maestre3d/coinlog/entity"
	"github.com/maestre3d/coinlog/repository"
)

type User struct {
	repository repository.User
}

func NewUser(r repository.User) User {
	return User{repository: r}
}

func (u User) Create(ctx context.Context, args entity.UserArgs) error {
	usr, err := entity.NewUser(args)
	if err != nil {
		return err
	}
	return u.repository.Save(ctx, usr)
}

func (u User) getByID(ctx context.Context, id string) (entity.User, error) {
	usr := entity.User{
		ID: id,
	}
	ok, err := u.repository.Get(ctx, &usr)
	if err != nil {
		return entity.User{}, err
	} else if !ok {
		return entity.User{}, exception.ResourceNotFound{
			Resource: "user",
		}
	}
	return usr, nil
}

func (u User) Update(ctx context.Context, args entity.UserArgs) error {
	log.Info().Msg(args.ID)
	usr, err := u.getByID(ctx, args.ID)
	if err != nil {
		return err
	}
	if err = usr.Update(args); err != nil {
		return err
	}
	return u.repository.Save(ctx, usr)
}

func (u User) Search(ctx context.Context, criteria valueobject.Criteria) ([]view.User, valueobject.PageToken, error) {
	users, token, err := u.repository.Search(ctx, criteria)
	if err != nil {
		return nil, nil, err
	}
	return view.NewCollection[entity.User, view.User](users, view.NewUser), token, nil
}

func (u User) GetByID(ctx context.Context, id string) (view.User, error) {
	usr, err := u.getByID(ctx, id)
	if err != nil {
		return view.User{}, err
	}
	return view.NewUser(usr), nil
}

func (u User) DeleteByID(ctx context.Context, id string) error {
	usr, err := u.getByID(ctx, id)
	if err != nil {
		return err
	}
	return u.repository.Remove(ctx, usr)
}
