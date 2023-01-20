package financialaccount

import (
	"context"

	"github.com/maestre3d/coinlog/domain"
	"github.com/maestre3d/coinlog/parser"
	"github.com/maestre3d/coinlog/storage"
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
	cmd := args.(CreateCommand)
	acc, err := newFinancialAccount(cmd)
	if err != nil {
		return err
	}

	return s.repo.Save(ctx, acc)
}

func (s Service) getByID(ctx context.Context, id string) (FinancialAccount, error) {
	acc, err := s.repo.Get(ctx, id)
	if err != nil {
		return FinancialAccount{}, err
	} else if acc == nil {
		return FinancialAccount{}, ErrNotFound
	}
	return *acc, nil
}

func (s Service) Update(ctx context.Context, args any) error {
	cmd := args.(UpdateCommand)
	acc, err := s.getByID(ctx, cmd.AccountID)
	if err != nil {
		return err
	}
	if err = acc.update(cmd); err != nil {
		return err
	}
	return s.repo.Save(ctx, acc)
}

func (s Service) Delete(ctx context.Context, id string) error {
	return s.repo.Remove(ctx, id)
}

func (s Service) GetByID(ctx context.Context, id string) (View, error) {
	acc, err := s.getByID(ctx, id)
	if err != nil {
		return View{}, err
	}

	return NewView(acc), nil
}

func (s Service) List(ctx context.Context, cr storage.Criteria) ([]View, storage.PageToken, error) {
	cc, nextPage, err := s.repo.Find(ctx, cr)
	if err != nil {
		return nil, nil, err
	} else if len(cc) == 0 {
		return nil, nil, ErrNotFound
	}

	return parser.NewCollection(cc, NewView), nextPage, nil
}
