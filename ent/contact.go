// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/maestre3d/coinlog/ent/contact"
	"github.com/maestre3d/coinlog/ent/user"
)

// Contact is the model entity for the Contact schema.
type Contact struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
	// Version holds the value of the "version" field.
	Version uint32 `json:"version,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DisplayName holds the value of the "display_name" field.
	DisplayName string `json:"display_name,omitempty"`
	// ImageURL holds the value of the "image_url" field.
	ImageURL string `json:"image_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ContactQuery when eager-loading is set.
	Edges              ContactEdges `json:"edges"`
	user_contacts      *string
	user_contact_links *string
}

// ContactEdges holds the relations/edges for other nodes in the graph.
type ContactEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// LinkedTo holds the value of the linked_to edge.
	LinkedTo *User `json:"linked_to,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ContactEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// LinkedToOrErr returns the LinkedTo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ContactEdges) LinkedToOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.LinkedTo == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.LinkedTo, nil
	}
	return nil, &NotLoadedError{edge: "linked_to"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Contact) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case contact.FieldIsActive:
			values[i] = new(sql.NullBool)
		case contact.FieldVersion:
			values[i] = new(sql.NullInt64)
		case contact.FieldID, contact.FieldDisplayName, contact.FieldImageURL:
			values[i] = new(sql.NullString)
		case contact.FieldCreatedAt, contact.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case contact.ForeignKeys[0]: // user_contacts
			values[i] = new(sql.NullString)
		case contact.ForeignKeys[1]: // user_contact_links
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Contact", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Contact fields.
func (c *Contact) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case contact.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = value.String
			}
		case contact.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				c.IsActive = value.Bool
			}
		case contact.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				c.Version = uint32(value.Int64)
			}
		case contact.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case contact.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case contact.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				c.DisplayName = value.String
			}
		case contact.FieldImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_url", values[i])
			} else if value.Valid {
				c.ImageURL = value.String
			}
		case contact.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_contacts", values[i])
			} else if value.Valid {
				c.user_contacts = new(string)
				*c.user_contacts = value.String
			}
		case contact.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_contact_links", values[i])
			} else if value.Valid {
				c.user_contact_links = new(string)
				*c.user_contact_links = value.String
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Contact entity.
func (c *Contact) QueryOwner() *UserQuery {
	return (&ContactClient{config: c.config}).QueryOwner(c)
}

// QueryLinkedTo queries the "linked_to" edge of the Contact entity.
func (c *Contact) QueryLinkedTo() *UserQuery {
	return (&ContactClient{config: c.config}).QueryLinkedTo(c)
}

// Update returns a builder for updating this Contact.
// Note that you need to call Contact.Unwrap() before calling this method if this Contact
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Contact) Update() *ContactUpdateOne {
	return (&ContactClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Contact entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Contact) Unwrap() *Contact {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Contact is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Contact) String() string {
	var builder strings.Builder
	builder.WriteString("Contact(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", c.IsActive))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", c.Version))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(c.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("image_url=")
	builder.WriteString(c.ImageURL)
	builder.WriteByte(')')
	return builder.String()
}

// Contacts is a parsable slice of Contact.
type Contacts []*Contact

func (c Contacts) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}