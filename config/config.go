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

	err = godotenv.Load(".env")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	config = &Config{
		Server: ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
		},
	}

	return config, nil
}
