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
		DB     DBConfig
	}
	ServerConfig struct {
		Port string
	}
	DBConfig struct {
		Host     string
		Port     string
		Username string
		Password string
		Name     string
		SSL      string
		TIMEZONE string
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
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			SSL:      os.Getenv("DB_SSL"),
			TIMEZONE: os.Getenv("DB_TIMEZONE"),
		},
	}

	return config, nil
}
