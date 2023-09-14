package model

type AuthorizeRequest struct {
	// REQUIRED
	ClientID     string `schema:"client_id" json:"clientID" validate:"required,alphanum"`
	Nonce        string `schema:"nonce" json:"nonce" validate:"required,ascii"`
	RedirectURI  string `schema:"redirect_uri" json:"redirectURI" validate:"required,printascii"`
	ResponseType string `schema:"response_type" json:"responseType" validate:"required,response-type-validator"`
	Scope        string `schema:"scope" json:"scope" validate:"required,scope-validator"`
	State        string `schema:"state" json:"state" validate:"required,alphanum"`
	// OPTIONAL
	Display     string `schema:"display" json:"display" validate:"required"`
	IDTokenHint string `schema:"id_token_hint" json:"idTokenHint" validate:"alpha"`
	MaxAge      int    `schema:"max_age" json:"maxAge" validate:"required"`
	Prompt      string `schema:"prompt" json:"prompt" validate:"required"`
}
