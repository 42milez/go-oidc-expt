package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/42milez/go-oidc-server/app/api/oapigen"
	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/pkg/xrandom"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestAuthorizationCodeFlow(t *testing.T) {
	t.Parallel()

	const nonceLength = 30
	const stateLength = 30

	const idpBaseUrl = "http://localhost:8081"
	const registerEndpoint = idpBaseUrl + config.RegisterPath
	const authenticationEndpoint = idpBaseUrl + config.AuthenticationPath
	const consentEndpoint = idpBaseUrl + config.ConsentPath

	const responseType = "code"
	const scope = "openid profile email"
	const clientId = "CDcp9v3Nn4i70FqWig5AuohmorD6MG"
	const redirectUri = "https://swagger.example.com/cb"

	nonce, err := xrandom.MakeCryptoRandomString(nonceLength)
	if err != nil {
		t.Fatal(err)
	}

	state, err := xrandom.MakeCryptoRandomString(stateLength)
	if err != nil {
		t.Fatal(err)
	}

	params := url.Values{}
	params.Add("client_id", clientId)
	params.Add("nonce", nonce)
	params.Add("redirect_uri", redirectUri)
	params.Add("response_type", responseType)
	params.Add("scope", scope)
	params.Add("state", state)
	params.Add("display", "page")
	params.Add("max_age", "86400")
	params.Add("prompt", "consent")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	username := fmt.Sprintf("user%d", rand.Uint64())
	password := "password"

	post := func(c *http.Client, u *url.URL, ck []*http.Cookie, data []byte) (*http.Response, error) {
		var payload io.Reader = nil

		if data != nil {
			payload = bytes.NewReader(data)
		}

		req, e := http.NewRequest("POST", u.String(), payload)
		if e != nil {
			return nil, e
		}

		req.Header.Add("Content-Type", "application/json")

		if ck != nil {
			for _, v := range ck {
				req.AddCookie(v)
			}
		}

		resp, e := c.Do(req)
		if e != nil {
			return nil, e
		}

		return resp, nil
	}

	closeRespBody := func(resp *http.Response) {
		if err = resp.Body.Close(); err != nil {
			t.Fatal(err)
		}
	}

	//  Registration
	// --------------------------------------------------

	regUrl, err := url.Parse(registerEndpoint)
	if err != nil {
		t.Fatal(err)
	}

	regReqBody := &oapigen.RegisterJSONRequestBody{
		Name:     username,
		Password: password,
	}

	regData, err := json.Marshal(regReqBody)
	if err != nil {
		t.Fatal(err)
	}

	regResp, err := post(client, regUrl, nil, regData)
	if err != nil {
		t.Fatal(err)
	}
	defer closeRespBody(regResp)

	if regResp.StatusCode != http.StatusOK {
		t.Fatalf("registration failed: want = %d; got = %d", http.StatusOK, regResp.StatusCode)
	}

	//  Authentication
	// --------------------------------------------------

	autheUrl, err := url.Parse(fmt.Sprintf("%s?%s", authenticationEndpoint, params.Encode()))
	if err != nil {
		t.Fatal(err)
	}

	autheReqBody := &oapigen.AuthenticateJSONRequestBody{
		Name:     username,
		Password: password,
	}

	autheData, err := json.Marshal(autheReqBody)
	if err != nil {
		t.Fatal(err)
	}

	autheResp, err := post(client, autheUrl, nil, autheData)
	defer closeRespBody(autheResp)
	if err != nil {
		t.Fatal(err)
	}

	if autheResp.StatusCode != http.StatusFound {
		t.Fatalf("authentication failed: want = %d; got = %d", http.StatusFound, autheResp.StatusCode)
	}

	cookies := autheResp.Cookies()

	if len(cookies) == 0 {
		t.Fatal("cookie not exist")
	}

	//  Consent
	// --------------------------------------------------

	consentUrl, err := url.Parse(fmt.Sprintf("%s?%s", consentEndpoint, params.Encode()))
	if err != nil {
		t.Fatal(err)
	}

	consentResp, err := post(client, consentUrl, cookies, nil)
	defer closeRespBody(consentResp)
	if err != nil {
		t.Fatal(err)
	}

	if consentResp.StatusCode != http.StatusFound {
		t.Fatalf("consent failed: want = %d; got = %d", http.StatusFound, consentResp.StatusCode)
	}
}
