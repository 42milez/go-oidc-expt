package model

type ErrorResponse struct {
	Error string `json:"error" validate:"required"`
}
