package main

import (
	"pack-items/internal/components/cache"
	"pack-items/internal/controller"
	"pack-items/internal/routing"
	"pack-items/internal/service"

	"github.com/gin-gonic/gin"
)

const (
	service_port = ":8080"
)

func main() {
	router := gin.Default()

	cache := cache.NewCache()
	calculatorService := service.NewService(cache)
	calculatorController := controller.NewCalculatorController(calculatorService)
	routing.CreateCalculatorRoutes(calculatorController, router)
	router.Run(service_port)
}
