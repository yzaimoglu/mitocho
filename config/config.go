package config

import (
	"github.com/joho/godotenv"
	"os"
)

func Load() {
	LoadEnv()
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func Debug() bool {
	return os.Getenv("DEBUG") == "true"
}

// Host returns the domain name of the website.
func Host() string {
	return os.Getenv("HOST")
}

func HTTPPort() string {
	return os.Getenv("HTTP_PORT")
}

func HTTPSPort() string {
	return os.Getenv("HTTPS_PORT")
}
