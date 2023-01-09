package persistence

import (
	"context"

	"github.com/maestre3d/coinlog/ent"
	"github.com/maestre3d/coinlog/ent/contact"
	"github.com/maestre3d/coinlog/ent/predicate"
	"github.com/maestre3d/coinlog/ent/user"
	"github.com/maestre3d/coinlog/entity"
	"github.com/maestre3d/coinlog/model"
	"github.com/maestre3d/coinlog/repository"
	"github.com/maestre3d/coinlog/valueobject"
)

type ContactSQL struct {
	client *ent.Client
}

var _ repository.Contact = ContactSQL{}

func NewContactSQL(c *ent.Client) ContactSQL {
	return ContactSQL{
		client: c,
	}
}

func (c ContactSQL) Save(ctx context.Context, v entity.Contact) error {
	stmt := c.client.Contact.Create().
		SetID(v.ID).
		SetDisplayName(v.DisplayName).
		SetOwnerID(v.User.ID).
		SetIsActive(v.Auditable.IsActive).
		SetImageURL(v.ImageURL).
		SetVersion(v.Auditable.Version).
		SetCreatedAt(v.Auditable.CreatedAt).
		SetUpdatedAt(v.Auditable.UpdatedAt)
	if v.LinkedTo.ID != "" {
		stmt = stmt.SetLinkedToID(v.LinkedTo.ID)
	}
	return stmt.OnConflictColumns(contact.FieldID).UpdateNewValues().Exec(ctx)
}

func (c ContactSQL) Get(ctx context.Context, k string) (entity.Contact, error) {
	out, err := c.client.Contact.Query().Where(contact.ID(k)).WithOwner().WithLinkedTo().Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return entity.Contact{}, entity.ErrContactNotFound
	} else if err != nil {
		return entity.Contact{}, err
	}

	return model.NewContactFromSQL(out), nil
}

func (c ContactSQL) search(ctx context.Context, criteria valueobject.Criteria,
	predicates ...predicate.Contact) ([]entity.Contact, valueobject.PageToken, error) {
	return paginateSQLFunc[*ent.Contact, entity.Contact](ctx, criteria, model.NewContactFromSQL,
		func(ctx context.Context, limit, offset int) ([]*ent.Contact, error) {
			return c.client.Contact.Query().
				Where(predicates...).
				Limit(limit).
				Offset(offset).
				WithOwner().
				WithLinkedTo().
				All(ctx)
		})
}

func (c ContactSQL) Search(ctx context.Context, criteria valueobject.Criteria) ([]entity.Contact,
	valueobject.PageToken, error) {
	return c.search(ctx, criteria, contact.IsActive(true))
}

func (c ContactSQL) GetUserContacts(ctx context.Context, criteria valueobject.Criteria, userID string) ([]entity.Contact,
	valueobject.PageToken, error) {
	return c.search(ctx, criteria, contact.And(contact.HasOwnerWith(user.ID(userID)), contact.IsActive(true)))
}

func (c ContactSQL) Remove(ctx context.Context, k string) error {
	return c.client.Contact.DeleteOneID(k).Exec(ctx)
}
