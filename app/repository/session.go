package repository

import (
	"context"
	"time"

	"github.com/42milez/go-oidc-server/app/ent/typedef"

	"github.com/redis/go-redis/v9"
)

const sessionTTL = 30 * time.Minute

type Session struct {
	Cache *redis.Client
}

func (p *Session) SaveUserID(ctx context.Context, key string, id typedef.UserID) error {
	if err := p.Cache.Set(ctx, key, id, sessionTTL).Err(); err != nil {
		return err
	}
	return nil
}

func (p *Session) LoadUserID(ctx context.Context, key string) (typedef.UserID, error) {
	ret, err := p.Cache.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return typedef.UserID(ret), nil
}

func (p *Session) Delete(ctx context.Context, key string) error {
	if err := p.Cache.Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}
