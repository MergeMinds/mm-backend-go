package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	LogLevel            string   `envconfig:"LOG_LEVEL" default:"debug"`
	PostgresUrl         string   `envconfig:"POSTGRES_URL" default:"postgres://t3m8ch@localhost/productsdb"`
	RedisUrl            string   `envconfig:"REDIS_URL" default:"redis://dragonfly:6379/0"`
	SessionCookieSecure bool     `envconfig:"SESSION_COOKIE_SECURE" default:"false"`
	SessionCookieDomain string   `envconfig:"SESSION_COOKIE_DOMAIN" default:"localhost:5173"`
	AllowOrigins        []string `envconfig:"ALLOW_ORIGINS" default:"http://localhost:5173"`
	AdminUsername       string   `envconfig:"ADMIN_USERNAME" default:"admin"`
	AdminPassword       string   `envconfig:"ADMIN_PASSWORD" default:"123456"`
	AdminEmail          string   `envconfig:"ADMIN_EMAIL" default:"admin@admin.ru"`
}

func LoadFromEnv() (*Config, error) {
	var config_env Config
	err := envconfig.Process("", &config_env)

	if err != nil {
		fmt.Println("Failed to process env vars: %v", err)
	}

	sessionCookieSecure, err := strconv.ParseBool(os.Getenv("SESSION_COOKIE_SECURE"))
	config_env.SessionCookieSecure = sessionCookieSecure
	if err != nil {
		return nil, err
	}

	return &config_env, nil
}
