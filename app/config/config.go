package config

import (
	"time"

	"github.com/rs/zerolog"

	"github.com/caarlos0/env/v8"
)

const LoggerTagKey = "log"
const AppLoggerTagValue = "app"
const MWLoggerTagValue = "mw"

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
	LogLevel      zerolog.Level `env:"ZEROLOG_LOG_LEVEL" envDefault:"0"` // debug
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

//  OIDC
// --------------------------------------------------

const (
	AuthenticationEndpoint = "/authenticate"
	ConsentEndpoint        = "/consent"
	AuthorizationEndpoint  = "/authorize"
)

const (
	AuthCodeLength   = 30
	AuthCodeLifetime = 10 * time.Minute
)

const (
	ClientIdLength     = 30
	ClientSecretLength = 30
)
