package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ssaini24/Go-Project/internal/middleware"
)

func SetUpRoutes(r *gin.Engine) {
	routerGroup := r.Group("/api/v1")
	{
		routerGroup.Use(middleware.InitAuthMiddleware().Auth())
		routerGroup.GET("/healthcheck", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
	}
}
