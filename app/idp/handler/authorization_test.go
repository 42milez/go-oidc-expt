package handler

//import (
//	"github.com/42milez/go-oidc-server/pkg/xtestutil"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//const (
//	tdAuthorizationDir = "testdata/authorization/"
//	tdAuthorizationRequest200 = tdAuthorizationDir + "200_req.txt"
//	tdAuthorizationResponse200 = tdAuthorizationDir + "200_resp.txt"
//)
//
//const dummyLocation = "TBD"
//
//func TestAuthorizeGet_ServeHTTP(t *testing.T) {
//	type mockResp struct {
//		location string
//		err      error
//	}
//
//	type want struct {
//		statusCode int
//		respFile   string
//	}
//
//	tests := map[string]struct {
//		reqFile string
//		resp    mockResp
//		want    want
//	}{
//		"OK": {
//			reqFile: tdAuthorizationRequest200,
//			resp: mockResp{
//				location: dummyLocation,
//				err: nil,
//			},
//			want: want{
//				statusCode: http.StatusOK,
//				respFile: tdAuthorizationResponse200,
//			},
//		},
//	}
//
//	for n, tt := range tests {
//		tt := tt
//
//		t.Run(n, func(t *testing.T) {
//			t.Parallel()
//
//			w := httptest.NewRecorder()
//
//			q := xtestutil.LoadFile(t, tt.reqFile)
//			r := httptest.NewRequest(
//				http.MethodPost,
//				"/authorize" + ,
//				nil,
//			)
//		})
//	}
//}
