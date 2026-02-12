package private

import (
	"multishop/controllers"
	"multishop/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {

	users := rg.Group("/users")

	users.Use(middleware.RequireRole("SuperAdmin"))

	users.GET("", controllers.GetUsers)
	users.PUT("/:id", controllers.UpdateUserRole)
	users.DELETE("/:id", controllers.DeleteUser)
	users.POST("", controllers.Register) // création Admin
}
