package handlers

import (
	"net/http"
	"strconv"

	"github.com/cropflow/api/internal/adapters/http/dto"
	"github.com/cropflow/api/internal/domain/entities"
	"github.com/cropflow/api/internal/usecases"
	"github.com/gin-gonic/gin"
)

// FertilizerHandler handles fertilizer HTTP requests
type FertilizerHandler struct {
	fertilizerUseCase *usecases.FertilizerUseCase
}

// NewFertilizerHandler creates a new fertilizer handler
func NewFertilizerHandler(fertilizerUseCase *usecases.FertilizerUseCase) *FertilizerHandler {
	return &FertilizerHandler{
		fertilizerUseCase: fertilizerUseCase,
	}
}

// CreateFertilizer handles POST /fertilizers
func (h *FertilizerHandler) CreateFertilizer(c *gin.Context) {
	var body dto.FertilizerBodyDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fertilizer := &entities.Fertilizer{
		Name:        body.Name,
		Brand:       body.Brand,
		Composition: body.Composition,
	}

	if err := h.fertilizerUseCase.CreateFertilizer(fertilizer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.FertilizerDTO{
		ID:          fertilizer.ID,
		Name:        fertilizer.Name,
		Brand:       fertilizer.Brand,
		Composition: fertilizer.Composition,
	}

	c.JSON(http.StatusCreated, response)
}

// GetAllFertilizers handles GET /fertilizers
func (h *FertilizerHandler) GetAllFertilizers(c *gin.Context) {
	fertilizers, err := h.fertilizerUseCase.GetAllFertilizers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]dto.FertilizerDTO, len(fertilizers))
	for i, fertilizer := range fertilizers {
		response[i] = dto.FertilizerDTO{
			ID:          fertilizer.ID,
			Name:        fertilizer.Name,
			Brand:       fertilizer.Brand,
			Composition: fertilizer.Composition,
		}
	}

	c.JSON(http.StatusOK, response)
}

// GetFertilizerByID handles GET /fertilizers/:id
func (h *FertilizerHandler) GetFertilizerByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	fertilizer, err := h.fertilizerUseCase.GetFertilizerByID(id)
	if err != nil {
		if err == usecases.ErrFertilizerNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "fertilizer not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.FertilizerDTO{
		ID:          fertilizer.ID,
		Name:        fertilizer.Name,
		Brand:       fertilizer.Brand,
		Composition: fertilizer.Composition,
	}

	c.JSON(http.StatusOK, response)
}