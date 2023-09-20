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
	"github.com/42milez/go-oidc-server/app/ent/ent/relyingparty"
	"github.com/42milez/go-oidc-server/app/typedef"
)

// AuthCodeUpdate is the builder for updating AuthCode entities.
type AuthCodeUpdate struct {
	config
	hooks    []Hook
	mutation *AuthCodeMutation
}

// Where appends a list predicates to the AuthCodeUpdate builder.
func (acu *AuthCodeUpdate) Where(ps ...predicate.AuthCode) *AuthCodeUpdate {
	acu.mutation.Where(ps...)
	return acu
}

// SetUsedAt sets the "used_at" field.
func (acu *AuthCodeUpdate) SetUsedAt(t time.Time) *AuthCodeUpdate {
	acu.mutation.SetUsedAt(t)
	return acu
}

// SetNillableUsedAt sets the "used_at" field if the given value is not nil.
func (acu *AuthCodeUpdate) SetNillableUsedAt(t *time.Time) *AuthCodeUpdate {
	if t != nil {
		acu.SetUsedAt(*t)
	}
	return acu
}

// ClearUsedAt clears the value of the "used_at" field.
func (acu *AuthCodeUpdate) ClearUsedAt() *AuthCodeUpdate {
	acu.mutation.ClearUsedAt()
	return acu
}

// SetRelyingPartyID sets the "relying_party" edge to the RelyingParty entity by ID.
func (acu *AuthCodeUpdate) SetRelyingPartyID(id typedef.RelyingPartyID) *AuthCodeUpdate {
	acu.mutation.SetRelyingPartyID(id)
	return acu
}

// SetNillableRelyingPartyID sets the "relying_party" edge to the RelyingParty entity by ID if the given value is not nil.
func (acu *AuthCodeUpdate) SetNillableRelyingPartyID(id *typedef.RelyingPartyID) *AuthCodeUpdate {
	if id != nil {
		acu = acu.SetRelyingPartyID(*id)
	}
	return acu
}

// SetRelyingParty sets the "relying_party" edge to the RelyingParty entity.
func (acu *AuthCodeUpdate) SetRelyingParty(r *RelyingParty) *AuthCodeUpdate {
	return acu.SetRelyingPartyID(r.ID)
}

// Mutation returns the AuthCodeMutation object of the builder.
func (acu *AuthCodeUpdate) Mutation() *AuthCodeMutation {
	return acu.mutation
}

