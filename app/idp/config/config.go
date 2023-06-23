package config

import "github.com/caarlos0/env/v8"

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
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
