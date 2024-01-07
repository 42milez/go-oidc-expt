// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/42milez/go-oidc-server/app/pkg/ent/ent/redirecturi"
	"github.com/42milez/go-oidc-server/app/pkg/ent/ent/relyingparty"
	"github.com/42milez/go-oidc-server/app/pkg/typedef"
)

// RedirectURI is the model entity for the RedirectURI schema.
type RedirectURI struct {
	config `json:"-"`
	// ID of the ent.
	ID typedef.RedirectURIID `json:"id,omitempty"`
	// URI holds the value of the "uri" field.
	URI string `json:"uri,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// ModifiedAt holds the value of the "modified_at" field.
	ModifiedAt time.Time `json:"modified_at,omitempty"`
	// RelyingPartyID holds the value of the "relying_party_id" field.
	RelyingPartyID typedef.RelyingPartyID `json:"relying_party_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RedirectURIQuery when eager-loading is set.
	Edges        RedirectURIEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RedirectURIEdges holds the relations/edges for other nodes in the graph.
type RedirectURIEdges struct {
	// RelyingParty holds the value of the relying_party edge.
	RelyingParty *RelyingParty `json:"relying_party,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RelyingPartyOrErr returns the RelyingParty value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RedirectURIEdges) RelyingPartyOrErr() (*RelyingParty, error) {
	if e.loadedTypes[0] {
		if e.RelyingParty == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: relyingparty.Label}
		}
		return e.RelyingParty, nil
	}
	return nil, &NotLoadedError{edge: "relying_party"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RedirectURI) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case redirecturi.FieldID, redirecturi.FieldRelyingPartyID:
			values[i] = new(sql.NullInt64)
		case redirecturi.FieldURI:
			values[i] = new(sql.NullString)
		case redirecturi.FieldCreatedAt, redirecturi.FieldModifiedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RedirectURI fields.
func (ru *RedirectURI) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case redirecturi.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ru.ID = typedef.RedirectURIID(value.Int64)
			}
		case redirecturi.FieldURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uri", values[i])
			} else if value.Valid {
				ru.URI = value.String
			}
		case redirecturi.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ru.CreatedAt = value.Time
			}
		case redirecturi.FieldModifiedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified_at", values[i])
			} else if value.Valid {
				ru.ModifiedAt = value.Time
			}
		case redirecturi.FieldRelyingPartyID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field relying_party_id", values[i])
			} else if value.Valid {
				ru.RelyingPartyID = typedef.RelyingPartyID(value.Int64)
			}
		default:
			ru.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RedirectURI.
// This includes values selected through modifiers, order, etc.
func (ru *RedirectURI) Value(name string) (ent.Value, error) {
	return ru.selectValues.Get(name)
}

// QueryRelyingParty queries the "relying_party" edge of the RedirectURI entity.
func (ru *RedirectURI) QueryRelyingParty() *RelyingPartyQuery {
	return NewRedirectURIClient(ru.config).QueryRelyingParty(ru)
}

// Update returns a builder for updating this RedirectURI.
// Note that you need to call RedirectURI.Unwrap() before calling this method if this RedirectURI
// was returned from a transaction, and the transaction was committed or rolled back.
func (ru *RedirectURI) Update() *RedirectURIUpdateOne {
	return NewRedirectURIClient(ru.config).UpdateOne(ru)
}

// Unwrap unwraps the RedirectURI entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ru *RedirectURI) Unwrap() *RedirectURI {
	_tx, ok := ru.config.driver.(*txDriver)
	if !ok {
		panic("ent: RedirectURI is not a transactional entity")
	}
	ru.config.driver = _tx.drv
	return ru
}

// String implements the fmt.Stringer.
func (ru *RedirectURI) String() string {
	var builder strings.Builder
	builder.WriteString("RedirectURI(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ru.ID))
	builder.WriteString("uri=")
	builder.WriteString(ru.URI)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ru.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("modified_at=")
	builder.WriteString(ru.ModifiedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("relying_party_id=")
	builder.WriteString(fmt.Sprintf("%v", ru.RelyingPartyID))
	builder.WriteByte(')')
	return builder.String()
}

// RedirectURIs is a parsable slice of RedirectURI.
type RedirectURIs []*RedirectURI
