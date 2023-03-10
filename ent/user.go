// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/maestre3d/coinlog/ent/user"
)

// User is the model entity for the User schema.
type User struct {
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
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Contacts holds the value of the contacts edge.
	Contacts []*Contact `json:"contacts,omitempty"`
	// ContactLinks holds the value of the contact_links edge.
	ContactLinks []*Contact `json:"contact_links,omitempty"`
	// FinancialAccounts holds the value of the financial_accounts edge.
	FinancialAccounts []*FinancialAccount `json:"financial_accounts,omitempty"`
	// Cards holds the value of the cards edge.
	Cards []*Card `json:"cards,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ContactsOrErr returns the Contacts value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ContactsOrErr() ([]*Contact, error) {
	if e.loadedTypes[0] {
		return e.Contacts, nil
	}
	return nil, &NotLoadedError{edge: "contacts"}
}

// ContactLinksOrErr returns the ContactLinks value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ContactLinksOrErr() ([]*Contact, error) {
	if e.loadedTypes[1] {
		return e.ContactLinks, nil
	}
	return nil, &NotLoadedError{edge: "contact_links"}
}

// FinancialAccountsOrErr returns the FinancialAccounts value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FinancialAccountsOrErr() ([]*FinancialAccount, error) {
	if e.loadedTypes[2] {
		return e.FinancialAccounts, nil
	}
	return nil, &NotLoadedError{edge: "financial_accounts"}
}

// CardsOrErr returns the Cards value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CardsOrErr() ([]*Card, error) {
	if e.loadedTypes[3] {
		return e.Cards, nil
	}
	return nil, &NotLoadedError{edge: "cards"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldIsActive:
			values[i] = new(sql.NullBool)
		case user.FieldVersion:
			values[i] = new(sql.NullInt64)
		case user.FieldID, user.FieldDisplayName:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				u.ID = value.String
			}
		case user.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				u.IsActive = value.Bool
			}
		case user.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				u.Version = uint32(value.Int64)
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				u.DisplayName = value.String
			}
		}
	}
	return nil
}

// QueryContacts queries the "contacts" edge of the User entity.
func (u *User) QueryContacts() *ContactQuery {
	return (&UserClient{config: u.config}).QueryContacts(u)
}

// QueryContactLinks queries the "contact_links" edge of the User entity.
func (u *User) QueryContactLinks() *ContactQuery {
	return (&UserClient{config: u.config}).QueryContactLinks(u)
}

// QueryFinancialAccounts queries the "financial_accounts" edge of the User entity.
func (u *User) QueryFinancialAccounts() *FinancialAccountQuery {
	return (&UserClient{config: u.config}).QueryFinancialAccounts(u)
}

// QueryCards queries the "cards" edge of the User entity.
func (u *User) QueryCards() *CardQuery {
	return (&UserClient{config: u.config}).QueryCards(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", u.IsActive))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", u.Version))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(u.DisplayName)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
