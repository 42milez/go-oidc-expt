package repository

import (
	"context"
	"errors"

	"github.com/42milez/go-oidc-server/pkg/typedef"

	"github.com/42milez/go-oidc-server/pkg/ent/ent/redirecturi"
	"github.com/42milez/go-oidc-server/pkg/ent/ent/relyingparty"

	"github.com/42milez/go-oidc-server/cmd/datastore"
	"github.com/42milez/go-oidc-server/cmd/entity"

	"entgo.io/ent/dialect/sql"
	"github.com/42milez/go-oidc-server/pkg/xerr"
)

func NewRedirectURI(db *datastore.Database) *RedirectURI {
	return &RedirectURI{
		db: db,
	}
}

type RedirectURI struct {
	db *datastore.Database
}

func (ru *RedirectURI) ReadRedirectURI(ctx context.Context, clientID typedef.ClientID) (*entity.RedirectURI, error) {
	v, err := ru.db.Client.RedirectURI.Query().
		Where(func(s *sql.Selector) {
			t := sql.Table(relyingparty.Table)
			s.Where(
				sql.In(
					s.C(redirecturi.FieldRelyingPartyID),
					sql.Select(t.C(relyingparty.FieldID)).From(t).Where(sql.EQ(t.C(relyingparty.FieldClientID), clientID)),
				),
			)
		}).
		Only(ctx)
	if err != nil {
		if errors.As(err, &errEntNotFoundError) {
			return nil, xerr.RedirectURINotFound
		} else {
			return nil, err
		}
	}
	return entity.NewRedirectURI(v), nil
}
