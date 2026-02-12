package controllers

import (
	"net/http"
	"strconv"

	"multishop/services"
	"multishop/utils"

	"github.com/gin-gonic/gin"
)

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
