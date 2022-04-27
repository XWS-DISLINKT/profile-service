package main

import (
	"profile-service/startup"
	cfg "profile-service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
