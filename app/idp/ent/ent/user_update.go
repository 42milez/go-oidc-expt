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
	"github.com/42milez/go-oidc-server/app/idp/ent/ent/authcode"
	"github.com/42milez/go-oidc-server/app/idp/ent/ent/predicate"
	"github.com/42milez/go-oidc-server/app/idp/ent/ent/redirecturi"
	"github.com/42milez/go-oidc-server/app/idp/ent/ent/user"
	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetPasswordHash sets the "password_hash" field.
func (uu *UserUpdate) SetPasswordHash(th typedef.PasswordHash) *UserUpdate {
	uu.mutation.SetPasswordHash(th)
	return uu
}

// SetTotpSecret sets the "totp_secret" field.
func (uu *UserUpdate) SetTotpSecret(s string) *UserUpdate {
	uu.mutation.SetTotpSecret(s)
	return uu
}

// SetNillableTotpSecret sets the "totp_secret" field if the given value is not nil.
func (uu *UserUpdate) SetNillableTotpSecret(s *string) *UserUpdate {
	if s != nil {
		uu.SetTotpSecret(*s)
	}
	return uu
}

// ClearTotpSecret clears the value of the "totp_secret" field.
func (uu *UserUpdate) ClearTotpSecret() *UserUpdate {
	uu.mutation.ClearTotpSecret()
	return uu
}

// SetModifiedAt sets the "modified_at" field.
func (uu *UserUpdate) SetModifiedAt(t time.Time) *UserUpdate {
	uu.mutation.SetModifiedAt(t)
	return uu
}

// AddAuthCodeIDs adds the "auth_codes" edge to the AuthCode entity by IDs.
func (uu *UserUpdate) AddAuthCodeIDs(ids ...int) *UserUpdate {
	uu.mutation.AddAuthCodeIDs(ids...)
	return uu
}

// AddAuthCodes adds the "auth_codes" edges to the AuthCode entity.
func (uu *UserUpdate) AddAuthCodes(a ...*AuthCode) *UserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.AddAuthCodeIDs(ids...)
}

// AddRedirectURIIDs adds the "redirect_uris" edge to the RedirectURI entity by IDs.
func (uu *UserUpdate) AddRedirectURIIDs(ids ...int) *UserUpdate {
	uu.mutation.AddRedirectURIIDs(ids...)
	return uu
}

// AddRedirectUris adds the "redirect_uris" edges to the RedirectURI entity.
func (uu *UserUpdate) AddRedirectUris(r ...*RedirectURI) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.AddRedirectURIIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearAuthCodes clears all "auth_codes" edges to the AuthCode entity.
func (uu *UserUpdate) ClearAuthCodes() *UserUpdate {
	uu.mutation.ClearAuthCodes()
	return uu
}

// RemoveAuthCodeIDs removes the "auth_codes" edge to AuthCode entities by IDs.
func (uu *UserUpdate) RemoveAuthCodeIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveAuthCodeIDs(ids...)
	return uu
}

// RemoveAuthCodes removes "auth_codes" edges to AuthCode entities.
func (uu *UserUpdate) RemoveAuthCodes(a ...*AuthCode) *UserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.RemoveAuthCodeIDs(ids...)
}

// ClearRedirectUris clears all "redirect_uris" edges to the RedirectURI entity.
func (uu *UserUpdate) ClearRedirectUris() *UserUpdate {
	uu.mutation.ClearRedirectUris()
	return uu
}

// RemoveRedirectURIIDs removes the "redirect_uris" edge to RedirectURI entities by IDs.
func (uu *UserUpdate) RemoveRedirectURIIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveRedirectURIIDs(ids...)
	return uu
}

