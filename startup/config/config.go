package config

type Config struct {
	Port          string
	ProfileDBHost string
	ProfileDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:          "8002",      //os.Getenv("PROFILE_SERVICE_PORT"),
		ProfileDBHost: "localhost", //os.Getenv("PROFILE_DB_HOST"),
		ProfileDBPort: "27017",     //os.Getenv("PROFILE_DB_PORT"),
	}
}
