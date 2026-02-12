package controllers

import (
	"net/http"

	"multishop/services"
	"multishop/utils"

	"github.com/gin-gonic/gin"
)

type UpdateWhatsAppInput struct {
	WhatsAppNumber string `json:"whatsapp_number" binding:"required"`
}

func UpdateShopWhatsApp(c *gin.Context) {

	shopID := c.GetUint("shop_id")

	var input UpdateWhatsAppInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := services.UpdateShopWhatsApp(shopID, input.WhatsAppNumber); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, gin.H{
		"message": "whatsapp number updated",
	})
}
