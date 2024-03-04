package routing

import (
	"net/http"

	c "pack-items/internal/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CreateCalculatorRoutes(calculatorController c.CalculatorController, engine *gin.Engine) {
	engine.Use(cors.Default())
	engine.GET("/health", healthCheck)

	calculator := engine.Group("/calculator")
	calculator.Handle(http.MethodPost, "/calculate", calculatorController.CalculateBoxItems)

	calculator.Handle(http.MethodOptions, "/calculate", optionsHandler)
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "UP"})
}

func optionsHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Content-Type", "application/json")
	c.AbortWithStatus(http.StatusNoContent)
}
