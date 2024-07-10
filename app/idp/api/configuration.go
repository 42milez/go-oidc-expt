package api

import "net/http"

var configuration *Configuration

func InitConfiguration() {
	if configuration == nil {
		configuration = &Configuration{}
	}
}

type Configuration struct{}

func (c *Configuration) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	respBody := &ConfigurationResponse{
		AuthorizationEndpoint: "https://localhost:4443/connect/authorization",
		DisplayValuesSupported: []DisplayValuesSupported{
			Page,
		},
		IDTokenSigningAlgValuesSupported: []IDTokenSigningAlgValuesSupported{
			ES256,
		},
		Issuer:  "https://localhost:4443",
		JWKsURI: "https://host.docker.internal:4443/connect/jwks",
		ResponseTypesSupported: []ResponseTypesSupported{
			ResponseTypesSupportedCode,
		},
		ScopesSupported: []ScopesSupported{
			Openid,
			Profile,
			Email,
		},
		SubjectTypesSupported: []SubjectTypesSupported{
			Public,
		},
		TokenEndpoint: "https://host.docker.internal:4443/connect/token",
		TokenEndpointAuthMethodsSupported: []TokenEndpointAuthMethodsSupported{
			ClientSecretBasic,
		},
		UILocalesSupported: []UILocalesSupported{
			JaJP,
		},
		UserInfoEndpoint: "https://host.docker.internal:4443/connect/userinfo",
	}

	RespondJSON(w, r, http.StatusOK, nil, respBody)
}
