package services

import (
	"errors"
	"time"

	"multishop/config"
	"multishop/models"
)

type TransactionResponse struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	ProductID uint      `json:"product_id,omitempty"`
	Quantity  int       `json:"quantity,omitempty"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func GetTransactions(shopID uint) ([]TransactionResponse, error) {

	var transactions []models.Transaction

	if err := config.DB.
		Where("shop_id = ?", shopID).
		Find(&transactions).Error; err != nil {
		return nil, err
	}

	var response []TransactionResponse

	for _, t := range transactions {

		var productID uint
		if t.ProductID != nil {
			productID = *t.ProductID
		}
		response = append(response, TransactionResponse{
			ID:        t.ID,
			Type:      t.Type,
			ProductID: productID,
			Quantity:  t.Quantity,
			Amount:    t.Amount,
			CreatedAt: t.CreatedAt,
		})
	}

	return response, nil
}

func CreateTransaction(shopID uint, input models.Transaction) error {

	input.ShopID = shopID

	// Vérification type valide
	if input.Type != "Sale" && input.Type != "Expense" && input.Type != "Withdrawal" {
		return errors.New("invalid transaction type")
	}

	// Si Sale → gérer stock
	if input.Type == "Sale" {

		var product models.Product

		if err := config.DB.
			Where("id = ? AND shop_id = ?", input.ProductID, shopID).
			First(&product).Error; err != nil {

			return errors.New("product not found")
		}

		if product.Stock < input.Quantity {
			return errors.New("insufficient stock")
		}

		// Décrémenter stock
		product.Stock -= input.Quantity

		if err := config.DB.Save(&product).Error; err != nil {
			return err
		}

		// Calcul automatique montant
		input.Amount = float64(input.Quantity) * product.SellingPrice
	}

	if err := config.DB.Create(&input).Error; err != nil {
		return err
	}

	return nil
}

func DeleteTransaction(shopID uint, id string) error {

	var transaction models.Transaction

	if err := config.DB.
		Where("id = ? AND shop_id = ?", id, shopID).
		First(&transaction).Error; err != nil {

		return errors.New("transaction not found")
	}

	// Si c'était une Sale → remettre le stock
	if transaction.Type == "Sale" {

		var product models.Product

		if err := config.DB.First(&product, transaction.ProductID).Error; err == nil {

			product.Stock += transaction.Quantity
			config.DB.Save(&product)
		}
	}

	if err := config.DB.Delete(&transaction).Error; err != nil {
		return err
	}

	return nil
}
