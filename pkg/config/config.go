// Package to keep track and handle configurations
package config

import "github.com/ilyakaznacheev/cleanenv"

type ConfigDatabase struct {
	Port int `env:"PORT" env-default:"3000"`

	// Youtube Related Configs
	YoutubeApiKeys       string `env:"YOUTUBE_API_KEYS"`
	YoutubeFetchInterval int    `env:"YOUTUBE_FETCH_INTERVAL" env-default:"10"` // seconds

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
