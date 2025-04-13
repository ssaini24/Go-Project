package main

import (
	"github.com/ssaini24/Go-Project/internal/config"
	"github.com/ssaini24/Go-Project/internal/logger"
	"github.com/ssaini24/Go-Project/internal/server"
)

func main() {
	config.Init()
	logger.Init()
	server.Init()
}
