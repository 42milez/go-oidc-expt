package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/42milez/go-oidc-server/app/idp/repository"

	"github.com/42milez/go-oidc-server/pkg/xutil"

	"github.com/42milez/go-oidc-server/app/idp/model"
)

const authCodeLen = 20

type Authorize struct {
	Repo *repository.User
}

func (p *Authorize) Authorize(ctx context.Context, param *model.AuthorizeRequest) (string, error) {
	// TODO: Generate authorization code
	code, err := xutil.MakeCryptoRandomString(authCodeLen)

	if err != nil {
		return "", err
	}

	userID, ok := xutil.GetUserID(ctx)

	if !ok {
		return "", errors.New("user id not found")
	}

	if _, err = p.Repo.SaveAuthorizationCode(ctx, userID, code); err != nil {
		return "", err
	}

	// TODO: Read redirect uri from database and verify it
	// ...

	// TODO: Return the authorization code and state
	return fmt.Sprintf("http://client.example.org/cb?code=%s&state=af0ifjsldk", code), nil
}
