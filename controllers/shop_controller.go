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

// UpdateShopWhatsApp godoc
// @Summary Update shop WhatsApp number
// @Description Allows a SuperAdmin to update the WhatsApp number of their shop.
// @Tags Shop
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body UpdateWhatsAppInput true "New WhatsApp number"
// @Success 200 {object} map[string]interface{} "WhatsApp number updated"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden - SuperAdmin only"
// @Router /api/shop/whatsapp [patch]
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
