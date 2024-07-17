package api

import (
	"strings"

	"github.com/42milez/go-oidc-expt/pkg/xerr"

	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slices"
)

func NewOIDCRequestParamValidator() (*validator.Validate, error) {
	ret := validator.New()

	if err := ret.RegisterValidation("display-validator", validateDisplay); err != nil {
		return nil, err
	}

	if err := ret.RegisterValidation("grant-type-validator", validateGrantType); err != nil {
		return nil, err
	}

	if err := ret.RegisterValidation("prompt-validator", validatePrompt); err != nil {
		return nil, err
	}

	if err := ret.RegisterValidation("response-type-validator", validateResponseType); err != nil {
		return nil, err
	}

	if err := ret.RegisterValidation("scope-validator", validateScope); err != nil {
		return nil, err
	}

	return ret, nil
}

func validateDisplay(fl validator.FieldLevel) bool {
	switch fl.Field().String() {
	case "page":
		return true
	case "popup":
		return true
	case "touch":
		return true
	case "wap":
		return true
	default:
		return false
	}
}

func validateGrantType(fl validator.FieldLevel) bool {
	switch fl.Field().String() {
	case "authorization_code":
		return true
	case "refresh_token":
		return true
	default:
		return false
	}
}

func validatePrompt(fl validator.FieldLevel) bool {
	switch fl.Field().String() {
	case "none":
		return true
	case "login":
		return true
	case "consent":
		return true
	case "select_account":
		return true
	default:
		return false
	}
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

const openIDScope = "openid"
const profileScope = "profile"
const emailScope = "email"

var validScopes = []string{
	openIDScope,
	profileScope,
	emailScope,
}

func validateScope(fl validator.FieldLevel) bool {
	scopes := strings.Split(fl.Field().String(), " ")
	validOpenIDScope := false

	for _, v := range scopes {
		if !slices.Contains(validScopes, v) {
			return false
		}
		if v == openIDScope {
			validOpenIDScope = true
		}
	}

	if !validOpenIDScope {
		return false
	}

	return true
}
