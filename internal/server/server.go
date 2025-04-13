package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ssaini24/Go-Project/internal/logger"
	"github.com/ssaini24/Go-Project/internal/router"
)

func Init() {
	gin.SetMode(gin.ReleaseMode)

	server := gin.New()
	server.Use(gin.Recovery())

	router.SetUpRoutes(server)
	if err := server.Run(":8081"); err != nil {
		panic(err)
	}

	logger.Infof("Server is starting at port %s", "8081")
}
