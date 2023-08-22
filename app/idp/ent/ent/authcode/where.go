// Code generated by ent, DO NOT EDIT.

package authcode

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/42milez/go-oidc-server/app/idp/ent/ent/predicate"
	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldLTE(FieldID, id))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldEQ(FieldCode, v))
}

// ExpireAt applies equality check predicate on the "expire_at" field. It's identical to ExpireAtEQ.
func ExpireAt(v time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldEQ(FieldExpireAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldEQ(FieldCreatedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v typedef.UserID) predicate.AuthCode {
	vc := string(v)
	return predicate.AuthCode(sql.FieldEQ(FieldUserID, vc))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldContainsFold(FieldCode, v))
}

// ExpireAtEQ applies the EQ predicate on the "expire_at" field.
func ExpireAtEQ(v time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldEQ(FieldExpireAt, v))
}

// ExpireAtNEQ applies the NEQ predicate on the "expire_at" field.
func ExpireAtNEQ(v time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldNEQ(FieldExpireAt, v))
}

// ExpireAtIn applies the In predicate on the "expire_at" field.
func ExpireAtIn(vs ...time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldIn(FieldExpireAt, vs...))
}

// ExpireAtNotIn applies the NotIn predicate on the "expire_at" field.
func ExpireAtNotIn(vs ...time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldNotIn(FieldExpireAt, vs...))
}

// ExpireAtGT applies the GT predicate on the "expire_at" field.
func ExpireAtGT(v time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldGT(FieldExpireAt, v))
}

// ExpireAtGTE applies the GTE predicate on the "expire_at" field.
func ExpireAtGTE(v time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldGTE(FieldExpireAt, v))
}

// ExpireAtLT applies the LT predicate on the "expire_at" field.
func ExpireAtLT(v time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldLT(FieldExpireAt, v))
}

// ExpireAtLTE applies the LTE predicate on the "expire_at" field.
func ExpireAtLTE(v time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldLTE(FieldExpireAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldLTE(FieldCreatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v typedef.UserID) predicate.AuthCode {
	vc := string(v)
	return predicate.AuthCode(sql.FieldEQ(FieldUserID, vc))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v typedef.UserID) predicate.AuthCode {
	vc := string(v)
	return predicate.AuthCode(sql.FieldNEQ(FieldUserID, vc))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...typedef.UserID) predicate.AuthCode {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.AuthCode(sql.FieldIn(FieldUserID, v...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...typedef.UserID) predicate.AuthCode {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.AuthCode(sql.FieldNotIn(FieldUserID, v...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v typedef.UserID) predicate.AuthCode {
	vc := string(v)
	return predicate.AuthCode(sql.FieldGT(FieldUserID, vc))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v typedef.UserID) predicate.AuthCode {
	vc := string(v)
	return predicate.AuthCode(sql.FieldGTE(FieldUserID, vc))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v typedef.UserID) predicate.AuthCode {
	vc := string(v)
	return predicate.AuthCode(sql.FieldLT(FieldUserID, vc))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v typedef.UserID) predicate.AuthCode {
	vc := string(v)
	return predicate.AuthCode(sql.FieldLTE(FieldUserID, vc))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v typedef.UserID) predicate.AuthCode {
	vc := string(v)
	return predicate.AuthCode(sql.FieldContains(FieldUserID, vc))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v typedef.UserID) predicate.AuthCode {
	vc := string(v)
	return predicate.AuthCode(sql.FieldHasPrefix(FieldUserID, vc))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v typedef.UserID) predicate.AuthCode {
	vc := string(v)
	return predicate.AuthCode(sql.FieldHasSuffix(FieldUserID, vc))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v typedef.UserID) predicate.AuthCode {
	vc := string(v)
	return predicate.AuthCode(sql.FieldEqualFold(FieldUserID, vc))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v typedef.UserID) predicate.AuthCode {
	vc := string(v)
	return predicate.AuthCode(sql.FieldContainsFold(FieldUserID, vc))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AuthCode) predicate.AuthCode {
	return predicate.AuthCode(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AuthCode) predicate.AuthCode {
	return predicate.AuthCode(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AuthCode) predicate.AuthCode {
	return predicate.AuthCode(func(s *sql.Selector) {
		p(s.Not())
	})
}
