package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func get() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_HOST"),
		},
		Database: Database{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
			Tz:   os.Getenv("DB_TZ"),
		},
	}
}