// ClearRelyingParty clears the "relying_party" edge to the RelyingParty entity.
func (acu *AuthCodeUpdate) ClearRelyingParty() *AuthCodeUpdate {
	acu.mutation.ClearRelyingParty()
	return acu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (acu *AuthCodeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, acu.sqlSave, acu.mutation, acu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (acu *AuthCodeUpdate) SaveX(ctx context.Context) int {
	affected, err := acu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (acu *AuthCodeUpdate) Exec(ctx context.Context) error {
	_, err := acu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acu *AuthCodeUpdate) ExecX(ctx context.Context) {
	if err := acu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (acu *AuthCodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(authcode.Table, authcode.Columns, sqlgraph.NewFieldSpec(authcode.FieldID, field.TypeUint64))
	if ps := acu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acu.mutation.UsedAt(); ok {
		_spec.SetField(authcode.FieldUsedAt, field.TypeTime, value)
	}
	if acu.mutation.UsedAtCleared() {
		_spec.ClearField(authcode.FieldUsedAt, field.TypeTime)
	}
	if acu.mutation.RelyingPartyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authcode.RelyingPartyTable,
			Columns: []string{authcode.RelyingPartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(relyingparty.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := acu.mutation.RelyingPartyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authcode.RelyingPartyTable,
			Columns: []string{authcode.RelyingPartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(relyingparty.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, acu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authcode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	acu.mutation.done = true
	return n, nil
}

// AuthCodeUpdateOne is the builder for updating a single AuthCode entity.
type AuthCodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AuthCodeMutation
}

// SetUsedAt sets the "used_at" field.
func (acuo *AuthCodeUpdateOne) SetUsedAt(t time.Time) *AuthCodeUpdateOne {
	acuo.mutation.SetUsedAt(t)
	return acuo
}

// SetNillableUsedAt sets the "used_at" field if the given value is not nil.
func (acuo *AuthCodeUpdateOne) SetNillableUsedAt(t *time.Time) *AuthCodeUpdateOne {
	if t != nil {
		acuo.SetUsedAt(*t)
	}
	return acuo
}

// ClearUsedAt clears the value of the "used_at" field.
func (acuo *AuthCodeUpdateOne) ClearUsedAt() *AuthCodeUpdateOne {
	acuo.mutation.ClearUsedAt()
	return acuo
}

// SetRelyingPartyID sets the "relying_party" edge to the RelyingParty entity by ID.
func (acuo *AuthCodeUpdateOne) SetRelyingPartyID(id typedef.RelyingPartyID) *AuthCodeUpdateOne {
	acuo.mutation.SetRelyingPartyID(id)
	return acuo
}

// SetNillableRelyingPartyID sets the "relying_party" edge to the RelyingParty entity by ID if the given value is not nil.
func (acuo *AuthCodeUpdateOne) SetNillableRelyingPartyID(id *typedef.RelyingPartyID) *AuthCodeUpdateOne {
	if id != nil {
		acuo = acuo.SetRelyingPartyID(*id)
	}
	return acuo
}

// SetRelyingParty sets the "relying_party" edge to the RelyingParty entity.
func (acuo *AuthCodeUpdateOne) SetRelyingParty(r *RelyingParty) *AuthCodeUpdateOne {
	return acuo.SetRelyingPartyID(r.ID)
}

// Mutation returns the AuthCodeMutation object of the builder.
func (acuo *AuthCodeUpdateOne) Mutation() *AuthCodeMutation {
	return acuo.mutation
}

// ClearRelyingParty clears the "relying_party" edge to the RelyingParty entity.
func (acuo *AuthCodeUpdateOne) ClearRelyingParty() *AuthCodeUpdateOne {
	acuo.mutation.ClearRelyingParty()
	return acuo
}

// Where appends a list predicates to the AuthCodeUpdate builder.
func (acuo *AuthCodeUpdateOne) Where(ps ...predicate.AuthCode) *AuthCodeUpdateOne {
	acuo.mutation.Where(ps...)
	return acuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (acuo *AuthCodeUpdateOne) Select(field string, fields ...string) *AuthCodeUpdateOne {
	acuo.fields = append([]string{field}, fields...)
	return acuo
}

// Save executes the query and returns the updated AuthCode entity.
func (acuo *AuthCodeUpdateOne) Save(ctx context.Context) (*AuthCode, error) {
	return withHooks(ctx, acuo.sqlSave, acuo.mutation, acuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (acuo *AuthCodeUpdateOne) SaveX(ctx context.Context) *AuthCode {
	node, err := acuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (acuo *AuthCodeUpdateOne) Exec(ctx context.Context) error {
	_, err := acuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acuo *AuthCodeUpdateOne) ExecX(ctx context.Context) {
	if err := acuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (acuo *AuthCodeUpdateOne) sqlSave(ctx context.Context) (_node *AuthCode, err error) {
	_spec := sqlgraph.NewUpdateSpec(authcode.Table, authcode.Columns, sqlgraph.NewFieldSpec(authcode.FieldID, field.TypeUint64))
	id, ok := acuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AuthCode.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := acuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authcode.FieldID)
		for _, f := range fields {
			if !authcode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != authcode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := acuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acuo.mutation.UsedAt(); ok {
		_spec.SetField(authcode.FieldUsedAt, field.TypeTime, value)
	}
	if acuo.mutation.UsedAtCleared() {
		_spec.ClearField(authcode.FieldUsedAt, field.TypeTime)
	}
	if acuo.mutation.RelyingPartyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authcode.RelyingPartyTable,
			Columns: []string{authcode.RelyingPartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(relyingparty.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := acuo.mutation.RelyingPartyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authcode.RelyingPartyTable,
			Columns: []string{authcode.RelyingPartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(relyingparty.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AuthCode{config: acuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, acuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authcode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	acuo.mutation.done = true
	return _node, nil
}
