package controllers

import (
	"net/http"

	"multishop/services"
	"multishop/utils"

	"github.com/gin-gonic/gin"
)

func GetDashboard(c *gin.Context) {

	shopID := c.GetUint("shop_id")

	dashboard, err := services.GetDashboard(shopID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, dashboard)
}
