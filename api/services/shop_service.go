package services

import (
	"errors"
	"time"

	"multishop/config"
	"multishop/models"

	"golang.org/x/crypto/bcrypt"
)

type ShopResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	Active         bool      `json:"active"`
	WhatsAppNumber string    `json:"whatsapp_number"`
	CreatedAt      time.Time `json:"created_at"`
}

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

func CreateShopWithSuperAdmin(
	name string,
	whatsApp string,
	email string,
	password string,
) (uint, error) {

	tx := config.DB.Begin()

	var existing models.User
	if err := tx.Where("email = ?", email).First(&existing).Error; err == nil {
		tx.Rollback()
		return 0, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		tx.Rollback()
		return 0, errors.New("failed to hash password")
	}

	shop := models.Shop{
		Name:           name,
		WhatsAppNumber: whatsApp,
		Active:         true,
	}

	if err := tx.Create(&shop).Error; err != nil {
		tx.Rollback()
		return 0, errors.New("failed to create shop")
	}

	superAdmin := models.User{
		Name:     "SuperAdmin",
		Email:    email,
		Password: string(hashedPassword),
		Role:     "SuperAdmin",
		ShopID:   shop.ID,
	}

	if err := tx.Create(&superAdmin).Error; err != nil {
		tx.Rollback()
		return 0, errors.New("failed to create superadmin")
	}

	tx.Commit()

	return shop.ID, nil
}

func GetShops() ([]ShopResponse, error) {

	var shops []models.Shop

	if err := config.DB.Find(&shops).Error; err != nil {
		return nil, errors.New("failed to retrieve shops")
	}

	var response []ShopResponse

	for _, s := range shops {
		response = append(response, ShopResponse{
			ID:             s.ID,
			Name:           s.Name,
			Active:         s.Active,
			WhatsAppNumber: s.WhatsAppNumber,
			CreatedAt:      s.CreatedAt,
		})
	}

	return response, nil
}
