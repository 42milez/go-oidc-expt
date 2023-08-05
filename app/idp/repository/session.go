package repository

import (
	"github.com/42milez/go-oidc-server/app/idp/auth"
	"github.com/redis/go-redis/v9"
)

//const sessionTTL = 30 * time.Minute

//const (
//	ErrFailedToDeleteItem   xerr.Err = "failed to delete item"
//	ErrFailedToExtractToken xerr.Err = "failed to extract token"
//	ErrFailedToSaveItem     xerr.Err = "failed to save item"
//	ErrFailedToLoadItem     xerr.Err = "failed to load item"
//)

type Session struct {
	Cache *redis.Client
	JWT   *auth.JWTUtil
}

//func (p *Session) SaveID(ctx context.Context) error {
//
//}

//func (p *Session) LoadID(ctx context.Context) {
//
//}
