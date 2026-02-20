package public

import (
	"multishop/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(rg *gin.RouterGroup) {

	rg.GET("public/:shopID/products", controllers.GetPublicProducts)
	rg.GET("public/:shopID/products/:productID/whatsapp", controllers.GetWhatsAppLink)
	rg.GET("public/categories", controllers.GetCategories)
}
