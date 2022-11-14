package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// get all env
type (
	Config struct {
		Server ServerConfig
	}
	ServerConfig struct {
		Port string
	}
)

func LoadConfig() (config *Config, err error) {

	err = godotenv.Load()
	if err != nil {
		log.Println(err.Error())
		return config, err
	}

	config = &Config{
		Server: ServerConfig{
			Port: os.Getenv("PORT"),
		},
	}

	return config, nil
}
