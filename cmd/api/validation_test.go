package api

import (
	"testing"
)

func TestAuthorizeParamsValidation(t *testing.T) {
	t.Parallel()

	v, err := NewOIDCRequestParamValidator()
	if err != nil {
		t.Fatal(err)
	}

	p := &AuthorizeParams{
		ClientID:     "CDcp9v3Nn4i70FqWig5AuohmorD6MG",
		Nonce:        "jAhe7E9oa0",
		RedirectURI:  "https://example.com/cb",
		ResponseType: "code",
		Scope:        "openid profile email",
		State:        "oKMFV15FcJ",
		Display:      nil,
		MaxAge:       nil,
		Prompt:       nil,
		Sid:          nil,
	}
	setAuthorizeParamsDefault(p)

	if err = v.Struct(p); err != nil {
		t.Errorf(err.Error())
	}
}
