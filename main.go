// @title MultiShop API
// @version 1.0
// @description Backend multi-tenant for electronics shops
// @host localhost:8080
// @BasePath /

package main

import (
	"multishop/config"
	_ "multishop/docs"
	"multishop/models"
	"multishop/router"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDatabase()

	config.DB.AutoMigrate(
		&models.Shop{},
		&models.User{},
		&models.Product{},
		&models.Transaction{},
	)

	config.SeedShops()
	config.SeedSuperAdmins()

	gin.SetMode(gin.ReleaseMode)

	r := router.SetupRouter()

	r.Run(":8080")
}
