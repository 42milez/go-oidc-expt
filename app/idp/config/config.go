package config

import (
	"github.com/42milez/go-oidc-server/pkg/xerr"
	"github.com/caarlos0/env/v8"
)

const (
	errFailedToParseEnvVal xerr.Err = "failed to parse environment variable"
)

type Config struct {
	Env           string `env:"ENV" envDefault:"dev"`
	Port          int    `env:"PORT" envDefault:"80"`
	DBHost        string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DBPort        int    `env:"DB_PORT" envDefault:"3306"`
	DBAdmin       string `env:"DB_USER" envDefault:"idp"`
	DBPassword    string `env:"DB_PASSWORD" envDefault:"idp"`
	DBName        string `env:"DB_NAME" envDefault:"idp"`
	RedisHost     string `env:"REDIS_HOST" envDefault:"127.0.0.1"`
	RedisPort     int    `env:"REDIS_PORT" envDefault:"6379"`
	RedisPassword string `env:"REDIS_PASSWORD" envDefault:""`
	RedisDB       int    `env:"REDIS_DB" envDefault:"0"`
	SwaggerHost   string `env:"SWAGGER_HOST" envDefault:"localhost"`
	SwaggerPort   int    `env:"SWAGGER_PORT" envDefault:"8080"`
	SwaggerPath   string `env:"SWAGGER_PATH" envDefault:"swagger"`
}

func (p *Config) IsDevelopment() bool {
	return p.Env == "dev"
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, xerr.Wrap(errFailedToParseEnvVal, err)
	}
	return cfg, nil
}
