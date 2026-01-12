package dto

// FertilizerBodyDTO represents the request body for fertilizer creation
type FertilizerBodyDTO struct {
	Name        string `json:"name" binding:"required"`
	Brand       string `json:"brand" binding:"required"`
	Composition string `json:"composition" binding:"required"`
}

// FertilizerDTO represents the response for fertilizer data
type FertilizerDTO struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Brand       string `json:"brand"`
	Composition string `json:"composition"`
}
