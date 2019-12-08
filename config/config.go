package config

import (
	"github.com/joho/godotenv"
	"os"
)

//AppConfig ...
type AppConfig struct {
	SendGridAPI string
	RabbitMQURL string
}

//LoadEnv ...
func LoadEnv() (AppConfig, error) {
	err := godotenv.Load()

	if err != nil {
		return AppConfig{}, err
	}

	return AppConfig{
		SendGridAPI: os.Getenv("SENDGRID_KEY"),
		RabbitMQURL: os.Getenv("RABBITMQ_URL"),
	}, nil

}
