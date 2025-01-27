package config

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	PORT               string
	DBURI              string
	SECRET             string
	JWT_KEY_SECRET     string
	EXCHANGE_API_TOKEN string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		return nil, errors.New("PORT tidak diatur di .env")
	}

	dbURI := os.Getenv("DBURI")
	if dbURI == "" {
		return nil, errors.New("DBURI tidak diatur di .env")
	}

	secret := os.Getenv("SECRET")
	if len(secret) != 48 {
		return nil, errors.New("SECRET harus berupa 48 karakter")
	}

	jwtKey := os.Getenv("JWT_KEY_SECRET")
	if len(jwtKey) < 20 {
		return nil, errors.New("JWT_KEY_SECRET harus lebih dari 20 karakter")
	}

	exchangeKey := os.Getenv("EXCHANGE_API_TOKEN")
	if len(exchangeKey) < 1 {
		return nil, errors.New("EXCHANGE_API_TOKEN harus diisi")
	}

	return &Config{
		PORT:               port,
		DBURI:              dbURI,
		SECRET:             secret,
		JWT_KEY_SECRET:     jwtKey,
		EXCHANGE_API_TOKEN: exchangeKey,
	}, nil
}
