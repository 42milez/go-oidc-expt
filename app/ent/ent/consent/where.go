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

// RelyingPartyID applies equality check predicate on the "relying_party_id" field. It's identical to RelyingPartyIDEQ.
func RelyingPartyID(v int) predicate.Consent {
	return predicate.Consent(sql.FieldEQ(FieldRelyingPartyID, v))
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

// RelyingPartyIDEQ applies the EQ predicate on the "relying_party_id" field.
func RelyingPartyIDEQ(v int) predicate.Consent {
	return predicate.Consent(sql.FieldEQ(FieldRelyingPartyID, v))
}

// RelyingPartyIDNEQ applies the NEQ predicate on the "relying_party_id" field.
func RelyingPartyIDNEQ(v int) predicate.Consent {
	return predicate.Consent(sql.FieldNEQ(FieldRelyingPartyID, v))
}

// RelyingPartyIDIn applies the In predicate on the "relying_party_id" field.
func RelyingPartyIDIn(vs ...int) predicate.Consent {
	return predicate.Consent(sql.FieldIn(FieldRelyingPartyID, vs...))
}

// RelyingPartyIDNotIn applies the NotIn predicate on the "relying_party_id" field.
func RelyingPartyIDNotIn(vs ...int) predicate.Consent {
	return predicate.Consent(sql.FieldNotIn(FieldRelyingPartyID, vs...))
}

// RelyingPartyIDGT applies the GT predicate on the "relying_party_id" field.
func RelyingPartyIDGT(v int) predicate.Consent {
	return predicate.Consent(sql.FieldGT(FieldRelyingPartyID, v))
}

// RelyingPartyIDGTE applies the GTE predicate on the "relying_party_id" field.
func RelyingPartyIDGTE(v int) predicate.Consent {
	return predicate.Consent(sql.FieldGTE(FieldRelyingPartyID, v))
}

// RelyingPartyIDLT applies the LT predicate on the "relying_party_id" field.
func RelyingPartyIDLT(v int) predicate.Consent {
	return predicate.Consent(sql.FieldLT(FieldRelyingPartyID, v))
}

// RelyingPartyIDLTE applies the LTE predicate on the "relying_party_id" field.
func RelyingPartyIDLTE(v int) predicate.Consent {
	return predicate.Consent(sql.FieldLTE(FieldRelyingPartyID, v))
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
