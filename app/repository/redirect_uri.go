package repository

import (
	"context"
	"errors"

	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/ent/ent/redirecturi"
	"github.com/42milez/go-oidc-server/app/ent/ent/relyingparty"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
)

func NewRedirectUri(db *datastore.Database) *RedirectUri {
	return &RedirectUri{
		db: db,
	}
}

type RedirectUri struct {
	db *datastore.Database
}

func (ru *RedirectUri) ReadRedirectUri(ctx context.Context, uri, clientId string) (*ent.RedirectURI, error) {
	ret, err := ru.db.Client.RedirectURI.Query().Where(redirecturi.URI(uri)).WithRelyingParty(func(q *ent.RelyingPartyQuery) {
		q.Where(relyingparty.ClientID(clientId))
	}).Only(ctx)

	if err != nil {
		if errors.As(err, &errEntNotFoundError) {
			return nil, xerr.RedirectUriNotFound
		} else {
			return nil, err
		}
	}

	return ret, nil
}
