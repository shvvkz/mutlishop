package models

import "time"

type Transaction struct {
	ID        uint   `gorm:"primaryKey"`
	Type      string `gorm:"not null"` // Sale | Expense | Withdrawal
	ProductID *uint
	Quantity  int
	Amount    float64 `gorm:"not null"`
	ShopID    uint    `gorm:"not null"`
	CreatedAt time.Time

	Product Product `gorm:"foreignKey:ProductID;constraint:OnDelete:SET NULL"`
	Shop    Shop    `gorm:"foreignKey:ShopID;constraint:OnDelete:CASCADE"`
}
