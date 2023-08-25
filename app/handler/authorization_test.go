package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/42milez/go-oidc-server/app/ent/typedef"

	"github.com/42milez/go-oidc-server/app/testutil"

	"github.com/42milez/go-oidc-server/app/validation"

	"github.com/42milez/go-oidc-server/pkg/xerr"
	"github.com/42milez/go-oidc-server/pkg/xutil"
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

	const userID typedef.UserID = "01H8P9NBW77WMNN2ZTNACGZ19X"

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
			r.URL.RawQuery = strings.Replace(xutil.ByteToString(testutil.LoadFile(t, tt.reqFile)), "\n", "", -1)
			r = r.Clone(context.WithValue(r.Context(), UserIDKey{}, userID))

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
				service:   svcMock,
				validator: v,
			}
			hdlr.ServeHTTP(w, r)
			resp := w.Result()

			wantResp := &testutil.Response{
				StatusCode: http.StatusFound,
				Location:   tt.resp.location,
				Body:       testutil.LoadFile(t, tt.want.respFile),
			}

			testutil.AssertResponse(t, resp, wantResp)
		})
	}
}
