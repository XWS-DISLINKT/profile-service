package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port                     string
	ProfileDBHost            string
	ProfileDBPort            string
	ConnectionHost           string
	ConnectionPort           string
	NatsHost                 string
	NatsPort                 string
	NatsUser                 string
	NatsPass                 string
	CreateUserCommandSubject string
	CreateUserReplySubject   string
}

func NewConfig() *Config {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		fmt.Println("docker")

		return &Config{
			Port:                     os.Getenv("PROFILE_SERVICE_PORT"),
			ProfileDBHost:            os.Getenv("PROFILE_DB_HOST"),
			ProfileDBPort:            os.Getenv("PROFILE_DB_PORT"),
			ConnectionHost:           os.Getenv("CONNECTION_SERVICE_HOST"),
			ConnectionPort:           os.Getenv("CONNECTION_SERVICE_PORT"),
			NatsHost:                 os.Getenv("NATS_HOST"),
			NatsPort:                 os.Getenv("NATS_PORT"),
			NatsUser:                 os.Getenv("NATS_USER"),
			NatsPass:                 os.Getenv("NATS_PASS"),
			CreateUserCommandSubject: os.Getenv("CREATE_USER_COMMAND_SUBJECT"),
			CreateUserReplySubject:   os.Getenv("CREATE_USER_REPLY_SUBJECT"),
		}
	} else {
		fmt.Println("local")

		return &Config{
			Port:                     "8001",
			ProfileDBHost:            "localhost",
			ProfileDBPort:            "27017",
			ConnectionHost:           "localhost",
			ConnectionPort:           "8004",
			NatsHost:                 "localhost",
			NatsPort:                 "4222",
			NatsUser:                 "ruser",
			NatsPass:                 "T0pS3cr3t",
			CreateUserCommandSubject: "user.create.command",
			CreateUserReplySubject:   "user.create.reply",
		}
	}
}
