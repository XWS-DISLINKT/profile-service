package config

import "os"

type Config struct {
	Port          string
	ProfileDBHost string
	ProfileDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:          os.Getenv("PROFILE_SERVICE_PORT"),
		ProfileDBHost: os.Getenv("PROFILE_DB_HOST"),
		ProfileDBPort: os.Getenv("PROFILE_DB_PORT"),
	}
}
