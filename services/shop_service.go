package services

import (
	"errors"

	"multishop/config"
	"multishop/models"
)

func UpdateShopWhatsApp(shopID uint, newNumber string) error {

	if newNumber == "" {
		return errors.New("whatsapp number cannot be empty")
	}

	var shop models.Shop

	if err := config.DB.
		Where("id = ?", shopID).
		First(&shop).Error; err != nil {

		return errors.New("shop not found")
	}

	shop.WhatsAppNumber = newNumber

	return config.DB.Save(&shop).Error
}
