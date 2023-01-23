package card

import (
	"context"

	"github.com/maestre3d/coinlog/domain/financialaccount"

	"github.com/maestre3d/coinlog/domain"
	"github.com/maestre3d/coinlog/parser"
	"github.com/maestre3d/coinlog/storage"
)

type Service struct {
	repo    Repository
	accRepo financialaccount.Repository
}

var _ domain.BasicService[View] = Service{}

func NewService(r Repository, accRepo financialaccount.Repository) Service {
	return Service{
		repo:    r,
		accRepo: accRepo,
	}
}

func (s Service) Create(ctx context.Context, args any) error {
	cmd := args.(CreateCommand)
	card, err := newCard(cmd)
	if err != nil {
		return err
	}

	if cmd.FinancialAccountID != "" {
		if acc, _ := s.accRepo.Get(ctx, cmd.FinancialAccountID); acc == nil {
			return financialaccount.ErrNotFound
		}
	}

	return s.repo.Save(ctx, card)
}

func (s Service) getByID(ctx context.Context, id string) (Card, error) {
	card, err := s.repo.Get(ctx, id)
	if err != nil {
		return Card{}, err
	} else if card == nil {
		return Card{}, ErrNotFound
	}
	return *card, nil
}

func (s Service) Update(ctx context.Context, args any) error {
	cmd := args.(UpdateCommand)
	card, err := s.getByID(ctx, cmd.CardID)
	if err != nil {
		return err
	}
	if err = card.update(cmd); err != nil {
		return err
	}
	return s.repo.Save(ctx, card)
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

func (s Service) ListUserCards(ctx context.Context, cr storage.Criteria, userID string) ([]View, storage.PageToken, error) {
	cc, nextPage, err := s.repo.GetUserCards(ctx, cr, userID)
	if err != nil {
		return nil, nil, err
	} else if len(cc) == 0 {
		return nil, nil, ErrNotFound
	}

	return parser.NewCollection(cc, NewView), nextPage, nil
}
