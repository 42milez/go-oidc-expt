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
		AuthorizationEndpoint: "https://host.docker.internal:4443/api/connect/authorization",
		DisplayValuesSupported: []DisplayValuesSupported{
			Page,
		},
		IDTokenSigningAlgValuesSupported: []IDTokenSigningAlgValuesSupported{
			ES256,
		},
		Issuer:  "https://host.docker.internal:4443/api/connect/",
		JWKsURI: "https://host.docker.internal:4443/api/connect/jwks",
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
		TokenEndpoint: "https://host.docker.internal:4443/api/connect/token",
		TokenEndpointAuthMethodsSupported: []TokenEndpointAuthMethodsSupported{
			ClientSecretBasic,
		},
		UILocalesSupported: []UILocalesSupported{
			JaJP,
		},
		UserInfoEndpoint: "https://host.docker.internal:4443/api/connect/userinfo",
	}

	RespondJSON(w, r, http.StatusOK, nil, respBody)
}
