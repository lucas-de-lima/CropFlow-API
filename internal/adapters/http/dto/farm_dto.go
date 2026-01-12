package dto

// FarmBodyDTO represents the request body for farm creation
type FarmBodyDTO struct {
	Name string  `json:"name" binding:"required"`
	Size float64 `json:"size" binding:"required"`
}

// FarmDTO represents the response for farm data
type FarmDTO struct {
	ID   int64   `json:"id"`
	Name string  `json:"name"`
	Size float64 `json:"size"`
}
