package main

import (
	"os"

	"github.com/InTeam-Russia/go-backend-template/internal/applogger"
	"github.com/InTeam-Russia/go-backend-template/internal/config"
	"github.com/InTeam-Russia/go-backend-template/internal/db"
)

func main() {
	config, err := config.LoadFromEnv()
	if err != nil {
		panic(err)
	}

	logger := applogger.Create(config.LogLevel)

	pgPool, err := db.DropDb(config.PostgresUrl, os.Getenv("SQL_FILE"), logger)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer pgPool.Close()
}
