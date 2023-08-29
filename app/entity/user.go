package entity

import (
	"encoding/json"

	"github.com/42milez/go-oidc-server/app/ent/typedef"
)

type UserSession struct {
	ID typedef.UserID
}

func (p *UserSession) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}
