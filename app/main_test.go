package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/42milez/go-oidc-server/app/api/oapigen"
	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/pkg/xrandom"
	"github.com/42milez/go-oidc-server/app/pkg/xtestutil"
)

func TestAuthorizationCodeFlow(t *testing.T) {
	t.Parallel()

	const nonceLength = 30
	const stateLength = 30

	const idpBaseUrl = "http://localhost:8081"
	const registerEndpoint = idpBaseUrl + config.RegisterPath
	const authenticationEndpoint = idpBaseUrl + config.AuthenticationPath
	const consentEndpoint = idpBaseUrl + config.ConsentPath
	const authorizationEndpoint = idpBaseUrl + config.AuthorizationPath
	const tokenEndpoint = idpBaseUrl + config.TokenPath

	const responseType = "code"
	const scope = "openid profile email"
	const redirectUri = "https://swagger.example.com/cb"

	clientId, err := xrandom.MakeCryptoRandomStringNoCache(config.ClientIdLength)
	xtestutil.ExitOnError(t, err)

	clientSecret, err := xrandom.MakeCryptoRandomStringNoCache(config.ClientIdLength)
	xtestutil.ExitOnError(t, err)

	ctx := context.Background()

	cfg, err := config.New()
	xtestutil.ExitOnError(t, err)

	cfg.DBAdmin = xtestutil.TestDBUser
	cfg.DBPassword = xtestutil.TestDBPassword
	cfg.DBHost = xtestutil.TestDBHost
	cfg.DBPort = xtestutil.TestDBPort
	cfg.DBName = "idp_integ_test"
	cfg.Debug = false

	db := xtestutil.NewDatabase(t, cfg)

	rp, err := db.Client.RelyingParty.Create().SetClientID(clientId).SetClientSecret(clientSecret).Save(ctx)
	xtestutil.ExitOnError(t, err)

	_, err = db.Client.RedirectURI.Create().SetURI(redirectUri).SetRelyingParty(rp).Save(ctx)
	xtestutil.ExitOnError(t, err)

	nonce, err := xrandom.MakeCryptoRandomString(nonceLength)
	xtestutil.ExitOnError(t, err)

	state, err := xrandom.MakeCryptoRandomString(stateLength)
	xtestutil.ExitOnError(t, err)

	authoParam := url.Values{}
	authoParam.Add("client_id", clientId)
	authoParam.Add("nonce", nonce)
	authoParam.Add("redirect_uri", redirectUri)
	authoParam.Add("response_type", responseType)
	authoParam.Add("scope", scope)
	authoParam.Add("state", state)
	authoParam.Add("display", "page")
	authoParam.Add("max_age", "86400")
	authoParam.Add("prompt", "consent")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	username := fmt.Sprintf("user%d", rand.Uint64())
	password := "password"

	//  Registration
	// --------------------------------------------------

	regUrl, err := url.Parse(registerEndpoint)
	xtestutil.ExitOnError(t, err)

	regReqBody := &oapigen.RegisterJSONRequestBody{
		Name:     username,
		Password: password,
	}

	regData, err := json.Marshal(regReqBody)
	xtestutil.ExitOnError(t, err)

	regResp, err := xtestutil.Request(t, client, http.MethodPost, regUrl, nil, regData)
	defer xtestutil.CloseResponseBody(t, regResp)
	xtestutil.ExitOnError(t, err)

	if regResp.StatusCode != http.StatusOK {
		t.Fatalf("POST /register failed: want = %d; got = %d", http.StatusOK, regResp.StatusCode)
	}

	//  Authentication
	// --------------------------------------------------

	autheUrl, err := url.Parse(fmt.Sprintf("%s?%s", authenticationEndpoint, authoParam.Encode()))
	xtestutil.ExitOnError(t, err)

	autheReqBody := &oapigen.AuthenticateJSONRequestBody{
		Name:     username,
		Password: password,
	}

	autheData, err := json.Marshal(autheReqBody)
	xtestutil.ExitOnError(t, err)

	autheResp, err := xtestutil.Request(t, client, http.MethodPost, autheUrl, nil, autheData)
	defer xtestutil.CloseResponseBody(t, autheResp)
	xtestutil.ExitOnError(t, err)

	if autheResp.StatusCode != http.StatusFound {
		t.Fatalf("POST /authenticate failed: want = %d; got = %d", http.StatusFound, autheResp.StatusCode)
	}

	cookies := autheResp.Cookies()
	if len(cookies) == 0 {
		t.Fatal("cookie not exist")
	}

	//  Consent
	// --------------------------------------------------

	consentUrl, err := url.Parse(fmt.Sprintf("%s?%s", consentEndpoint, authoParam.Encode()))
	xtestutil.ExitOnError(t, err)

	consentReqParam := &xtestutil.RequestParam{
		Cookies: cookies,
	}

	consentResp, err := xtestutil.Request(t, client, http.MethodPost, consentUrl, consentReqParam, nil)
	defer xtestutil.CloseResponseBody(t, consentResp)
	xtestutil.ExitOnError(t, err)

	if consentResp.StatusCode != http.StatusFound {
		t.Fatalf("POST /consent failed: want = %d; got = %d", http.StatusFound, consentResp.StatusCode)
	}

	//  Authorization
	// --------------------------------------------------

	authoUrl, err := url.Parse(fmt.Sprintf("%s?%s", authorizationEndpoint, authoParam.Encode()))
	xtestutil.ExitOnError(t, err)

	authoReqParam := &xtestutil.RequestParam{
		Cookies: cookies,
	}

	authoResp, err := xtestutil.Request(t, client, http.MethodGet, authoUrl, authoReqParam, nil)
	defer xtestutil.CloseResponseBody(t, authoResp)
	xtestutil.ExitOnError(t, err)

	if authoResp.StatusCode != http.StatusFound {
		t.Fatalf("GET /authorize failed: want = %d; got = %d", http.StatusFound, authoResp.StatusCode)
	}

	//  Token
	// --------------------------------------------------

	tokenUrl, err := url.Parse(fmt.Sprintf("%s?%s", tokenEndpoint, authoParam.Encode()))
	xtestutil.ExitOnError(t, err)

	credential := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientId, clientSecret)))

	tokenReqParam := &xtestutil.RequestParam{
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Basic %s", credential),
			"Content-Type":  "application/x-www-form-urlencoded",
		},
		Cookies: cookies,
	}

	cbUrl, err := authoResp.Location()
	xtestutil.ExitOnError(t, err)

	cbQuery := cbUrl.Query()

	tokenParam := url.Values{}
	tokenParam.Add("grant_type", "authorization_code")
	tokenParam.Add("code", cbQuery.Get("code"))
	tokenParam.Add("redirect_uri", redirectUri)
	tokenReqBody := []byte(tokenParam.Encode())

	tokenResp, err := xtestutil.Request(t, client, http.MethodPost, tokenUrl, tokenReqParam, tokenReqBody)
	defer xtestutil.CloseResponseBody(t, tokenResp)
	xtestutil.ExitOnError(t, err)

	if tokenResp.StatusCode != http.StatusOK {
		t.Fatalf("POST /token failed: want = %d; got = %d", http.StatusOK, tokenResp.StatusCode)
	}
}
