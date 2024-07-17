// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/42milez/go-oidc-expt/pkg/ent/ent/authcode"
	"github.com/42milez/go-oidc-expt/pkg/ent/ent/predicate"
	"github.com/42milez/go-oidc-expt/pkg/ent/ent/relyingparty"
	"github.com/42milez/go-oidc-expt/pkg/typedef"
)

// AuthCodeQuery is the builder for querying AuthCode entities.
type AuthCodeQuery struct {
	config
	ctx              *QueryContext
	order            []authcode.OrderOption
	inters           []Interceptor
	predicates       []predicate.AuthCode
	withRelyingParty *RelyingPartyQuery
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AuthCodeQuery builder.
func (acq *AuthCodeQuery) Where(ps ...predicate.AuthCode) *AuthCodeQuery {
	acq.predicates = append(acq.predicates, ps...)
	return acq
}

// Limit the number of records to be returned by this query.
func (acq *AuthCodeQuery) Limit(limit int) *AuthCodeQuery {
	acq.ctx.Limit = &limit
	return acq
}

// Offset to start from.
func (acq *AuthCodeQuery) Offset(offset int) *AuthCodeQuery {
	acq.ctx.Offset = &offset
	return acq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (acq *AuthCodeQuery) Unique(unique bool) *AuthCodeQuery {
	acq.ctx.Unique = &unique
	return acq
}

// Order specifies how the records should be ordered.
func (acq *AuthCodeQuery) Order(o ...authcode.OrderOption) *AuthCodeQuery {
	acq.order = append(acq.order, o...)
	return acq
}

// QueryRelyingParty chains the current query on the "relying_party" edge.
func (acq *AuthCodeQuery) QueryRelyingParty() *RelyingPartyQuery {
	query := (&RelyingPartyClient{config: acq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := acq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := acq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(authcode.Table, authcode.FieldID, selector),
			sqlgraph.To(relyingparty.Table, relyingparty.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, authcode.RelyingPartyTable, authcode.RelyingPartyColumn),
		)
		fromU = sqlgraph.SetNeighbors(acq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AuthCode entity from the query.
// Returns a *NotFoundError when no AuthCode was found.
func (acq *AuthCodeQuery) First(ctx context.Context) (*AuthCode, error) {
	nodes, err := acq.Limit(1).All(setContextOp(ctx, acq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{authcode.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (acq *AuthCodeQuery) FirstX(ctx context.Context) *AuthCode {
	node, err := acq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AuthCode ID from the query.
// Returns a *NotFoundError when no AuthCode ID was found.
func (acq *AuthCodeQuery) FirstID(ctx context.Context) (id typedef.AuthCodeID, err error) {
	var ids []typedef.AuthCodeID
	if ids, err = acq.Limit(1).IDs(setContextOp(ctx, acq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{authcode.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (acq *AuthCodeQuery) FirstIDX(ctx context.Context) typedef.AuthCodeID {
	id, err := acq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AuthCode entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AuthCode entity is found.
// Returns a *NotFoundError when no AuthCode entities are found.
func (acq *AuthCodeQuery) Only(ctx context.Context) (*AuthCode, error) {
	nodes, err := acq.Limit(2).All(setContextOp(ctx, acq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{authcode.Label}
	default:
		return nil, &NotSingularError{authcode.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (acq *AuthCodeQuery) OnlyX(ctx context.Context) *AuthCode {
	node, err := acq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AuthCode ID in the query.
// Returns a *NotSingularError when more than one AuthCode ID is found.
// Returns a *NotFoundError when no entities are found.
func (acq *AuthCodeQuery) OnlyID(ctx context.Context) (id typedef.AuthCodeID, err error) {
	var ids []typedef.AuthCodeID
	if ids, err = acq.Limit(2).IDs(setContextOp(ctx, acq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{authcode.Label}
	default:
		err = &NotSingularError{authcode.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (acq *AuthCodeQuery) OnlyIDX(ctx context.Context) typedef.AuthCodeID {
	id, err := acq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AuthCodes.
func (acq *AuthCodeQuery) All(ctx context.Context) ([]*AuthCode, error) {
	ctx = setContextOp(ctx, acq.ctx, "All")
	if err := acq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AuthCode, *AuthCodeQuery]()
	return withInterceptors[[]*AuthCode](ctx, acq, qr, acq.inters)
}

// AllX is like All, but panics if an error occurs.
func (acq *AuthCodeQuery) AllX(ctx context.Context) []*AuthCode {
	nodes, err := acq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AuthCode IDs.
func (acq *AuthCodeQuery) IDs(ctx context.Context) (ids []typedef.AuthCodeID, err error) {
	if acq.ctx.Unique == nil && acq.path != nil {
		acq.Unique(true)
	}
	ctx = setContextOp(ctx, acq.ctx, "IDs")
	if err = acq.Select(authcode.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (acq *AuthCodeQuery) IDsX(ctx context.Context) []typedef.AuthCodeID {
	ids, err := acq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (acq *AuthCodeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, acq.ctx, "Count")
	if err := acq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, acq, querierCount[*AuthCodeQuery](), acq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (acq *AuthCodeQuery) CountX(ctx context.Context) int {
	count, err := acq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (acq *AuthCodeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, acq.ctx, "Exist")
	switch _, err := acq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (acq *AuthCodeQuery) ExistX(ctx context.Context) bool {
	exist, err := acq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AuthCodeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (acq *AuthCodeQuery) Clone() *AuthCodeQuery {
	if acq == nil {
		return nil
	}
	return &AuthCodeQuery{
		config:           acq.config,
		ctx:              acq.ctx.Clone(),
		order:            append([]authcode.OrderOption{}, acq.order...),
		inters:           append([]Interceptor{}, acq.inters...),
		predicates:       append([]predicate.AuthCode{}, acq.predicates...),
		withRelyingParty: acq.withRelyingParty.Clone(),
		// clone intermediate query.
		sql:  acq.sql.Clone(),
		path: acq.path,
	}
}

// WithRelyingParty tells the query-builder to eager-load the nodes that are connected to
// the "relying_party" edge. The optional arguments are used to configure the query builder of the edge.
func (acq *AuthCodeQuery) WithRelyingParty(opts ...func(*RelyingPartyQuery)) *AuthCodeQuery {
	query := (&RelyingPartyClient{config: acq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	acq.withRelyingParty = query
	return acq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Code string `json:"code,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AuthCode.Query().
//		GroupBy(authcode.FieldCode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (acq *AuthCodeQuery) GroupBy(field string, fields ...string) *AuthCodeGroupBy {
	acq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AuthCodeGroupBy{build: acq}
	grbuild.flds = &acq.ctx.Fields
	grbuild.label = authcode.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Code string `json:"code,omitempty"`
//	}
//
//	client.AuthCode.Query().
//		Select(authcode.FieldCode).
//		Scan(ctx, &v)
func (acq *AuthCodeQuery) Select(fields ...string) *AuthCodeSelect {
	acq.ctx.Fields = append(acq.ctx.Fields, fields...)
	sbuild := &AuthCodeSelect{AuthCodeQuery: acq}
	sbuild.label = authcode.Label
	sbuild.flds, sbuild.scan = &acq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AuthCodeSelect configured with the given aggregations.
func (acq *AuthCodeQuery) Aggregate(fns ...AggregateFunc) *AuthCodeSelect {
	return acq.Select().Aggregate(fns...)
}

func (acq *AuthCodeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range acq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, acq); err != nil {
				return err
			}
		}
	}
	for _, f := range acq.ctx.Fields {
		if !authcode.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if acq.path != nil {
		prev, err := acq.path(ctx)
		if err != nil {
			return err
		}
		acq.sql = prev
	}
	return nil
}

func (acq *AuthCodeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AuthCode, error) {
	var (
		nodes       = []*AuthCode{}
		_spec       = acq.querySpec()
		loadedTypes = [1]bool{
			acq.withRelyingParty != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AuthCode).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AuthCode{config: acq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(acq.modifiers) > 0 {
		_spec.Modifiers = acq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, acq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := acq.withRelyingParty; query != nil {
		if err := acq.loadRelyingParty(ctx, query, nodes, nil,
			func(n *AuthCode, e *RelyingParty) { n.Edges.RelyingParty = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (acq *AuthCodeQuery) loadRelyingParty(ctx context.Context, query *RelyingPartyQuery, nodes []*AuthCode, init func(*AuthCode), assign func(*AuthCode, *RelyingParty)) error {
	ids := make([]typedef.RelyingPartyID, 0, len(nodes))
	nodeids := make(map[typedef.RelyingPartyID][]*AuthCode)
	for i := range nodes {
		fk := nodes[i].RelyingPartyID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(relyingparty.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "relying_party_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (acq *AuthCodeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := acq.querySpec()
	if len(acq.modifiers) > 0 {
		_spec.Modifiers = acq.modifiers
	}
	_spec.Node.Columns = acq.ctx.Fields
	if len(acq.ctx.Fields) > 0 {
		_spec.Unique = acq.ctx.Unique != nil && *acq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, acq.driver, _spec)
}

func (acq *AuthCodeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(authcode.Table, authcode.Columns, sqlgraph.NewFieldSpec(authcode.FieldID, field.TypeUint64))
	_spec.From = acq.sql
	if unique := acq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if acq.path != nil {
		_spec.Unique = true
	}
	if fields := acq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authcode.FieldID)
		for i := range fields {
			if fields[i] != authcode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if acq.withRelyingParty != nil {
			_spec.Node.AddColumnOnce(authcode.FieldRelyingPartyID)
		}
	}
	if ps := acq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := acq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := acq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := acq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (acq *AuthCodeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(acq.driver.Dialect())
	t1 := builder.Table(authcode.Table)
	columns := acq.ctx.Fields
	if len(columns) == 0 {
		columns = authcode.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if acq.sql != nil {
		selector = acq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if acq.ctx.Unique != nil && *acq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range acq.modifiers {
		m(selector)
	}
	for _, p := range acq.predicates {
		p(selector)
	}
	for _, p := range acq.order {
		p(selector)
	}
	if offset := acq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := acq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (acq *AuthCodeQuery) ForUpdate(opts ...sql.LockOption) *AuthCodeQuery {
	if acq.driver.Dialect() == dialect.Postgres {
		acq.Unique(false)
	}
	acq.modifiers = append(acq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return acq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (acq *AuthCodeQuery) ForShare(opts ...sql.LockOption) *AuthCodeQuery {
	if acq.driver.Dialect() == dialect.Postgres {
		acq.Unique(false)
	}
	acq.modifiers = append(acq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return acq
}

// AuthCodeGroupBy is the group-by builder for AuthCode entities.
type AuthCodeGroupBy struct {
	selector
	build *AuthCodeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (acgb *AuthCodeGroupBy) Aggregate(fns ...AggregateFunc) *AuthCodeGroupBy {
	acgb.fns = append(acgb.fns, fns...)
	return acgb
}

// Scan applies the selector query and scans the result into the given value.
func (acgb *AuthCodeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acgb.build.ctx, "GroupBy")
	if err := acgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AuthCodeQuery, *AuthCodeGroupBy](ctx, acgb.build, acgb, acgb.build.inters, v)
}

func (acgb *AuthCodeGroupBy) sqlScan(ctx context.Context, root *AuthCodeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(acgb.fns))
	for _, fn := range acgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*acgb.flds)+len(acgb.fns))
		for _, f := range *acgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*acgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AuthCodeSelect is the builder for selecting fields of AuthCode entities.
type AuthCodeSelect struct {
	*AuthCodeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (acs *AuthCodeSelect) Aggregate(fns ...AggregateFunc) *AuthCodeSelect {
	acs.fns = append(acs.fns, fns...)
	return acs
}

// Scan applies the selector query and scans the result into the given value.
func (acs *AuthCodeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acs.ctx, "Select")
	if err := acs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AuthCodeQuery, *AuthCodeSelect](ctx, acs.AuthCodeQuery, acs, acs.inters, v)
}

func (acs *AuthCodeSelect) sqlScan(ctx context.Context, root *AuthCodeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(acs.fns))
	for _, fn := range acs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*acs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
