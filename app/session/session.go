package session

import (
	"context"
	"net/http"

	"github.com/42milez/go-oidc-server/app/auth"
	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/redis/go-redis/v9"

	"github.com/42milez/go-oidc-server/app/ent/typedef"

	"github.com/42milez/go-oidc-server/app/entity"

	"github.com/42milez/go-oidc-server/pkg/xerr"
	"github.com/google/uuid"
)

const nRetrySaveSession = 3

type UserIDKey struct{}

type Session struct {
	Repo     ReadWriter
	TokenExt TokenExtractor
}

func (p *Session) Create(ctx context.Context, sess *entity.UserSession) (string, error) {
	var sessionID uuid.UUID
	var ok bool
	var err error

	for i := 0; i < nRetrySaveSession; i++ {
		if sessionID, err = uuid.NewRandom(); err != nil {
			return "", err
		}
		if ok, err = p.Repo.Write(ctx, sessionID.String(), sess); err != nil {
			return "", err
		}
		if ok {
			break
		}
	}

	if !ok {
		return "", xerr.SessionIDAlreadyExists
	}

	return sessionID.String(), nil
}

func (p *Session) Restore(r *http.Request, sessionID string) (*http.Request, error) {
	sess, err := p.Repo.Read(r.Context(), sessionID)

	if err != nil {
		return nil, err
	}

	ctx := context.WithValue(r.Context(), UserIDKey{}, sess.ID)

	return r.Clone(ctx), nil
}

func NewSession(rc *redis.Client, jwt *auth.JWT) *Session {
	return &Session{
		Repo: &repository.Session{
			Cache: rc,
		},
		TokenExt: jwt,
	}
}

func GetUserID(ctx context.Context) (typedef.UserID, bool) {
	ret, ok := ctx.Value(UserIDKey{}).(typedef.UserID)
	return ret, ok
}
