package sql

import (
	"context"

	"github.com/maestre3d/coinlog/domain"
	"github.com/maestre3d/coinlog/domain/contact"
	"github.com/maestre3d/coinlog/ent"
	entcontact "github.com/maestre3d/coinlog/ent/contact"
	"github.com/maestre3d/coinlog/ent/predicate"
	entuser "github.com/maestre3d/coinlog/ent/user"
	"github.com/maestre3d/coinlog/parser"
)

type ContactStorage struct {
	db *ent.Client
}

var _ contact.Repository = ContactStorage{}

var _ parser.ParseFunc[*ent.Contact, contact.Contact] = newContactFromEnt

func NewContactStorage(e *ent.Client) ContactStorage {
	return ContactStorage{
		db: e,
	}
}

func newContactFromEnt(src *ent.Contact) contact.Contact {
	if src == nil {
		return contact.Contact{}
	}

	return contact.Contact{
		ID:          src.ID,
		User:        newUserFromEnt(src.Edges.Owner),
		LinkedTo:    domain.PtrIfNotEmpty(newUserFromEnt(src.Edges.LinkedTo)),
		DisplayName: src.DisplayName,
		ImageURL:    src.ImageURL,
		Auditable: domain.Auditable{
			IsActive:  src.IsActive,
			Version:   src.Version,
			CreatedAt: src.CreatedAt,
			UpdatedAt: src.UpdatedAt,
		},
	}
}

func (c ContactStorage) create(ctx context.Context, v contact.Contact) error {
	stmt := c.db.Contact.Create().
		SetID(v.ID).
		SetDisplayName(v.DisplayName).
		SetOwnerID(v.User.ID).
		SetImageURL(v.ImageURL).
		SetIsActive(v.IsActive).
		SetVersion(v.Version).
		SetCreatedAt(v.CreatedAt).
		SetUpdatedAt(v.UpdatedAt)

	if v.LinkedTo != nil {
		stmt = stmt.SetLinkedToUser(v.LinkedTo.ID)
	}
	return stmt.Exec(ctx)
}

func (c ContactStorage) update(ctx context.Context, v contact.Contact) error {
	stmt := c.db.Contact.Update().
		Where(entcontact.IDEQ(v.ID)).
		SetDisplayName(v.DisplayName).
		SetOwnerID(v.User.ID).
		SetImageURL(v.ImageURL).
		SetIsActive(v.Auditable.IsActive).
		SetVersion(v.Auditable.Version).
		SetUpdatedAt(v.UpdatedAt)

	if v.LinkedTo == nil {
		stmt.ClearLinkedToUser()
	} else {
		stmt = stmt.SetLinkedToUser(v.LinkedTo.ID)
	}
	return stmt.Exec(ctx)
}

func (c ContactStorage) Save(ctx context.Context, v contact.Contact) error {
	if v.Version == 1 {
		return c.create(ctx, v)
	}

	return c.update(ctx, v)
}

func (c ContactStorage) Get(ctx context.Context, id string) (*contact.Contact, error) {
	out, err := c.db.Contact.Query().Where(entcontact.ID(id)).WithOwner().WithLinkedTo().Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	cp := newContactFromEnt(out)
	return &cp, nil
}

func (c ContactStorage) buildQueryFunc(pred ...predicate.Contact) querySQLFunc[*ent.Contact] {
	return func(ctx context.Context, limit, offset int) ([]*ent.Contact, error) {
		return c.db.Contact.Query().
			Where(pred...).
			Limit(limit).
			Offset(offset).
			WithOwner().
			WithLinkedTo().
			All(ctx)
	}
}

func (c ContactStorage) find(ctx context.Context, cr domain.Criteria, pred ...predicate.Contact) ([]contact.Contact,
	domain.PageToken, error) {
	return paginateSQLFunc(ctx, cr, newContactFromEnt, c.buildQueryFunc(pred...))
}

func (c ContactStorage) Find(ctx context.Context, cr domain.Criteria) (items []contact.Contact, nextPage domain.PageToken, err error) {
	return c.find(ctx, cr, entcontact.IsActive(true))
}

func (c ContactStorage) GetUserContacts(ctx context.Context, cr domain.Criteria,
	userID string) ([]contact.Contact, domain.PageToken, error) {
	return c.find(ctx, cr, entcontact.And(
		entcontact.HasOwnerWith(entuser.ID(userID)),
		entcontact.IsActive(true)),
	)
}

func (c ContactStorage) Remove(ctx context.Context, id string) error {
	err := c.db.Contact.DeleteOneID(id).Exec(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil
	}

	return err
}
