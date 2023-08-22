package main

import (
	"context"
	"github.com/42milez/go-oidc-server/app/idp/auth"
	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
	"github.com/42milez/go-oidc-server/app/idp/repository"
	"github.com/42milez/go-oidc-server/pkg/xutil"
	"log"
	"time"
)

func insertUsers(ctx context.Context, client *ent.Client) ([]*ent.User, error) {
	params := []struct{
		name string
		pwHash typedef.PasswordHash
	}{
		{name: "user01"},
		{name: "user02"},
	}

	for i, param := range params {
		pwHash, err := auth.GeneratePasswordHash(param.name)
		if err != nil {
			return nil, err
		}
		params[i].pwHash = pwHash
	}

	builders := make([]*ent.UserCreate, len(params))

	for i, param := range params {
		builders[i] = client.User.Create().SetName(param.name).SetPasswordHash(param.pwHash)
	}

	return client.User.CreateBulk(builders...).Save(ctx)
}

func insertAuthCodes(ctx context.Context, client *ent.Client, users []*ent.User) ([]*ent.AuthCode, error) {
	params := []struct{
		code string
		expireAt time.Time
		userID typedef.UserID
	}{
		// TODO: Use const variable
		{expireAt: time.Now().Add(10*time.Minute)},
		{expireAt: time.Now().Add(10*time.Minute)},
	}

	for i, user := range users {
		params[i].userID = user.ID
		// TODO: Use const variable
		code, err := xutil.MakeCryptoRandomString(20)
		if err != nil {
			return nil, err
		}
		params[i].code = code
	}

	builders := make([]*ent.AuthCodeCreate, len(params))

	for i, param := range params {
		builders[i] = client.AuthCode.Create().SetCode(param.code).SetExpireAt(param.expireAt).SetUserID(param.userID)
	}

	return client.AuthCode.CreateBulk(builders...).Save(ctx)
}

func insertRedirectURIs(ctx context.Context, client *ent.Client, users []*ent.User) ([]*ent.RedirectURI, error) {
	params := []struct{
		uri string
		userID typedef.UserID
	}{
		{uri: "http://example.com/cb1"},
		{uri: "http://example.com/cb2"},
	}

	for i, user := range users {
		params[i].userID = user.ID
	}

	builders := make([]*ent.RedirectURICreate, len(params))

	for i, param := range params {
		builders[i] = client.RedirectURI.Create().SetURI(param.uri).SetUserID(param.userID)
	}

	return client.RedirectURI.CreateBulk(builders...).Save(ctx)
}

func run(ctx context.Context, client *ent.Client) error {
	users, err := insertUsers(ctx, client)

	if err != nil {
		return err
	}

	_, err = insertAuthCodes(ctx, client, users)

	if err != nil {
		return err
	}

	_, err = insertRedirectURIs(ctx, client, users)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	ctx := context.Background()
	cfg, err := config.New()

	if err != nil {
		log.Fatal(err)
	}

	cfg.DBPort = 13306
	dbClient, err := repository.NewDBClient(ctx, cfg)

	if err != nil {
		log.Fatal(err)
	}

	entClient := repository.NewEntClient(dbClient)

	if err = run(ctx, entClient); err != nil {
		log.Fatal(err)
	}
}
