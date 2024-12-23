package envs

import (
	"os"
	"test-task/internal/config"
)

const dsn = "PG_DSN"

type postgresConfig struct {
	dsn string
}

func (p *postgresConfig) DSN() string {
	ds := os.Getenv(dsn)
	return ds
}

func NewPostgresConfig() config.PostgresConfig {
	return &postgresConfig{
		dsn: os.Getenv(dsn),
	}
}
