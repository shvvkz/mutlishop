package controllers

import (
	"net/http"

	"multishop/models"
	"multishop/services"
	"multishop/utils"

	"github.com/gin-gonic/gin"
)

// GetTransactions godoc
// @Summary Get all transactions of the current shop
// @Description Returns all transactions (Sale, Expense, Withdrawal) belonging to the authenticated user's shop.
// @Tags Transactions
// @Produce json
// @Security BearerAuth
// @Success 200 {array} services.TransactionResponse "List of transactions"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/transactions [get]
func GetTransactions(c *gin.Context) {

	shopID := c.GetUint("shop_id")

	transactions, err := services.GetTransactions(shopID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, transactions)
}

type TransactionSaleInput struct {
	ProductID *uint   `json:"product_id"`
	Quantity  int     `json:"quantity"`
}

type TransactionExpenseInput struct {
	Amount    float64 `json:"amount"`
}

type TransactionWithdrawalInput struct {
	Amount float64 `json:"amount"`
}

// CreateTransactionSale godoc
// @Summary Create a new Sale transaction
// @Description Allows Admin or SuperAdmin to create a Sale transaction. The request must include the product ID and quantity. The system will automatically calculate the total amount based on the product price and quantity, and reduce the product stock accordingly.
// @Tags Transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body controllers.TransactionSaleInput true "Transaction payload"
// @Success 201 {object} map[string]interface{} "Transaction created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input, insufficient stock, or product not found"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Router /api/transactions/sale [post]
func CreateTransactionSale(c *gin.Context) {

	shopID := c.GetUint("shop_id")

	var input TransactionSaleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	transaction := models.Transaction{
		Type:      "Sale",
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		Amount:    0,
	}

	if err := services.CreateTransaction(shopID, transaction); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(c, http.StatusCreated, gin.H{
		"message": "transaction created successfully",
	})
}

// CreateTransactionExpense godoc
// @Summary Create a new Expense transaction
// @Description Allows Admin or SuperAdmin to create an Expense transaction. The request must include the amount of the expense.
// @Tags Transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body controllers.TransactionExpenseInput true "Transaction payload"
// @Success 201 {object} map[string]interface{} "Transaction created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Router /api/transactions/expense [post]
func CreateTransactionExpense(c *gin.Context) {

	shopID := c.GetUint("shop_id")

	var input TransactionExpenseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	transaction := models.Transaction{
		Type:   "Expense",
		Amount: input.Amount,
	}

	if err := services.CreateTransaction(shopID, transaction); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(c, http.StatusCreated, gin.H{
		"message": "transaction created successfully",
	})
}

// CreateTransactionWithdrawal godoc
// @Summary Create a new Withdrawal transaction
// @Description Allows Admin or SuperAdmin to create a Withdrawal transaction. The request must include the amount of the withdrawal.
// @Tags Transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body controllers.TransactionWithdrawalInput true "Transaction payload"
// @Success 201 {object} map[string]interface{} "Transaction created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Router /api/transactions/withdrawal [post]
func CreateTransactionWithdrawal(c *gin.Context) {

	shopID := c.GetUint("shop_id")

	var input TransactionWithdrawalInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	transaction := models.Transaction{
		Type:   "Withdrawal",
		Amount: input.Amount,
	}

	if err := services.CreateTransaction(shopID, transaction); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(c, http.StatusCreated, gin.H{
		"message": "transaction created successfully",
	})
}

// DeleteTransaction godoc
// @Summary Delete a transaction
// @Description Allows Admin or SuperAdmin to delete a transaction belonging to their shop. If the transaction is a Sale, the product stock is restored automatically.
// @Tags Transactions
// @Produce json
// @Security BearerAuth
// @Param id path int true "Transaction ID"
// @Success 200 {object} map[string]interface{} "Transaction deleted successfully"
// @Failure 400 {object} map[string]interface{} "Transaction not found"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Router /api/transactions/{id} [delete]
func DeleteTransaction(c *gin.Context) {

	shopID := c.GetUint("shop_id")
	id := c.Param("id")

	if err := services.DeleteTransaction(shopID, id); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, gin.H{
		"message": "transaction deleted successfully",
	})
}
