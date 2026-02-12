package private

import (
	"multishop/controllers"
	"multishop/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(rg *gin.RouterGroup) {

	products := rg.Group("/products")

	products.GET("", controllers.GetProducts)

	products.POST("",
		middleware.RequireRole("SuperAdmin", "Admin"),
		controllers.CreateProduct,
	)

	products.PUT("/:id",
		middleware.RequireRole("SuperAdmin", "Admin"),
		controllers.UpdateProduct,
	)

	products.DELETE("/:id",
		middleware.RequireRole("SuperAdmin"),
		controllers.DeleteProduct,
	)
}
