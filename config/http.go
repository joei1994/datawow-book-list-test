package config

import "os"

type HttpConfig struct {
	Host string
	Port string
}

func LoadHttpConfig() HttpConfig {
	return HttpConfig{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}
}
