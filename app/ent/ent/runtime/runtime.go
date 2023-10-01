// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/42milez/go-oidc-server/app/ent/ent/authcode"
	"github.com/42milez/go-oidc-server/app/ent/ent/consent"
	"github.com/42milez/go-oidc-server/app/ent/ent/redirecturi"
	"github.com/42milez/go-oidc-server/app/ent/ent/relyingparty"
	"github.com/42milez/go-oidc-server/app/ent/ent/user"
	"github.com/42milez/go-oidc-server/app/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	authcodeFields := schema.AuthCode{}.Fields()
	_ = authcodeFields
	// authcodeDescCode is the schema descriptor for code field.
	authcodeDescCode := authcodeFields[2].Descriptor()
	// authcode.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	authcode.CodeValidator = func() func(string) error {
		validators := authcodeDescCode.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(code string) error {
			for _, fn := range fns {
				if err := fn(code); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// authcodeDescExpireAt is the schema descriptor for expire_at field.
	authcodeDescExpireAt := authcodeFields[3].Descriptor()
	// authcode.DefaultExpireAt holds the default value on creation for the expire_at field.
	authcode.DefaultExpireAt = authcodeDescExpireAt.Default.(func() time.Time)
	// authcodeDescCreatedAt is the schema descriptor for created_at field.
	authcodeDescCreatedAt := authcodeFields[5].Descriptor()
	// authcode.DefaultCreatedAt holds the default value on creation for the created_at field.
	authcode.DefaultCreatedAt = authcodeDescCreatedAt.Default.(func() time.Time)
	// authcodeDescModifiedAt is the schema descriptor for modified_at field.
	authcodeDescModifiedAt := authcodeFields[6].Descriptor()
	// authcode.DefaultModifiedAt holds the default value on creation for the modified_at field.
	authcode.DefaultModifiedAt = authcodeDescModifiedAt.Default.(func() time.Time)
	// authcode.UpdateDefaultModifiedAt holds the default value on update for the modified_at field.
	authcode.UpdateDefaultModifiedAt = authcodeDescModifiedAt.UpdateDefault.(func() time.Time)
	consentFields := schema.Consent{}.Fields()
	_ = consentFields
	// consentDescCreatedAt is the schema descriptor for created_at field.
	consentDescCreatedAt := consentFields[2].Descriptor()
	// consent.DefaultCreatedAt holds the default value on creation for the created_at field.
	consent.DefaultCreatedAt = consentDescCreatedAt.Default.(func() time.Time)
	redirecturiFields := schema.RedirectURI{}.Fields()
	_ = redirecturiFields
	// redirecturiDescURI is the schema descriptor for uri field.
	redirecturiDescURI := redirecturiFields[1].Descriptor()
	// redirecturi.URIValidator is a validator for the "uri" field. It is called by the builders before save.
	redirecturi.URIValidator = redirecturiDescURI.Validators[0].(func(string) error)
	// redirecturiDescCreatedAt is the schema descriptor for created_at field.
	redirecturiDescCreatedAt := redirecturiFields[2].Descriptor()
	// redirecturi.DefaultCreatedAt holds the default value on creation for the created_at field.
	redirecturi.DefaultCreatedAt = redirecturiDescCreatedAt.Default.(func() time.Time)
	// redirecturiDescModifiedAt is the schema descriptor for modified_at field.
	redirecturiDescModifiedAt := redirecturiFields[3].Descriptor()
	// redirecturi.DefaultModifiedAt holds the default value on creation for the modified_at field.
	redirecturi.DefaultModifiedAt = redirecturiDescModifiedAt.Default.(func() time.Time)
	// redirecturi.UpdateDefaultModifiedAt holds the default value on update for the modified_at field.
	redirecturi.UpdateDefaultModifiedAt = redirecturiDescModifiedAt.UpdateDefault.(func() time.Time)
	relyingpartyFields := schema.RelyingParty{}.Fields()
	_ = relyingpartyFields
	// relyingpartyDescClientSecret is the schema descriptor for client_secret field.
	relyingpartyDescClientSecret := relyingpartyFields[2].Descriptor()
	// relyingparty.ClientSecretValidator is a validator for the "client_secret" field. It is called by the builders before save.
	relyingparty.ClientSecretValidator = relyingpartyDescClientSecret.Validators[0].(func(string) error)
	// relyingpartyDescCreatedAt is the schema descriptor for created_at field.
	relyingpartyDescCreatedAt := relyingpartyFields[3].Descriptor()
	// relyingparty.DefaultCreatedAt holds the default value on creation for the created_at field.
	relyingparty.DefaultCreatedAt = relyingpartyDescCreatedAt.Default.(func() time.Time)
	// relyingpartyDescModifiedAt is the schema descriptor for modified_at field.
	relyingpartyDescModifiedAt := relyingpartyFields[4].Descriptor()
	// relyingparty.DefaultModifiedAt holds the default value on creation for the modified_at field.
	relyingparty.DefaultModifiedAt = relyingpartyDescModifiedAt.Default.(func() time.Time)
	// relyingparty.UpdateDefaultModifiedAt holds the default value on update for the modified_at field.
	relyingparty.UpdateDefaultModifiedAt = relyingpartyDescModifiedAt.UpdateDefault.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinHooks0 := userMixin[0].Hooks()
	user.Hooks[0] = userMixinHooks0[0]
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = func() func(string) error {
		validators := userDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
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
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescTotpSecret is the schema descriptor for totp_secret field.
	userDescTotpSecret := userFields[2].Descriptor()
	// user.TotpSecretValidator is a validator for the "totp_secret" field. It is called by the builders before save.
	user.TotpSecretValidator = userDescTotpSecret.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescModifiedAt is the schema descriptor for modified_at field.
	userDescModifiedAt := userFields[4].Descriptor()
	// user.DefaultModifiedAt holds the default value on creation for the modified_at field.
	user.DefaultModifiedAt = userDescModifiedAt.Default.(func() time.Time)
	// user.UpdateDefaultModifiedAt holds the default value on update for the modified_at field.
	user.UpdateDefaultModifiedAt = userDescModifiedAt.UpdateDefault.(func() time.Time)
}

const (
	Version = "v0.12.4"                                         // Version of ent codegen.
	Sum     = "h1:LddPnAyxls/O7DTXZvUGDj0NZIdGSu317+aoNLJWbD8=" // Sum of ent codegen.
)
