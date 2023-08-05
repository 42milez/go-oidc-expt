package repository

//import (
//	"context"
//	"fmt"
//	"github.com/42milez/go-oidc-server/app/idp/auth"
//	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
//	"github.com/42milez/go-oidc-server/pkg/xerr"
//	"github.com/redis/go-redis/v9"
//	"net/http"
//)
//
//type EpStore struct {
//	client *redis.Client
//	JWT    *auth.JWTUtil
//}
//
//func (p *EpStore) saveID(ctx context.Context, key string, id typedef.UserID) error {
//	if err := p.client.Set(ctx, key, id, sessionTTL).Err(); err != nil {
//		return xerr.Wrap(fmt.Errorf("%w : key=%s, id=%s", ErrFailedToSaveItem, key, id), err)
//	}
//	return nil
//}
//
//func (p *EpStore) load(ctx context.Context, key string) (typedef.UserID, error) {
//	ret, err := p.client.Get(ctx, key).Result()
//	if err != nil {
//		return "", xerr.Wrap(ErrFailedToLoadItem, err)
//	}
//	return typedef.UserID(ret), nil
//}
//
//func (p *EpStore) delete(ctx context.Context, key string) error {
//	if err := p.client.Del(ctx, key).Err(); err != nil {
//		return xerr.Wrap(fmt.Errorf("%w : key=%s", ErrFailedToDeleteItem, key), err)
//	}
//	return nil
//}
//
//type IDKey struct{}
//
//func (p *EpStore) setID(ctx context.Context, id typedef.UserID) context.Context {
//	return context.WithValue(ctx, IDKey{}, id)
//}
//
//func (p *EpStore) getID(ctx context.Context) (typedef.UserID, bool) {
//	id, ok := ctx.Value(IDKey{}).(typedef.UserID)
//	return id, ok
//}
//
//func (p *EpStore) FillContext(r *http.Request) (*http.Request, error) {
//	token, err := p.JWT.ExtractToken(r)
//
//	if err != nil {
//		return nil, xerr.Wrap(ErrFailedToExtractToken, err)
//	}
//
//	id, err := p.load(r.Context(), token.JwtID())
//
//	if err != nil {
//		return nil, xerr.Wrap(ErrFailedToLoadItem, err)
//	}
//
//	ctx := p.setID(r.Context(), id)
//
//	return r.Clone(ctx), nil
//}
