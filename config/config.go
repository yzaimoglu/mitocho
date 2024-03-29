package config

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"os"
)

func Load() {
	LoadEnv()
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Infof("No .env file found. Defaulting to system environment variables.")
	}
}

// Debug returns whether the website is in debug mode.
func Debug() bool {
	return os.Getenv("DEBUG") == "true"
}

// Host returns the domain name of the website.
func Host() string {
	return os.Getenv("HOST")
}

// Port returns the port number for the website.
func Port() string {
	return os.Getenv("PORT")
}

// HTTPPort returns the port number for HTTP.
func HTTPPort() string {
	return os.Getenv("HTTP_PORT")
}

// HTTPSPort returns the port number for HTTPS.
func HTTPSPort() string {
	return os.Getenv("HTTPS_PORT")
}

// SSL returns whether the website should create an SSL Server.
func SSL() bool {
	return os.Getenv("SSL") == "true"
}
