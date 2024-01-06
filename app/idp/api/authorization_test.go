package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/42milez/go-oidc-server/app/idp/httpstore"
	"github.com/42milez/go-oidc-server/app/pkg/typedef"

	"github.com/42milez/go-oidc-server/app/idp/iface"
	"github.com/42milez/go-oidc-server/app/pkg/xstring"
	"github.com/42milez/go-oidc-server/app/pkg/xtestutil"
	"go.uber.org/mock/gomock"
)

const (
	tdAuthorizationDir         = "testdata/authorization/"
	tdAuthorizationRequest200  = tdAuthorizationDir + "302_req.txt"
	tdAuthorizationResponse200 = tdAuthorizationDir + "302_resp.txt"
)

func TestAuthorizeGet_ServeHTTP(t *testing.T) {
	type mockResp struct {
		location *url.URL
		err      error
	}

	type want struct {
		statusCode int
		respFile   string
	}

	clock := xtestutil.FixedClocker{}

	userID := typedef.UserID(475924034190589956)
	authTime := clock.Now()

	location := func(uri string) *url.URL {
		ret, err := url.Parse(uri)
		if err != nil {
			t.Fatal(err)
		}
		return ret
	}

	tests := map[string]struct {
		reqFile string
		resp    mockResp
		want    want
	}{
		"ok": {
			reqFile: tdAuthorizationRequest200,
			resp: mockResp{
				location: location("https://client.example.com/cb?code=SplxlOBeZQQYbYS6WxSbIA&state=af0ifjsldk"),
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
				"https://example.com/authorization",
				nil,
			)
			r.URL.RawQuery = strings.Replace(xstring.ByteToString(xtestutil.LoadFile(t, tt.reqFile)), "\n", "", -1)

			sess := &httpstore.Session{
				UserID:   userID,
				AuthTime: authTime,
			}
			r = r.Clone(context.WithValue(r.Context(), httpstore.SessionKey{}, sess))

			svcMock := NewMockAuthorizer(gomock.NewController(t))
			svcMock.EXPECT().Authorize(r.Context(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(tt.resp.location, "", nil).AnyTimes()
			svcMock.EXPECT().SaveAuthorizationRequestFingerprint(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

			ctxMock := iface.NewMockContextReader(gomock.NewController(t))
			ctxMock.EXPECT().Read(gomock.Any(), typedef.RequestParamKey{}).Return(&AuthorizeParams{
				ClientID:     "9NXtT29fw2lvmQ5EA42htc8sfNRQYe",
				Nonce:        "K45zJFN4L7tXjlXpFtVRjqWbSnSCz6",
				RedirectUri:  "https://example.com/cb",
				ResponseType: "code",
				Scope:        "openid",
				State:        "lgwyrqpZ0jLGQI5Ftu94HytJRJOJSa",
				Display:      "page",
				MaxAge:       600,
				Prompt:       "consent",
			}).AnyTimes()
			ctxMock.EXPECT().Read(gomock.Any(), httpstore.SessionKey{}).Return(sess).AnyTimes()

			v, err := NewOIDCRequestParamValidator()
			if err != nil {
				t.Fatal(err)
			}

			hdlr := &AuthorizationGet{
				svc:     svcMock,
				context: ctxMock,
				v:       v,
			}
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			wantResp := &xtestutil.Response{
				StatusCode: http.StatusFound,
				Location:   tt.resp.location,
				Body:       xtestutil.LoadFile(t, tt.want.respFile),
			}

			xtestutil.AssertResponse(t, wantResp, resp)
		})
	}
}