// RemoveRedirectUris removes "redirect_uris" edges to RedirectURI entities.
func (uu *UserUpdate) RemoveRedirectUris(r ...*RedirectURI) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.RemoveRedirectURIIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.ModifiedAt(); !ok {
		v := user.UpdateDefaultModifiedAt()
		uu.mutation.SetModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.PasswordHash(); ok {
		if err := user.PasswordHashValidator(string(v)); err != nil {
			return &ValidationError{Name: "password_hash", err: fmt.Errorf(`ent: validator failed for field "User.password_hash": %w`, err)}
		}
	}
	if v, ok := uu.mutation.TotpSecret(); ok {
		if err := user.TotpSecretValidator(v); err != nil {
			return &ValidationError{Name: "totp_secret", err: fmt.Errorf(`ent: validator failed for field "User.totp_secret": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.PasswordHash(); ok {
		_spec.SetField(user.FieldPasswordHash, field.TypeString, value)
	}
	if value, ok := uu.mutation.TotpSecret(); ok {
		_spec.SetField(user.FieldTotpSecret, field.TypeString, value)
	}
	if uu.mutation.TotpSecretCleared() {
		_spec.ClearField(user.FieldTotpSecret, field.TypeString)
	}
	if value, ok := uu.mutation.ModifiedAt(); ok {
		_spec.SetField(user.FieldModifiedAt, field.TypeTime, value)
	}
	if uu.mutation.AuthCodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuthCodesTable,
			Columns: []string{user.AuthCodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authcode.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedAuthCodesIDs(); len(nodes) > 0 && !uu.mutation.AuthCodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuthCodesTable,
			Columns: []string{user.AuthCodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authcode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.AuthCodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuthCodesTable,
			Columns: []string{user.AuthCodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authcode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.RedirectUrisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RedirectUrisTable,
			Columns: []string{user.RedirectUrisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(redirecturi.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedRedirectUrisIDs(); len(nodes) > 0 && !uu.mutation.RedirectUrisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RedirectUrisTable,
			Columns: []string{user.RedirectUrisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(redirecturi.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RedirectUrisIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RedirectUrisTable,
			Columns: []string{user.RedirectUrisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(redirecturi.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetPasswordHash sets the "password_hash" field.
func (uuo *UserUpdateOne) SetPasswordHash(th typedef.PasswordHash) *UserUpdateOne {
	uuo.mutation.SetPasswordHash(th)
	return uuo
}

// SetTotpSecret sets the "totp_secret" field.
func (uuo *UserUpdateOne) SetTotpSecret(s string) *UserUpdateOne {
	uuo.mutation.SetTotpSecret(s)
	return uuo
}

// SetNillableTotpSecret sets the "totp_secret" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTotpSecret(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetTotpSecret(*s)
	}
	return uuo
}

// ClearTotpSecret clears the value of the "totp_secret" field.
func (uuo *UserUpdateOne) ClearTotpSecret() *UserUpdateOne {
	uuo.mutation.ClearTotpSecret()
	return uuo
}

// SetModifiedAt sets the "modified_at" field.
func (uuo *UserUpdateOne) SetModifiedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetModifiedAt(t)
	return uuo
}

// AddAuthCodeIDs adds the "auth_codes" edge to the AuthCode entity by IDs.
func (uuo *UserUpdateOne) AddAuthCodeIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddAuthCodeIDs(ids...)
	return uuo
}

// AddAuthCodes adds the "auth_codes" edges to the AuthCode entity.
func (uuo *UserUpdateOne) AddAuthCodes(a ...*AuthCode) *UserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.AddAuthCodeIDs(ids...)
}

// AddRedirectURIIDs adds the "redirect_uris" edge to the RedirectURI entity by IDs.
func (uuo *UserUpdateOne) AddRedirectURIIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddRedirectURIIDs(ids...)
	return uuo
}

// AddRedirectUris adds the "redirect_uris" edges to the RedirectURI entity.
func (uuo *UserUpdateOne) AddRedirectUris(r ...*RedirectURI) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.AddRedirectURIIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearAuthCodes clears all "auth_codes" edges to the AuthCode entity.
func (uuo *UserUpdateOne) ClearAuthCodes() *UserUpdateOne {
	uuo.mutation.ClearAuthCodes()
	return uuo
}

// RemoveAuthCodeIDs removes the "auth_codes" edge to AuthCode entities by IDs.
func (uuo *UserUpdateOne) RemoveAuthCodeIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveAuthCodeIDs(ids...)
	return uuo
}

// RemoveAuthCodes removes "auth_codes" edges to AuthCode entities.
func (uuo *UserUpdateOne) RemoveAuthCodes(a ...*AuthCode) *UserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.RemoveAuthCodeIDs(ids...)
}

// ClearRedirectUris clears all "redirect_uris" edges to the RedirectURI entity.
func (uuo *UserUpdateOne) ClearRedirectUris() *UserUpdateOne {
	uuo.mutation.ClearRedirectUris()
	return uuo
}

// RemoveRedirectURIIDs removes the "redirect_uris" edge to RedirectURI entities by IDs.
func (uuo *UserUpdateOne) RemoveRedirectURIIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveRedirectURIIDs(ids...)
	return uuo
}

// RemoveRedirectUris removes "redirect_uris" edges to RedirectURI entities.
func (uuo *UserUpdateOne) RemoveRedirectUris(r ...*RedirectURI) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.RemoveRedirectURIIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.ModifiedAt(); !ok {
		v := user.UpdateDefaultModifiedAt()
		uuo.mutation.SetModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.PasswordHash(); ok {
		if err := user.PasswordHashValidator(string(v)); err != nil {
			return &ValidationError{Name: "password_hash", err: fmt.Errorf(`ent: validator failed for field "User.password_hash": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.TotpSecret(); ok {
		if err := user.TotpSecretValidator(v); err != nil {
			return &ValidationError{Name: "totp_secret", err: fmt.Errorf(`ent: validator failed for field "User.totp_secret": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.PasswordHash(); ok {
		_spec.SetField(user.FieldPasswordHash, field.TypeString, value)
	}
	if value, ok := uuo.mutation.TotpSecret(); ok {
		_spec.SetField(user.FieldTotpSecret, field.TypeString, value)
	}
	if uuo.mutation.TotpSecretCleared() {
		_spec.ClearField(user.FieldTotpSecret, field.TypeString)
	}
	if value, ok := uuo.mutation.ModifiedAt(); ok {
		_spec.SetField(user.FieldModifiedAt, field.TypeTime, value)
	}
	if uuo.mutation.AuthCodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuthCodesTable,
			Columns: []string{user.AuthCodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authcode.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedAuthCodesIDs(); len(nodes) > 0 && !uuo.mutation.AuthCodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuthCodesTable,
			Columns: []string{user.AuthCodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authcode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.AuthCodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuthCodesTable,
			Columns: []string{user.AuthCodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authcode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.RedirectUrisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RedirectUrisTable,
			Columns: []string{user.RedirectUrisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(redirecturi.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedRedirectUrisIDs(); len(nodes) > 0 && !uuo.mutation.RedirectUrisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RedirectUrisTable,
			Columns: []string{user.RedirectUrisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(redirecturi.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RedirectUrisIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RedirectUrisTable,
			Columns: []string{user.RedirectUrisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(redirecturi.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
