package entity

type BoxItemsRequest struct {
	PackSizes []int `json:"packSizes,omitempty"`
	Quantity  int   `json:"quantity"`
}

type BoxItemsResponse struct {
	BoxItems map[int]int `json:"box_items"`
}
