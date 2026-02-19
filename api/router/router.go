package router

import (
	"multishop/router/private"
	"multishop/router/public"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	// Servir l'index à la racine et les assets statiques sous /static
	r.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})
	// Les assets (CSS/JS/images) produits par Vite seront dans ./public/static
	r.Static("/static", "./public/static")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	public.RegisterRoutes(r)
	private.RegisterRoutes(r)

	// SPA fallback: pour toutes les routes non-API/non-swagger, renvoyer index.html
	r.NoRoute(func(c *gin.Context) {
		p := c.Request.URL.Path
		if strings.HasPrefix(p, "/api") || strings.HasPrefix(p, "/swagger") || strings.HasPrefix(p, "/docs") {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.File("./public/index.html")
	})

	return r
}
