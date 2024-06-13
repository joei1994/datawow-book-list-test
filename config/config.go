package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	Http  HttpConfig
	MySql MySqlConfig
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		Http:  LoadHttpConfig(),
		MySql: LoadMySqlConfig(),
	}
}
