package model

type Authorize struct {
	ClientID     string `json:"clientID" validate:"required,alphanum"`
	Display      string `json:"display" validate:"required"`
	IDTokenHint  string `json:"idTokenHint" validate:"alpha"`
	MaxAge       int    `json:"maxAge" validate:"required"`
	Nonce        string `json:"nonce" validate:"required,alphanum"`
	Prompt       string `json:"prompt" validate:"required"`
	RedirectURI  string `json:"redirectURI" validate:"required,printascii"`
	ResponseType string `json:"responseType" validate:"required,response-type-validator"`
	Scope        string `json:"scope" validate:"required,scope-validator"`
	State        string `json:"state" validate:"required,alphanum"`
}
