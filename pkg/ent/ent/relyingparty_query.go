// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/42milez/go-oidc-expt/pkg/ent/ent/authcode"
	"github.com/42milez/go-oidc-expt/pkg/ent/ent/predicate"
	"github.com/42milez/go-oidc-expt/pkg/ent/ent/redirecturi"
	"github.com/42milez/go-oidc-expt/pkg/ent/ent/relyingparty"
	"github.com/42milez/go-oidc-expt/pkg/typedef"
)

// RelyingPartyQuery is the builder for querying RelyingParty entities.
type RelyingPartyQuery struct {
	config
	ctx              *QueryContext
	order            []relyingparty.OrderOption
	inters           []Interceptor
	predicates       []predicate.RelyingParty
	withAuthCodes    *AuthCodeQuery
	withRedirectUris *RedirectURIQuery
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RelyingPartyQuery builder.
func (rpq *RelyingPartyQuery) Where(ps ...predicate.RelyingParty) *RelyingPartyQuery {
	rpq.predicates = append(rpq.predicates, ps...)
	return rpq
}

// Limit the number of records to be returned by this query.
func (rpq *RelyingPartyQuery) Limit(limit int) *RelyingPartyQuery {
	rpq.ctx.Limit = &limit
	return rpq
}

// Offset to start from.
func (rpq *RelyingPartyQuery) Offset(offset int) *RelyingPartyQuery {
	rpq.ctx.Offset = &offset
	return rpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rpq *RelyingPartyQuery) Unique(unique bool) *RelyingPartyQuery {
	rpq.ctx.Unique = &unique
	return rpq
}

// Order specifies how the records should be ordered.
func (rpq *RelyingPartyQuery) Order(o ...relyingparty.OrderOption) *RelyingPartyQuery {
	rpq.order = append(rpq.order, o...)
	return rpq
}

// QueryAuthCodes chains the current query on the "auth_codes" edge.
func (rpq *RelyingPartyQuery) QueryAuthCodes() *AuthCodeQuery {
	query := (&AuthCodeClient{config: rpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(relyingparty.Table, relyingparty.FieldID, selector),
			sqlgraph.To(authcode.Table, authcode.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, relyingparty.AuthCodesTable, relyingparty.AuthCodesColumn),
		)
		fromU = sqlgraph.SetNeighbors(rpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRedirectUris chains the current query on the "redirect_uris" edge.
func (rpq *RelyingPartyQuery) QueryRedirectUris() *RedirectURIQuery {
	query := (&RedirectURIClient{config: rpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(relyingparty.Table, relyingparty.FieldID, selector),
			sqlgraph.To(redirecturi.Table, redirecturi.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, relyingparty.RedirectUrisTable, relyingparty.RedirectUrisColumn),
		)
		fromU = sqlgraph.SetNeighbors(rpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RelyingParty entity from the query.
// Returns a *NotFoundError when no RelyingParty was found.
func (rpq *RelyingPartyQuery) First(ctx context.Context) (*RelyingParty, error) {
	nodes, err := rpq.Limit(1).All(setContextOp(ctx, rpq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{relyingparty.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rpq *RelyingPartyQuery) FirstX(ctx context.Context) *RelyingParty {
	node, err := rpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RelyingParty ID from the query.
// Returns a *NotFoundError when no RelyingParty ID was found.
func (rpq *RelyingPartyQuery) FirstID(ctx context.Context) (id typedef.RelyingPartyID, err error) {
	var ids []typedef.RelyingPartyID
	if ids, err = rpq.Limit(1).IDs(setContextOp(ctx, rpq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{relyingparty.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rpq *RelyingPartyQuery) FirstIDX(ctx context.Context) typedef.RelyingPartyID {
	id, err := rpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RelyingParty entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RelyingParty entity is found.
// Returns a *NotFoundError when no RelyingParty entities are found.
func (rpq *RelyingPartyQuery) Only(ctx context.Context) (*RelyingParty, error) {
	nodes, err := rpq.Limit(2).All(setContextOp(ctx, rpq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{relyingparty.Label}
	default:
		return nil, &NotSingularError{relyingparty.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rpq *RelyingPartyQuery) OnlyX(ctx context.Context) *RelyingParty {
	node, err := rpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RelyingParty ID in the query.
// Returns a *NotSingularError when more than one RelyingParty ID is found.
// Returns a *NotFoundError when no entities are found.
func (rpq *RelyingPartyQuery) OnlyID(ctx context.Context) (id typedef.RelyingPartyID, err error) {
	var ids []typedef.RelyingPartyID
	if ids, err = rpq.Limit(2).IDs(setContextOp(ctx, rpq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{relyingparty.Label}
	default:
		err = &NotSingularError{relyingparty.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rpq *RelyingPartyQuery) OnlyIDX(ctx context.Context) typedef.RelyingPartyID {
	id, err := rpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RelyingParties.
func (rpq *RelyingPartyQuery) All(ctx context.Context) ([]*RelyingParty, error) {
	ctx = setContextOp(ctx, rpq.ctx, "All")
	if err := rpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RelyingParty, *RelyingPartyQuery]()
	return withInterceptors[[]*RelyingParty](ctx, rpq, qr, rpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rpq *RelyingPartyQuery) AllX(ctx context.Context) []*RelyingParty {
	nodes, err := rpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RelyingParty IDs.
func (rpq *RelyingPartyQuery) IDs(ctx context.Context) (ids []typedef.RelyingPartyID, err error) {
	if rpq.ctx.Unique == nil && rpq.path != nil {
		rpq.Unique(true)
	}
	ctx = setContextOp(ctx, rpq.ctx, "IDs")
	if err = rpq.Select(relyingparty.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rpq *RelyingPartyQuery) IDsX(ctx context.Context) []typedef.RelyingPartyID {
	ids, err := rpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rpq *RelyingPartyQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rpq.ctx, "Count")
	if err := rpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rpq, querierCount[*RelyingPartyQuery](), rpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rpq *RelyingPartyQuery) CountX(ctx context.Context) int {
	count, err := rpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rpq *RelyingPartyQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rpq.ctx, "Exist")
	switch _, err := rpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rpq *RelyingPartyQuery) ExistX(ctx context.Context) bool {
	exist, err := rpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RelyingPartyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rpq *RelyingPartyQuery) Clone() *RelyingPartyQuery {
	if rpq == nil {
		return nil
	}
	return &RelyingPartyQuery{
		config:           rpq.config,
		ctx:              rpq.ctx.Clone(),
		order:            append([]relyingparty.OrderOption{}, rpq.order...),
		inters:           append([]Interceptor{}, rpq.inters...),
		predicates:       append([]predicate.RelyingParty{}, rpq.predicates...),
		withAuthCodes:    rpq.withAuthCodes.Clone(),
		withRedirectUris: rpq.withRedirectUris.Clone(),
		// clone intermediate query.
		sql:  rpq.sql.Clone(),
		path: rpq.path,
	}
}

// WithAuthCodes tells the query-builder to eager-load the nodes that are connected to
// the "auth_codes" edge. The optional arguments are used to configure the query builder of the edge.
func (rpq *RelyingPartyQuery) WithAuthCodes(opts ...func(*AuthCodeQuery)) *RelyingPartyQuery {
	query := (&AuthCodeClient{config: rpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rpq.withAuthCodes = query
	return rpq
}

// WithRedirectUris tells the query-builder to eager-load the nodes that are connected to
// the "redirect_uris" edge. The optional arguments are used to configure the query builder of the edge.
func (rpq *RelyingPartyQuery) WithRedirectUris(opts ...func(*RedirectURIQuery)) *RelyingPartyQuery {
	query := (&RedirectURIClient{config: rpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rpq.withRedirectUris = query
	return rpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ClientID typedef.ClientID `json:"client_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RelyingParty.Query().
//		GroupBy(relyingparty.FieldClientID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rpq *RelyingPartyQuery) GroupBy(field string, fields ...string) *RelyingPartyGroupBy {
	rpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RelyingPartyGroupBy{build: rpq}
	grbuild.flds = &rpq.ctx.Fields
	grbuild.label = relyingparty.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ClientID typedef.ClientID `json:"client_id,omitempty"`
//	}
//
//	client.RelyingParty.Query().
//		Select(relyingparty.FieldClientID).
//		Scan(ctx, &v)
func (rpq *RelyingPartyQuery) Select(fields ...string) *RelyingPartySelect {
	rpq.ctx.Fields = append(rpq.ctx.Fields, fields...)
	sbuild := &RelyingPartySelect{RelyingPartyQuery: rpq}
	sbuild.label = relyingparty.Label
	sbuild.flds, sbuild.scan = &rpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RelyingPartySelect configured with the given aggregations.
func (rpq *RelyingPartyQuery) Aggregate(fns ...AggregateFunc) *RelyingPartySelect {
	return rpq.Select().Aggregate(fns...)
}

func (rpq *RelyingPartyQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rpq); err != nil {
				return err
			}
		}
	}
	for _, f := range rpq.ctx.Fields {
		if !relyingparty.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rpq.path != nil {
		prev, err := rpq.path(ctx)
		if err != nil {
			return err
		}
		rpq.sql = prev
	}
	return nil
}

func (rpq *RelyingPartyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RelyingParty, error) {
	var (
		nodes       = []*RelyingParty{}
		_spec       = rpq.querySpec()
		loadedTypes = [2]bool{
			rpq.withAuthCodes != nil,
			rpq.withRedirectUris != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RelyingParty).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RelyingParty{config: rpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(rpq.modifiers) > 0 {
		_spec.Modifiers = rpq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rpq.withAuthCodes; query != nil {
		if err := rpq.loadAuthCodes(ctx, query, nodes,
			func(n *RelyingParty) { n.Edges.AuthCodes = []*AuthCode{} },
			func(n *RelyingParty, e *AuthCode) { n.Edges.AuthCodes = append(n.Edges.AuthCodes, e) }); err != nil {
			return nil, err
		}
	}
	if query := rpq.withRedirectUris; query != nil {
		if err := rpq.loadRedirectUris(ctx, query, nodes,
			func(n *RelyingParty) { n.Edges.RedirectUris = []*RedirectURI{} },
			func(n *RelyingParty, e *RedirectURI) { n.Edges.RedirectUris = append(n.Edges.RedirectUris, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rpq *RelyingPartyQuery) loadAuthCodes(ctx context.Context, query *AuthCodeQuery, nodes []*RelyingParty, init func(*RelyingParty), assign func(*RelyingParty, *AuthCode)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[typedef.RelyingPartyID]*RelyingParty)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(authcode.FieldRelyingPartyID)
	}
	query.Where(predicate.AuthCode(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(relyingparty.AuthCodesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.RelyingPartyID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "relying_party_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (rpq *RelyingPartyQuery) loadRedirectUris(ctx context.Context, query *RedirectURIQuery, nodes []*RelyingParty, init func(*RelyingParty), assign func(*RelyingParty, *RedirectURI)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[typedef.RelyingPartyID]*RelyingParty)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(redirecturi.FieldRelyingPartyID)
	}
	query.Where(predicate.RedirectURI(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(relyingparty.RedirectUrisColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.RelyingPartyID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "relying_party_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rpq *RelyingPartyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rpq.querySpec()
	if len(rpq.modifiers) > 0 {
		_spec.Modifiers = rpq.modifiers
	}
	_spec.Node.Columns = rpq.ctx.Fields
	if len(rpq.ctx.Fields) > 0 {
		_spec.Unique = rpq.ctx.Unique != nil && *rpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rpq.driver, _spec)
}

func (rpq *RelyingPartyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(relyingparty.Table, relyingparty.Columns, sqlgraph.NewFieldSpec(relyingparty.FieldID, field.TypeUint64))
	_spec.From = rpq.sql
	if unique := rpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rpq.path != nil {
		_spec.Unique = true
	}
	if fields := rpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, relyingparty.FieldID)
		for i := range fields {
			if fields[i] != relyingparty.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rpq *RelyingPartyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rpq.driver.Dialect())
	t1 := builder.Table(relyingparty.Table)
	columns := rpq.ctx.Fields
	if len(columns) == 0 {
		columns = relyingparty.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rpq.sql != nil {
		selector = rpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rpq.ctx.Unique != nil && *rpq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range rpq.modifiers {
		m(selector)
	}
	for _, p := range rpq.predicates {
		p(selector)
	}
	for _, p := range rpq.order {
		p(selector)
	}
	if offset := rpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (rpq *RelyingPartyQuery) ForUpdate(opts ...sql.LockOption) *RelyingPartyQuery {
	if rpq.driver.Dialect() == dialect.Postgres {
		rpq.Unique(false)
	}
	rpq.modifiers = append(rpq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return rpq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (rpq *RelyingPartyQuery) ForShare(opts ...sql.LockOption) *RelyingPartyQuery {
	if rpq.driver.Dialect() == dialect.Postgres {
		rpq.Unique(false)
	}
	rpq.modifiers = append(rpq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return rpq
}

// RelyingPartyGroupBy is the group-by builder for RelyingParty entities.
type RelyingPartyGroupBy struct {
	selector
	build *RelyingPartyQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rpgb *RelyingPartyGroupBy) Aggregate(fns ...AggregateFunc) *RelyingPartyGroupBy {
	rpgb.fns = append(rpgb.fns, fns...)
	return rpgb
}

// Scan applies the selector query and scans the result into the given value.
func (rpgb *RelyingPartyGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rpgb.build.ctx, "GroupBy")
	if err := rpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RelyingPartyQuery, *RelyingPartyGroupBy](ctx, rpgb.build, rpgb, rpgb.build.inters, v)
}

func (rpgb *RelyingPartyGroupBy) sqlScan(ctx context.Context, root *RelyingPartyQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rpgb.fns))
	for _, fn := range rpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rpgb.flds)+len(rpgb.fns))
		for _, f := range *rpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RelyingPartySelect is the builder for selecting fields of RelyingParty entities.
type RelyingPartySelect struct {
	*RelyingPartyQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rps *RelyingPartySelect) Aggregate(fns ...AggregateFunc) *RelyingPartySelect {
	rps.fns = append(rps.fns, fns...)
	return rps
}

// Scan applies the selector query and scans the result into the given value.
func (rps *RelyingPartySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rps.ctx, "Select")
	if err := rps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RelyingPartyQuery, *RelyingPartySelect](ctx, rps.RelyingPartyQuery, rps, rps.inters, v)
}

func (rps *RelyingPartySelect) sqlScan(ctx context.Context, root *RelyingPartyQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rps.fns))
	for _, fn := range rps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
