package config

import (
	"time"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/rs/zerolog"

	"github.com/caarlos0/env/v8"
)

const LoggerTagKey = "type"
const AppLoggerTagValue = "app"
const AccessLoggerTagValue = "access"

//  ENVIRONMENT VARIABLE
// --------------------------------------------------

type Config struct {
	Env           string        `env:"ENV" envDefault:"dev"`
	Port          int           `env:"PORT" envDefault:"80"`
	DBHost        string        `env:"DB_HOST" envDefault:"127.0.0.1"`
	DBPort        int           `env:"DB_PORT" envDefault:"3306"`
	DBAdmin       string        `env:"DB_USER" envDefault:"idp"`
	DBPassword    string        `env:"DB_PASSWORD" envDefault:"idp"`
	DBName        string        `env:"DB_NAME" envDefault:"idp"`
	RedisHost     string        `env:"REDIS_HOST" envDefault:"127.0.0.1"`
	RedisPort     int           `env:"REDIS_PORT" envDefault:"6379"`
	RedisPassword string        `env:"REDIS_PASSWORD" envDefault:""`
	RedisDB       int           `env:"REDIS_DB" envDefault:"0"`
	LogLevel      zerolog.Level `env:"ZEROLOG_LEVEL" envDefault:"0"` // debug
	IdpHost       string        `env:"IDP_HOST" envDefault:"http://localhost:8080"`
	Debug         bool          `env:"DEBUG" envDefault:"true"`
}

func (p *Config) IsDevelopment() bool {
	return p.Env == "dev"
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

//  APPLICATION
// --------------------------------------------------

const (
	AppName = "idp"
)

//  COOKIE
// --------------------------------------------------

const (
	SessionIDCookieName = "sid"
	SessionIDCookieTTL  = SessionTTL
)

//  SESSION
// --------------------------------------------------

const (
	SessionTTL = 24 * time.Hour * 30 // 30 days
)

//  User
// --------------------------------------------------

const (
	RegisterPath = "/user/register"
)

//  OIDC
// --------------------------------------------------

const (
	AuthenticationPath = "/authenticate"
	AuthorizationPath  = "/authorize"
	ConsentPath        = "/consent"
	TokenPath          = "/token"
)

const (
	AuthCodeLength   = 30
	AuthCodeLifetime = 10 * time.Minute
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
