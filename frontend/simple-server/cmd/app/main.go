package main

import (
	"fmt"
	"simple-server/internal/config"
	"simple-server/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.InitConfig("APP")
	if err != nil {
		panic(fmt.Sprintf("error initializing config %s", err))
	}

	router := gin.Default()
	handlers := handlers.NewHandlers(router, cfg.AppHost)
	handlers.Registry()
	router.Run(":8080")
}
