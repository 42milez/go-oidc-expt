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
	"github.com/42milez/go-oidc-server/app/ent/ent/authcode"
	"github.com/42milez/go-oidc-server/app/ent/ent/predicate"
	"github.com/42milez/go-oidc-server/app/ent/ent/redirecturi"
	"github.com/42milez/go-oidc-server/app/ent/ent/relyingparty"
	"github.com/42milez/go-oidc-server/app/typedef"
)

// RelyingPartyUpdate is the builder for updating RelyingParty entities.
type RelyingPartyUpdate struct {
	config
	hooks    []Hook
	mutation *RelyingPartyMutation
}

// Where appends a list predicates to the RelyingPartyUpdate builder.
func (rpu *RelyingPartyUpdate) Where(ps ...predicate.RelyingParty) *RelyingPartyUpdate {
	rpu.mutation.Where(ps...)
	return rpu
}

// SetClientSecret sets the "client_secret" field.
func (rpu *RelyingPartyUpdate) SetClientSecret(s string) *RelyingPartyUpdate {
	rpu.mutation.SetClientSecret(s)
	return rpu
}

// SetModifiedAt sets the "modified_at" field.
func (rpu *RelyingPartyUpdate) SetModifiedAt(t time.Time) *RelyingPartyUpdate {
	rpu.mutation.SetModifiedAt(t)
	return rpu
}

// AddAuthCodeIDs adds the "auth_codes" edge to the AuthCode entity by IDs.
func (rpu *RelyingPartyUpdate) AddAuthCodeIDs(ids ...typedef.AuthCodeID) *RelyingPartyUpdate {
	rpu.mutation.AddAuthCodeIDs(ids...)
	return rpu
}

