package repository

import (
	"context"
	"encoding/json"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/entity"

	"github.com/redis/go-redis/v9"
)

type Session struct {
	Cache *redis.Client
}

func (p *Session) Create(ctx context.Context, sid typedef.SessionID, sess *entity.Session) (bool, error) {
	return p.Cache.SetNX(ctx, string(sid), sess, config.SessionTTL).Result()
}

func (p *Session) Read(ctx context.Context, sid typedef.SessionID) (*entity.Session, error) {
	v, err := p.Cache.Get(ctx, string(sid)).Result()

	if err != nil {
		return nil, err
	}

	ret := &entity.Session{}

	if err = json.Unmarshal([]byte(v), ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (p *Session) Update(ctx context.Context, sid typedef.SessionID, sess *entity.Session) (string, error) {
	return p.Cache.Set(ctx, string(sid), sess, redis.KeepTTL).Result()
}

func (p *Session) Delete(ctx context.Context, sid typedef.SessionID) error {
	if err := p.Cache.Del(ctx, string(sid)).Err(); err != nil {
		return err
	}
	return nil
}
