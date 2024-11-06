package config
// Secure Configuration setup (environment variables)

import (
    "log"
    "github.com/joho/godotenv"
    "os"
)

func LoadEnv() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }
}

func GetEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
