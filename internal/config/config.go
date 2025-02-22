package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	LogLevel            string
	PostgresUrl         string
	RedisUrl            string
	SessionCookieSecure bool
	SessionCookieDomain string
	AllowOrigins        []string
	AdminUsername       string
	AdminPassword       string
	AdminEmail          string
}

func LoadFromEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found")
	}

	sessionCookieSecure, err := strconv.ParseBool(os.Getenv("SESSION_COOKIE_SECURE"))
	if err != nil {
		return nil, err
	}

	return &Config{
		LogLevel:            os.Getenv("LOG_LEVEL"),
		PostgresUrl:         os.Getenv("POSTGRES_URL"),
		RedisUrl:            os.Getenv("REDIS_URL"),
		SessionCookieSecure: sessionCookieSecure,
		SessionCookieDomain: os.Getenv("SESSION_COOKIE_DOMAIN"),
		AllowOrigins:        strings.Split(os.Getenv("ALLOW_ORIGINS"), ","),
		AdminUsername:       os.Getenv("ADMIN_USERNAME"),
		AdminPassword:       os.Getenv("ADMIN_PASSWORD"),
		AdminEmail:          os.Getenv("ADMIN_EMAIL"),
	}, nil
}
