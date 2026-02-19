package controllers

import (
	"net/http"
	"strconv"

	"multishop/services"
	"multishop/utils"

	"github.com/gin-gonic/gin"
)

// GetPublicProducts godoc
// @Summary Get public products of a shop
// @Description Returns the list of publicly visible products for a specific shop. PurchasePrice is never exposed. Accessible without authentication.
// @Tags Public
// @Produce json
// @Param shopID path int true "Shop ID"
// @Success 200 {object} map[string]interface{} "List of public products"
// @Failure 400 {object} map[string]interface{} "Invalid shop ID"
// @Failure 404 {object} map[string]interface{} "Shop not found or inactive"
// @Router /api/public/{shopID}/products [get]
func GetPublicProducts(c *gin.Context) {

	shopIDParam := c.Param("shopID")

	shopID, err := strconv.ParseUint(shopIDParam, 10, 64)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid shop id")
		return
	}

	products, err := services.GetPublicProducts(uint(shopID))
	if err != nil {
		utils.Error(c, http.StatusNotFound, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, products)
}

// GetWhatsAppLink godoc
// @Summary Generate WhatsApp link for a product
// @Description Generates a dynamic WhatsApp link for a specific product of a given shop. Accessible without authentication.
// @Tags Public
// @Produce json
// @Param shopID path int true "Shop ID"
// @Param productID path int true "Product ID"
// @Success 200 {object} map[string]string "Generated WhatsApp link"
// @Failure 400 {object} map[string]interface{} "Invalid shop or product ID"
// @Failure 404 {object} map[string]interface{} "Shop or product not found"
// @Router /api/public/{shopID}/products/{productID}/whatsapp [get]
func GetWhatsAppLink(c *gin.Context) {

	shopIDParam := c.Param("shopID")
	productIDParam := c.Param("productID")

	shopID, err := strconv.ParseUint(shopIDParam, 10, 64)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid shop id")
		return
	}

	productID, err := strconv.ParseUint(productIDParam, 10, 64)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid product id")
		return
	}

	link, err := services.GenerateWhatsAppLink(uint(shopID), uint(productID))
	if err != nil {
		utils.Error(c, http.StatusNotFound, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, gin.H{
		"whatsapp_url": link,
	})
}
