package controllers

import (
	"strconv"

	"multishop/services"
	"multishop/utils"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	shopID := c.GetUint("shop_id")

	users, err := services.GetUsers(shopID)
	if err != nil {
		utils.Error(c, 500, err.Error())
		return
	}

	utils.JSON(c, 200, users)
}

func UpdateUserRole(c *gin.Context) {

	shopID := c.GetUint("shop_id")

	idParam := c.Param("id")
	userID, _ := strconv.ParseUint(idParam, 10, 64)

	var input struct {
		Role string `json:"role"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, 400, err.Error())
		return
	}

	if err := services.UpdateUserRole(shopID, uint(userID), input.Role); err != nil {
		utils.Error(c, 400, err.Error())
		return
	}

	utils.JSON(c, 200, gin.H{"message": "role updated"})
}

func DeleteUser(c *gin.Context) {

	shopID := c.GetUint("shop_id")

	idParam := c.Param("id")
	userID, _ := strconv.ParseUint(idParam, 10, 64)

	if err := services.DeleteUser(shopID, uint(userID)); err != nil {
		utils.Error(c, 400, err.Error())
		return
	}

	utils.JSON(c, 200, gin.H{"message": "user deleted"})
}
