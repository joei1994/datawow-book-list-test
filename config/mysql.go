package config

import "os"

type MySqlConfig struct {
	User     string
	Password string
	Drive    string
	Name     string
	Host     string
	Port     string
}

func LoadMySqlConfig() MySqlConfig {
	return MySqlConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Drive:    os.Getenv("DB_DRIVE"),
		Name:     os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}
}
