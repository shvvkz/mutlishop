package config

import (
	"log"
	"time"

	"multishop/models"
)

func SeedShops() {

	var count int64
	DB.Model(&models.Shop{}).Count(&count)

	if count > 0 {
		return
	}

	shops := []models.Shop{
		{
			ID:             1,
			Name:           "Shop One",
			Active:         true,
			WhatsAppNumber: "212600000001",
			CreatedAt:      time.Now(),
		},
		{
			ID:             2,
			Name:           "Shop Two",
			Active:         true,
			WhatsAppNumber: "212600000002",
			CreatedAt:      time.Now(),
		},
	}

	if err := DB.Create(&shops).Error; err != nil {
		log.Fatal("Failed to seed shops:", err)
	}

	log.Println("Seeded 2 default shops")
}
