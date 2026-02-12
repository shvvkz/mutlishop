package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	ShopID uint   `json:"shop_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint, shopID uint, role string) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))

	expireHours := 24
	expirationTime := time.Now().Add(time.Hour * time.Duration(expireHours))

	claims := &JWTClaims{
		UserID: userID,
		ShopID: shopID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}
