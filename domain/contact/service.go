package contact

import (
	"context"

	"github.com/maestre3d/coinlog/domain"
	"github.com/maestre3d/coinlog/parser"
)

type Service struct {
	repo Repository
}

var _ domain.BasicService[View] = Service{}

func NewService(r Repository) Service {
	return Service{
		repo: r,
	}
}

func (s Service) Create(ctx context.Context, args any) error {
	a := args.(CreateArgs)
	if err := domain.Validate.Struct(a); err != nil {
		return err
	}

	return s.repo.Save(ctx, newContact(a))
}

func (s Service) getByID(ctx context.Context, id string) (Contact, error) {
	usr, err := s.repo.Get(ctx, id)
	if err != nil {
		return Contact{}, err
	} else if usr == nil {
		return Contact{}, ErrNotFound
	}
	return *usr, nil
}

func (s Service) Update(ctx context.Context, args any) error {
	a := args.(UpdateArgs)
	if err := domain.Validate.Struct(a); err != nil {
		return err
	}
	ctc, err := s.getByID(ctx, a.ID)
	if err != nil {
		return err
	}
	ctc.Update(a)
	return s.repo.Save(ctx, ctc)
}

func (s Service) GetByID(ctx context.Context, id string) (View, error) {
	usr, err := s.getByID(ctx, id)
	if err != nil {
		return View{}, err
	}
	return NewView(usr), nil
}

func (s Service) List(ctx context.Context, cr domain.Criteria) ([]View, domain.PageToken, error) {
	cc, nextPage, err := s.repo.Find(ctx, cr)
	if err != nil {
		return nil, nil, err
	} else if len(cc) == 0 {
		return nil, nil, ErrNotFound
	}

	return parser.NewCollection(cc, NewView), nextPage, nil
}

func (s Service) ListUserContacts(ctx context.Context, cr domain.Criteria, userID string) ([]View,
	domain.PageToken, error) {
	cc, nextPage, err := s.repo.GetUserContacts(ctx, cr, userID)
	if err != nil {
		return nil, nil, err
	} else if len(cc) == 0 {
		return nil, nil, ErrNotFound
	}

	return parser.NewCollection(cc, NewView), nextPage, nil
}

func (s Service) Delete(ctx context.Context, id string) error {
	return s.repo.Remove(ctx, id)
}
