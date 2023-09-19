// Code generated by ent, DO NOT EDIT.

package consent

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/42milez/go-oidc-server/app/ent/ent/predicate"
	"github.com/42milez/go-oidc-server/app/typedef"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Consent {
	return predicate.Consent(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Consent {
	return predicate.Consent(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Consent {
	return predicate.Consent(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Consent {
	return predicate.Consent(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Consent {
	return predicate.Consent(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Consent {
	return predicate.Consent(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Consent {
	return predicate.Consent(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Consent {
	return predicate.Consent(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Consent {
	return predicate.Consent(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v typedef.UserID) predicate.Consent {
	vc := uint64(v)
	return predicate.Consent(sql.FieldEQ(FieldUserID, vc))
}

// ClientID applies equality check predicate on the "client_id" field. It's identical to ClientIDEQ.
func ClientID(v typedef.ClientId) predicate.Consent {
	vc := string(v)
	return predicate.Consent(sql.FieldEQ(FieldClientID, vc))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Consent {
	return predicate.Consent(sql.FieldEQ(FieldCreatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v typedef.UserID) predicate.Consent {
	vc := uint64(v)
	return predicate.Consent(sql.FieldEQ(FieldUserID, vc))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v typedef.UserID) predicate.Consent {
	vc := uint64(v)
	return predicate.Consent(sql.FieldNEQ(FieldUserID, vc))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...typedef.UserID) predicate.Consent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.Consent(sql.FieldIn(FieldUserID, v...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...typedef.UserID) predicate.Consent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.Consent(sql.FieldNotIn(FieldUserID, v...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v typedef.UserID) predicate.Consent {
	vc := uint64(v)
	return predicate.Consent(sql.FieldGT(FieldUserID, vc))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v typedef.UserID) predicate.Consent {
	vc := uint64(v)
	return predicate.Consent(sql.FieldGTE(FieldUserID, vc))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v typedef.UserID) predicate.Consent {
	vc := uint64(v)
	return predicate.Consent(sql.FieldLT(FieldUserID, vc))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v typedef.UserID) predicate.Consent {
	vc := uint64(v)
	return predicate.Consent(sql.FieldLTE(FieldUserID, vc))
}

// ClientIDEQ applies the EQ predicate on the "client_id" field.
func ClientIDEQ(v typedef.ClientId) predicate.Consent {
	vc := string(v)
	return predicate.Consent(sql.FieldEQ(FieldClientID, vc))
}

// ClientIDNEQ applies the NEQ predicate on the "client_id" field.
func ClientIDNEQ(v typedef.ClientId) predicate.Consent {
	vc := string(v)
	return predicate.Consent(sql.FieldNEQ(FieldClientID, vc))
}

// ClientIDIn applies the In predicate on the "client_id" field.
func ClientIDIn(vs ...typedef.ClientId) predicate.Consent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.Consent(sql.FieldIn(FieldClientID, v...))
}

// ClientIDNotIn applies the NotIn predicate on the "client_id" field.
func ClientIDNotIn(vs ...typedef.ClientId) predicate.Consent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.Consent(sql.FieldNotIn(FieldClientID, v...))
}

// ClientIDGT applies the GT predicate on the "client_id" field.
func ClientIDGT(v typedef.ClientId) predicate.Consent {
	vc := string(v)
	return predicate.Consent(sql.FieldGT(FieldClientID, vc))
}

// ClientIDGTE applies the GTE predicate on the "client_id" field.
func ClientIDGTE(v typedef.ClientId) predicate.Consent {
	vc := string(v)
	return predicate.Consent(sql.FieldGTE(FieldClientID, vc))
}

// ClientIDLT applies the LT predicate on the "client_id" field.
func ClientIDLT(v typedef.ClientId) predicate.Consent {
	vc := string(v)
	return predicate.Consent(sql.FieldLT(FieldClientID, vc))
}

// ClientIDLTE applies the LTE predicate on the "client_id" field.
func ClientIDLTE(v typedef.ClientId) predicate.Consent {
	vc := string(v)
	return predicate.Consent(sql.FieldLTE(FieldClientID, vc))
}

// ClientIDContains applies the Contains predicate on the "client_id" field.
func ClientIDContains(v typedef.ClientId) predicate.Consent {
	vc := string(v)
	return predicate.Consent(sql.FieldContains(FieldClientID, vc))
}

// ClientIDHasPrefix applies the HasPrefix predicate on the "client_id" field.
func ClientIDHasPrefix(v typedef.ClientId) predicate.Consent {
	vc := string(v)
	return predicate.Consent(sql.FieldHasPrefix(FieldClientID, vc))
}

// ClientIDHasSuffix applies the HasSuffix predicate on the "client_id" field.
func ClientIDHasSuffix(v typedef.ClientId) predicate.Consent {
	vc := string(v)
	return predicate.Consent(sql.FieldHasSuffix(FieldClientID, vc))
}

// ClientIDEqualFold applies the EqualFold predicate on the "client_id" field.
func ClientIDEqualFold(v typedef.ClientId) predicate.Consent {
	vc := string(v)
	return predicate.Consent(sql.FieldEqualFold(FieldClientID, vc))
}

// ClientIDContainsFold applies the ContainsFold predicate on the "client_id" field.
func ClientIDContainsFold(v typedef.ClientId) predicate.Consent {
	vc := string(v)
	return predicate.Consent(sql.FieldContainsFold(FieldClientID, vc))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Consent {
	return predicate.Consent(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Consent {
	return predicate.Consent(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Consent {
	return predicate.Consent(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Consent {
	return predicate.Consent(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Consent {
	return predicate.Consent(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Consent {
	return predicate.Consent(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Consent {
	return predicate.Consent(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Consent {
	return predicate.Consent(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Consent) predicate.Consent {
	return predicate.Consent(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Consent) predicate.Consent {
	return predicate.Consent(func(s *sql.Selector) {
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
func Not(p predicate.Consent) predicate.Consent {
	return predicate.Consent(func(s *sql.Selector) {
		p(s.Not())
	})
}
