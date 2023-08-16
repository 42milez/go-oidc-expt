package validation

import (
	"strings"

	"github.com/42milez/go-oidc-server/pkg/xerr"
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slices"
)

func NewAuthorizeValidator() (*validator.Validate, error) {
	ret := validator.New()

	if err := ret.RegisterValidation("scope-validator", validateScope); err != nil {
		return nil, xerr.FailedToInitialize
	}

	if err := ret.RegisterValidation("response-type-validator", validateResponseType); err != nil {
		return nil, xerr.FailedToInitialize
	}

	return ret, nil
}

var validScopes = []string{
	"openid",
	"profile",
	"email",
}

func validateScope(fl validator.FieldLevel) bool {
	scopes := strings.Split(fl.Field().String(), " ")

	for _, v := range scopes {
		if !slices.Contains(validScopes, v) {
			return false
		}
	}

	return true
}

type responseTypeNumber uint64

const (
	rtCode responseTypeNumber = 1 << iota
	rtIDToken
	rtToken
)

var validResponseTypeCombinations = []responseTypeNumber{
	// Authorization Code Flow
	rtCode,
	// Implicit Flow
	rtIDToken | rtToken,
	rtIDToken,
	// Hybrid Flow
	rtCode | rtIDToken,
	rtCode | rtToken,
	rtCode | rtIDToken | rtToken,
}

func convertResponseTypeToNumber(respTypes []string) (responseTypeNumber, error) {
	var ret responseTypeNumber
	for _, v := range respTypes {
		switch v {
		case "code":
			ret |= rtCode
		case "id_token":
			ret |= rtIDToken
		case "token":
			ret |= rtToken
		default:
			return 0, xerr.InvalidResponseType
		}
	}
	return ret, nil
}

func validateResponseType(fl validator.FieldLevel) bool {
	respTypes := strings.Split(fl.Field().String(), " ")
	respTypeComb, err := convertResponseTypeToNumber(respTypes)

	if err != nil {
		return false
	}

	if !slices.Contains(validResponseTypeCombinations, respTypeComb) {
		return false
	}

	return true
}
