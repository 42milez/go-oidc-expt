package handler

import (
	"bytes"
	"github.com/42milez/go-oidc-server/pkg/xerr"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/42milez/go-oidc-server/pkg/testutil"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
)

const (
	testDataSignInDir         = "testdata/signin/"
	testDataSignInRequest200  = testDataSignInDir + "200_req.json"
	testDataSignInResponse200 = testDataSignInDir + "200_resp.json"
	testDataSignInRequest400  = testDataSignInDir + "400_req.json"
	testDataSignInResponse400 = testDataSignInDir + "400_resp.json"
	testDataSignInResponse500 = testDataSignInDir + "500_resp.json"
)

const dummyAccessToken = "DUMMY ACCESS TOKEN"

func TestSignIn_ServeHTTP(t *testing.T) {
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
			reqFile: testDataSignInRequest200,
			resp: mockResp{
				token: dummyAccessToken,
				err:   nil,
			},
			want: want{
				statusCode: http.StatusOK,
				respFile:   testDataSignInResponse200,
			},
		},
		"BadRequest": {
			reqFile: testDataSignInRequest400,
			want: want{
				statusCode: http.StatusBadRequest,
				respFile:   testDataSignInResponse400,
			},
		},
		"InternalServerError": {
			reqFile: testDataSignInRequest200,
			resp: mockResp{
				token: "",
				err:   xerr.ErrInternalServerError,
			},
			want: want{
				statusCode: http.StatusInternalServerError,
				respFile:   testDataSignInResponse500,
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
				"/signin",
				bytes.NewReader(testutil.LoadFile(t, tt.reqFile)))

			svcMock := NewMockSignInService(gomock.NewController(t))
			svcMock.
				EXPECT().
				SignIn(r.Context(), gomock.Any(), gomock.Any()).
				Return(tt.resp.token, tt.resp.err).
				AnyTimes()

			sut := SignIn{
				Service:   svcMock,
				Validator: validator.New(),
			}
			sut.ServeHTTP(w, r)

			resp := w.Result()

			testutil.AssertResponse(t, resp, tt.want.statusCode, testutil.LoadFile(t, tt.want.respFile))
		})
	}
}
