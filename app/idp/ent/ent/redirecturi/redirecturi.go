// Code generated by ent, DO NOT EDIT.

package redirecturi

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the redirecturi type in the database.
	Label = "redirect_uri"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldURI holds the string denoting the uri field in the database.
	FieldURI = "uri"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldModifiedAt holds the string denoting the modified_at field in the database.
	FieldModifiedAt = "modified_at"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// Table holds the table name of the redirecturi in the database.
	Table = "redirect_uris"
)

// Columns holds all SQL columns for redirecturi fields.
var Columns = []string{
	FieldID,
	FieldURI,
	FieldCreatedAt,
	FieldModifiedAt,
	FieldUserID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "redirect_uris"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_id",
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
	// URIValidator is a validator for the "uri" field. It is called by the builders before save.
	URIValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultModifiedAt holds the default value on creation for the "modified_at" field.
	DefaultModifiedAt func() time.Time
	// UpdateDefaultModifiedAt holds the default value on update for the "modified_at" field.
	UpdateDefaultModifiedAt func() time.Time
	// UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	UserIDValidator func(string) error
)

// OrderOption defines the ordering options for the RedirectURI queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByURI orders the results by the uri field.
func ByURI(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURI, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByModifiedAt orders the results by the modified_at field.
func ByModifiedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModifiedAt, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}
