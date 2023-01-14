// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/maestre3d/coinlog/ent/contact"
	"github.com/maestre3d/coinlog/ent/predicate"
	"github.com/maestre3d/coinlog/ent/user"
)

// ContactUpdate is the builder for updating Contact entities.
type ContactUpdate struct {
	config
	hooks    []Hook
	mutation *ContactMutation
}

// Where appends a list predicates to the ContactUpdate builder.
func (cu *ContactUpdate) Where(ps ...predicate.Contact) *ContactUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetIsActive sets the "is_active" field.
func (cu *ContactUpdate) SetIsActive(b bool) *ContactUpdate {
	cu.mutation.SetIsActive(b)
	return cu
}

// SetVersion sets the "version" field.
func (cu *ContactUpdate) SetVersion(u uint32) *ContactUpdate {
	cu.mutation.ResetVersion()
	cu.mutation.SetVersion(u)
	return cu
}

// AddVersion adds u to the "version" field.
func (cu *ContactUpdate) AddVersion(u int32) *ContactUpdate {
	cu.mutation.AddVersion(u)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *ContactUpdate) SetUpdatedAt(t time.Time) *ContactUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetDisplayName sets the "display_name" field.
func (cu *ContactUpdate) SetDisplayName(s string) *ContactUpdate {
	cu.mutation.SetDisplayName(s)
	return cu
}

// SetLinkedToUser sets the "linked_to_user" field.
func (cu *ContactUpdate) SetLinkedToUser(s string) *ContactUpdate {
	cu.mutation.SetLinkedToUser(s)
	return cu
}

// SetNillableLinkedToUser sets the "linked_to_user" field if the given value is not nil.
func (cu *ContactUpdate) SetNillableLinkedToUser(s *string) *ContactUpdate {
	if s != nil {
		cu.SetLinkedToUser(*s)
	}
	return cu
}

// ClearLinkedToUser clears the value of the "linked_to_user" field.
func (cu *ContactUpdate) ClearLinkedToUser() *ContactUpdate {
	cu.mutation.ClearLinkedToUser()
	return cu
}

// SetImageURL sets the "image_url" field.
func (cu *ContactUpdate) SetImageURL(s string) *ContactUpdate {
	cu.mutation.SetImageURL(s)
	return cu
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (cu *ContactUpdate) SetNillableImageURL(s *string) *ContactUpdate {
	if s != nil {
		cu.SetImageURL(*s)
	}
	return cu
}

// ClearImageURL clears the value of the "image_url" field.
func (cu *ContactUpdate) ClearImageURL() *ContactUpdate {
	cu.mutation.ClearImageURL()
	return cu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cu *ContactUpdate) SetOwnerID(id string) *ContactUpdate {
	cu.mutation.SetOwnerID(id)
	return cu
}

// SetOwner sets the "owner" edge to the User entity.
func (cu *ContactUpdate) SetOwner(u *User) *ContactUpdate {
	return cu.SetOwnerID(u.ID)
}

// SetLinkedToID sets the "linked_to" edge to the User entity by ID.
func (cu *ContactUpdate) SetLinkedToID(id string) *ContactUpdate {
	cu.mutation.SetLinkedToID(id)
	return cu
}

// SetNillableLinkedToID sets the "linked_to" edge to the User entity by ID if the given value is not nil.
func (cu *ContactUpdate) SetNillableLinkedToID(id *string) *ContactUpdate {
	if id != nil {
		cu = cu.SetLinkedToID(*id)
	}
	return cu
}

// SetLinkedTo sets the "linked_to" edge to the User entity.
func (cu *ContactUpdate) SetLinkedTo(u *User) *ContactUpdate {
	return cu.SetLinkedToID(u.ID)
}

// Mutation returns the ContactMutation object of the builder.
func (cu *ContactUpdate) Mutation() *ContactMutation {
	return cu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (cu *ContactUpdate) ClearOwner() *ContactUpdate {
	cu.mutation.ClearOwner()
	return cu
}

// ClearLinkedTo clears the "linked_to" edge to the User entity.
func (cu *ContactUpdate) ClearLinkedTo() *ContactUpdate {
	cu.mutation.ClearLinkedTo()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ContactUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ContactMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ContactUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ContactUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ContactUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ContactUpdate) check() error {
	if v, ok := cu.mutation.DisplayName(); ok {
		if err := contact.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`ent: validator failed for field "Contact.display_name": %w`, err)}
		}
	}
	if _, ok := cu.mutation.OwnerID(); cu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Contact.owner"`)
	}
	return nil
}

