package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/42milez/go-oidc-server/app/pkg/xtestutil"

	"github.com/42milez/go-oidc-server/app/api/cookie"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
)

const (
	tdAuthenticationDir         = "testdata/authentication/"
	tdAuthenticationRequest200  = tdAuthenticationDir + "200_req.json"
	tdAuthenticationResponse200 = tdAuthenticationDir + "200_resp.json"
	tdAuthenticationRequest400  = tdAuthenticationDir + "400_req.json"
	tdAuthenticationResponse400 = tdAuthenticationDir + "400_resp.json"
	tdAuthenticationResponse500 = tdAuthenticationDir + "500_resp.json"
)

const dummyUserID typedef.UserID = 475924035230777348

func TestAuthentication_ServeHTTP(t *testing.T) {
	type serviceMockResp struct {
		userID typedef.UserID
		err    error
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
	cookieHashKey := "nlmUN8ccpAIgCFWtminsNkr6uJU0YrPquFE7eqbXAH1heOYddNjV1Ni3YSZWdpob"
	cookieBlockKey := "aMe6Jbqnnee4lXR0PHC2Eg5gaB5Mv5p5"

	tests := map[string]struct {
		reqFile      string
		respSVCMock  serviceMockResp
		respSessMock sessionMockResp
		want         want
	}{
		"OK": {
			reqFile: tdAuthenticationRequest200,
			respSVCMock: serviceMockResp{
				userID: dummyUserID,
				err:    nil,
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
		"BadRequest": {
			reqFile: tdAuthenticationRequest400,
			want: want{
				statusCode: http.StatusBadRequest,
				resp:       xtestutil.LoadFile(t, tdAuthenticationResponse400),
			},
		},
		"InternalServerError": {
			reqFile: tdAuthenticationRequest200,
			respSVCMock: serviceMockResp{
				err: xtestutil.DummyError,
			},
			want: want{
				statusCode: http.StatusInternalServerError,
				resp:       xtestutil.LoadFile(t, tdAuthenticationResponse500),
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
				"/authentication",
				bytes.NewReader(xtestutil.LoadFile(t, tt.reqFile)))

			svcMock := NewMockAuthenticator(gomock.NewController(t))
			svcMock.
				EXPECT().
				Authenticate(r.Context(), gomock.Any(), gomock.Any()).
				Return(tt.respSVCMock.userID, tt.respSVCMock.err).
				AnyTimes()

			sessMock := NewMockSessionCreator(gomock.NewController(t))
			sessMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(tt.respSessMock.sessionID, tt.respSessMock.err).AnyTimes()

			sut := AuthenticateUser{
				Service:   svcMock,
				Session:   sessMock,
				Cookie:    cookie.NewCookie(cookieHashKey, cookieBlockKey),
				validator: validator.New(),
			}
			sut.ServeHTTP(w, r, nil)

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
