package model

type Authorize struct {
	Scope        string `json:"scope" validate:"ScopeValidator"`
	ResponseType string `json:"responseType" validate:"required,alpha"`
	ClientID     string `json:"clientID" validate:"required,alphanum"`
	RedirectURI  string `json:"redirectURI" validate:"required,printascii"`
	State        string `json:"state" validate:"required,alphanum"`
	ResponseMode string `json:"responseMode" validate:"required,alpha"`
	Nonce        string `json:"nonce" validate:"required,alphanum"`
	Display      string `json:"display" validate:"required"`
	Prompt       string `json:"prompt" validate:"required"`
	MaxAge       int    `json:"maxAge" validate:"required"`
	UILocales    string `json:"uiLocales" validate:"alpha"`
	IDTokenHint  string `json:"idTokenHint" validate:"alpha"`
}
