package auth

import (
	"github.com/MergeMinds/mm-backend-go/internal/auth/cookie"
	"github.com/MergeMinds/mm-backend-go/internal/auth/session"
	"github.com/MergeMinds/mm-backend-go/internal/auth/user"
	"github.com/MergeMinds/mm-backend-go/internal/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"go.uber.org/zap"
)

// Swagger parameters
// @title           MergeMinds Web-API
// @version         1.0
// @host      localhost:8080
// @BasePath  /api/v1
// @securityDefinitions.basic  BasicAuth

func SetupRoutes(
	r *gin.RouterGroup,
	userRepo user.Repo,
	sessionRepo session.Repo,
	logger *zap.Logger,
	cookieConfig *cookie.CookieConfig,
) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/login", func(ctx *gin.Context) {
		routes.Login(ctx, userRepo, sessionRepo, logger, cookieConfig)
	})

	r.POST("/register", func(ctx *gin.Context) {
		routes.Register(ctx, userRepo, sessionRepo, logger, cookieConfig)
	})

	r.POST("/logout", func(ctx *gin.Context) {
		routes.Logout(ctx, userRepo, sessionRepo, logger, cookieConfig)
	})

	r.GET("/session", func(ctx *gin.Context) {
		routes.Session(ctx, userRepo, sessionRepo, logger, cookieConfig)
	})

	r.GET("/block/:id", func(ctx *gin.Context) {
		blockId := ctx.Param("id")
		routes.GetBlock(ctx, blockId)
	})

	r.POST("/block", func(ctx *gin.Context) {
		routes.CreateBlock(ctx)
	})

	r.PATCH("/block/:id", func(ctx *gin.Context) {
		blockId := ctx.Param("id")
		routes.PatchBlock(ctx, blockId)
	})

	r.DELETE("/block/:id", func(ctx *gin.Context) {
		blockId := ctx.Param("id")
		routes.DeleteBlock(ctx, blockId)
	})

}
