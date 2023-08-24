package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/42milez/go-oidc-server/pkg/xutil"

	"github.com/42milez/go-oidc-server/app/idp/auth"
	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
	"github.com/42milez/go-oidc-server/app/idp/repository"
)

func insertUsers(ctx context.Context, client *ent.Client) ([]*ent.User, error) {
	params := []struct {
		name   string
		pwHash string
	}{
		{name: "user01"},
		{name: "user02"},
	}

	for i, v := range params {
		pwHash, err := auth.HashPassword(v.name)
		if err != nil {
			return nil, err
		}
		params[i].pwHash = pwHash
	}

	builders := make([]*ent.UserCreate, len(params))

	for i, v := range params {
		builders[i] = client.User.Create().SetName(v.name).SetPassword(v.pwHash)
	}

	return client.User.CreateBulk(builders...).Save(ctx)
}

func insertAuthCodes(ctx context.Context, client *ent.Client, users []*ent.User) ([]*ent.AuthCode, error) {
	type param struct {
		code     string
		expireAt time.Time
		userID   typedef.UserID
	}

	nCodeByUser := 2
	params := make([]*param, len(users)*nCodeByUser)
	expireAt := time.Now().Add(config.AuthCodeLifetime)

	for i := range params {
		code, err := xutil.MakeCryptoRandomString(config.AuthCodeLength)
		if err != nil {
			return nil, err
		}
		params[i].code = code
		params[i].expireAt = expireAt
		params[i].userID = users[i%nCodeByUser].ID
	}

	builders := make([]*ent.AuthCodeCreate, len(params))

	for i, v := range params {
		builders[i] = client.AuthCode.Create().SetCode(v.code).SetExpireAt(v.expireAt).SetUserID(v.userID)
	}

	return client.AuthCode.CreateBulk(builders...).Save(ctx)
}

func insertRedirectURIs(ctx context.Context, client *ent.Client, users []*ent.User) ([]*ent.RedirectURI, error) {
	type param struct {
		uri    string
		userID typedef.UserID
	}

	nURIByUser := 2
	params := make([]*param, len(users)*nURIByUser)

	for i := range params {
		params[i].uri = fmt.Sprintf("http://example.com/cb%d", i)
		params[i].userID = users[i%nURIByUser].ID
	}

	builders := make([]*ent.RedirectURICreate, len(params))

	for i, v := range params {
		builders[i] = client.RedirectURI.Create().SetURI(v.uri).SetUserID(v.userID)
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
