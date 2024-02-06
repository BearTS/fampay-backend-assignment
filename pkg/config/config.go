// Package to keep track and handle configurations
package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

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

var Config ConfigDatabase

func ReadFromEnv() ConfigDatabase {
	file, err := os.OpenFile(".env", os.O_RDONLY|os.O_SYNC, 0)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No .env file found, reading from environment")
			err = cleanenv.ReadEnv(&Config)
			if err != nil {
				panic(err)
			}
			return Config
		}
		panic(err)
	}

	defer file.Close()

	err = cleanenv.ReadConfig(".env", &Config)
	if err != nil {
		panic(err)
	}

	return Config
}
