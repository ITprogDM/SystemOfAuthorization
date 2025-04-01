package config

import (
	"errors"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	DBUser    string
	Password  string
	Host      string
	Port      string
	DBName    string
	SSLMode   string
	JWTSecret string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		return nil, err
	}

	cfg := &Config{
		DBUser:    os.Getenv("DB_USER"),
		Password:  os.Getenv("DB_PASSWORD"),
		Host:      os.Getenv("DB_HOST"),
		Port:      os.Getenv("DB_PORT"),
		DBName:    os.Getenv("DB_NAME"),
		SSLMode:   os.Getenv("SSLMode"),
		JWTSecret: os.Getenv("JWTSecret"),
	}

	if cfg.DBUser == "" || cfg.Password == "" || cfg.Host == "" || cfg.Port == "" || cfg.DBName == "" || cfg.SSLMode == "" {
		log.Println("Ошибка в конфигурации, проверьте конфиг!")
		return nil, errors.New("Ошибка в конфигурации, проверьте конфиг!")
	}

	return cfg, nil
}
