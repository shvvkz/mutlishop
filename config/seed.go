package config

import (
	"errors"
	"fmt"
	"log"
	"time"

	"multishop/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

func SeedSuperAdmins() {

	var shops []models.Shop
	if err := DB.Find(&shops).Error; err != nil {
		log.Println("Error fetching shops:", err)
		return
	}

	for _, shop := range shops {

		var existing models.User

		err := DB.
			Where("shop_id = ? AND role = ?", shop.ID, "SuperAdmin").
			First(&existing).Error

		// ✅ SuperAdmin déjà existant
		if err == nil {
			continue
		}

		// ❌ Vraie erreur DB
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Database error while checking SuperAdmin:", err)
			continue
		}

		// 🔐 Création du SuperAdmin
		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte("superadmin"),
			bcrypt.DefaultCost,
		)
		if err != nil {
			log.Println("Error hashing password:", err)
			continue
		}

		admin := models.User{
			Name:     "SuperAdmin",
			Email:    fmt.Sprintf("super%d@admin.com", shop.ID),
			Password: string(hashedPassword),
			Role:     "SuperAdmin",
			ShopID:   shop.ID,
		}

		if err := DB.Create(&admin).Error; err != nil {
			log.Println("Error creating SuperAdmin:", err)
			continue
		}

		log.Printf("SuperAdmin created for shop %d\n", shop.ID)
	}
}
