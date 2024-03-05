package service

import (
	"math"
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

	sort.Sort(sort.Reverse(sort.IntSlice(boxItemsRequest.PackSizes)))

	leastBoxes := math.MaxInt64
	leastExtra := math.MaxInt64

	c.calculateRecursivellyBoxes(0, 0, make(map[int]int), 0, &result, &leastBoxes, &leastExtra, boxItemsRequest)

	return e.BoxItemsResponse{
		BoxItems: result,
	}
}

func (c *calculatorService) calculateRecursivellyBoxes(index, quantity int, packConfig map[int]int, currentBoxes int, result *map[int]int, leastBoxes *int, leastExtra *int, boxItemsRequest e.BoxItemsRequest) {
	if quantity >= boxItemsRequest.Quantity {
		extra := quantity - boxItemsRequest.Quantity
		if extra < *leastExtra || (extra == *leastExtra && currentBoxes < *leastBoxes) {
			*leastExtra = extra
			*leastBoxes = currentBoxes
			*result = copyMap(packConfig)
		}
		return
	}

	if index == len(boxItemsRequest.PackSizes) {
		return
	}

	c.calculateRecursivellyBoxes(index+1, quantity, packConfig, currentBoxes, result, leastBoxes, leastExtra, boxItemsRequest)

	updatedPackConfig := copyMap(packConfig)
	updatedPackConfig[boxItemsRequest.PackSizes[index]]++
	c.calculateRecursivellyBoxes(index, quantity+boxItemsRequest.PackSizes[index], updatedPackConfig, currentBoxes+1, result, leastBoxes, leastExtra, boxItemsRequest)
}

func copyMap(original map[int]int) map[int]int {
	copy := make(map[int]int)
	for k, v := range original {
		copy[k] = v
	}
	return copy
}
