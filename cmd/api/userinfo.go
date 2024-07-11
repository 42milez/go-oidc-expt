package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/cmd/service"

	"github.com/42milez/go-oidc-server/cmd/iface"
	"github.com/42milez/go-oidc-server/cmd/option"
)

var userinfo *UserInfo

func InitUserInfo(opt *option.Option) {
	if userinfo == nil {
		userinfo = &UserInfo{
			svc:   service.NewUserInfo(opt),
			token: opt.Token,
		}
	}
}

type UserInfo struct {
	svc   UserInfoReader
	token iface.TokenParser
}

func (ui *UserInfo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tokenRaw, err := extractAccessToken(r)
	if err != nil {
		RespondServerError(w, r, err)
		return
	}

	t, err := ui.token.Parse(tokenRaw)
	if err != nil {
		RespondServerError(w, r, err)
		return
	}

	ctx := r.Context()
	userInfo, err := ui.svc.ReadUserInfo(ctx, t)
	if err != nil {
		RespondServerError(w, r, err)
		return
	}

	RespondJSON(w, r, http.StatusOK, nil, userInfo)
}
