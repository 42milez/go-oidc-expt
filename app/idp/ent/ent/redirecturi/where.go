// Code generated by ent, DO NOT EDIT.

package redirecturi

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/42milez/go-oidc-server/app/idp/ent/ent/predicate"
	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.RedirectURI {
	return predicate.RedirectURI(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.RedirectURI {
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

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v typedef.UserID) predicate.RedirectURI {
	vc := string(v)
	return predicate.RedirectURI(sql.FieldEQ(FieldUserID, vc))
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

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v typedef.UserID) predicate.RedirectURI {
	vc := string(v)
	return predicate.RedirectURI(sql.FieldEQ(FieldUserID, vc))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v typedef.UserID) predicate.RedirectURI {
	vc := string(v)
	return predicate.RedirectURI(sql.FieldNEQ(FieldUserID, vc))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...typedef.UserID) predicate.RedirectURI {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.RedirectURI(sql.FieldIn(FieldUserID, v...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...typedef.UserID) predicate.RedirectURI {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.RedirectURI(sql.FieldNotIn(FieldUserID, v...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v typedef.UserID) predicate.RedirectURI {
	vc := string(v)
	return predicate.RedirectURI(sql.FieldGT(FieldUserID, vc))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v typedef.UserID) predicate.RedirectURI {
	vc := string(v)
	return predicate.RedirectURI(sql.FieldGTE(FieldUserID, vc))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v typedef.UserID) predicate.RedirectURI {
	vc := string(v)
	return predicate.RedirectURI(sql.FieldLT(FieldUserID, vc))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v typedef.UserID) predicate.RedirectURI {
	vc := string(v)
	return predicate.RedirectURI(sql.FieldLTE(FieldUserID, vc))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v typedef.UserID) predicate.RedirectURI {
	vc := string(v)
	return predicate.RedirectURI(sql.FieldContains(FieldUserID, vc))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v typedef.UserID) predicate.RedirectURI {
	vc := string(v)
	return predicate.RedirectURI(sql.FieldHasPrefix(FieldUserID, vc))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v typedef.UserID) predicate.RedirectURI {
	vc := string(v)
	return predicate.RedirectURI(sql.FieldHasSuffix(FieldUserID, vc))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v typedef.UserID) predicate.RedirectURI {
	vc := string(v)
	return predicate.RedirectURI(sql.FieldEqualFold(FieldUserID, vc))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v typedef.UserID) predicate.RedirectURI {
	vc := string(v)
	return predicate.RedirectURI(sql.FieldContainsFold(FieldUserID, vc))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RedirectURI) predicate.RedirectURI {
	return predicate.RedirectURI(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RedirectURI) predicate.RedirectURI {
	return predicate.RedirectURI(func(s *sql.Selector) {
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
func Not(p predicate.RedirectURI) predicate.RedirectURI {
	return predicate.RedirectURI(func(s *sql.Selector) {
		p(s.Not())
	})
}
