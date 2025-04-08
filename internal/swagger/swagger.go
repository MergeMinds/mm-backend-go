package swagger

import (
	docs "github.com/MergeMinds/mm-backend-go/docs"
	"github.com/gin-gonic/gin"
)

func InitSwagger(apiPath *gin.RouterGroup) {
	docs.SwaggerInfo.Title = "MergeMinds Web-API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "0.0.0.0:8080"
	docs.SwaggerInfo.BasePath = apiPath.BasePath()
}
