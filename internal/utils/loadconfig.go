package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() Config {
	godotenv.Load()

	return Config{
		ENV:  fallback(os.Getenv("ENV"), "development"),
		CRT:  os.Getenv("CRT"),
		KEY:  os.Getenv("KEY"),
		PORT: fallback(os.Getenv("PORT"), "5099"),
	}
}

type Config struct {
	ENV   string
	CRT   string
	KEY   string
	PORT  string
}

func fallback(value, def string) string {
	if value == "" {
		return def
	}
	return value
}
