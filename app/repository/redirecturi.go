package repository

import (
	"context"
	"errors"

	"github.com/42milez/go-oidc-server/app/entity"

	"entgo.io/ent/dialect/sql"
	"github.com/42milez/go-oidc-server/app/datastore"
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

func (ru *RedirectUri) ReadRedirectUri(ctx context.Context, clientId string) (*entity.RedirectUri, error) {
	v, err := ru.db.Client.RedirectUri.Query().
		Where(func(s *sql.Selector) {
			t := sql.Table(relyingparty.Table)
			s.Where(
				sql.In(
					s.C(redirecturi.FieldRelyingPartyID),
					sql.Select(t.C(relyingparty.FieldID)).From(t).Where(sql.EQ(t.C(relyingparty.FieldClientID), clientId)),
				),
			)
		}).
		Only(ctx)
	if err != nil {
		if errors.As(err, &errEntNotFoundError) {
			return nil, xerr.RedirectUriNotFound
		} else {
			return nil, err
		}
	}
	return entity.NewRedirectUri(v), nil
}