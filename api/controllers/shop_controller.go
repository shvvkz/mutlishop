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

type ShopInput struct {
	Name           string `json:"name" binding:"required"`
	WhatsAppNumber string `json:"whatsapp_number" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required,min=6"`
}

// CreateShop godoc
// @Summary Create a new shop with SuperAdmin
// @Description Creates a new shop along with a SuperAdmin user. The provided email must be unique across all users.
// @Tags Public
// @Accept json
// @Produce json
// @Param input body ShopInput true "Shop creation payload"
// @Success 201 {object} map[string]interface{} "Shop and SuperAdmin created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input or email already exists"
// @Router /api/public/shop [post]
func CreateShop(c *gin.Context) {

	var input ShopInput

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	shopID, err := services.CreateShopWithSuperAdmin(
		input.Name,
		input.WhatsAppNumber,
		input.Email,
		input.Password,
	)

	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSON(c, http.StatusCreated, gin.H{
		"shop_id": shopID,
		"message": "shop and superadmin created successfully",
	})
}

// @GetShops godoc
// @Summary Get list of all shops
// @Description Retrieves a list of all shops with their basic details. Accessible without authentication.
// @Tags Public
// @Produce json
// @Success 200 {array} map[string]interface{} "List of shops"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/public/shop [get]
func GetShops(c *gin.Context) {
	shops, err := services.GetShops()
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, shops)
}
