// Code generated by ent, DO NOT EDIT.

package authcode

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/42milez/go-oidc-server/app/ent/ent/predicate"
	"github.com/42milez/go-oidc-server/app/typedef"
)

// ID filters vertices based on their ID field.
func ID(id typedef.AuthCodeID) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id typedef.AuthCodeID) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id typedef.AuthCodeID) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...typedef.AuthCodeID) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...typedef.AuthCodeID) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id typedef.AuthCodeID) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id typedef.AuthCodeID) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id typedef.AuthCodeID) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id typedef.AuthCodeID) predicate.AuthCode {
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

// UsedAt applies equality check predicate on the "used_at" field. It's identical to UsedAtEQ.
func UsedAt(v time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldEQ(FieldUsedAt, v))
}

// RelyingPartyAuthCodes applies equality check predicate on the "relying_party_auth_codes" field. It's identical to RelyingPartyAuthCodesEQ.
func RelyingPartyAuthCodes(v typedef.RelyingPartyID) predicate.AuthCode {
	vc := uint64(v)
	return predicate.AuthCode(sql.FieldEQ(FieldRelyingPartyAuthCodes, vc))
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

// UsedAtEQ applies the EQ predicate on the "used_at" field.
func UsedAtEQ(v time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldEQ(FieldUsedAt, v))
}

// UsedAtNEQ applies the NEQ predicate on the "used_at" field.
func UsedAtNEQ(v time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldNEQ(FieldUsedAt, v))
}

// UsedAtIn applies the In predicate on the "used_at" field.
func UsedAtIn(vs ...time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldIn(FieldUsedAt, vs...))
}

// UsedAtNotIn applies the NotIn predicate on the "used_at" field.
func UsedAtNotIn(vs ...time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldNotIn(FieldUsedAt, vs...))
}

// UsedAtGT applies the GT predicate on the "used_at" field.
func UsedAtGT(v time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldGT(FieldUsedAt, v))
}

// UsedAtGTE applies the GTE predicate on the "used_at" field.
func UsedAtGTE(v time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldGTE(FieldUsedAt, v))
}

// UsedAtLT applies the LT predicate on the "used_at" field.
func UsedAtLT(v time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldLT(FieldUsedAt, v))
}

// UsedAtLTE applies the LTE predicate on the "used_at" field.
func UsedAtLTE(v time.Time) predicate.AuthCode {
	return predicate.AuthCode(sql.FieldLTE(FieldUsedAt, v))
}

// UsedAtIsNil applies the IsNil predicate on the "used_at" field.
func UsedAtIsNil() predicate.AuthCode {
	return predicate.AuthCode(sql.FieldIsNull(FieldUsedAt))
}

// UsedAtNotNil applies the NotNil predicate on the "used_at" field.
func UsedAtNotNil() predicate.AuthCode {
	return predicate.AuthCode(sql.FieldNotNull(FieldUsedAt))
}

// RelyingPartyAuthCodesEQ applies the EQ predicate on the "relying_party_auth_codes" field.
func RelyingPartyAuthCodesEQ(v typedef.RelyingPartyID) predicate.AuthCode {
	vc := uint64(v)
	return predicate.AuthCode(sql.FieldEQ(FieldRelyingPartyAuthCodes, vc))
}

// RelyingPartyAuthCodesNEQ applies the NEQ predicate on the "relying_party_auth_codes" field.
func RelyingPartyAuthCodesNEQ(v typedef.RelyingPartyID) predicate.AuthCode {
	vc := uint64(v)
	return predicate.AuthCode(sql.FieldNEQ(FieldRelyingPartyAuthCodes, vc))
}

// RelyingPartyAuthCodesIn applies the In predicate on the "relying_party_auth_codes" field.
func RelyingPartyAuthCodesIn(vs ...typedef.RelyingPartyID) predicate.AuthCode {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.AuthCode(sql.FieldIn(FieldRelyingPartyAuthCodes, v...))
}

// RelyingPartyAuthCodesNotIn applies the NotIn predicate on the "relying_party_auth_codes" field.
func RelyingPartyAuthCodesNotIn(vs ...typedef.RelyingPartyID) predicate.AuthCode {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.AuthCode(sql.FieldNotIn(FieldRelyingPartyAuthCodes, v...))
}

// RelyingPartyAuthCodesGT applies the GT predicate on the "relying_party_auth_codes" field.
func RelyingPartyAuthCodesGT(v typedef.RelyingPartyID) predicate.AuthCode {
	vc := uint64(v)
	return predicate.AuthCode(sql.FieldGT(FieldRelyingPartyAuthCodes, vc))
}

// RelyingPartyAuthCodesGTE applies the GTE predicate on the "relying_party_auth_codes" field.
func RelyingPartyAuthCodesGTE(v typedef.RelyingPartyID) predicate.AuthCode {
	vc := uint64(v)
	return predicate.AuthCode(sql.FieldGTE(FieldRelyingPartyAuthCodes, vc))
}

// RelyingPartyAuthCodesLT applies the LT predicate on the "relying_party_auth_codes" field.
func RelyingPartyAuthCodesLT(v typedef.RelyingPartyID) predicate.AuthCode {
	vc := uint64(v)
	return predicate.AuthCode(sql.FieldLT(FieldRelyingPartyAuthCodes, vc))
}

// RelyingPartyAuthCodesLTE applies the LTE predicate on the "relying_party_auth_codes" field.
func RelyingPartyAuthCodesLTE(v typedef.RelyingPartyID) predicate.AuthCode {
	vc := uint64(v)
	return predicate.AuthCode(sql.FieldLTE(FieldRelyingPartyAuthCodes, vc))
}

// HasRelyingParty applies the HasEdge predicate on the "relying_party" edge.
func HasRelyingParty() predicate.AuthCode {
	return predicate.AuthCode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RelyingPartyTable, RelyingPartyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRelyingPartyWith applies the HasEdge predicate on the "relying_party" edge with a given conditions (other predicates).
func HasRelyingPartyWith(preds ...predicate.RelyingParty) predicate.AuthCode {
	return predicate.AuthCode(func(s *sql.Selector) {
		step := newRelyingPartyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
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
