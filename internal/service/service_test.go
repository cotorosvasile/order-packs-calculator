package service

import (
	"pack-items/internal/components/cache"
	"pack-items/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() CalculatorService {
	cacheService := cache.NewCache()
	return NewService(cacheService)
}
func TestCalculateBoxes(t *testing.T) {
	calculator := setup()
	tests := []struct {
		boxItemsRequest entity.BoxItemsRequest
		expectedMap     entity.BoxItemsResponse
	}{
		{
			boxItemsRequest: entity.BoxItemsRequest{
				PackSizes: []int{250, 500, 1000, 2000, 5000},
				Quantity:  25,
			},
			expectedMap: entity.BoxItemsResponse{BoxItems: map[int]int{250: 1}},
		},
		{
			boxItemsRequest: entity.BoxItemsRequest{
				PackSizes: []int{250, 500, 1000, 2000, 5000},
				Quantity:  12001,
			},
			expectedMap: entity.BoxItemsResponse{BoxItems: map[int]int{250: 1, 2000: 1, 5000: 2}},
		},
		{
			boxItemsRequest: entity.BoxItemsRequest{
				PackSizes: []int{250, 500, 1000, 2000, 5000},
				Quantity:  251,
			},
			expectedMap: entity.BoxItemsResponse{BoxItems: map[int]int{500: 1}},
		},
		{
			boxItemsRequest: entity.BoxItemsRequest{
				PackSizes: []int{5, 12},
				Quantity:  14,
			},
			expectedMap: entity.BoxItemsResponse{BoxItems: map[int]int{5: 3}},
		},
	}

	for _, test := range tests {
		result := calculator.CalculateBoxes(test.boxItemsRequest)
		assert.Equal(t, result, test.expectedMap)
	}
}
