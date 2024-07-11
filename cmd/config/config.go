package config

import (
	"fmt"
	"time"

	"github.com/42milez/go-oidc-server/pkg/typedef"

	"github.com/rs/zerolog"

	"github.com/caarlos0/env/v8"
)

const LoggerTagKey = "type"
const AppLoggerTagValue = "app"
const AccessLoggerTagValue = "access"

type Config struct {
	Port                int           `env:"PORT" envDefault:"80"`
	DBHost              string        `env:"DB_HOST" envDefault:"127.0.0.1"`
	DBPort              int           `env:"DB_PORT" envDefault:"3306"`
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

const (
	AppName = "idp"
)

const (
	SessionIDCookieName = "sid"
	SessionTTL          = 24 * time.Hour * 30 // 30 days
	SessionIDCookieTTL  = SessionTTL
)

const OIDCRootPath = "connect"
const UserRootPath = "user"

func AuthorizationPath() string {
	return fmt.Sprintf("/%s", "authorization")
}

func TokenPath() string {
	return fmt.Sprintf("/%s", "token")
}

func UserInfoPath() string {
	return fmt.Sprintf("/%s", "userinfo")
}

func UserRegistrationPath() string {
	return fmt.Sprintf("/%s/%s", UserRootPath, "registration")
}

func UserAuthenticationPath() string {
	return fmt.Sprintf("/%s/%s", UserRootPath, "authentication")
}

func UserConsentPath() string {
	return fmt.Sprintf("/%s/%s", UserRootPath, "consent")
}

const (
	Issuer = "https://localhost:4443"
)

const (
	AuthCodeTTL     = 10 * time.Minute
	AccessTokenTTL  = 30 * time.Minute
	RefreshTokenTTL = SessionTTL
	IDTokenTTL      = SessionTTL
)

const (
	AuthCodeLength     = 30
	ClientIDLength     = 30
	ClientSecretLength = 30
)

const (
	BearerTokenType typedef.TokenType = "Bearer"
)

const (
	AuthorizationCodeGrantType typedef.GrantType = "authorization_code"
	RefreshTokenGrantType      typedef.GrantType = "refresh_token"
)
