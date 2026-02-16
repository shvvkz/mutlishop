package controllers

import (
	"net/http"

	"multishop/services"
	"multishop/utils"

	"github.com/gin-gonic/gin"
)

// GetDashboard godoc
// @Summary Get financial dashboard of the shop
// @Description Returns financial metrics for the authenticated SuperAdmin's shop including total sales, expenses, withdrawals, net profit and low stock products.
// @Tags Reports
// @Produce json
// @Security BearerAuth
// @Success 200 {object} services.DashboardResponse "Financial dashboard data"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden - SuperAdmin only"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/reports/dashboard [get]
func GetDashboard(c *gin.Context) {

	shopID := c.GetUint("shop_id")

	dashboard, err := services.GetDashboard(shopID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, dashboard)
}
