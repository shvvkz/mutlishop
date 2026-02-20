package public

import (
	"multishop/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterShopRoutes(rg *gin.RouterGroup) {
	rg.GET("public/shop", controllers.GetShops)
	rg.POST("public/shop", controllers.CreateShop)
}
