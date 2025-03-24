package auth

import (
	"github.com/MergeMinds/mm-backend-go/internal/auth/cookie"
	"github.com/MergeMinds/mm-backend-go/internal/auth/session"
	"github.com/MergeMinds/mm-backend-go/internal/auth/user"
	"github.com/MergeMinds/mm-backend-go/internal/routes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupRoutes(
	r *gin.Engine,
	userRepo user.Repo,
	sessionRepo session.Repo,
	logger *zap.Logger,
	cookieConfig *cookie.CookieConfig,
) {
	r.POST("/login", func(ctx *gin.Context) {
		routes.Api_login(ctx, userRepo, sessionRepo, logger, cookieConfig)
	})

	r.POST("/register", func(ctx *gin.Context) {
		routes.Api_register(ctx, userRepo, sessionRepo, logger, cookieConfig)
	})

	r.POST("/logout", func(ctx *gin.Context) {
		routes.Api_logout(ctx, userRepo, sessionRepo, logger, cookieConfig)
	})

	r.GET("/session", func(ctx *gin.Context) {
		routes.Api_session(ctx, userRepo, sessionRepo, logger, cookieConfig)
	})
}
