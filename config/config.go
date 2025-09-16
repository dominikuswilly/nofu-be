package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	DatabaseURL string
	Port        int
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	if port == 0 {
		port = 8080
	}

	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        port,
	}
}