// AddAuthCodes adds the "auth_codes" edges to the AuthCode entity.
func (rpu *RelyingPartyUpdate) AddAuthCodes(a ...*AuthCode) *RelyingPartyUpdate {
	ids := make([]typedef.AuthCodeID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rpu.AddAuthCodeIDs(ids...)
}

// AddRedirectURIIDs adds the "redirect_uris" edge to the RedirectUri entity by IDs.
func (rpu *RelyingPartyUpdate) AddRedirectURIIDs(ids ...typedef.RedirectUriID) *RelyingPartyUpdate {
	rpu.mutation.AddRedirectURIIDs(ids...)
	return rpu
}

// AddRedirectUris adds the "redirect_uris" edges to the RedirectUri entity.
func (rpu *RelyingPartyUpdate) AddRedirectUris(r ...*RedirectUri) *RelyingPartyUpdate {
	ids := make([]typedef.RedirectUriID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rpu.AddRedirectURIIDs(ids...)
}

// Mutation returns the RelyingPartyMutation object of the builder.
func (rpu *RelyingPartyUpdate) Mutation() *RelyingPartyMutation {
	return rpu.mutation
}

// ClearAuthCodes clears all "auth_codes" edges to the AuthCode entity.
func (rpu *RelyingPartyUpdate) ClearAuthCodes() *RelyingPartyUpdate {
	rpu.mutation.ClearAuthCodes()
	return rpu
}

// RemoveAuthCodeIDs removes the "auth_codes" edge to AuthCode entities by IDs.
func (rpu *RelyingPartyUpdate) RemoveAuthCodeIDs(ids ...typedef.AuthCodeID) *RelyingPartyUpdate {
	rpu.mutation.RemoveAuthCodeIDs(ids...)
	return rpu
}

// RemoveAuthCodes removes "auth_codes" edges to AuthCode entities.
func (rpu *RelyingPartyUpdate) RemoveAuthCodes(a ...*AuthCode) *RelyingPartyUpdate {
	ids := make([]typedef.AuthCodeID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rpu.RemoveAuthCodeIDs(ids...)
}

// ClearRedirectUris clears all "redirect_uris" edges to the RedirectUri entity.
func (rpu *RelyingPartyUpdate) ClearRedirectUris() *RelyingPartyUpdate {
	rpu.mutation.ClearRedirectUris()
	return rpu
}

// RemoveRedirectURIIDs removes the "redirect_uris" edge to RedirectUri entities by IDs.
func (rpu *RelyingPartyUpdate) RemoveRedirectURIIDs(ids ...typedef.RedirectUriID) *RelyingPartyUpdate {
	rpu.mutation.RemoveRedirectURIIDs(ids...)
	return rpu
}

// RemoveRedirectUris removes "redirect_uris" edges to RedirectUri entities.
func (rpu *RelyingPartyUpdate) RemoveRedirectUris(r ...*RedirectUri) *RelyingPartyUpdate {
	ids := make([]typedef.RedirectUriID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rpu.RemoveRedirectURIIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rpu *RelyingPartyUpdate) Save(ctx context.Context) (int, error) {
	rpu.defaults()
	return withHooks(ctx, rpu.sqlSave, rpu.mutation, rpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rpu *RelyingPartyUpdate) SaveX(ctx context.Context) int {
	affected, err := rpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rpu *RelyingPartyUpdate) Exec(ctx context.Context) error {
	_, err := rpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rpu *RelyingPartyUpdate) ExecX(ctx context.Context) {
	if err := rpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rpu *RelyingPartyUpdate) defaults() {
	if _, ok := rpu.mutation.ModifiedAt(); !ok {
		v := relyingparty.UpdateDefaultModifiedAt()
		rpu.mutation.SetModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rpu *RelyingPartyUpdate) check() error {
	if v, ok := rpu.mutation.ClientSecret(); ok {
		if err := relyingparty.ClientSecretValidator(v); err != nil {
			return &ValidationError{Name: "client_secret", err: fmt.Errorf(`ent: validator failed for field "RelyingParty.client_secret": %w`, err)}
		}
	}
	return nil
}

func (rpu *RelyingPartyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rpu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(relyingparty.Table, relyingparty.Columns, sqlgraph.NewFieldSpec(relyingparty.FieldID, field.TypeUint64))
	if ps := rpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rpu.mutation.ClientSecret(); ok {
		_spec.SetField(relyingparty.FieldClientSecret, field.TypeString, value)
	}
	if value, ok := rpu.mutation.ModifiedAt(); ok {
		_spec.SetField(relyingparty.FieldModifiedAt, field.TypeTime, value)
	}
	if rpu.mutation.AuthCodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   relyingparty.AuthCodesTable,
			Columns: []string{relyingparty.AuthCodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authcode.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rpu.mutation.RemovedAuthCodesIDs(); len(nodes) > 0 && !rpu.mutation.AuthCodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   relyingparty.AuthCodesTable,
			Columns: []string{relyingparty.AuthCodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authcode.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rpu.mutation.AuthCodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   relyingparty.AuthCodesTable,
			Columns: []string{relyingparty.AuthCodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authcode.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rpu.mutation.RedirectUrisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   relyingparty.RedirectUrisTable,
			Columns: []string{relyingparty.RedirectUrisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(redirecturi.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rpu.mutation.RemovedRedirectUrisIDs(); len(nodes) > 0 && !rpu.mutation.RedirectUrisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   relyingparty.RedirectUrisTable,
			Columns: []string{relyingparty.RedirectUrisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(redirecturi.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rpu.mutation.RedirectUrisIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   relyingparty.RedirectUrisTable,
			Columns: []string{relyingparty.RedirectUrisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(redirecturi.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{relyingparty.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rpu.mutation.done = true
	return n, nil
}

// RelyingPartyUpdateOne is the builder for updating a single RelyingParty entity.
type RelyingPartyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RelyingPartyMutation
}

// SetClientSecret sets the "client_secret" field.
func (rpuo *RelyingPartyUpdateOne) SetClientSecret(s string) *RelyingPartyUpdateOne {
	rpuo.mutation.SetClientSecret(s)
	return rpuo
}

// SetModifiedAt sets the "modified_at" field.
func (rpuo *RelyingPartyUpdateOne) SetModifiedAt(t time.Time) *RelyingPartyUpdateOne {
	rpuo.mutation.SetModifiedAt(t)
	return rpuo
}

// AddAuthCodeIDs adds the "auth_codes" edge to the AuthCode entity by IDs.
func (rpuo *RelyingPartyUpdateOne) AddAuthCodeIDs(ids ...typedef.AuthCodeID) *RelyingPartyUpdateOne {
	rpuo.mutation.AddAuthCodeIDs(ids...)
	return rpuo
}

// AddAuthCodes adds the "auth_codes" edges to the AuthCode entity.
func (rpuo *RelyingPartyUpdateOne) AddAuthCodes(a ...*AuthCode) *RelyingPartyUpdateOne {
	ids := make([]typedef.AuthCodeID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rpuo.AddAuthCodeIDs(ids...)
}

// AddRedirectURIIDs adds the "redirect_uris" edge to the RedirectUri entity by IDs.
func (rpuo *RelyingPartyUpdateOne) AddRedirectURIIDs(ids ...typedef.RedirectUriID) *RelyingPartyUpdateOne {
	rpuo.mutation.AddRedirectURIIDs(ids...)
	return rpuo
}

// AddRedirectUris adds the "redirect_uris" edges to the RedirectUri entity.
func (rpuo *RelyingPartyUpdateOne) AddRedirectUris(r ...*RedirectUri) *RelyingPartyUpdateOne {
	ids := make([]typedef.RedirectUriID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rpuo.AddRedirectURIIDs(ids...)
}

// Mutation returns the RelyingPartyMutation object of the builder.
func (rpuo *RelyingPartyUpdateOne) Mutation() *RelyingPartyMutation {
	return rpuo.mutation
}

// ClearAuthCodes clears all "auth_codes" edges to the AuthCode entity.
func (rpuo *RelyingPartyUpdateOne) ClearAuthCodes() *RelyingPartyUpdateOne {
	rpuo.mutation.ClearAuthCodes()
	return rpuo
}

// RemoveAuthCodeIDs removes the "auth_codes" edge to AuthCode entities by IDs.
func (rpuo *RelyingPartyUpdateOne) RemoveAuthCodeIDs(ids ...typedef.AuthCodeID) *RelyingPartyUpdateOne {
	rpuo.mutation.RemoveAuthCodeIDs(ids...)
	return rpuo
}

// RemoveAuthCodes removes "auth_codes" edges to AuthCode entities.
func (rpuo *RelyingPartyUpdateOne) RemoveAuthCodes(a ...*AuthCode) *RelyingPartyUpdateOne {
	ids := make([]typedef.AuthCodeID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rpuo.RemoveAuthCodeIDs(ids...)
}

// ClearRedirectUris clears all "redirect_uris" edges to the RedirectUri entity.
func (rpuo *RelyingPartyUpdateOne) ClearRedirectUris() *RelyingPartyUpdateOne {
	rpuo.mutation.ClearRedirectUris()
	return rpuo
}

// RemoveRedirectURIIDs removes the "redirect_uris" edge to RedirectUri entities by IDs.
func (rpuo *RelyingPartyUpdateOne) RemoveRedirectURIIDs(ids ...typedef.RedirectUriID) *RelyingPartyUpdateOne {
	rpuo.mutation.RemoveRedirectURIIDs(ids...)
	return rpuo
}

// RemoveRedirectUris removes "redirect_uris" edges to RedirectUri entities.
func (rpuo *RelyingPartyUpdateOne) RemoveRedirectUris(r ...*RedirectUri) *RelyingPartyUpdateOne {
	ids := make([]typedef.RedirectUriID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rpuo.RemoveRedirectURIIDs(ids...)
}

// Where appends a list predicates to the RelyingPartyUpdate builder.
func (rpuo *RelyingPartyUpdateOne) Where(ps ...predicate.RelyingParty) *RelyingPartyUpdateOne {
	rpuo.mutation.Where(ps...)
	return rpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rpuo *RelyingPartyUpdateOne) Select(field string, fields ...string) *RelyingPartyUpdateOne {
	rpuo.fields = append([]string{field}, fields...)
	return rpuo
}

// Save executes the query and returns the updated RelyingParty entity.
func (rpuo *RelyingPartyUpdateOne) Save(ctx context.Context) (*RelyingParty, error) {
	rpuo.defaults()
	return withHooks(ctx, rpuo.sqlSave, rpuo.mutation, rpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rpuo *RelyingPartyUpdateOne) SaveX(ctx context.Context) *RelyingParty {
	node, err := rpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rpuo *RelyingPartyUpdateOne) Exec(ctx context.Context) error {
	_, err := rpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rpuo *RelyingPartyUpdateOne) ExecX(ctx context.Context) {
	if err := rpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rpuo *RelyingPartyUpdateOne) defaults() {
	if _, ok := rpuo.mutation.ModifiedAt(); !ok {
		v := relyingparty.UpdateDefaultModifiedAt()
		rpuo.mutation.SetModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rpuo *RelyingPartyUpdateOne) check() error {
	if v, ok := rpuo.mutation.ClientSecret(); ok {
		if err := relyingparty.ClientSecretValidator(v); err != nil {
			return &ValidationError{Name: "client_secret", err: fmt.Errorf(`ent: validator failed for field "RelyingParty.client_secret": %w`, err)}
		}
	}
	return nil
}

func (rpuo *RelyingPartyUpdateOne) sqlSave(ctx context.Context) (_node *RelyingParty, err error) {
	if err := rpuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(relyingparty.Table, relyingparty.Columns, sqlgraph.NewFieldSpec(relyingparty.FieldID, field.TypeUint64))
	id, ok := rpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RelyingParty.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, relyingparty.FieldID)
		for _, f := range fields {
			if !relyingparty.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != relyingparty.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rpuo.mutation.ClientSecret(); ok {
		_spec.SetField(relyingparty.FieldClientSecret, field.TypeString, value)
	}
	if value, ok := rpuo.mutation.ModifiedAt(); ok {
		_spec.SetField(relyingparty.FieldModifiedAt, field.TypeTime, value)
	}
	if rpuo.mutation.AuthCodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   relyingparty.AuthCodesTable,
			Columns: []string{relyingparty.AuthCodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authcode.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rpuo.mutation.RemovedAuthCodesIDs(); len(nodes) > 0 && !rpuo.mutation.AuthCodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   relyingparty.AuthCodesTable,
			Columns: []string{relyingparty.AuthCodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authcode.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rpuo.mutation.AuthCodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   relyingparty.AuthCodesTable,
			Columns: []string{relyingparty.AuthCodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authcode.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rpuo.mutation.RedirectUrisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   relyingparty.RedirectUrisTable,
			Columns: []string{relyingparty.RedirectUrisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(redirecturi.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rpuo.mutation.RemovedRedirectUrisIDs(); len(nodes) > 0 && !rpuo.mutation.RedirectUrisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   relyingparty.RedirectUrisTable,
			Columns: []string{relyingparty.RedirectUrisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(redirecturi.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rpuo.mutation.RedirectUrisIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   relyingparty.RedirectUrisTable,
			Columns: []string{relyingparty.RedirectUrisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(redirecturi.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &RelyingParty{config: rpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{relyingparty.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rpuo.mutation.done = true
	return _node, nil
}
