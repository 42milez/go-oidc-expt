// Code generated by ent, DO NOT EDIT.

package authcode

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the authcode type in the database.
	Label = "auth_code"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldExpireAt holds the string denoting the expire_at field in the database.
	FieldExpireAt = "expire_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUsedAt holds the string denoting the used_at field in the database.
	FieldUsedAt = "used_at"
	// FieldRelyingPartyAuthCodes holds the string denoting the relying_party_auth_codes field in the database.
	FieldRelyingPartyAuthCodes = "relying_party_auth_codes"
	// EdgeRelyingParty holds the string denoting the relying_party edge name in mutations.
	EdgeRelyingParty = "relying_party"
	// Table holds the table name of the authcode in the database.
	Table = "auth_codes"
	// RelyingPartyTable is the table that holds the relying_party relation/edge.
	RelyingPartyTable = "auth_codes"
	// RelyingPartyInverseTable is the table name for the RelyingParty entity.
	// It exists in this package in order to avoid circular dependency with the "relyingparty" package.
	RelyingPartyInverseTable = "relying_parties"
	// RelyingPartyColumn is the table column denoting the relying_party relation/edge.
	RelyingPartyColumn = "relying_party_auth_codes"
)

// Columns holds all SQL columns for authcode fields.
var Columns = []string{
	FieldID,
	FieldCode,
	FieldExpireAt,
	FieldCreatedAt,
	FieldUsedAt,
	FieldRelyingPartyAuthCodes,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "auth_codes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"relying_party_auth_codes",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// DefaultExpireAt holds the default value on creation for the "expire_at" field.
	DefaultExpireAt func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the AuthCode queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByExpireAt orders the results by the expire_at field.
func ByExpireAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpireAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUsedAt orders the results by the used_at field.
func ByUsedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsedAt, opts...).ToFunc()
}

// ByRelyingPartyAuthCodes orders the results by the relying_party_auth_codes field.
func ByRelyingPartyAuthCodes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRelyingPartyAuthCodes, opts...).ToFunc()
}

// ByRelyingPartyField orders the results by relying_party field.
func ByRelyingPartyField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRelyingPartyStep(), sql.OrderByField(field, opts...))
	}
}
func newRelyingPartyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RelyingPartyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, RelyingPartyTable, RelyingPartyColumn),
	)
}
