// Code generated by app/entity/gen/gen.go; DO NOT EDIT.
package entity

import (
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/typedef"
	"time"
)

func NewRedirectUri(entity *ent.RedirectUri) *RedirectUri {
	return &RedirectUri{
		entity: entity,
	}
}

type RedirectUri struct {
	entity *ent.RedirectUri
}

func (ru *RedirectUri) ID() typedef.RedirectURIID {
	return ru.entity.ID
}

func (ru *RedirectUri) URI() string {
	return ru.entity.URI
}

func (ru *RedirectUri) CreatedAt() time.Time {
	return ru.entity.CreatedAt
}

func (ru *RedirectUri) ModifiedAt() time.Time {
	return ru.entity.ModifiedAt
}

func (ru *RedirectUri) RelyingPartyID() typedef.RelyingPartyID {
	return ru.entity.RelyingPartyID
}
