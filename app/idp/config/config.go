package config

import (
	"time"

	"github.com/42milez/go-oidc-server/app/pkg/typedef"

	"github.com/rs/zerolog"

	"github.com/caarlos0/env/v8"
)

const LoggerTagKey = "type"
const AppLoggerTagValue = "app"
const AccessLoggerTagValue = "access"

//  Environment Variable
// --------------------------------------------------

type Config struct {
	Port                int           `env:"PORT" envDefault:"80"`
	DB1Host             string        `env:"DB1_HOST" envDefault:"127.0.0.1"`
	DB1Port             int           `env:"DB1_PORT" envDefault:"3306"`
	DBAdmin             string        `env:"DB_USER" envDefault:"idp"`
	DBPassword          string        `env:"DB_PASSWORD" envDefault:"idp"`
	DBName              string        `env:"DB_NAME" envDefault:"idp"`
	RedisHost           string        `env:"REDIS_HOST" envDefault:"127.0.0.1"`
	RedisPort           int           `env:"REDIS_PORT" envDefault:"6379"`
	RedisPassword       string        `env:"REDIS_PASSWORD" envDefault:""`
	RedisDB             int           `env:"REDIS_DB" envDefault:"0"`
	LogLevel            zerolog.Level `env:"ZEROLOG_LEVEL" envDefault:"0"` // debug
	IdpHost             string        `env:"IDP_HOST" envDefault:"http://localhost:8080"`
	EnableDebugDBClient bool          `env:"ENABLE_DEBUG_DB_CLIENT" envDefault:"true"`
	EnableProfiler      bool          `env:"ENABLE_PROFILER" envDefault:"true"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

//  Application
// --------------------------------------------------

const (
	AppName = "idp"
)

//  Cache
// --------------------------------------------------

const (
	SessionTTL = 24 * time.Hour * 30 // 30 days
)

//  Cookie
// --------------------------------------------------

const (
	SessionIDCookieName = "sid"
	SessionIDCookieTTL  = SessionTTL
)

//  Identity Provider
// --------------------------------------------------

const (
	RegisterPath = "/user/register"
)

const (
	AuthenticationPath = "/authenticate"
	AuthorizationPath  = "/authorize"
	ConsentPath        = "/consent"
	TokenPath          = "/token"
)

const (
	Issuer = "42milez.dev"
)

const (
	AuthCodeLength     = 30
	AuthCodeTTL        = 10 * time.Minute
	AccessTokenLength  = 30
	RefreshTokenLength = 30
)

const (
	ClientIDLength     = 30
	ClientSecretLength = 30
)

const (
	AuthorizationCodeGrantType = "authorization_code"
	RefreshTokenGrantType      = "refresh_token"
)

const (
	BearerTokenType typedef.TokenType = "Bearer"
)

const (
	AccessTokenTTL  = 30 * time.Minute
	RefreshTokenTTL = SessionTTL
	IDTokenTTL      = SessionTTL
)
