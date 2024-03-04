package controller

import (
	e "pack-items/internal/entity"
	s "pack-items/internal/service"

	"github.com/gin-gonic/gin"
)

type CalculatorController interface {
	CalculateBoxItems(ctx *gin.Context)
}

type calculatorController struct {
	calculatorService s.CalculatorService
}

func NewCalculatorController(calculatorService s.CalculatorService) CalculatorController {
	return &calculatorController{
		calculatorService: calculatorService,
	}
}

func (c *calculatorController) CalculateBoxItems(ctx *gin.Context) {
	var boxItemsRequest e.BoxItemsRequest
	if err := ctx.BindJSON(&boxItemsRequest); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := c.calculatorService.CalculateBoxes(boxItemsRequest)
	ctx.JSON(200, result)
}
