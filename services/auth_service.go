package services

import (
	"errors"

	"multishop/config"
	"multishop/models"
	"multishop/utils"
)

func Register(name, email, password, role string, shopID uint) error {

	// Vérifier que le shop existe
	var shop models.Shop
	if err := config.DB.First(&shop, shopID).Error; err != nil {
		return errors.New("shop does not exist")
	}

	// Vérifier rôle valide
	if role != "SuperAdmin" && role != "Admin" {
		return errors.New("invalid role")
	}

	// Vérifier si email existe déjà
	var existing models.User
	if err := config.DB.Where("email = ?", email).First(&existing).Error; err == nil {
		return errors.New("email already exists")
	}

	hashed, err := utils.HashPassword(password)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: hashed,
		Role:     role,
		ShopID:   shopID,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func Login(email, password string) (string, error) {

	var user models.User

	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := utils.CheckPassword(user.Password, password); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.ID, user.ShopID, user.Role)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
