package private

import (
	"multishop/controllers"
	"multishop/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterShopRoutes(rg *gin.RouterGroup) {

	shop := rg.Group("/shop")

	shop.PATCH("/whatsapp",
		middleware.RequireRole("SuperAdmin"),
		controllers.UpdateShopWhatsApp,
	)
}
