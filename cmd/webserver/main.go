package main

import (
	"os"
	"time"

	"github.com/MergeMinds/mm-backend-go/internal/applogger"
	"github.com/MergeMinds/mm-backend-go/internal/auth"
	"github.com/MergeMinds/mm-backend-go/internal/auth/session"
	"github.com/MergeMinds/mm-backend-go/internal/auth/user"
	"github.com/MergeMinds/mm-backend-go/internal/config"
	"github.com/MergeMinds/mm-backend-go/internal/cors"
	"github.com/MergeMinds/mm-backend-go/internal/db"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	config, err := config.LoadFromEnv()
	if err != nil {
		panic(err)
	}

	logger := applogger.Create(config.LogLevel)

	dbConn, err := db.CreateDb(config.PostgresUrl, logger)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer dbConn.Close()

	redisOpts, err := redis.ParseURL(config.RedisUrl)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	redisClient := redis.NewClient(redisOpts)
	defer redisClient.Close()

	r := gin.New()

	cors.Setup(r, config)
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	cookieConfig := auth.DefaultCookieConfig()
	cookieConfig.Secure = config.SessionCookieSecure
	cookieConfig.Domain = config.SessionCookieDomain

	userRepo := user.NewPGRepo(dbConn, logger)
	sessionRepo := session.NewRedisRepo(redisClient, logger)

	auth.SetupRoutes(r, userRepo, sessionRepo, logger, cookieConfig)

	err = r.Run()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
