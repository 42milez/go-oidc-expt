package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/42milez/go-oidc-server/app/pkg/xstring"

	"github.com/42milez/go-oidc-server/app/service"

	"github.com/42milez/go-oidc-server/app/pkg/xtestutil"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/golang/mock/gomock"
)

const (
	tdAuthenticationDir         = "testdata/authentication/"
	tdAuthenticationReqBody200  = tdAuthenticationDir + "200_req_body.json"
	tdAuthenticationReqParam200 = tdAuthenticationDir + "200_req_param.txt"
	tdAuthenticationResponse200 = tdAuthenticationDir + "200_resp.json"
	tdAuthenticationRequest400  = tdAuthenticationDir + "400_req.json"
	tdAuthenticationResponse400 = tdAuthenticationDir + "400_resp.json"
	tdAuthenticationResponse500 = tdAuthenticationDir + "500_resp.json"
)

const dummyUserID typedef.UserID = 475924035230777348

func TestAuthentication_ServeHTTP(t *testing.T) {
	type verifyPasswordMockResp struct {
		userID typedef.UserID
		err    error
	}

	type verifyConsentMockResp struct {
		ok  bool
		err error
	}

	type sessionMockResp struct {
		sessionID string
		err       error
	}

	type want struct {
		statusCode int
		resp       []byte
	}

	sessionID := "dd9a0158-092c-4dc2-b470-7e68c97bfdb0"

	tests := map[string]struct {
		reqBodyFile  string
		reqParamFile string
		respVPMock   verifyPasswordMockResp
		respVCMock   verifyConsentMockResp
		respSessMock sessionMockResp
		want         want
	}{
		"ok": {
			reqBodyFile:  tdAuthenticationReqBody200,
			reqParamFile: tdAuthenticationReqParam200,
			respVPMock: verifyPasswordMockResp{
				userID: dummyUserID,
				err:    nil,
			},
			respVCMock: verifyConsentMockResp{
				ok:  false,
				err: nil,
			},
			respSessMock: sessionMockResp{
				sessionID: sessionID,
				err:       nil,
			},
			want: want{
				statusCode: http.StatusFound,
				resp:       nil,
			},
		},
		//"BadRequest": {
		//	reqFile: tdAuthenticationRequest400,
		//	want: want{
		//		statusCode: http.StatusBadRequest,
		//		resp:       xtestutil.LoadFile(t, tdAuthenticationResponse400),
		//	},
		//},
		//"InternalServerError": {
		//	reqFile: tdAuthenticationRequest200,
		//	respVPMock: verifyPasswordMockResp{
		//		err: xtestutil.DummyError,
		//	},
		//	want: want{
		//		statusCode: http.StatusInternalServerError,
		//		resp:       xtestutil.LoadFile(t, tdAuthenticationResponse500),
		//	},
		//},
	}

	for n, tt := range tests {
		tt := tt

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPost,
				"/authentication",
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

			sessMock := NewMockSessionCreator(gomock.NewController(t))
			sessMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(tt.respSessMock.sessionID, tt.respSessMock.err).AnyTimes()

			v, err := NewAuthorizeParamValidator()

			if err != nil {
				t.Error(err)
			}

			sut := AuthenticateHdlr{
				service:   svcMock,
				session:   sessMock,
				cookie:    service.NewCookie(rawHashKey, rawBlockKey, xtestutil.FixedClocker{}),
				validator: v,
			}
			sut.ServeHTTP(w, r)

			resp := w.Result()

			wantResp := &xtestutil.Response{
				StatusCode: tt.want.statusCode,
				Body:       tt.want.resp,
			}

			if tt.want.resp != nil {
				xtestutil.AssertResponseJSON(t, resp, wantResp)
			} else {
				xtestutil.AssertResponse(t, resp, wantResp)
			}
		})
	}
}
