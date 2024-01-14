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
		AuthorizationEndpoint: "https://auth.42milez.com/connect/authorization",
		DisplayValuesSupported: []DisplayValuesSupported{
			Page,
		},
		IDTokenSigningAlgValuesSupported: []IDTokenSigningAlgValuesSupported{
			ES256,
		},
		Issuer:  "https://auth.42milez.com/connect/",
		JWKsURI: "https://auth.42milez.com/connect/jwks",
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
		TokenEndpoint: "https://auth.42milez.com/connect/token",
		TokenEndpointAuthMethodsSupported: []TokenEndpointAuthMethodsSupported{
			ClientSecretBasic,
		},
		UILocalesSupported: []UILocalesSupported{
			JaJP,
		},
		UserInfoEndpoint: "https://auth.42milez.com/connect/userinfo",
	}

	RespondJSON(w, r, http.StatusOK, nil, respBody)
}
