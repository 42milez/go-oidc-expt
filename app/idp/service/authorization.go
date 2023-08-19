package service

import (
	"context"

	"github.com/42milez/go-oidc-server/app/idp/model"
)

type Authorize struct {
	Endpoint string
}

func (p *Authorize) Authorize(ctx context.Context, param *model.AuthorizeRequest) (string, error) {
	// TODO: Generate authorization code
	// ...

	// TODO: Save authorization code
	// ...

	// TODO: Read redirect uri from database
	// ...

	// TODO: Return the authorization code and state
	return "http://client.example.org/cb?code=SplxlOBeZQQYbYS6WxSbIA&state=af0ifjsldk", nil
}
