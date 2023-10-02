package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/42milez/go-oidc-server/app/api"
	"github.com/42milez/go-oidc-server/app/api/oapigen"
	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/pkg/xrandom"
)

const nonceLength = 30
const stateLength = 30

const host = "http://localhost:8080"
const registerEndpoint = host + config.RegisterPath
const authenticationEndpoint = host + config.AuthenticationPath

//const authorizationEndpoint = host + config.AuthorizationPath

const responseType = "code"
const scope = "openid profile email"
const clientId = "CDcp9v3Nn4i70FqWig5AuohmorD6MG"
const redirectUri = "https://swagger.example.com/cb"

func main() {
	nonce, err := xrandom.MakeCryptoRandomString(nonceLength)
	if err != nil {
		log.Fatal(err)
	}

	state, err := xrandom.MakeCryptoRandomString(stateLength)
	if err != nil {
		log.Fatal(err)
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

	registerUrl, err := url.Parse(registerEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	err = register(client, registerUrl, &oapigen.RegisterJSONRequestBody{
		Name:     username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}

	authenticationUrl, err := url.Parse(fmt.Sprintf("%s?%s", authenticationEndpoint, params.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	consentUrl, err := authenticate(client, authenticationUrl, &oapigen.AuthenticateJSONRequestBody{
		Name:     username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Print(consentUrl)
}

func post(c *http.Client, u *url.URL, data []byte) (*http.Response, error) {
	log.Printf("RequestTo: %s\n", u.String())
	resp, err := c.Post(u.String(), "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func register(c *http.Client, u *url.URL, data *oapigen.RegisterJSONRequestBody) error {
	reqBody, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := post(c, u, reqBody)
	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Print(err)
		}
	}()
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		var respBody api.Response
		if err = json.NewDecoder(resp.Body).Decode(&respBody); (err != nil) && (err != io.EOF) {
			return err
		}
		return respBody.Summary
	}

	var respBody oapigen.User
	if err = json.NewDecoder(resp.Body).Decode(&respBody); (err != nil) && (err != io.EOF) {
		return err
	}
	log.Print(respBody)

	return nil
}

func authenticate(c *http.Client, u *url.URL, data *oapigen.AuthenticateJSONRequestBody) (*url.URL, error) {
	reqBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := post(c, u, reqBody)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusFound {
		var respBody api.Response
		if err = json.NewDecoder(resp.Body).Decode(&respBody); (err != nil) && (err != io.EOF) {
			return nil, err
		}
		return nil, respBody.Summary
	}

	l, err := resp.Location()
	if err != nil {
		return nil, err
	}

	return l, nil
}
