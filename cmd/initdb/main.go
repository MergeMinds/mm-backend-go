package main

import (
	"os"

	"github.com/InTeam-Russia/go-backend-template/internal/applogger"
	"github.com/InTeam-Russia/go-backend-template/internal/auth/user"
	"github.com/InTeam-Russia/go-backend-template/internal/config"
	"github.com/InTeam-Russia/go-backend-template/internal/db"
)

func main() {
	config, err := config.LoadFromEnv()
	if err != nil {
		panic(err)
	}

	logger := applogger.Create(config.LogLevel)

	pgPool, err := db.InitDb(config.PostgresUrl, os.Getenv("SQL_FILE"), logger)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer pgPool.Close()

	userRepo := user.NewPGRepo(pgPool, logger)
	_, err = userRepo.Create(&user.CreateModel{
		FirstName: "Admin",
		LastName:  "Admin",
		Username:  config.AdminUsername,
		Role:      "ADMIN",
		Password:  config.AdminPassword,
		Email:     config.AdminEmail,
	})

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	logger.Info("Admin created!")
}
