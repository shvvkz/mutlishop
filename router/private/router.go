package private

import (
	"multishop/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	privateGroup := r.Group("/api/")
	privateGroup.Use(middleware.AuthMiddleware())

	RegisterProductRoutes(privateGroup)
	RegisterTransactionRoutes(privateGroup)
	RegisterReportRoutes(privateGroup)
}
