package config

import "os"

type PostgresConnection struct {
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
}

type PostgresConfig struct {
	PostgresConnection PostgresConnection
}

func GetPostgresConnection() PostgresConfig {
	return PostgresConfig{
		PostgresConnection: PostgresConnection{
			DatabaseHost:     os.Getenv("DB_HOST"),
			DatabasePort:     os.Getenv("DB_PORT"),
			DatabaseName:     os.Getenv("DB_NAME"),
			DatabaseUser:     os.Getenv("DB_USER"),
			DatabasePassword: os.Getenv("DB_PASSWORD"),
		},
	}
}
