package config

import "os"

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
}

func LoadConfig() *Config {
    return &Config{
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
        DBUser:     getEnv("DB_USER", "user"),
        DBPassword: getEnv("DB_PASSWORD", "password"),
        DBName:     getEnv("DB_NAME", "database"),
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
