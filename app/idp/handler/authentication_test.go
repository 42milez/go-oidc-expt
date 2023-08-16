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
	testDataAuthenticationDir         = "testdata/authentication/"
	testDataAuthenticationRequest200  = testDataAuthenticationDir + "200_req.json"
	testDataAuthenticationResponse200 = testDataAuthenticationDir + "200_resp.json"
	testDataAuthenticationRequest400  = testDataAuthenticationDir + "400_req.json"
	testDataAuthenticationResponse400 = testDataAuthenticationDir + "400_resp.json"
	testDataAuthenticationResponse500 = testDataAuthenticationDir + "500_resp.json"
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
			reqFile: testDataAuthenticationRequest200,
			resp: mockResp{
				token: dummyAccessToken,
				err:   nil,
			},
			want: want{
				statusCode: http.StatusOK,
				respFile:   testDataAuthenticationResponse200,
			},
		},
		"BadRequest": {
			reqFile: testDataAuthenticationRequest400,
			want: want{
				statusCode: http.StatusBadRequest,
				respFile:   testDataAuthenticationResponse400,
			},
		},
		"InternalServerError": {
			reqFile: testDataAuthenticationRequest200,
			resp: mockResp{
				token: "",
				err:   xtestutil.DummyError,
			},
			want: want{
				statusCode: http.StatusInternalServerError,
				respFile:   testDataAuthenticationResponse500,
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
