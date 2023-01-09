package appservice

import (
	"context"

	"github.com/maestre3d/coinlog/domainutil"
	"github.com/maestre3d/coinlog/entity"
	"github.com/maestre3d/coinlog/repository"
	"github.com/maestre3d/coinlog/valueobject"
	"github.com/maestre3d/coinlog/view"
)

type User struct {
	repository repository.User
}

func NewUser(r repository.User) User {
	return User{repository: r}
}

func (u User) Create(ctx context.Context, args entity.UserArgs) error {
	out, err := entity.NewUser(args)
	if err != nil {
		return err
	}
	return u.repository.Save(ctx, out)
}

func (u User) getByID(ctx context.Context, id string) (entity.User, error) {
	out, err := u.repository.Get(ctx, id)
	if err != nil {
		return entity.User{}, err
	}
	return out, nil
}

func (u User) Update(ctx context.Context, args entity.UserArgs) error {
	out, err := u.getByID(ctx, args.ID)
	if err != nil {
		return err
	}
	if err = out.Update(args); err != nil {
		return err
	}
	return u.repository.Save(ctx, out)
}

func (u User) Search(ctx context.Context, criteria valueobject.Criteria) ([]view.User, valueobject.PageToken, error) {
	out, token, err := u.repository.Search(ctx, criteria)
	if err != nil {
		return nil, nil, err
	}
	return domainutil.NewCollection[entity.User, view.User](out, view.NewUser), token, nil
}

func (u User) GetByID(ctx context.Context, id string) (view.User, error) {
	out, err := u.getByID(ctx, id)
	if err != nil {
		return view.User{}, err
	}
	return view.NewUser(out), nil
}

func (u User) DeleteByID(ctx context.Context, id string) error {
	return u.repository.Remove(ctx, id)
}
