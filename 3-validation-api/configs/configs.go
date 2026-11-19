package configs

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EmailConfig struct {
	Email    string `json: email`
	Password string `json: password`
	Address  string `json: address`
}

func LoadConfig() (*EmailConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
		return nil, errors.New("Error loading .env file, using default config")
	}
	return &EmailConfig{
		Email:    os.Getenv("Email"),
		Password: os.Getenv("Password"),
		Address:  os.Getenv("Address"),
	}, nil

}
