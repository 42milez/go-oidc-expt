// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/42milez/go-oidc-server/app/ent/ent/relyingparty"
	"github.com/42milez/go-oidc-server/app/typedef"
)

// RelyingParty is the model entity for the RelyingParty schema.
type RelyingParty struct {
	config `json:"-"`
	// ID of the ent.
	ID typedef.RelyingPartyID `json:"id,omitempty"`
	// ClientID holds the value of the "client_id" field.
	ClientID typedef.ClientID `json:"client_id,omitempty"`
	// ClientSecret holds the value of the "client_secret" field.
	ClientSecret typedef.ClientSecret `json:"client_secret,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// ModifiedAt holds the value of the "modified_at" field.
	ModifiedAt time.Time `json:"modified_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RelyingPartyQuery when eager-loading is set.
	Edges        RelyingPartyEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RelyingPartyEdges holds the relations/edges for other nodes in the graph.
type RelyingPartyEdges struct {
	// AuthCodes holds the value of the auth_codes edge.
	AuthCodes []*AuthCode `json:"auth_codes,omitempty"`
	// RedirectUris holds the value of the redirect_uris edge.
	RedirectUris []*RedirectURI `json:"redirect_uris,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AuthCodesOrErr returns the AuthCodes value or an error if the edge
// was not loaded in eager-loading.
func (e RelyingPartyEdges) AuthCodesOrErr() ([]*AuthCode, error) {
	if e.loadedTypes[0] {
		return e.AuthCodes, nil
	}
	return nil, &NotLoadedError{edge: "auth_codes"}
}

// RedirectUrisOrErr returns the RedirectUris value or an error if the edge
// was not loaded in eager-loading.
func (e RelyingPartyEdges) RedirectUrisOrErr() ([]*RedirectURI, error) {
	if e.loadedTypes[1] {
		return e.RedirectUris, nil
	}
	return nil, &NotLoadedError{edge: "redirect_uris"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RelyingParty) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case relyingparty.FieldID:
			values[i] = new(sql.NullInt64)
		case relyingparty.FieldClientID, relyingparty.FieldClientSecret:
			values[i] = new(sql.NullString)
		case relyingparty.FieldCreatedAt, relyingparty.FieldModifiedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RelyingParty fields.
func (rp *RelyingParty) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case relyingparty.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				rp.ID = typedef.RelyingPartyID(value.Int64)
			}
		case relyingparty.FieldClientID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_id", values[i])
			} else if value.Valid {
				rp.ClientID = typedef.ClientID(value.String)
			}
		case relyingparty.FieldClientSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_secret", values[i])
			} else if value.Valid {
				rp.ClientSecret = typedef.ClientSecret(value.String)
			}
		case relyingparty.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				rp.CreatedAt = value.Time
			}
		case relyingparty.FieldModifiedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified_at", values[i])
			} else if value.Valid {
				rp.ModifiedAt = value.Time
			}
		default:
			rp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RelyingParty.
// This includes values selected through modifiers, order, etc.
func (rp *RelyingParty) Value(name string) (ent.Value, error) {
	return rp.selectValues.Get(name)
}

// QueryAuthCodes queries the "auth_codes" edge of the RelyingParty entity.
func (rp *RelyingParty) QueryAuthCodes() *AuthCodeQuery {
	return NewRelyingPartyClient(rp.config).QueryAuthCodes(rp)
}

// QueryRedirectUris queries the "redirect_uris" edge of the RelyingParty entity.
func (rp *RelyingParty) QueryRedirectUris() *RedirectURIQuery {
	return NewRelyingPartyClient(rp.config).QueryRedirectUris(rp)
}

// Update returns a builder for updating this RelyingParty.
// Note that you need to call RelyingParty.Unwrap() before calling this method if this RelyingParty
// was returned from a transaction, and the transaction was committed or rolled back.
func (rp *RelyingParty) Update() *RelyingPartyUpdateOne {
	return NewRelyingPartyClient(rp.config).UpdateOne(rp)
}

// Unwrap unwraps the RelyingParty entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rp *RelyingParty) Unwrap() *RelyingParty {
	_tx, ok := rp.config.driver.(*txDriver)
	if !ok {
		panic("ent: RelyingParty is not a transactional entity")
	}
	rp.config.driver = _tx.drv
	return rp
}

// String implements the fmt.Stringer.
func (rp *RelyingParty) String() string {
	var builder strings.Builder
	builder.WriteString("RelyingParty(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rp.ID))
	builder.WriteString("client_id=")
	builder.WriteString(fmt.Sprintf("%v", rp.ClientID))
	builder.WriteString(", ")
	builder.WriteString("client_secret=")
	builder.WriteString(fmt.Sprintf("%v", rp.ClientSecret))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(rp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("modified_at=")
	builder.WriteString(rp.ModifiedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// RelyingParties is a parsable slice of RelyingParty.
type RelyingParties []*RelyingParty
