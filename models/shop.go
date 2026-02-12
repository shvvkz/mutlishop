package models

import "time"

type Shop struct {
	ID             uint      `gorm:"primaryKey"`
	Name           string    `gorm:"not null"`
	Active         bool      `gorm:"default:true"`
	WhatsAppNumber string    `gorm:"not null"`
	CreatedAt      time.Time

	Users        []User
	Products     []Product
	Transactions []Transaction
}
