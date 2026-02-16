package controllers

import (
	"strconv"

	"multishop/services"
	"multishop/utils"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary Get all users of the current shop
// @Description Allows a SuperAdmin to retrieve all users belonging to their shop. Only id, email and role are returned.
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "List of users"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden - SuperAdmin only"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/users [get]
func GetUsers(c *gin.Context) {

	shopID := c.GetUint("shop_id")

	users, err := services.GetUsers(shopID)
	if err != nil {
		utils.Error(c, 500, err.Error())
		return
	}

	utils.JSON(c, 200, users)
}

// UpdateUserRole godoc
// @Summary Update user role
// @Description Allows a SuperAdmin to update the role of a user belonging to their shop. Only Admin or SuperAdmin roles are allowed.
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Param input body object{role=string} true "Role update payload (Admin or SuperAdmin)"
// @Success 200 {object} map[string]interface{} "Role updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid role or user not found"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden - SuperAdmin only"
// @Router /api/users/{id} [patch]
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

// DeleteUser godoc
// @Summary Delete a user
// @Description Allows a SuperAdmin to delete a user belonging to their shop.
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{} "User deleted successfully"
// @Failure 400 {object} map[string]interface{} "User not found"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden - SuperAdmin only"
// @Router /api/users/{id} [delete]
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
