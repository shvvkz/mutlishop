package models

import "time"

type Product struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"not null;uniqueIndex:idx_shop_product_name"`
	Description   string
	Category      ProductCategory
	PurchasePrice float64 `gorm:"not null"`
	SellingPrice  float64 `gorm:"not null"`
	Stock         int     `gorm:"not null"`
	ImageURL      string
	ShopID        uint `gorm:"uniqueIndex:idx_shop_product_name"`
	CreatedAt     time.Time

	Shop Shop `gorm:"foreignKey:ShopID;constraint:OnDelete:CASCADE"`
}

type ProductCategory string

const (
	Phone      ProductCategory = "Phone"
	Television ProductCategory = "Television"
	Charger    ProductCategory = "Charger"
	Computer   ProductCategory = "Computer"
	Peripheral ProductCategory = "Peripheral"
)

var AllProductCategories = []ProductCategory{
	Phone,
	Television,
	Charger,
	Computer,
	Peripheral,
}

func IsValidCategory(cat ProductCategory) bool {
	for _, c := range AllProductCategories {
		if c == cat {
			return true
		}
	}
	return false
}
