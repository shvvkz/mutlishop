package controllers

import (
	"net/http"

	"multishop/models"
	"multishop/services"
	"multishop/utils"

	"github.com/gin-gonic/gin"
)

func GetTransactions(c *gin.Context) {

	shopID := c.GetUint("shop_id")

	transactions, err := services.GetTransactions(shopID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, transactions)
}

type TransactionInput struct {
	Type      string  `json:"type" binding:"required"`
	ProductID *uint   `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Amount    float64 `json:"amount"`
}

func CreateTransaction(c *gin.Context) {

	shopID := c.GetUint("shop_id")

	var input TransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	transaction := models.Transaction{
		Type:      input.Type,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		Amount:    input.Amount,
	}

	if err := services.CreateTransaction(shopID, transaction); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(c, http.StatusCreated, gin.H{
		"message": "transaction created successfully",
	})
}

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
