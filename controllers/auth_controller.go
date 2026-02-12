package controllers

import (
	"net/http"

	"multishop/models"
	"multishop/services"
	"multishop/utils"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {

	shopID := c.GetUint("shop_id")

	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, 400, err.Error())
		return
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     "Admin",
		ShopID:   shopID,
	}

	if err := services.Register(user.Name, user.Email, user.Password, user.Role, user.ShopID); err != nil {
		utils.Error(c, 400, err.Error())
		return
	}

	utils.JSON(c, 201, gin.H{"message": "admin created"})
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body LoginInput true "Login data"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /api/login [post]
func Login(c *gin.Context) {

	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := services.Login(input.Email, input.Password)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, gin.H{
		"token": token,
	})
}
