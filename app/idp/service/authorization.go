package service

import (
	"context"
	"fmt"

	"github.com/42milez/go-oidc-server/pkg/xutil"

	"github.com/42milez/go-oidc-server/app/idp/model"
)

const authCodeLen = 20

type Authorize struct {
	Endpoint string
}

func (p *Authorize) Authorize(ctx context.Context, param *model.AuthorizeRequest) (string, error) {
	// TODO: Generate authorization code
	code, err := xutil.MakeCryptoRandomString(authCodeLen)

	if err != nil {
		return "", err
	}

	// TODO: Save authorization code into database
	// ...

	// TODO: Read redirect uri from database and verify it
	// ...

	// TODO: Return the authorization code and state
	return fmt.Sprintf("http://client.example.org/cb?code=%s&state=af0ifjsldk", code), nil
}
