package services

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"multishop/config"
	"multishop/models"
)

type PublicProductResponse struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Category     string    `json:"category"`
	SellingPrice float64   `json:"selling_price"`
	Stock        int       `json:"stock"`
	ImageURL     string    `json:"image_url"`
	CreatedAt    time.Time `json:"created_at"`
	Status       string    `json:"status"`
}

func GetPublicProducts(shopID uint) ([]PublicProductResponse, error) {

	var shop models.Shop
	if err := config.DB.First(&shop, shopID).Error; err != nil {
		return nil, errors.New("shop not found")
	}

	if !shop.Active {
		return nil, errors.New("shop is not active")
	}

	var products []models.Product

	if err := config.DB.
		Where("shop_id = ?", shopID).
		Find(&products).Error; err != nil {
		return nil, err
	}

	var response []PublicProductResponse

	for _, p := range products {

		status := "Available"
		if p.Stock == 0 {
			status = "Out of stock"
		}

		response = append(response, PublicProductResponse{
			ID:           p.ID,
			Name:         p.Name,
			Description:  p.Description,
			Category:     p.Category,
			SellingPrice: p.SellingPrice,
			Stock:        p.Stock,
			ImageURL:     p.ImageURL,
			CreatedAt:    p.CreatedAt,
			Status:       status,
		})
	}

	return response, nil
}

func GenerateWhatsAppLink(shopID uint, productID uint) (string, error) {

	var shop models.Shop
	if err := config.DB.First(&shop, shopID).Error; err != nil {
		return "", errors.New("shop not found")
	}

	var product models.Product
	if err := config.DB.
		Where("id = ? AND shop_id = ?", productID, shopID).
		First(&product).Error; err != nil {
		return "", errors.New("product not found")
	}

	message := url.QueryEscape(
		fmt.Sprintf("Bonjour je veux plus d'information sur %s", product.Name),
	)

	link := fmt.Sprintf(
		"https://wa.me/%s?text=%s",
		shop.WhatsAppNumber,
		message,
	)

	return link, nil
}
