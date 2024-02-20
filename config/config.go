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

// HTTPPort returns the port number for HTTP.
func HTTPPort() string {
	return os.Getenv("HTTP_PORT")
}

// HTTPSPort returns the port number for HTTPS.
func HTTPSPort() string {
	return os.Getenv("HTTPS_PORT")
}