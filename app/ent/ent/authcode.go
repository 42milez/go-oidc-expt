// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/42milez/go-oidc-server/app/ent/ent/authcode"
)

// AuthCode is the model entity for the AuthCode schema.
type AuthCode struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// ExpireAt holds the value of the "expire_at" field.
	ExpireAt time.Time `json:"expire_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UsedAt holds the value of the "used_at" field.
	UsedAt       time.Time `json:"used_at,omitempty"`
	client_id    *int
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AuthCode) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case authcode.FieldID:
			values[i] = new(sql.NullInt64)
		case authcode.FieldCode:
			values[i] = new(sql.NullString)
		case authcode.FieldExpireAt, authcode.FieldCreatedAt, authcode.FieldUsedAt:
			values[i] = new(sql.NullTime)
		case authcode.ForeignKeys[0]: // client_id
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AuthCode fields.
func (ac *AuthCode) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case authcode.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ac.ID = int(value.Int64)
		case authcode.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				ac.Code = value.String
			}
		case authcode.FieldExpireAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expire_at", values[i])
			} else if value.Valid {
				ac.ExpireAt = value.Time
			}
		case authcode.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ac.CreatedAt = value.Time
			}
		case authcode.FieldUsedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field used_at", values[i])
			} else if value.Valid {
				ac.UsedAt = value.Time
			}
		case authcode.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field client_id", value)
			} else if value.Valid {
				ac.client_id = new(int)
				*ac.client_id = int(value.Int64)
			}
		default:
			ac.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AuthCode.
// This includes values selected through modifiers, order, etc.
func (ac *AuthCode) Value(name string) (ent.Value, error) {
	return ac.selectValues.Get(name)
}

// Update returns a builder for updating this AuthCode.
// Note that you need to call AuthCode.Unwrap() before calling this method if this AuthCode
// was returned from a transaction, and the transaction was committed or rolled back.
func (ac *AuthCode) Update() *AuthCodeUpdateOne {
	return NewAuthCodeClient(ac.config).UpdateOne(ac)
}

// Unwrap unwraps the AuthCode entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ac *AuthCode) Unwrap() *AuthCode {
	_tx, ok := ac.config.driver.(*txDriver)
	if !ok {
		panic("ent: AuthCode is not a transactional entity")
	}
	ac.config.driver = _tx.drv
	return ac
}

// String implements the fmt.Stringer.
func (ac *AuthCode) String() string {
	var builder strings.Builder
	builder.WriteString("AuthCode(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ac.ID))
	builder.WriteString("code=")
	builder.WriteString(ac.Code)
	builder.WriteString(", ")
	builder.WriteString("expire_at=")
	builder.WriteString(ac.ExpireAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ac.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("used_at=")
	builder.WriteString(ac.UsedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// AuthCodes is a parsable slice of AuthCode.
type AuthCodes []*AuthCode
