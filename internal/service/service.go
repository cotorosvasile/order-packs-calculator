package service

import (
	cache "pack-items/internal/components/cache"
	e "pack-items/internal/entity"
	"sort"
)

type CalculatorService interface {
	CalculateBoxes(boxItemsRequest e.BoxItemsRequest) e.BoxItemsResponse
}

type calculatorService struct {
	cache cache.Cache
}

func NewService(cache cache.Cache) CalculatorService {
	return &calculatorService{cache: cache}
}

func (c *calculatorService) CalculateBoxes(boxItemsRequest e.BoxItemsRequest) e.BoxItemsResponse {
	if boxItemsRequest.PackSizes == nil || len(boxItemsRequest.PackSizes) == 0 {
		boxItemsRequest.PackSizes = c.cache.Get()
	} else {
		c.cache.Set(boxItemsRequest.PackSizes)
	}

	result := make(map[int]int)

	sort.Sort(sort.Reverse(sort.IntSlice(boxItemsRequest.PackSizes))) // Sort items in descending order

	for _, packSize := range boxItemsRequest.PackSizes {
		boxCount := boxItemsRequest.Quantity / packSize
		boxItemsRequest.Quantity -= packSize * boxCount
		result[packSize] = boxCount
	}

	// If there's any remaining quantity, pack it in the smallest available box
	if boxItemsRequest.Quantity > 0 {
		smallestBox := boxItemsRequest.PackSizes[len(boxItemsRequest.PackSizes)-1] // Smallest box size
		result[smallestBox]++
	}

	return e.BoxItemsResponse{BoxItems: result}
}