func (cu *ContactUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   contact.Table,
			Columns: contact.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: contact.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.IsActive(); ok {
		_spec.SetField(contact.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := cu.mutation.Version(); ok {
		_spec.SetField(contact.FieldVersion, field.TypeUint32, value)
	}
	if value, ok := cu.mutation.AddedVersion(); ok {
		_spec.AddField(contact.FieldVersion, field.TypeUint32, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(contact.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.DisplayName(); ok {
		_spec.SetField(contact.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := cu.mutation.ImageURL(); ok {
		_spec.SetField(contact.FieldImageURL, field.TypeString, value)
	}
	if cu.mutation.ImageURLCleared() {
		_spec.ClearField(contact.FieldImageURL, field.TypeString)
	}
	if cu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contact.OwnerTable,
			Columns: []string{contact.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contact.OwnerTable,
			Columns: []string{contact.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.LinkedToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contact.LinkedToTable,
			Columns: []string{contact.LinkedToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.LinkedToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contact.LinkedToTable,
			Columns: []string{contact.LinkedToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contact.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ContactUpdateOne is the builder for updating a single Contact entity.
type ContactUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ContactMutation
}

// SetIsActive sets the "is_active" field.
func (cuo *ContactUpdateOne) SetIsActive(b bool) *ContactUpdateOne {
	cuo.mutation.SetIsActive(b)
	return cuo
}

// SetVersion sets the "version" field.
func (cuo *ContactUpdateOne) SetVersion(u uint32) *ContactUpdateOne {
	cuo.mutation.ResetVersion()
	cuo.mutation.SetVersion(u)
	return cuo
}

// AddVersion adds u to the "version" field.
func (cuo *ContactUpdateOne) AddVersion(u int32) *ContactUpdateOne {
	cuo.mutation.AddVersion(u)
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *ContactUpdateOne) SetUpdatedAt(t time.Time) *ContactUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetDisplayName sets the "display_name" field.
func (cuo *ContactUpdateOne) SetDisplayName(s string) *ContactUpdateOne {
	cuo.mutation.SetDisplayName(s)
	return cuo
}

// SetLinkedToUser sets the "linked_to_user" field.
func (cuo *ContactUpdateOne) SetLinkedToUser(s string) *ContactUpdateOne {
	cuo.mutation.SetLinkedToUser(s)
	return cuo
}

// SetNillableLinkedToUser sets the "linked_to_user" field if the given value is not nil.
func (cuo *ContactUpdateOne) SetNillableLinkedToUser(s *string) *ContactUpdateOne {
	if s != nil {
		cuo.SetLinkedToUser(*s)
	}
	return cuo
}

// ClearLinkedToUser clears the value of the "linked_to_user" field.
func (cuo *ContactUpdateOne) ClearLinkedToUser() *ContactUpdateOne {
	cuo.mutation.ClearLinkedToUser()
	return cuo
}

// SetImageURL sets the "image_url" field.
func (cuo *ContactUpdateOne) SetImageURL(s string) *ContactUpdateOne {
	cuo.mutation.SetImageURL(s)
	return cuo
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (cuo *ContactUpdateOne) SetNillableImageURL(s *string) *ContactUpdateOne {
	if s != nil {
		cuo.SetImageURL(*s)
	}
	return cuo
}

// ClearImageURL clears the value of the "image_url" field.
func (cuo *ContactUpdateOne) ClearImageURL() *ContactUpdateOne {
	cuo.mutation.ClearImageURL()
	return cuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cuo *ContactUpdateOne) SetOwnerID(id string) *ContactUpdateOne {
	cuo.mutation.SetOwnerID(id)
	return cuo
}

// SetOwner sets the "owner" edge to the User entity.
func (cuo *ContactUpdateOne) SetOwner(u *User) *ContactUpdateOne {
	return cuo.SetOwnerID(u.ID)
}

// SetLinkedToID sets the "linked_to" edge to the User entity by ID.
func (cuo *ContactUpdateOne) SetLinkedToID(id string) *ContactUpdateOne {
	cuo.mutation.SetLinkedToID(id)
	return cuo
}

// SetNillableLinkedToID sets the "linked_to" edge to the User entity by ID if the given value is not nil.
func (cuo *ContactUpdateOne) SetNillableLinkedToID(id *string) *ContactUpdateOne {
	if id != nil {
		cuo = cuo.SetLinkedToID(*id)
	}
	return cuo
}

// SetLinkedTo sets the "linked_to" edge to the User entity.
func (cuo *ContactUpdateOne) SetLinkedTo(u *User) *ContactUpdateOne {
	return cuo.SetLinkedToID(u.ID)
}

// Mutation returns the ContactMutation object of the builder.
func (cuo *ContactUpdateOne) Mutation() *ContactMutation {
	return cuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (cuo *ContactUpdateOne) ClearOwner() *ContactUpdateOne {
	cuo.mutation.ClearOwner()
	return cuo
}

// ClearLinkedTo clears the "linked_to" edge to the User entity.
func (cuo *ContactUpdateOne) ClearLinkedTo() *ContactUpdateOne {
	cuo.mutation.ClearLinkedTo()
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ContactUpdateOne) Select(field string, fields ...string) *ContactUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Contact entity.
func (cuo *ContactUpdateOne) Save(ctx context.Context) (*Contact, error) {
	return withHooks[*Contact, ContactMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ContactUpdateOne) SaveX(ctx context.Context) *Contact {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ContactUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ContactUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ContactUpdateOne) check() error {
	if v, ok := cuo.mutation.DisplayName(); ok {
		if err := contact.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`ent: validator failed for field "Contact.display_name": %w`, err)}
		}
	}
	if _, ok := cuo.mutation.OwnerID(); cuo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Contact.owner"`)
	}
	return nil
}

func (cuo *ContactUpdateOne) sqlSave(ctx context.Context) (_node *Contact, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   contact.Table,
			Columns: contact.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: contact.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Contact.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contact.FieldID)
		for _, f := range fields {
			if !contact.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != contact.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.IsActive(); ok {
		_spec.SetField(contact.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := cuo.mutation.Version(); ok {
		_spec.SetField(contact.FieldVersion, field.TypeUint32, value)
	}
	if value, ok := cuo.mutation.AddedVersion(); ok {
		_spec.AddField(contact.FieldVersion, field.TypeUint32, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(contact.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.DisplayName(); ok {
		_spec.SetField(contact.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.ImageURL(); ok {
		_spec.SetField(contact.FieldImageURL, field.TypeString, value)
	}
	if cuo.mutation.ImageURLCleared() {
		_spec.ClearField(contact.FieldImageURL, field.TypeString)
	}
	if cuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contact.OwnerTable,
			Columns: []string{contact.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contact.OwnerTable,
			Columns: []string{contact.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.LinkedToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contact.LinkedToTable,
			Columns: []string{contact.LinkedToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.LinkedToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contact.LinkedToTable,
			Columns: []string{contact.LinkedToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Contact{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contact.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
