package config

import "os"

type Config struct {
	DBDriver      string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	ServerAddress string
}

func LoadConfig() *Config {
	return &Config{
		DBDriver:      getEnv("DB_DRIVER", "postgres"),
		DBHost:        getEnv("DB_HOST", "localhost"),
		DBPort:        getEnv("DB_PORT", "5432"),
		DBUser:        getEnv("DB_USER", "root"),
		DBPassword:    getEnv("DB_PASSWORD", "password"),
		DBName:        getEnv("DB_NAME", "simple_bank"),
		ServerAddress: getEnv("SERVER_ADDRESS", "0.0.0.0:8081"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
