// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent/admin"
	"github.com/42milez/go-oidc-server/app/idp/ent/schema"
	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	adminFields := schema.Admin{}.Fields()
	_ = adminFields
	// adminDescName is the schema descriptor for name field.
	adminDescName := adminFields[1].Descriptor()
	// admin.NameValidator is a validator for the "name" field. It is called by the builders before save.
	admin.NameValidator = func() func(string) error {
		validators := adminDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
			validators[3].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// adminDescPasswordHash is the schema descriptor for password_hash field.
	adminDescPasswordHash := adminFields[2].Descriptor()
	// admin.PasswordHashValidator is a validator for the "password_hash" field. It is called by the builders before save.
	admin.PasswordHashValidator = func() func(string) error {
		validators := adminDescPasswordHash.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(password_hash string) error {
			for _, fn := range fns {
				if err := fn(password_hash); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// adminDescTotpSecret is the schema descriptor for totp_secret field.
	adminDescTotpSecret := adminFields[3].Descriptor()
	// admin.TotpSecretValidator is a validator for the "totp_secret" field. It is called by the builders before save.
	admin.TotpSecretValidator = func() func(string) error {
		validators := adminDescTotpSecret.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(totp_secret string) error {
			for _, fn := range fns {
				if err := fn(totp_secret); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// adminDescCreatedAt is the schema descriptor for created_at field.
	adminDescCreatedAt := adminFields[4].Descriptor()
	// admin.DefaultCreatedAt holds the default value on creation for the created_at field.
	admin.DefaultCreatedAt = adminDescCreatedAt.Default.(func() time.Time)
	// adminDescModifiedAt is the schema descriptor for modified_at field.
	adminDescModifiedAt := adminFields[5].Descriptor()
	// admin.DefaultModifiedAt holds the default value on creation for the modified_at field.
	admin.DefaultModifiedAt = adminDescModifiedAt.Default.(func() time.Time)
	// admin.UpdateDefaultModifiedAt holds the default value on update for the modified_at field.
	admin.UpdateDefaultModifiedAt = adminDescModifiedAt.UpdateDefault.(func() time.Time)
	// adminDescID is the schema descriptor for id field.
	adminDescID := adminFields[0].Descriptor()
	// admin.DefaultID holds the default value on creation for the id field.
	admin.DefaultID = adminDescID.Default.(func() typedef.AdminID)
}
