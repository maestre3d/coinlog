package user

import (
	"context"

	"github.com/maestre3d/coinlog/domain"
	"github.com/maestre3d/coinlog/messaging"
	"github.com/maestre3d/coinlog/parser"
	"github.com/maestre3d/coinlog/storage"
)

type Service struct {
	repo Repository
	bus  *messaging.Bus
}

var _ domain.BasicService[View] = Service{}

func NewService(r Repository, b *messaging.Bus) (Service, error) {
	if err := b.Register(Stream, Event{}); err != nil {
		return Service{}, err
	}
	return Service{
		repo: r,
		bus:  b,
	}, nil
}

func (s Service) Create(ctx context.Context, args any) error {
	a := args.(CreateCommand)
	if err := domain.Validate.Struct(a); err != nil {
		return err
	}

	usr := newUser(a)
	if err := s.repo.Save(ctx, usr); err != nil {
		return nil
	}
	return s.bus.Publish(ctx, usr.PullEvents())
}

func (s Service) getByID(ctx context.Context, id string) (User, error) {
	usr, err := s.repo.Get(ctx, id)
	if err != nil {
		return User{}, err
	} else if usr == nil {
		return User{}, ErrNotFound
	}
	return *usr, nil
}

func (s Service) Update(ctx context.Context, args any) error {
	cmd := args.(UpdateCommand)
	if err := domain.Validate.Struct(cmd); err != nil {
		return err
	}
	usr, err := s.getByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	usr.update(cmd)
	if err = s.repo.Save(ctx, usr); err != nil {
		return nil
	}
	return s.bus.Publish(ctx, usr.PullEvents())
}

func (s Service) GetByID(ctx context.Context, id string) (View, error) {
	usr, err := s.getByID(ctx, id)
	if err != nil {
		return View{}, err
	}
	return NewView(usr), nil
}

func (s Service) List(ctx context.Context, cr storage.Criteria) ([]View, storage.PageToken, error) {
	uu, nextPage, err := s.repo.Find(ctx, cr)
	if err != nil {
		return nil, nil, err
	} else if len(uu) == 0 {
		return nil, nil, ErrNotFound
	}

	return parser.NewCollection(uu, NewView), nextPage, nil
}

func (s Service) Delete(ctx context.Context, id string) error {
	if err := s.repo.Remove(ctx, id); err != nil {
		return err
	}
	usr := User{}
	usr.delete(id)
	return s.bus.Publish(ctx, usr.PullEvents())
}
