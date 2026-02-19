package public

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {

	publicGroup := r.Group("/api/")

	RegisterAuthRoutes(publicGroup)
	RegisterPublicRoutes(publicGroup)
}
