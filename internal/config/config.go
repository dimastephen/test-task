package config

import "github.com/joho/godotenv"

func Load(configPath string) error {
	err := godotenv.Load(configPath)
	if err != nil {
		return err
	}
	return nil
}

type HTTPConfig interface {
	Address() string
}

type PostgresConfig interface {
	DSN() string
}
