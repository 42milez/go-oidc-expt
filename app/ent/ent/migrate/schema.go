// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AuthCodesColumns holds the columns for the "auth_codes" table.
	AuthCodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "code", Type: field.TypeString, SchemaType: map[string]string{"mysql": "CHAR(10)"}},
		{Name: "expire_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUint64},
	}
	// AuthCodesTable holds the schema information for the "auth_codes" table.
	AuthCodesTable = &schema.Table{
		Name:       "auth_codes",
		Columns:    AuthCodesColumns,
		PrimaryKey: []*schema.Column{AuthCodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "auth_codes_users_auth_codes",
				Columns:    []*schema.Column{AuthCodesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "authcode_code",
				Unique:  true,
				Columns: []*schema.Column{AuthCodesColumns[1]},
			},
		},
	}
	// RedirectUrisColumns holds the columns for the "redirect_uris" table.
	RedirectUrisColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uri", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "modified_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUint64},
	}
	// RedirectUrisTable holds the schema information for the "redirect_uris" table.
	RedirectUrisTable = &schema.Table{
		Name:       "redirect_uris",
		Columns:    RedirectUrisColumns,
		PrimaryKey: []*schema.Column{RedirectUrisColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "redirect_uris_users_redirect_uris",
				Columns:    []*schema.Column{RedirectUrisColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "redirecturi_user_id_uri",
				Unique:  true,
				Columns: []*schema.Column{RedirectUrisColumns[4], RedirectUrisColumns[1]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(284)"}},
		{Name: "totp_secret", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "CHAR(160)"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "modified_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AuthCodesTable,
		RedirectUrisTable,
		UsersTable,
	}
)

func init() {
	AuthCodesTable.ForeignKeys[0].RefTable = UsersTable
	RedirectUrisTable.ForeignKeys[0].RefTable = UsersTable
	RedirectUrisTable.Annotation = &entsql.Annotation{
		Table: "redirect_uris",
	}
}
