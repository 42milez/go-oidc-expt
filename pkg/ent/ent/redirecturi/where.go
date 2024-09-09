// Code generated by ent, DO NOT EDIT.

package redirecturi

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/42milez/go-oidc-expt/pkg/ent/ent/predicate"
	"github.com/42milez/go-oidc-expt/pkg/typedef"
)

// ID filters vertices based on their ID field.
func ID(id typedef.RedirectURIID) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id typedef.RedirectURIID) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id typedef.RedirectURIID) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...typedef.RedirectURIID) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...typedef.RedirectURIID) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id typedef.RedirectURIID) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id typedef.RedirectURIID) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id typedef.RedirectURIID) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id typedef.RedirectURIID) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldLTE(FieldID, id))
}

// URI applies equality check predicate on the "uri" field. It's identical to URIEQ.
func URI(v string) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldEQ(FieldURI, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldEQ(FieldCreatedAt, v))
}

// ModifiedAt applies equality check predicate on the "modified_at" field. It's identical to ModifiedAtEQ.
func ModifiedAt(v time.Time) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldEQ(FieldModifiedAt, v))
}

// RelyingPartyID applies equality check predicate on the "relying_party_id" field. It's identical to RelyingPartyIDEQ.
func RelyingPartyID(v typedef.RelyingPartyID) predicate.RedirectURI {
	vc := uint64(v)
	return predicate.RedirectURI(sql.FieldEQ(FieldRelyingPartyID, vc))
}

// URIEQ applies the EQ predicate on the "uri" field.
func URIEQ(v string) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldEQ(FieldURI, v))
}

// URINEQ applies the NEQ predicate on the "uri" field.
func URINEQ(v string) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldNEQ(FieldURI, v))
}

// URIIn applies the In predicate on the "uri" field.
func URIIn(vs ...string) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldIn(FieldURI, vs...))
}

// URINotIn applies the NotIn predicate on the "uri" field.
func URINotIn(vs ...string) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldNotIn(FieldURI, vs...))
}

// URIGT applies the GT predicate on the "uri" field.
func URIGT(v string) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldGT(FieldURI, v))
}

// URIGTE applies the GTE predicate on the "uri" field.
func URIGTE(v string) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldGTE(FieldURI, v))
}

// URILT applies the LT predicate on the "uri" field.
func URILT(v string) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldLT(FieldURI, v))
}

// URILTE applies the LTE predicate on the "uri" field.
func URILTE(v string) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldLTE(FieldURI, v))
}

// URIContains applies the Contains predicate on the "uri" field.
func URIContains(v string) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldContains(FieldURI, v))
}

// URIHasPrefix applies the HasPrefix predicate on the "uri" field.
func URIHasPrefix(v string) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldHasPrefix(FieldURI, v))
}

// URIHasSuffix applies the HasSuffix predicate on the "uri" field.
func URIHasSuffix(v string) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldHasSuffix(FieldURI, v))
}

// URIEqualFold applies the EqualFold predicate on the "uri" field.
func URIEqualFold(v string) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldEqualFold(FieldURI, v))
}

// URIContainsFold applies the ContainsFold predicate on the "uri" field.
func URIContainsFold(v string) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldContainsFold(FieldURI, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldLTE(FieldCreatedAt, v))
}

// ModifiedAtEQ applies the EQ predicate on the "modified_at" field.
func ModifiedAtEQ(v time.Time) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldEQ(FieldModifiedAt, v))
}

// ModifiedAtNEQ applies the NEQ predicate on the "modified_at" field.
func ModifiedAtNEQ(v time.Time) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldNEQ(FieldModifiedAt, v))
}

// ModifiedAtIn applies the In predicate on the "modified_at" field.
func ModifiedAtIn(vs ...time.Time) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldIn(FieldModifiedAt, vs...))
}

// ModifiedAtNotIn applies the NotIn predicate on the "modified_at" field.
func ModifiedAtNotIn(vs ...time.Time) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldNotIn(FieldModifiedAt, vs...))
}

// ModifiedAtGT applies the GT predicate on the "modified_at" field.
func ModifiedAtGT(v time.Time) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldGT(FieldModifiedAt, v))
}

// ModifiedAtGTE applies the GTE predicate on the "modified_at" field.
func ModifiedAtGTE(v time.Time) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldGTE(FieldModifiedAt, v))
}

// ModifiedAtLT applies the LT predicate on the "modified_at" field.
func ModifiedAtLT(v time.Time) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldLT(FieldModifiedAt, v))
}

// ModifiedAtLTE applies the LTE predicate on the "modified_at" field.
func ModifiedAtLTE(v time.Time) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldLTE(FieldModifiedAt, v))
}

// RelyingPartyIDEQ applies the EQ predicate on the "relying_party_id" field.
func RelyingPartyIDEQ(v typedef.RelyingPartyID) predicate.RedirectURI {
	vc := uint64(v)
	return predicate.RedirectURI(sql.FieldEQ(FieldRelyingPartyID, vc))
}

// RelyingPartyIDNEQ applies the NEQ predicate on the "relying_party_id" field.
func RelyingPartyIDNEQ(v typedef.RelyingPartyID) predicate.RedirectURI {
	vc := uint64(v)
	return predicate.RedirectURI(sql.FieldNEQ(FieldRelyingPartyID, vc))
}

// RelyingPartyIDIn applies the In predicate on the "relying_party_id" field.
func RelyingPartyIDIn(vs ...typedef.RelyingPartyID) predicate.RedirectURI {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.RedirectURI(sql.FieldIn(FieldRelyingPartyID, v...))
}

// RelyingPartyIDNotIn applies the NotIn predicate on the "relying_party_id" field.
func RelyingPartyIDNotIn(vs ...typedef.RelyingPartyID) predicate.RedirectURI {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.RedirectURI(sql.FieldNotIn(FieldRelyingPartyID, v...))
}

// HasRelyingParty applies the HasEdge predicate on the "relying_party" edge.
func HasRelyingParty() predicate.RedirectURI {
	return predicate.RedirectURI(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RelyingPartyTable, RelyingPartyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRelyingPartyWith applies the HasEdge predicate on the "relying_party" edge with a given conditions (other predicates).
func HasRelyingPartyWith(preds ...predicate.RelyingParty) predicate.RedirectURI {
	return predicate.RedirectURI(func(s *sql.Selector) {
		step := newRelyingPartyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RedirectURI) predicate.RedirectURI {
	return predicate.RedirectURI(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RedirectURI) predicate.RedirectURI {
	return predicate.RedirectURI(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RedirectURI) predicate.RedirectURI {
	return predicate.RedirectURI(sql.NotPredicates(p))
}