// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/42milez/go-oidc-server/app/idp/ent/ent/user"
	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID typedef.UserID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// PasswordHash holds the value of the "password_hash" field.
	PasswordHash typedef.PasswordHash `json:"password_hash,omitempty"`
	// TotpSecret holds the value of the "totp_secret" field.
	TotpSecret string `json:"totp_secret,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// ModifiedAt holds the value of the "modified_at" field.
	ModifiedAt time.Time `json:"modified_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
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
func (e UserEdges) AuthCodesOrErr() ([]*AuthCode, error) {
	if e.loadedTypes[0] {
		return e.AuthCodes, nil
	}
	return nil, &NotLoadedError{edge: "auth_codes"}
}

// RedirectUrisOrErr returns the RedirectUris value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RedirectUrisOrErr() ([]*RedirectURI, error) {
	if e.loadedTypes[1] {
		return e.RedirectUris, nil
	}
	return nil, &NotLoadedError{edge: "redirect_uris"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldName, user.FieldPasswordHash, user.FieldTotpSecret:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldModifiedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				u.ID = typedef.UserID(value.String)
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldPasswordHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password_hash", values[i])
			} else if value.Valid {
				u.PasswordHash = typedef.PasswordHash(value.String)
			}
		case user.FieldTotpSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field totp_secret", values[i])
			} else if value.Valid {
				u.TotpSecret = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldModifiedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified_at", values[i])
			} else if value.Valid {
				u.ModifiedAt = value.Time
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryAuthCodes queries the "auth_codes" edge of the User entity.
func (u *User) QueryAuthCodes() *AuthCodeQuery {
	return NewUserClient(u.config).QueryAuthCodes(u)
}

// QueryRedirectUris queries the "redirect_uris" edge of the User entity.
func (u *User) QueryRedirectUris() *RedirectURIQuery {
	return NewUserClient(u.config).QueryRedirectUris(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("password_hash=")
	builder.WriteString(fmt.Sprintf("%v", u.PasswordHash))
	builder.WriteString(", ")
	builder.WriteString("totp_secret=")
	builder.WriteString(u.TotpSecret)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("modified_at=")
	builder.WriteString(u.ModifiedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
