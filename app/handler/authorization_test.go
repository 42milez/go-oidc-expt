package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/42milez/go-oidc-server/app/handler/session"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/pkg/xtestutil"

	"github.com/42milez/go-oidc-server/pkg/xstring"

	"github.com/42milez/go-oidc-server/app/validation"

	"github.com/42milez/go-oidc-server/pkg/xerr"
	"github.com/golang/mock/gomock"
)

const (
	tdAuthorizationDir         = "testdata/authorization/"
	tdAuthorizationRequest200  = tdAuthorizationDir + "302_req.txt"
	tdAuthorizationResponse200 = tdAuthorizationDir + "302_resp.txt"
)

func TestAuthorizeGet_ServeHTTP(t *testing.T) {
	type mockResp struct {
		location string
		err      error
	}

	type want struct {
		statusCode int
		respFile   string
	}

	const userID typedef.UserID = 475924034190589956

	tests := map[string]struct {
		reqFile string
		resp    mockResp
		want    want
	}{
		"OK": {
			reqFile: tdAuthorizationRequest200,
			resp: mockResp{
				location: "https://client.example.com/cb?code=SplxlOBeZQQYbYS6WxSbIA&state=af0ifjsldk",
				err:      nil,
			},
			want: want{
				statusCode: http.StatusOK,
				respFile:   tdAuthorizationResponse200,
			},
		},
	}

	for n, tt := range tests {
		tt := tt

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodGet,
				"http://idp/v1/authorize",
				nil,
			)
			r.URL.RawQuery = strings.Replace(xstring.ByteToString(xtestutil.LoadFile(t, tt.reqFile)), "\n", "", -1)
			r = r.Clone(context.WithValue(r.Context(), session.UserIDKey{}, userID))

			svcMock := NewMockAuthorizer(gomock.NewController(t))
			svcMock.
				EXPECT().
				Authorize(r.Context(), gomock.Any(), gomock.Any()).
				Return(tt.resp.location, nil).
				AnyTimes()

			v, err := validation.NewAuthorizeValidator()

			if err != nil {
				t.Error(xerr.FailedToInitialize)
			}

			hdlr := &AuthorizeGet{
				Service:   svcMock,
				validator: v,
			}
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			wantResp := &xtestutil.Response{
				StatusCode: http.StatusFound,
				Location:   tt.resp.location,
				Body:       xtestutil.LoadFile(t, tt.want.respFile),
			}

			xtestutil.AssertResponse(t, resp, wantResp)
		})
	}
}