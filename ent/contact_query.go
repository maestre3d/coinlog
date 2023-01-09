// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/maestre3d/coinlog/ent/contact"
	"github.com/maestre3d/coinlog/ent/predicate"
	"github.com/maestre3d/coinlog/ent/user"
)

// ContactQuery is the builder for querying Contact entities.
type ContactQuery struct {
	config
	limit        *int
	offset       *int
	unique       *bool
	order        []OrderFunc
	fields       []string
	inters       []Interceptor
	predicates   []predicate.Contact
	withOwner    *UserQuery
	withLinkedTo *UserQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ContactQuery builder.
func (cq *ContactQuery) Where(ps ...predicate.Contact) *ContactQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *ContactQuery) Limit(limit int) *ContactQuery {
	cq.limit = &limit
	return cq
}

// Offset to start from.
func (cq *ContactQuery) Offset(offset int) *ContactQuery {
	cq.offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ContactQuery) Unique(unique bool) *ContactQuery {
	cq.unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *ContactQuery) Order(o ...OrderFunc) *ContactQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryOwner chains the current query on the "owner" edge.
func (cq *ContactQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(contact.Table, contact.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, contact.OwnerTable, contact.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLinkedTo chains the current query on the "linked_to" edge.
func (cq *ContactQuery) QueryLinkedTo() *UserQuery {
	query := (&UserClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(contact.Table, contact.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, contact.LinkedToTable, contact.LinkedToColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Contact entity from the query.
// Returns a *NotFoundError when no Contact was found.
func (cq *ContactQuery) First(ctx context.Context) (*Contact, error) {
	nodes, err := cq.Limit(1).All(newQueryContext(ctx, TypeContact, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{contact.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ContactQuery) FirstX(ctx context.Context) *Contact {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Contact ID from the query.
// Returns a *NotFoundError when no Contact ID was found.
func (cq *ContactQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cq.Limit(1).IDs(newQueryContext(ctx, TypeContact, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{contact.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ContactQuery) FirstIDX(ctx context.Context) string {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Contact entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Contact entity is found.
// Returns a *NotFoundError when no Contact entities are found.
func (cq *ContactQuery) Only(ctx context.Context) (*Contact, error) {
	nodes, err := cq.Limit(2).All(newQueryContext(ctx, TypeContact, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{contact.Label}
	default:
		return nil, &NotSingularError{contact.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ContactQuery) OnlyX(ctx context.Context) *Contact {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Contact ID in the query.
// Returns a *NotSingularError when more than one Contact ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ContactQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cq.Limit(2).IDs(newQueryContext(ctx, TypeContact, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{contact.Label}
	default:
		err = &NotSingularError{contact.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ContactQuery) OnlyIDX(ctx context.Context) string {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Contacts.
func (cq *ContactQuery) All(ctx context.Context) ([]*Contact, error) {
	ctx = newQueryContext(ctx, TypeContact, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Contact, *ContactQuery]()
	return withInterceptors[[]*Contact](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *ContactQuery) AllX(ctx context.Context) []*Contact {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Contact IDs.
func (cq *ContactQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	ctx = newQueryContext(ctx, TypeContact, "IDs")
	if err := cq.Select(contact.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ContactQuery) IDsX(ctx context.Context) []string {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ContactQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeContact, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*ContactQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ContactQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ContactQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeContact, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ContactQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ContactQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ContactQuery) Clone() *ContactQuery {
	if cq == nil {
		return nil
	}
	return &ContactQuery{
		config:       cq.config,
		limit:        cq.limit,
		offset:       cq.offset,
		order:        append([]OrderFunc{}, cq.order...),
		inters:       append([]Interceptor{}, cq.inters...),
		predicates:   append([]predicate.Contact{}, cq.predicates...),
		withOwner:    cq.withOwner.Clone(),
		withLinkedTo: cq.withLinkedTo.Clone(),
		// clone intermediate query.
		sql:    cq.sql.Clone(),
		path:   cq.path,
		unique: cq.unique,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ContactQuery) WithOwner(opts ...func(*UserQuery)) *ContactQuery {
	query := (&UserClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withOwner = query
	return cq
}

// WithLinkedTo tells the query-builder to eager-load the nodes that are connected to
// the "linked_to" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ContactQuery) WithLinkedTo(opts ...func(*UserQuery)) *ContactQuery {
	query := (&UserClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withLinkedTo = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		IsActive bool `json:"is_active,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Contact.Query().
//		GroupBy(contact.FieldIsActive).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *ContactQuery) GroupBy(field string, fields ...string) *ContactGroupBy {
	cq.fields = append([]string{field}, fields...)
	grbuild := &ContactGroupBy{build: cq}
	grbuild.flds = &cq.fields
	grbuild.label = contact.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		IsActive bool `json:"is_active,omitempty"`
//	}
//
//	client.Contact.Query().
//		Select(contact.FieldIsActive).
//		Scan(ctx, &v)
func (cq *ContactQuery) Select(fields ...string) *ContactSelect {
	cq.fields = append(cq.fields, fields...)
	sbuild := &ContactSelect{ContactQuery: cq}
	sbuild.label = contact.Label
	sbuild.flds, sbuild.scan = &cq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ContactSelect configured with the given aggregations.
func (cq *ContactQuery) Aggregate(fns ...AggregateFunc) *ContactSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ContactQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.fields {
		if !contact.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ContactQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Contact, error) {
	var (
		nodes       = []*Contact{}
		withFKs     = cq.withFKs
		_spec       = cq.querySpec()
		loadedTypes = [2]bool{
			cq.withOwner != nil,
			cq.withLinkedTo != nil,
		}
	)
	if cq.withOwner != nil || cq.withLinkedTo != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, contact.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Contact).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Contact{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withOwner; query != nil {
		if err := cq.loadOwner(ctx, query, nodes, nil,
			func(n *Contact, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withLinkedTo; query != nil {
		if err := cq.loadLinkedTo(ctx, query, nodes, nil,
			func(n *Contact, e *User) { n.Edges.LinkedTo = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *ContactQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*Contact, init func(*Contact), assign func(*Contact, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Contact)
	for i := range nodes {
		if nodes[i].user_contacts == nil {
			continue
		}
		fk := *nodes[i].user_contacts
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_contacts" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *ContactQuery) loadLinkedTo(ctx context.Context, query *UserQuery, nodes []*Contact, init func(*Contact), assign func(*Contact, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Contact)
	for i := range nodes {
		if nodes[i].user_contact_links == nil {
			continue
		}
		fk := *nodes[i].user_contact_links
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_contact_links" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cq *ContactQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.fields
	if len(cq.fields) > 0 {
		_spec.Unique = cq.unique != nil && *cq.unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ContactQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   contact.Table,
			Columns: contact.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: contact.FieldID,
			},
		},
		From:   cq.sql,
		Unique: true,
	}
	if unique := cq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contact.FieldID)
		for i := range fields {
			if fields[i] != contact.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ContactQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(contact.Table)
	columns := cq.fields
	if len(columns) == 0 {
		columns = contact.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.unique != nil && *cq.unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ContactGroupBy is the group-by builder for Contact entities.
type ContactGroupBy struct {
	selector
	build *ContactQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ContactGroupBy) Aggregate(fns ...AggregateFunc) *ContactGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *ContactGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeContact, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ContactQuery, *ContactGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *ContactGroupBy) sqlScan(ctx context.Context, root *ContactQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ContactSelect is the builder for selecting fields of Contact entities.
type ContactSelect struct {
	*ContactQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ContactSelect) Aggregate(fns ...AggregateFunc) *ContactSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ContactSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeContact, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ContactQuery, *ContactSelect](ctx, cs.ContactQuery, cs, cs.inters, v)
}

func (cs *ContactSelect) sqlScan(ctx context.Context, root *ContactQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
