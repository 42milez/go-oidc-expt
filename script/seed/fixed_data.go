package main

import (
	"context"

	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/pkg/xargon2"
	"github.com/42milez/go-oidc-server/app/typedef"
)

func insertFixedData(ctx context.Context, db *datastore.Database) error {
	username := "swagger"
	password := "swagger"

	var user *ent.User
	var err error

	if user, err = insertFixedUser(ctx, db, username, password); err != nil {
		return err
	}

	clientID := "CDcp9v3Nn4i70FqWig5AuohmorD6MG"
	clientSecret := "nZ83cfW2yPmIItORmzYH9XdM5oLE7t"

	var relyingParty *ent.RelyingParty

	if relyingParty, err = insertFixedRelyingParty(ctx, db, clientID, clientSecret); err != nil {
		return err
	}

	code := "EYdxIU30xstnWZKxgA54RJMz1YUR0J"

	if _, err = insertAuthCode(ctx, db, code, user.ID, relyingParty.ID); err != nil {
		return nil
	}

	uri := "https://swagger.example.com/cb"

	if _, err = insertRedirectUri(ctx, db, uri, relyingParty.ID); err != nil {
		return err
	}

	return nil
}

func insertFixedUser(ctx context.Context, db *datastore.Database, username, password string) (*ent.User, error) {
	hashedPassword, err := xargon2.HashPassword(username)

	if err != nil {
		return nil, err
	}

	return db.Client.User.Create().SetName(username).SetPassword(hashedPassword).Save(ctx)
}

func insertFixedRelyingParty(ctx context.Context, db *datastore.Database, clientID, clientSecret string) (*ent.RelyingParty, error) {
	return db.Client.RelyingParty.Create().SetClientID(clientID).SetClientSecret(clientSecret).Save(ctx)
}

func insertAuthCode(ctx context.Context, db *datastore.Database, code string, userID typedef.UserID, relyingPartyID typedef.RelyingPartyID) (*ent.AuthCode, error) {
	return db.Client.AuthCode.Create().SetCode(code).SetUserID(userID).SetRelyingPartyID(relyingPartyID).Save(ctx)
}

func insertRedirectUri(ctx context.Context, db *datastore.Database, uri string, rpID typedef.RelyingPartyID) (*ent.RedirectUri, error) {
	return db.Client.RedirectUri.Create().SetURI(uri).SetRelyingPartyID(rpID).Save(ctx)
}
