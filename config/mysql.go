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
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Drive:    os.Getenv("MYSQL_DRIVE"),
		Name:     os.Getenv("MYSQL_NAME"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
	}
}
