package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port          string
	ProfileDBHost string
	ProfileDBPort string
}

func NewConfig() *Config {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		fmt.Println("docker")

		return &Config{
			Port:          os.Getenv("PROFILE_SERVICE_PORT"),
			ProfileDBHost: os.Getenv("PROFILE_DB_HOST"),
			ProfileDBPort: os.Getenv("PROFILE_DB_PORT"),
		}
	} else {
		fmt.Println("local")

		return &Config{
			Port:          "8001",
			ProfileDBHost: "localhost",
			ProfileDBPort: "27017",
		}
	}
}
