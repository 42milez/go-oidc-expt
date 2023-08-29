package repository

import (
	"context"
	"encoding/json"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/entity"

	"github.com/redis/go-redis/v9"
)

type Session struct {
	Cache *redis.Client
}

func (p *Session) Write(ctx context.Context, key string, sess *entity.UserSession) (bool, error) {
	return p.Cache.SetNX(ctx, key, sess, config.SessionTTL).Result()
}

func (p *Session) Read(ctx context.Context, key string) (*entity.UserSession, error) {
	v, err := p.Cache.Get(ctx, key).Result()

	if err != nil {
		return nil, err
	}

	ret := &entity.UserSession{}

	if err = json.Unmarshal([]byte(v), ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (p *Session) Delete(ctx context.Context, key string) error {
	if err := p.Cache.Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}
