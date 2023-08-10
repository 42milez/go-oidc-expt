package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
	"github.com/42milez/go-oidc-server/pkg/xerr"
	"github.com/redis/go-redis/v9"
)

const sessionTTL = 30 * time.Minute

const (
	ErrFailedToDeleteItem xerr.Err = "failed to delete item"
	ErrFailedToSaveItem   xerr.Err = "failed to save item"
	ErrFailedToLoadItem   xerr.Err = "failed to load item"
)

type Session struct {
	Cache *redis.Client
}

func (p *Session) SaveID(ctx context.Context, key string, id typedef.UserID) error {
	if err := p.Cache.Set(ctx, key, id, sessionTTL).Err(); err != nil {
		return xerr.Wrap(fmt.Errorf("%w : key=%s, id=%s", ErrFailedToSaveItem, key, id), err)
	}
	return nil
}

func (p *Session) LoadID(ctx context.Context, key string) (typedef.UserID, error) {
	ret, err := p.Cache.Get(ctx, key).Result()
	if err != nil {
		return "", xerr.Wrap(ErrFailedToLoadItem, err)
	}
	return typedef.UserID(ret), nil
}

func (p *Session) Delete(ctx context.Context, key string) error {
	if err := p.Cache.Del(ctx, key).Err(); err != nil {
		return xerr.Wrap(fmt.Errorf("%w : key=%s", ErrFailedToDeleteItem, key), err)
	}
	return nil
}
