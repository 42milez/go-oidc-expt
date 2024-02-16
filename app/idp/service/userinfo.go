package service

import (
	"context"
	"strconv"

	"github.com/42milez/go-oidc-server/app/idp/entity"
	"github.com/42milez/go-oidc-server/app/idp/option"
	"github.com/42milez/go-oidc-server/app/idp/repository"
	"github.com/42milez/go-oidc-server/app/pkg/typedef"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type UserInfo struct {
	repo UserReader
}

func NewUserInfo(opt *option.Option) *UserInfo {
	return &UserInfo{
		repo: repository.NewUser(opt),
	}
}

func (ui *UserInfo) ReadUserInfo(ctx context.Context, accessToken jwt.Token) (*entity.UserInfo, error) {
	sub := accessToken.Subject()

	userID, err := strconv.Atoi(sub)
	if err != nil {
		return nil, err
	}

	user, err := ui.repo.ReadUserByID(ctx, typedef.UserID(userID))
	if err != nil {
		return nil, err
	}

	return &entity.UserInfo{
		ID:   user.ID(),
		Name: user.Name(),
	}, nil
}