package api

import (
	"encoding/base64"
	"net/http"

	"github.com/42milez/go-oidc-server/app/idp/security"
)

var jwks *Jwks

func InitJwks() {
	if jwks == nil {
		jwks = &Jwks{}
	}
}

type Jwks struct{}

func (j *Jwks) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pub, fp, err := security.DecodePublicKey()
	if err != nil {
		RespondJSON500(w, r, err)
		return
	}

	b64X := base64.RawURLEncoding.EncodeToString(pub.X.Bytes())
	b64Y := base64.RawURLEncoding.EncodeToString(pub.Y.Bytes())
	key := JWK{
		Kty: EC,
		Crv: P256,
		Use: "sig",
		X:   b64X,
		Y:   b64Y,
		Kid: fp,
	}
	respBody := &JwksResponse{
		Keys: []JWK{
			key,
		},
	}
	RespondJSON(w, r, http.StatusOK, nil, respBody)
}
