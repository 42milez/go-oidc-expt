package model

type AuthenticateRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthenticateResponse struct {
	Error string `json:"error"`
}
