package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/42milez/go-oidc-server/app/pkg/typedef"

	"github.com/42milez/go-oidc-server/app/pkg/ent/ent"
	"github.com/42milez/go-oidc-server/app/pkg/ent/ent/relyingparty"
	"github.com/42milez/go-oidc-server/app/pkg/ent/ent/user"

	"github.com/42milez/go-oidc-server/app/idp/entity"

	"github.com/42milez/go-oidc-server/app/idp/api"
	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/google/go-querystring/query"

	"github.com/42milez/go-oidc-server/app/pkg/xrandom"
	"github.com/42milez/go-oidc-server/app/pkg/xtestutil"
)

func TestAuthorizationCodeFlow(t *testing.T) {
	t.Parallel()

	const baseUrl = "http://127.0.0.1:8081"
	userRegistrationEndpoint := baseUrl + config.UserRegistrationPath()
	userAuthenticationEndpoint := baseUrl + config.UserAuthenticationPath()
	userConsentEndpoint := baseUrl + config.UserConsentPath()
	authorizationEndpoint := baseUrl + config.AuthorizationPath()
	tokenEndpoint := baseUrl + config.TokenPath()

	const responseType = "code"
	const scope = "openid profile email"
	const redirectUri = "https://swagger.example.com/cb"
	const nonceLength = 30
	const stateLength = 30

	ctx := context.Background()
	db := xtestutil.NewDatabase(t, nil)
	httpClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	newAuthorizeParam := func(clientID typedef.ClientID) string {
		nonce, err := xrandom.GenerateCryptoRandomString(nonceLength)
		xtestutil.ExitOnError(t, err)

		state, err := xrandom.GenerateCryptoRandomString(stateLength)
		xtestutil.ExitOnError(t, err)

		authoParam := &api.AuthorizeParams{
			ClientID:     clientID,
			Display:      "page",
			MaxAge:       86400,
			Nonce:        nonce,
			Prompt:       "consent",
			RedirectURI:  redirectUri,
			ResponseType: responseType,
			Scope:        scope,
			State:        state,
		}

		v, err := query.Values(authoParam)
		xtestutil.ExitOnError(t, err)

		v.Del("sid")

		return v.Encode()
	}

	registerRelyingParty := func() *entity.RelyingParty {
		clientID, err := xrandom.GenerateCryptoRandomString(config.ClientIDLength)
		xtestutil.ExitOnError(t, err)

		clientSecret, err := xrandom.GenerateCryptoRandomString(config.ClientIDLength)
		xtestutil.ExitOnError(t, err)

		rp, err := db.Client.RelyingParty.Create().SetClientID(typedef.ClientID(clientID)).SetClientSecret(clientSecret).Save(ctx)
		xtestutil.ExitOnError(t, err)

		t.Cleanup(func() {
			_, err = db.Client.RelyingParty.Delete().Where(relyingparty.ID(rp.ID)).Exec(ctx)
			xtestutil.ExitOnError(t, err)
		})

		_, err = db.Client.RedirectURI.Create().SetURI(redirectUri).SetRelyingParty(rp).Save(ctx)
		xtestutil.ExitOnError(t, err)

		return entity.NewRelyingParty(rp)
	}

	registerUser := func() *entity.User {
		regUrl, err := url.Parse(userRegistrationEndpoint)
		xtestutil.ExitOnError(t, err)

		rand.New(rand.NewSource(time.Now().UnixNano()))
		username := fmt.Sprintf("user%d", rand.Uint64())
		password := "password"

		regReqBody := &api.RegisterJSONRequestBody{
			Name:     username,
			Password: password,
		}

		regData, err := json.Marshal(regReqBody)
		xtestutil.ExitOnError(t, err)

		regResp, err := xtestutil.Request(t, httpClient, http.MethodPost, regUrl, nil, regData)
		defer xtestutil.CloseResponseBody(t, regResp)
		xtestutil.ExitOnError(t, err)

		if regResp.StatusCode != http.StatusOK {
			t.Fatalf("POST /register failed: want = %d; got = %d", http.StatusOK, regResp.StatusCode)
		}

		var regRespBody []byte
		var u api.User

		regRespBody, err = io.ReadAll(regResp.Body)
		xtestutil.ExitOnError(t, err)

		err = json.Unmarshal(regRespBody, &u)
		xtestutil.ExitOnError(t, err)

		t.Cleanup(func() {
			_, err = db.Client.User.Delete().Where(user.ID(u.ID)).Exec(ctx)
			xtestutil.ExitOnError(t, err)
		})

		return entity.NewUser(&ent.User{
			Name:     username,
			Password: password,
		})
	}

	authenticate := func(u *entity.User, authoParam string) []*http.Cookie {
		autheUrl, err := url.Parse(fmt.Sprintf("%s?%s", userAuthenticationEndpoint, authoParam))
		xtestutil.ExitOnError(t, err)

		autheReqBody := &api.AuthenticateJSONRequestBody{
			Name:     u.Name(),
			Password: u.Password(),
		}

		autheData, err := json.Marshal(autheReqBody)
		xtestutil.ExitOnError(t, err)

		autheResp, err := xtestutil.Request(t, httpClient, http.MethodPost, autheUrl, nil, autheData)
		defer xtestutil.CloseResponseBody(t, autheResp)
		xtestutil.ExitOnError(t, err)

		if autheResp.StatusCode != http.StatusFound {
			t.Fatalf("POST /authentication failed: want = %d; got = %d", http.StatusFound, autheResp.StatusCode)
		}

		cookies := autheResp.Cookies()
		if len(cookies) == 0 {
			t.Fatal("cookie not exist")
		}

		return cookies
	}

	consent := func(cookies []*http.Cookie, authoParam string) {
		consentUrl, err := url.Parse(fmt.Sprintf("%s?%s", userConsentEndpoint, authoParam))
		xtestutil.ExitOnError(t, err)

		consentReqParam := &xtestutil.RequestParam{
			Cookies: cookies,
		}

		consentResp, err := xtestutil.Request(t, httpClient, http.MethodPost, consentUrl, consentReqParam, nil)
		defer xtestutil.CloseResponseBody(t, consentResp)
		xtestutil.ExitOnError(t, err)

		if consentResp.StatusCode != http.StatusFound {
			t.Fatalf("POST /consent failed: want = %d; got = %d", http.StatusFound, consentResp.StatusCode)
		}
	}

	authorize := func(cookies []*http.Cookie, authoParam string) *url.URL {
		authoUrl, err := url.Parse(fmt.Sprintf("%s?%s", authorizationEndpoint, authoParam))
		xtestutil.ExitOnError(t, err)

		authoReqParam := &xtestutil.RequestParam{
			Cookies: cookies,
		}

		authoResp, err := xtestutil.Request(t, httpClient, http.MethodGet, authoUrl, authoReqParam, nil)
		defer xtestutil.CloseResponseBody(t, authoResp)
		xtestutil.ExitOnError(t, err)

		if authoResp.StatusCode != http.StatusFound {
			t.Fatalf("GET /authorization failed: want = %d; got = %d", http.StatusFound, authoResp.StatusCode)
		}

		cbUrl, err := authoResp.Location()
		xtestutil.ExitOnError(t, err)

		return cbUrl
	}

	initialRequestToken := func(rp *entity.RelyingParty, cookies []*http.Cookie, authoParam string, cbUrl *url.URL) string {
		tokenUrl, err := url.Parse(fmt.Sprintf("%s?%s", tokenEndpoint, authoParam))
		xtestutil.ExitOnError(t, err)

		credential := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", rp.ClientID(), rp.ClientSecret())))

		reqParam := &xtestutil.RequestParam{
			Headers: map[string]string{
				"Authorization": fmt.Sprintf("Basic %s", credential),
				"Content-Type":  "application/x-www-form-urlencoded",
			},
			Cookies: cookies,
		}

		cbQuery := cbUrl.Query()

		// TODO: Handle error if the query contains error parameter
		// ...

		param := url.Values{}
		param.Add("grant_type", "authorization_code")
		param.Add("code", cbQuery.Get("code"))
		param.Add("redirect_uri", redirectUri)
		reqBody := []byte(param.Encode())

		resp, err := xtestutil.Request(t, httpClient, http.MethodPost, tokenUrl, reqParam, reqBody)
		defer xtestutil.CloseResponseBody(t, resp)
		xtestutil.ExitOnError(t, err)

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("POST /token failed: want = %d; got = %d", http.StatusOK, resp.StatusCode)
		}

		respBody, err := io.ReadAll(resp.Body)
		xtestutil.ExitOnError(t, err)

		var tokenResp api.TokenResponse

		err = json.Unmarshal(respBody, &tokenResp)
		xtestutil.ExitOnError(t, err)

		return tokenResp.RefreshToken
	}

	requestToken := func(rp *entity.RelyingParty, cookies []*http.Cookie, refreshToken string) {
		tokenUrl, err := url.Parse(fmt.Sprintf("%s", tokenEndpoint))
		xtestutil.ExitOnError(t, err)

		credential := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", rp.ClientID(), rp.ClientSecret())))

		reqParam := &xtestutil.RequestParam{
			Headers: map[string]string{
				"Authorization": fmt.Sprintf("Basic %s", credential),
				"Content-Type":  "application/x-www-form-urlencoded",
			},
			Cookies: cookies,
		}

		param := url.Values{}
		param.Add("client_id", rp.ClientID().String())
		param.Add("client_secret", rp.ClientSecret())
		param.Add("grant_type", "refresh_token")
		param.Add("refresh_token", refreshToken)
		param.Add("scope", "openid profile")
		reqBody := []byte(param.Encode())

		resp, err := xtestutil.Request(t, httpClient, http.MethodPost, tokenUrl, reqParam, reqBody)
		defer xtestutil.CloseResponseBody(t, resp)
		xtestutil.ExitOnError(t, err)

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("POST /token failed: want = %d; got = %d", http.StatusOK, resp.StatusCode)
		}
	}

	registeredRp := registerRelyingParty()
	registeredUser := registerUser()
	authoParam := newAuthorizeParam(registeredRp.ClientID())
	cookies := authenticate(registeredUser, authoParam)
	consent(cookies, authoParam)
	callbackUrl := authorize(cookies, authoParam)
	refreshToken := initialRequestToken(registeredRp, cookies, authoParam, callbackUrl)
	requestToken(registeredRp, cookies, refreshToken)
}
