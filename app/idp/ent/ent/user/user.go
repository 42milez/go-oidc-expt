// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPasswordHash holds the string denoting the password_hash field in the database.
	FieldPasswordHash = "password_hash"
	// FieldTotpSecret holds the string denoting the totp_secret field in the database.
	FieldTotpSecret = "totp_secret"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldModifiedAt holds the string denoting the modified_at field in the database.
	FieldModifiedAt = "modified_at"
	// EdgeAuthCodes holds the string denoting the auth_codes edge name in mutations.
	EdgeAuthCodes = "auth_codes"
	// EdgeRedirectUris holds the string denoting the redirect_uris edge name in mutations.
	EdgeRedirectUris = "redirect_uris"
	// Table holds the table name of the user in the database.
	Table = "users"
	// AuthCodesTable is the table that holds the auth_codes relation/edge.
	AuthCodesTable = "auth_codes"
	// AuthCodesInverseTable is the table name for the AuthCode entity.
	// It exists in this package in order to avoid circular dependency with the "authcode" package.
	AuthCodesInverseTable = "auth_codes"
	// AuthCodesColumn is the table column denoting the auth_codes relation/edge.
	AuthCodesColumn = "user_id"
	// RedirectUrisTable is the table that holds the redirect_uris relation/edge.
	RedirectUrisTable = "redirect_ur_is"
	// RedirectUrisInverseTable is the table name for the RedirectURI entity.
	// It exists in this package in order to avoid circular dependency with the "redirecturi" package.
	RedirectUrisInverseTable = "redirect_ur_is"
	// RedirectUrisColumn is the table column denoting the redirect_uris relation/edge.
	RedirectUrisColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPasswordHash,
	FieldTotpSecret,
	FieldCreatedAt,
	FieldModifiedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// PasswordHashValidator is a validator for the "password_hash" field. It is called by the builders before save.
	PasswordHashValidator func(string) error
	// TotpSecretValidator is a validator for the "totp_secret" field. It is called by the builders before save.
	TotpSecretValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultModifiedAt holds the default value on creation for the "modified_at" field.
	DefaultModifiedAt func() time.Time
	// UpdateDefaultModifiedAt holds the default value on update for the "modified_at" field.
	UpdateDefaultModifiedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() typedef.UserID
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByPasswordHash orders the results by the password_hash field.
func ByPasswordHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPasswordHash, opts...).ToFunc()
}

// ByTotpSecret orders the results by the totp_secret field.
func ByTotpSecret(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotpSecret, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByModifiedAt orders the results by the modified_at field.
func ByModifiedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModifiedAt, opts...).ToFunc()
}

// ByAuthCodesCount orders the results by auth_codes count.
func ByAuthCodesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAuthCodesStep(), opts...)
	}
}

// ByAuthCodes orders the results by auth_codes terms.
func ByAuthCodes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthCodesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRedirectUrisCount orders the results by redirect_uris count.
func ByRedirectUrisCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRedirectUrisStep(), opts...)
	}
}

// ByRedirectUris orders the results by redirect_uris terms.
func ByRedirectUris(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRedirectUrisStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAuthCodesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthCodesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AuthCodesTable, AuthCodesColumn),
	)
}
func newRedirectUrisStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RedirectUrisInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RedirectUrisTable, RedirectUrisColumn),
	)
}
