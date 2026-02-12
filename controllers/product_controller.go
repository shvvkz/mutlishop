package controllers

import (
	"net/http"

	"multishop/models"
	"multishop/services"
	"multishop/utils"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {

	shopID := c.GetUint("shop_id")
	role := c.GetString("role")

	products, err := services.GetProducts(shopID, role)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, products)
}

type ProductInput struct {
	Name          string  `json:"name" binding:"required"`
	Description   string  `json:"description"`
	Category      string  `json:"category"`
	PurchasePrice float64 `json:"purchase_price"`
	SellingPrice  float64 `json:"selling_price" binding:"required"`
	Stock         int     `json:"stock" binding:"required"`
	ImageURL      string  `json:"image_url"`
}

func CreateProduct(c *gin.Context) {

	shopID := c.GetUint("shop_id")

	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	product := models.Product{
		Name:          input.Name,
		Description:   input.Description,
		Category:      input.Category,
		PurchasePrice: input.PurchasePrice,
		SellingPrice:  input.SellingPrice,
		Stock:         input.Stock,
		ImageURL:      input.ImageURL,
	}

	if err := services.CreateProduct(shopID, product); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusCreated, gin.H{
		"message": "product created successfully",
	})
}

func UpdateProduct(c *gin.Context) {

	shopID := c.GetUint("shop_id")
	role := c.GetString("role")
	id := c.Param("id")

	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	product := models.Product{
		Name:          input.Name,
		Description:   input.Description,
		Category:      input.Category,
		PurchasePrice: input.PurchasePrice,
		SellingPrice:  input.SellingPrice,
		Stock:         input.Stock,
		ImageURL:      input.ImageURL,
	}

	if err := services.UpdateProduct(shopID, role, id, product); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, gin.H{
		"message": "product updated successfully",
	})
}

func DeleteProduct(c *gin.Context) {

	shopID := c.GetUint("shop_id")
	id := c.Param("id")

	if err := services.DeleteProduct(shopID, id); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, gin.H{
		"message": "product deleted successfully",
	})
}
