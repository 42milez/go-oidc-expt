package validation

import (
	"strings"

	"github.com/42milez/go-oidc-server/pkg/xerr"
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slices"
)

func NewAuthorizeValidator() (*validator.Validate, error) {
	ret := validator.New()

	if err := ret.RegisterValidation("ScopeValidator", validScope); err != nil {
		return nil, xerr.FailedToInitialize
	}

	return ret, nil
}

type scopeNumber uint64

const (
	scpOpenID scopeNumber = 1 << iota
	scpIDToken
	scpToken
	scpCode
)

var validScopeCombinations = []scopeNumber{
	// Authorization Code Flow
	scpOpenID,
	// Implicit Flow
	scpIDToken | scpToken,
	scpIDToken,
	// Hybrid Flow
	scpCode | scpIDToken,
	scpCode | scpToken,
	scpCode | scpIDToken | scpToken,
}

func convertScopeToNumber(scopes []string) (scopeNumber, error) {
	var ret scopeNumber
	for _, v := range scopes {
		switch v {
		case "openid":
			ret |= scpOpenID
		case "id_token":
			ret |= scpIDToken
		case "token":
			ret |= scpToken
		case "code":
			ret |= scpCode
		default:
			return 0, xerr.InvalidScope
		}
	}
	return ret, nil
}

func validScope(fl validator.FieldLevel) bool {
	scopes := strings.Split(fl.Field().String(), " ")
	scopeComb, err := convertScopeToNumber(scopes)

	if err != nil {
		return false
	}

	if !slices.Contains(validScopeCombinations, scopeComb) {
		return false
	}

	return true
}
