package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type JWTConfig struct {
	Secret string
	Issuer string
}

type MailConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	FromName string
}

type Config struct {
	App  AppConfig
	DB   DatabaseConfig
	JWT  JWTConfig
	Mail MailConfig
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	return &Config{
		App: AppConfig{
			AppName: os.Getenv("APP_NAME"),
			AppEnv:  os.Getenv("APP_ENV"),
			AppPort: os.Getenv("APP_PORT"),
		},
		DB: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		JWT: JWTConfig{
			Secret: os.Getenv("JWT_SECRET"),
			Issuer: os.Getenv("JWT_ISSUER"),
		},
		Mail: MailConfig{
			Host:     os.Getenv("SMTP_HOST"),
			Port:     os.Getenv("SMTP_PORT"),
			Username: os.Getenv("SMTP_USERNAME"),
			Password: os.Getenv("SMTP_PASSWORD"),
			FromName: os.Getenv("SMTP_FROM_NAME"),
		},
	}, nil
}
