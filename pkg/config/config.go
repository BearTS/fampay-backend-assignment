// Package to keep track and handle configurations
package config

import "github.com/ilyakaznacheev/cleanenv"

type ConfigDatabase struct {
	Environment string `env:"ENVIRONMENT" env-default:"development"`
	Port        int    `env:"PORT" env-default:"3000"`

	// Youtube Related Configs
	YoutubeApiKeys       string `env:"YOUTUBE_API_KEYS"`
	YoutubeFetchInterval int    `env:"YOUTUBE_FETCH_INTERVAL" env-default:"10"` // seconds
	YoutubeQuery         string `env:"YOUTUBE_QUERY" env-default:"cricket"`

	// Postgres Related Configs
	PostgresHost string `env:"POSTGRES_HOST" env-default:"localhost"`
	PostgresPort int    `env:"POSTGRES_PORT" env-default:"5432"`
	PostgresUser string `env:"POSTGRES_USER" env-default:"postgres"`
	PostgresPass string `env:"POSTGRES_PASS" env-default:"postgres"`
	PostgresDb   string `env:"POSTGRES_DB" env-default:"fampay"`
}

// TODO: Revisit this to make it follow singleton pattern
var Config ConfigDatabase

func ReadFromEnv() ConfigDatabase {
	err := cleanenv.ReadEnv(&Config)
	if err != nil {
		panic(err)
	}

	return Config
}
