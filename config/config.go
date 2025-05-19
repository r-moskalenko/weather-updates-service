package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	FromEmail           string
	SendGridApiKey      string
	MailVerifTemplateID string
	Scheme              string
	Host                string
}

// New returns a new Config struct
func New() *Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Config{
		FromEmail:           getEnv("FROM_EMAIL", ""),
		SendGridApiKey:      getEnv("SENDGRID_API_KEY", ""),
		MailVerifTemplateID: getEnv("MAIL_VERIF_TEMPLATE_ID", ""),
		Scheme:              getEnv("SCHEME", "http"),
		Host:                getEnv("HOST", "localhost:8080"),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	value := os.Getenv(key)

	if value != "" {
		return value
	}

	return defaultVal
}
