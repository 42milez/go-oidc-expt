package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/42milez/go-oidc-server/app/testutil"

	"github.com/42milez/go-oidc-server/app/ent/typedef"

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

const dummyUserID typedef.UserID = "user01"

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
		respFile   string
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
				statusCode: http.StatusOK,
				respFile:   tdAuthenticationResponse200,
			},
		},
		"BadRequest": {
			reqFile: tdAuthenticationRequest400,
			want: want{
				statusCode: http.StatusBadRequest,
				respFile:   tdAuthenticationResponse400,
			},
		},
		"InternalServerError": {
			reqFile: tdAuthenticationRequest200,
			respSVCMock: serviceMockResp{
				userID: "",
				err:    testutil.DummyError,
			},
			want: want{
				statusCode: http.StatusInternalServerError,
				respFile:   tdAuthenticationResponse500,
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
				bytes.NewReader(testutil.LoadFile(t, tt.reqFile)))

			svcMock := NewMockAuthenticator(gomock.NewController(t))
			svcMock.
				EXPECT().
				Authenticate(r.Context(), gomock.Any(), gomock.Any()).
				Return(tt.respSVCMock.userID, tt.respSVCMock.err).
				AnyTimes()

			sessMock := NewMockSessionCreator(gomock.NewController(t))
			sessMock.EXPECT().Create(gomock.Any()).Return(tt.respSessMock.sessionID, tt.respSessMock.err).AnyTimes()

			sut := Authenticate{
				service:   svcMock,
				session:   sessMock,
				cookie:    NewCookie(cookieHashKey, cookieBlockKey),
				validator: validator.New(),
			}
			sut.ServeHTTP(w, r)

			resp := w.Result()

			wantResp := &testutil.Response{
				StatusCode: tt.want.statusCode,
				Body:       testutil.LoadFile(t, tt.want.respFile),
			}

			testutil.AssertResponseJSON(t, resp, wantResp)
		})
	}
}