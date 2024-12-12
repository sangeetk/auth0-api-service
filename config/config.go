package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Auth0Domain       string
	Auth0ClientID     string
	Auth0ClientSecret string
	Auth0Audience     string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Auth0Domain:       os.Getenv("AUTH0_DOMAIN"),
		Auth0ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		Auth0ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		Auth0Audience:     os.Getenv("AUTH0_AUDIENCE"),
	}
}
