package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"time"
)

type Config struct {
	Listener ListenerConfig `envPrefix:"LISTENER_"`
	Postgres PostgresConfig `envPrefix:"PG_"`
}

type ListenerConfig struct {
	Host string `env:"HOST" envDefault:"0.0.0.0"`
	Port int32  `env:"PORT" envDefault:"8081"`
}

type PostgresConfig struct {
	HostPort          string        `env:"HOST_PORT,required"`
	Username          string        `env:"USERNAME,required"`
	Password          string        `env:"PASSWORD,required"`
	DBName            string        `env:"DB_NAME" envDefault:"maestro"`
	DBMaxConn         int           `env:"DB_MAX_CONN" envDefault:"10"`
	DBMaxConnLifeTime time.Duration `env:"DB_MAX_CONN_LIFE_TIME" envDefault:"5m"`
	DBMaxConnIdleTime time.Duration `env:"DB_MAX_CONN_IDLE_TIME" envDefault:"1m"`
	DBTimeout         time.Duration `env:"DB_TIMEOUT" envDefault:"5s"`
}

func New() (*Config, error) {
	cfg := &Config{}

	if err := godotenv.Load("./.env"); err != nil {
		return nil, err
	}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
