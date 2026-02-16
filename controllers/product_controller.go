package controllers

import (
	"net/http"

	"multishop/models"
	"multishop/services"
	"multishop/utils"

	"github.com/gin-gonic/gin"
)

// GetProducts godoc
// @Summary Get all products of the current shop
// @Description Returns all products belonging to the authenticated user's shop. SuperAdmin can see PurchasePrice, Admin cannot.
// @Tags Products
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "List of products"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/products [get]
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
	Name          string   `json:"name" binding:"required"`
	Description   string   `json:"description"`
	Category      string   `json:"category"`
	PurchasePrice *float64 `json:"purchase_price" binding:"required"`
	SellingPrice  *float64 `json:"selling_price" binding:"required"`
	Stock         *int     `json:"stock" binding:"required"`
	ImageURL      string   `json:"image_url"`
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Allows an Admin or SuperAdmin to create a new product in their shop. Product name must be unique per shop.
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body ProductInput true "Product creation payload"
// @Success 201 {object} map[string]interface{} "Product created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/products [post]
func CreateProduct(c *gin.Context) {

	shopID := c.GetUint("shop_id")

	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var purchasePrice float64
	switch {
	case input.PurchasePrice == nil:
		utils.Error(c, http.StatusBadRequest, "purchase_price is required")
		return
	case *input.PurchasePrice <= 0:
		utils.Error(c, http.StatusBadRequest, "purchase_price must be at least 1")
		return
	default:
		purchasePrice = *input.PurchasePrice
	}

	var sellingPrice float64
	switch {
	case input.SellingPrice == nil:
		utils.Error(c, http.StatusBadRequest, "selling_price is required")
		return
	case *input.SellingPrice <= 0:
		utils.Error(c, http.StatusBadRequest, "selling_price must be at least 1")
		return
	default:
		sellingPrice = *input.SellingPrice
	}

	var stock int
	switch {
	case input.Stock == nil:
		utils.Error(c, http.StatusBadRequest, "stock is required")
		return
	case *input.Stock <= 0:
		utils.Error(c, http.StatusBadRequest, "stock must be at least 1")
		return
	default:
		stock = *input.Stock
	}

	product := models.Product{
		Name:          input.Name,
		Description:   input.Description,
		Category:      input.Category,
		PurchasePrice: purchasePrice,
		SellingPrice:  sellingPrice,
		Stock:         stock,
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

// UpdateProduct godoc
// @Summary Update an existing product
// @Description Allows Admin or SuperAdmin to update a product belonging to their shop. Product name must remain unique. Only SuperAdmin can update PurchasePrice.
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Param input body ProductInput true "Updated product payload"
// @Success 200 {object} map[string]interface{} "Product updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input or product not found"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Router /api/products/{id} [put]
func UpdateProduct(c *gin.Context) {

	shopID := c.GetUint("shop_id")
	role := c.GetString("role")
	id := c.Param("id")

	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var purchasePrice float64
	switch {
	case input.PurchasePrice == nil:
		utils.Error(c, http.StatusBadRequest, "purchase_price is required")
		return
	case *input.PurchasePrice <= 0:
		utils.Error(c, http.StatusBadRequest, "purchase_price must be at least 1")
		return
	default:
		purchasePrice = *input.PurchasePrice
	}

	var sellingPrice float64
	switch {
	case input.SellingPrice == nil:
		utils.Error(c, http.StatusBadRequest, "selling_price is required")
		return
	case *input.SellingPrice <= 0:
		utils.Error(c, http.StatusBadRequest, "selling_price must be at least 1")
		return
	default:
		sellingPrice = *input.SellingPrice
	}

	var stock int
	switch {
	case input.Stock == nil:
		utils.Error(c, http.StatusBadRequest, "stock is required")
		return
	case *input.Stock <= 0:
		utils.Error(c, http.StatusBadRequest, "stock must be at least 1")
		return
	default:
		stock = *input.Stock
	}

	product := models.Product{
		Name:          input.Name,
		Description:   input.Description,
		Category:      input.Category,
		PurchasePrice: purchasePrice,
		SellingPrice:  sellingPrice,
		Stock:         stock,
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

// DeleteProduct godoc
// @Summary Delete a product
// @Description Allows a SuperAdmin to delete a product belonging to their shop.
// @Tags Products
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{} "Product deleted successfully"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden - SuperAdmin only"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/products/{id} [delete]
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
