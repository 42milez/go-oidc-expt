package entity

import (
	"encoding/json"

	"github.com/42milez/go-oidc-server/app/typedef"
)

type Session struct {
	RedirectUri string
	UserID      typedef.UserID
}

func (s *Session) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}
