package appservice

import (
	"context"

	"github.com/maestre3d/coinlog/domainutil"
	"github.com/maestre3d/coinlog/entity"
	"github.com/maestre3d/coinlog/repository"
	"github.com/maestre3d/coinlog/valueobject"
	"github.com/maestre3d/coinlog/view"
)

type Contact struct {
	repository repository.Contact
}

func NewContact(r repository.Contact) Contact {
	return Contact{repository: r}
}

func (u Contact) Create(ctx context.Context, args entity.ContactArgs) error {
	out, err := entity.NewContact(args)
	if err != nil {
		return err
	}
	return u.repository.Save(ctx, out)
}

func (u Contact) getByID(ctx context.Context, id string) (entity.Contact, error) {
	out, err := u.repository.Get(ctx, id)
	if err != nil {
		return entity.Contact{}, err
	}
	return out, nil
}

func (u Contact) Update(ctx context.Context, args entity.ContactArgs) error {
	out, err := u.getByID(ctx, args.ID)
	if err != nil {
		return err
	}
	if err = out.Update(args); err != nil {
		return err
	}
	return u.repository.Save(ctx, out)
}

func (u Contact) Search(ctx context.Context, criteria valueobject.Criteria) ([]view.Contact, valueobject.PageToken, error) {
	out, token, err := u.repository.Search(ctx, criteria)
	if err != nil {
		return nil, nil, err
	}
	return domainutil.NewCollection[entity.Contact, view.Contact](out, view.NewContact), token, nil
}

func (u Contact) GetByID(ctx context.Context, id string) (view.Contact, error) {
	out, err := u.getByID(ctx, id)
	if err != nil {
		return view.Contact{}, err
	}
	return view.NewContact(out), nil
}

func (u Contact) DeleteByID(ctx context.Context, id string) error {
	return u.repository.Remove(ctx, id)
}
