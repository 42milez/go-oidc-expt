package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/42milez/go-oidc-server/app/iface"

	"github.com/42milez/go-oidc-server/app/httpstore"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/pkg/xstring"

	"github.com/42milez/go-oidc-server/app/pkg/xtestutil"

	"github.com/42milez/go-oidc-server/app/typedef"

	"go.uber.org/mock/gomock"
)

func TestAuthentication_ServeHTTP(t *testing.T) {
	const (
		testdataDir   = "testdata/authentication/"
		td200ReqBody  = testdataDir + "200_req_body.json"
		td200ReqParam = testdataDir + "200_req_param.txt"
		td400ReqBody  = testdataDir + "400_req_body.json"
		td400ReqParam = testdataDir + "400_req_param.txt"
		td400RespBody = testdataDir + "400_resp_body.json"
		td401RespBody = testdataDir + "401_resp_body.json"
		td500ResBody  = testdataDir + "500_resp_body.json"
	)

	type verifyConsentMockResp struct {
		ok  bool
		err error
	}

	type verifyPasswordMockResp struct {
		userID typedef.UserID
		err    error
	}

	type sessionMockResp struct {
		sessionID typedef.SessionID
		err       error
	}

	type want struct {
		statusCode int
		resp       []byte
	}

	const userID = 475924035230777348
	const sessionID = typedef.SessionID(484481116225536365)

	tests := map[string]struct {
		reqBodyFile  string
		reqParamFile string
		respSessMock sessionMockResp
		respVCMock   verifyConsentMockResp
		respVPMock   verifyPasswordMockResp
		want         want
	}{
		"OK": {
			reqBodyFile:  td200ReqBody,
			reqParamFile: td200ReqParam,
			respSessMock: sessionMockResp{
				sessionID: sessionID,
				err:       nil,
			},
			respVCMock: verifyConsentMockResp{
				ok:  false,
				err: nil,
			},
			respVPMock: verifyPasswordMockResp{
				userID: userID,
				err:    nil,
			},
			want: want{
				statusCode: http.StatusFound,
				resp:       nil,
			},
		},
		"BadRequest_FailedToParseQueryParam": {
			reqBodyFile:  td200ReqBody,
			reqParamFile: td400ReqParam,
			want: want{
				statusCode: http.StatusBadRequest,
				resp:       xtestutil.LoadFile(t, td400RespBody),
			},
		},
		"BadRequest_FailedToParseRequestBody": {
			reqBodyFile:  td400ReqBody,
			reqParamFile: td200ReqParam,
			want: want{
				statusCode: http.StatusBadRequest,
				resp:       xtestutil.LoadFile(t, td400RespBody),
			},
		},
		"Unauthorized_PasswordNotMatched": {
			reqBodyFile:  td200ReqBody,
			reqParamFile: td200ReqParam,
			respVPMock: verifyPasswordMockResp{
				userID: 0,
				err:    xerr.PasswordNotMatched,
			},
			want: want{
				statusCode: http.StatusUnauthorized,
				resp:       xtestutil.LoadFile(t, td401RespBody),
			},
		},
		"Unauthorized_UserNotFound": {
			reqBodyFile:  td200ReqBody,
			reqParamFile: td200ReqParam,
			respVPMock: verifyPasswordMockResp{
				userID: 0,
				err:    xerr.UserNotFound,
			},
			want: want{
				statusCode: http.StatusUnauthorized,
				resp:       xtestutil.LoadFile(t, td401RespBody),
			},
		},
		"InternalServerError_FailedToCreateSession": {
			reqBodyFile:  td200ReqBody,
			reqParamFile: td200ReqParam,
			respSessMock: sessionMockResp{
				sessionID: 0,
				err:       xerr.FailedToWriteSession,
			},
			respVPMock: verifyPasswordMockResp{
				userID: userID,
				err:    nil,
			},
			want: want{
				statusCode: http.StatusInternalServerError,
				resp:       xtestutil.LoadFile(t, td500ResBody),
			},
		},
	}

	for n, tt := range tests {
		tt := tt

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPost,
				config.AuthenticationPath,
				bytes.NewReader(xtestutil.LoadFile(t, tt.reqBodyFile)))
			r.URL.RawQuery = strings.Replace(xstring.ByteToString(xtestutil.LoadFile(t, tt.reqParamFile)), "\n", "", -1)

			ctx := r.Context()

			svcMock := NewMockAuthenticator(gomock.NewController(t))
			svcMock.
				EXPECT().
				VerifyPassword(ctx, gomock.Any(), gomock.Any()).
				Return(tt.respVPMock.userID, tt.respVPMock.err).
				AnyTimes()
			svcMock.
				EXPECT().
				VerifyConsent(ctx, gomock.Any(), gomock.Any()).
				Return(tt.respVCMock.ok, tt.respVCMock.err).
				AnyTimes()

			sessMock := iface.NewMockUserInfoSessionWriter(gomock.NewController(t))
			sessMock.EXPECT().WriteUserInfo(gomock.Any(), gomock.Any()).Return(tt.respSessMock.sessionID, tt.respSessMock.err).AnyTimes()

			v, err := NewAuthorizeParamValidator()

			if err != nil {
				t.Error(err)
			}

			sut := AuthenticateHdlr{
				svc:  svcMock,
				sess: sessMock,
				ck:   httpstore.NewCookie(rawHashKey, rawBlockKey, xtestutil.FixedClocker{}),
				v:    v,
			}
			sut.ServeHTTP(w, r)

			resp := w.Result()

			wantResp := &xtestutil.Response{
				StatusCode: tt.want.statusCode,
				Body:       tt.want.resp,
			}

			if tt.want.resp != nil {
				xtestutil.AssertResponseJSON(t, wantResp, resp)
			} else {
				xtestutil.AssertResponse(t, wantResp, resp)
			}
		})
	}
}
