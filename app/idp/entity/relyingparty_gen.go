// Code generated by app/entity/gen/gen.go; DO NOT EDIT.
package entity

import (
	"time"

	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/pkg/typedef"
)

func NewRelyingParty(entity *ent.RelyingParty) *RelyingParty {
	return &RelyingParty{
		entity: entity,
	}
}

type RelyingParty struct {
	entity *ent.RelyingParty
}

func (rp *RelyingParty) ID() typedef.RelyingPartyID {
	return rp.entity.ID
}

func (rp *RelyingParty) ClientID() string {
	return rp.entity.ClientID
}

func (rp *RelyingParty) ClientSecret() string {
	return rp.entity.ClientSecret
}

func (rp *RelyingParty) CreatedAt() time.Time {
	return rp.entity.CreatedAt
}

func (rp *RelyingParty) ModifiedAt() time.Time {
	return rp.entity.ModifiedAt
}
