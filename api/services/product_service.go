package services

import (
	"errors"
	"time"

	"multishop/config"
	"multishop/models"
)

type ProductResponse struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Category      string    `json:"category"`
	PurchasePrice float64   `json:"purchase_price,omitempty"`
	SellingPrice  float64   `json:"selling_price"`
	Stock         int       `json:"stock"`
	ImageURL      string    `json:"image_url"`
	CreatedAt     time.Time `json:"created_at"`
}

func GetProducts(shopID uint, role string) ([]ProductResponse, error) {

	var products []models.Product

	if err := config.DB.
		Where("shop_id = ?", shopID).
		Find(&products).Error; err != nil {
		return nil, err
	}

	var response []ProductResponse

	for _, p := range products {

		product := ProductResponse{
			ID:           p.ID,
			Name:         p.Name,
			Description:  p.Description,
			Category:     p.Category,
			SellingPrice: p.SellingPrice,
			Stock:        p.Stock,
			ImageURL:     p.ImageURL,
			CreatedAt:    p.CreatedAt,
		}

		if role == "SuperAdmin" {
			product.PurchasePrice = p.PurchasePrice
		}

		response = append(response, product)
	}

	return response, nil
}

func CreateProduct(shopID uint, input models.Product) error {

	input.ShopID = shopID

	var existing models.Product

	if err := config.DB.
		Where("shop_id = ? AND name = ?", shopID, input.Name).
		First(&existing).Error; err == nil {

		return errors.New("product name already exists")
	}

	if err := config.DB.Create(&input).Error; err != nil {
		return err
	}

	return nil
}

func UpdateProduct(shopID uint, role string, id string, input models.Product) error {

	var product models.Product

	if err := config.DB.
		Where("id = ? AND shop_id = ?", id, shopID).
		First(&product).Error; err != nil {
		return errors.New("product not found")
	}

	var existing models.Product

	if err := config.DB.
		Where("shop_id = ? AND name = ? AND id <> ?", shopID, input.Name, id).
		First(&existing).Error; err == nil {

		return errors.New("product name already exists")
	}

	product.Name = input.Name
	product.Description = input.Description
	product.Category = input.Category
	product.SellingPrice = input.SellingPrice
	product.Stock = input.Stock
	product.ImageURL = input.ImageURL

	if role == "SuperAdmin" {
		product.PurchasePrice = input.PurchasePrice
	}

	if err := config.DB.Save(&product).Error; err != nil {
		return err
	}

	return nil
}

func DeleteProduct(shopID uint, id string) error {

	result := config.DB.
		Where("id = ? AND shop_id = ?", id, shopID).
		Delete(&models.Product{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}
