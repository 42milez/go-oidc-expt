package entity

import "github.com/42milez/go-oidc-server/app/ent/ent"

func NewRelyingParty(entity *ent.RelyingParty) *RelyingParty {
	return &RelyingParty{
		entity: entity,
	}
}

type RelyingParty struct {
	entity *ent.RelyingParty
}

func (r *RelyingParty) ClientId() string {
	return r.entity.ClientID
}

func (r *RelyingParty) ClientSecret() string {
	return r.entity.ClientSecret
}
