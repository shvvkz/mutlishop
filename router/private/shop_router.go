package private

import (
	"multishop/controllers"
	"multishop/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterShopRoutes(rg *gin.RouterGroup) {

	shop := rg.Group("/shop")

	shop.PUT("/whatsapp",
		middleware.RequireRole("SuperAdmin"),
		controllers.UpdateShopWhatsApp,
	)
}
