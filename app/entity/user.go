package entity

import (
	"encoding/json"

	"github.com/42milez/go-oidc-server/app/typedef"
)

type Session struct {
	UserID  typedef.UserID
	Consent bool
}

func (p *Session) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}
