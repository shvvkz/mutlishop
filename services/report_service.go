package services

import (
	"multishop/config"
	"multishop/models"
)

type DashboardResponse struct {
	TotalSales       float64 `json:"total_sales"`
	TotalExpenses    float64 `json:"total_expenses"`
	TotalWithdrawals float64 `json:"total_withdrawals"`
	NetProfit        float64 `json:"net_profit"`
	LowStockCount    int64   `json:"low_stock_products"`
}

func GetDashboard(shopID uint) (DashboardResponse, error) {

	var totalSales float64
	var totalExpenses float64
	var totalWithdrawals float64
	var lowStockCount int64

	config.DB.
		Model(&models.Transaction{}).
		Where("shop_id = ? AND type = ?", shopID, "Sale").
		Select("COALESCE(SUM(amount),0)").
		Scan(&totalSales)

	config.DB.
		Model(&models.Transaction{}).
		Where("shop_id = ? AND type = ?", shopID, "Expense").
		Select("COALESCE(SUM(amount),0)").
		Scan(&totalExpenses)

	config.DB.
		Model(&models.Transaction{}).
		Where("shop_id = ? AND type = ?", shopID, "Withdrawal").
		Select("COALESCE(SUM(amount),0)").
		Scan(&totalWithdrawals)

	config.DB.
		Model(&models.Product{}).
		Where("shop_id = ? AND stock < 5", shopID).
		Count(&lowStockCount)

	netProfit := totalSales - totalExpenses

	return DashboardResponse{
		TotalSales:       totalSales,
		TotalExpenses:    totalExpenses,
		TotalWithdrawals: totalWithdrawals,
		NetProfit:        netProfit,
		LowStockCount:    lowStockCount,
	}, nil
}
