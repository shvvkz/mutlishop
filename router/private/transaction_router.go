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

	transactions.POST("/sale",
		middleware.RequireRole("SuperAdmin", "Admin"),
		controllers.CreateTransactionSale,
	)

	transactions.POST("/expense",
		middleware.RequireRole("SuperAdmin", "Admin"),
		controllers.CreateTransactionExpense,
	)

	transactions.POST("/withdrawal",
		middleware.RequireRole("SuperAdmin", "Admin"),
		controllers.CreateTransactionWithdrawal,
	)

	transactions.DELETE("/:id",
		middleware.RequireRole("SuperAdmin", "Admin"),
		controllers.DeleteTransaction,
	)
}
