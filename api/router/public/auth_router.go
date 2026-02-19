package public

import (
	"multishop/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup) {

	rg.POST("/login", controllers.Login)
}
