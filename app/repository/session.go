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

type CreateSession struct {
	Cache *datastore.Cache
}

func (p *CreateSession) Create(ctx context.Context, sid typedef.SessionID, sess *entity.Session) (bool, error) {
	return p.Cache.Client.SetNX(ctx, string(sid), sess, config.SessionTTL).Result()
}

type ReadSession struct {
	Cache *datastore.Cache
}

func (p *ReadSession) Read(ctx context.Context, sid typedef.SessionID) (*entity.Session, error) {
	v, err := p.Cache.Client.Get(ctx, string(sid)).Result()

	if err != nil {
		return nil, err
	}

	ret := &entity.Session{}

	if err = json.Unmarshal([]byte(v), ret); err != nil {
		return nil, err
	}

	return ret, nil
}

type UpdateSession struct {
	Cache *datastore.Cache
}

func (p *UpdateSession) Update(ctx context.Context, sid typedef.SessionID, sess *entity.Session) (string, error) {
	return p.Cache.Client.Set(ctx, string(sid), sess, redis.KeepTTL).Result()
}

type DeleteSession struct {
	Cache *datastore.Cache
}

func (p *DeleteSession) Delete(ctx context.Context, sid typedef.SessionID) error {
	if err := p.Cache.Client.Del(ctx, string(sid)).Err(); err != nil {
		return err
	}
	return nil
}
