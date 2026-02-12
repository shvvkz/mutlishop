package services

import (
	"errors"

	"multishop/config"
	"multishop/models"

	"golang.org/x/crypto/bcrypt"
)

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func GetUsers(shopID uint) ([]UserResponse, error) {

	var users []models.User

	if err := config.DB.
		Where("shop_id = ?", shopID).
		Find(&users).Error; err != nil {

		return nil, err
	}

	var response []UserResponse

	for _, u := range users {
		response = append(response, UserResponse{
			ID:    u.ID,
			Email: u.Email,
			Role:  u.Role,
		})
	}

	return response, nil
}

func UpdateUserRole(shopID uint, userID uint, role string) error {

	if role != "Admin" && role != "SuperAdmin" {
		return errors.New("invalid role")
	}

	var user models.User

	if err := config.DB.
		Where("id = ? AND shop_id = ?", userID, shopID).
		First(&user).Error; err != nil {

		return errors.New("user not found")
	}

	user.Role = role

	return config.DB.Save(&user).Error
}

func DeleteUser(shopID uint, userID uint) error {

	result := config.DB.
		Where("id = ? AND shop_id = ?", userID, shopID).
		Delete(&models.User{})

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return result.Error
}

func RegisterUser(user models.User) error {

	hashed, _ := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	user.Password = string(hashed)

	return config.DB.Create(&user).Error
}
