package model

type AuthorizeRequest struct {
	ClientID     string `schema:"client_id" json:"clientID" validate:"required,alphanum"`
	Display      string `schema:"display" json:"display" validate:"required"`
	IDTokenHint  string `schema:"id_token_hint" json:"idTokenHint" validate:"required,alpha"`
	MaxAge       int    `schema:"max_age" json:"maxAge" validate:"required"`
	Nonce        string `schema:"nonce" json:"nonce" validate:"required,ascii"`
	Prompt       string `schema:"prompt" json:"prompt" validate:"required"`
	RedirectURI  string `schema:"redirect_uri" json:"redirectURI" validate:"required,printascii"`
	ResponseType string `schema:"response_type" json:"responseType" validate:"required,response-type-validator"`
	Scope        string `schema:"scope" json:"scope" validate:"required,scope-validator"`
	State        string `schema:"state" json:"state" validate:"required,alphanum"`
}
