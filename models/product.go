package models

import "time"

type Product struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"not null;uniqueIndex:idx_shop_product_name"`
	Description   string
	Category      string
	PurchasePrice float64 `gorm:"not null"`
	SellingPrice  float64 `gorm:"not null"`
	Stock         int     `gorm:"not null"`
	ImageURL      string
	ShopID        uint `gorm:"uniqueIndex:idx_shop_product_name"`
	CreatedAt     time.Time

	Shop Shop `gorm:"foreignKey:ShopID;constraint:OnDelete:CASCADE"`
}
