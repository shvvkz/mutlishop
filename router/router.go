package router

import (
	"multishop/router/private"
	"multishop/router/public"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	public.RegisterRoutes(r)
	private.RegisterRoutes(r)

	return r
}
