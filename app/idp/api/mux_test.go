package api

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/42milez/go-oidc-server/app/pkg/ent/ent/relyingparty"

	"github.com/42milez/go-oidc-server/app/idp/entity"
	"github.com/42milez/go-oidc-server/app/idp/option"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xtestutil"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"
)

func TestNewOapiErrorHandler(t *testing.T) {
	t.Parallel()

	var swag *openapi3.T
	var err error

	if swag, err = GetSwagger(); err != nil {
		t.Fatal(err)
	}
	swag.Servers = nil

	opt := &option.Option{
		DB: xtestutil.NewDatabase(t, nil),
	}

	ctx := context.Background()

	createRelyingParty := func() *entity.RelyingParty {
		clientID := "CDcp9v3Nn4i70FqWig5AuohmorD6MG"
		clientSecret := "whc5nzVjt7AQpTrAhUqVaGgV2PK4oo"

		rp, err := opt.DB.Client.RelyingParty.Create().SetClientID(clientID).SetClientSecret(clientSecret).Save(ctx)
		xtestutil.ExitOnError(t, err)

		t.Cleanup(func() {
			_, err = opt.DB.Client.RelyingParty.Delete().Where(relyingparty.ID(rp.ID)).Exec(ctx)
			xtestutil.ExitOnError(t, err)
		})

		return entity.NewRelyingParty(rp)
	}
	rp := createRelyingParty()

	mux := chi.NewRouter()
	mux.Use(nethttpmiddleware.OapiRequestValidatorWithOptions(swag, &nethttpmiddleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: NewOapiAuthentication(opt),
		},
		ErrorHandler: NewOapiErrorHandler(),
	}))
	mux.Post("/token", func(w http.ResponseWriter, r *http.Request) {
		RespondJSON200(w, r)
	})

	tests := map[string]struct {
		ClientID   string
		StatusCode int
		Error      *xerr.PublicError
	}{
		"OK": {
			ClientID: "CDcp9v3Nn4i70FqWig5AuohmorD6MG",
			// TODO: Compare response with JSON
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			createHTTPRequest := func() *http.Request {
				param := url.Values{}
				param.Add("grant_type", "authorization_code")
				param.Add("code", "EYdxIU30xstnWZKxgA54RJMz1YUR0J")
				param.Add("redirect_uri", "https://rp.example.com/cb")
				reqBody := bytes.NewReader([]byte(param.Encode()))
				req := httptest.NewRequest(http.MethodPost, "/token", reqBody)
				credential := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", tt.ClientID, rp.ClientSecret())))
				req.Header.Set("Authorization", fmt.Sprintf("Basic %s", credential))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				return req
			}
			req := createHTTPRequest()
			respRec := httptest.NewRecorder()

			mux.ServeHTTP(respRec, req)

			if respRec.Code != http.StatusOK {
				t.Fatalf("want = %d; got = %d", http.StatusOK, respRec.Code)
			}
		})
	}
}
