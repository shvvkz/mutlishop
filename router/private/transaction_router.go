package private

import (
	"multishop/controllers"
	"multishop/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterTransactionRoutes(rg *gin.RouterGroup) {

	transactions := rg.Group("/transactions")

	transactions.GET("",
		middleware.RequireRole("SuperAdmin", "Admin"),
		controllers.GetTransactions,
	)

	transactions.POST("",
		middleware.RequireRole("SuperAdmin", "Admin"),
		controllers.CreateTransaction,
	)

	transactions.DELETE("/:id",
		middleware.RequireRole("SuperAdmin", "Admin"),
		controllers.DeleteTransaction,
	)
}
