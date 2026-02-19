package private

import (
	"multishop/controllers"
	"multishop/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterReportRoutes(rg *gin.RouterGroup) {

	reports := rg.Group("/reports")

	reports.GET("/dashboard",
		middleware.RequireRole("SuperAdmin"),
		controllers.GetDashboard,
	)
}
