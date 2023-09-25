package api

import "net/http"

type TokenHdlr struct{}

func (th *TokenHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: Basic Authentication
	// ...

	// TODO: Generate Access Token, Refresh token and ID Token if grant_type is authorization_code.

	// Verify auth code（ owner, expiration date, duplication ）
	// ...

	// Verify redirect uri
	// ...

	// TODO: Generate Access Token if grant_type is refresh_token.
	// ...
}
