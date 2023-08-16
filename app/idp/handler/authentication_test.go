package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/42milez/go-oidc-server/pkg/xtestutil"
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

const dummyAccessToken = "DUMMY ACCESS TOKEN"

func TestAuthentication_ServeHTTP(t *testing.T) {
	type mockResp struct {
		token string
		err   error
	}

	type want struct {
		statusCode int
		respFile   string
	}

	tests := map[string]struct {
		reqFile string
		resp    mockResp
		want    want
	}{
		"OK": {
			reqFile: tdAuthenticationRequest200,
			resp: mockResp{
				token: dummyAccessToken,
				err:   nil,
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
			resp: mockResp{
				token: "",
				err:   xtestutil.DummyError,
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
				bytes.NewReader(xtestutil.LoadFile(t, tt.reqFile)))

			svcMock := NewMockAuthenticator(gomock.NewController(t))
			svcMock.
				EXPECT().
				Authenticate(r.Context(), gomock.Any(), gomock.Any()).
				Return(tt.resp.token, tt.resp.err).
				AnyTimes()

			sut := Authenticate{
				Service:   svcMock,
				Validator: validator.New(),
			}
			sut.ServeHTTP(w, r)

			resp := w.Result()

			xtestutil.AssertResponse(t, resp, tt.want.statusCode, xtestutil.LoadFile(t, tt.want.respFile))
		})
	}
}
