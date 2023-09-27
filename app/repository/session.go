package repository

import (
	"context"
	"encoding/json"

	"github.com/42milez/go-oidc-server/app/datastore"

	"github.com/42milez/go-oidc-server/app/config"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/entity"

	"github.com/redis/go-redis/v9"
)

func NewSession(cache *datastore.Cache) *Session {
	return &Session{
		cache: cache,
	}
}

type Session struct {
	cache *datastore.Cache
}

func (s *Session) Create(ctx context.Context, sid typedef.SessionID, sess *entity.Session) (bool, error) {
	return s.cache.Client.SetNX(ctx, string(sid), sess, config.SessionTTL).Result()
}

func (s *Session) Read(ctx context.Context, sid typedef.SessionID) (*entity.Session, error) {
	v, err := s.cache.Client.Get(ctx, string(sid)).Result()

	if err != nil {
		return nil, err
	}

	ret := &entity.Session{}

	if err = json.Unmarshal([]byte(v), ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *Session) Update(ctx context.Context, sid typedef.SessionID, sess *entity.Session) (string, error) {
	return s.cache.Client.Set(ctx, string(sid), sess, redis.KeepTTL).Result()
}

func (s *Session) Delete(ctx context.Context, sid typedef.SessionID) error {
	return s.cache.Client.Del(ctx, string(sid)).Err()
}
