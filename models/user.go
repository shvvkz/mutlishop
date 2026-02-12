package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"not null"` // SuperAdmin | Admin
	ShopID    uint      `gorm:"not null"`
	CreatedAt time.Time

	Shop Shop `gorm:"foreignKey:ShopID;constraint:OnDelete:CASCADE"`
}
