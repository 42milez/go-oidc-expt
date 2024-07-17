// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/42milez/go-oidc-expt/pkg/ent/ent/consent"
	"github.com/42milez/go-oidc-expt/pkg/ent/ent/user"
	"github.com/42milez/go-oidc-expt/pkg/typedef"
)

// Consent is the model entity for the Consent schema.
type Consent struct {
	config `json:"-"`
	// ID of the ent.
	ID typedef.ConsentID `json:"id,omitempty"`
	// ClientID holds the value of the "client_id" field.
	ClientID typedef.ClientID `json:"client_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID typedef.UserID `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ConsentQuery when eager-loading is set.
	Edges        ConsentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ConsentEdges holds the relations/edges for other nodes in the graph.
type ConsentEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ConsentEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Consent) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case consent.FieldID, consent.FieldUserID:
			values[i] = new(sql.NullInt64)
		case consent.FieldClientID:
			values[i] = new(sql.NullString)
		case consent.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Consent fields.
func (c *Consent) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case consent.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = typedef.ConsentID(value.Int64)
			}
		case consent.FieldClientID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_id", values[i])
			} else if value.Valid {
				c.ClientID = typedef.ClientID(value.String)
			}
		case consent.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case consent.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				c.UserID = typedef.UserID(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Consent.
// This includes values selected through modifiers, order, etc.
func (c *Consent) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Consent entity.
func (c *Consent) QueryUser() *UserQuery {
	return NewConsentClient(c.config).QueryUser(c)
}

// Update returns a builder for updating this Consent.
// Note that you need to call Consent.Unwrap() before calling this method if this Consent
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Consent) Update() *ConsentUpdateOne {
	return NewConsentClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Consent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Consent) Unwrap() *Consent {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Consent is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Consent) String() string {
	var builder strings.Builder
	builder.WriteString("Consent(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("client_id=")
	builder.WriteString(fmt.Sprintf("%v", c.ClientID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", c.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// Consents is a parsable slice of Consent.
type Consents []*Consent
