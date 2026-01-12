package handlers

import (
	"net/http"
	"strconv"

	"github.com/cropflow/api/internal/adapters/http/dto"
	"github.com/cropflow/api/internal/domain/entities"
	"github.com/cropflow/api/internal/usecases"
	"github.com/gin-gonic/gin"
)

// CropHandler handles crop HTTP requests
type CropHandler struct {
	cropUseCase *usecases.CropUseCase
}

// NewCropHandler creates a new crop handler
func NewCropHandler(cropUseCase *usecases.CropUseCase) *CropHandler {
	return &CropHandler{
		cropUseCase: cropUseCase,
	}
}

// CreateCrop handles POST /farms/:id/crops
func (h *CropHandler) CreateCrop(c *gin.Context) {
	farmID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid farm id"})
		return
	}

	var body dto.CropBodyDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	crop := &entities.Crop{
		Name:        body.Name,
		PlantedArea: body.PlantedArea,
		FarmID:      farmID,
		PlantedDate: body.PlantedDate,
		HarvestDate: body.HarvestDate,
	}

	if err := h.cropUseCase.CreateCrop(crop); err != nil {
		if err == usecases.ErrFarmNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "farm not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.CropDTO{
		ID:          crop.ID,
		Name:        crop.Name,
		PlantedArea: crop.PlantedArea,
		FarmID:      crop.FarmID,
		PlantedDate: crop.PlantedDate,
		HarvestDate: crop.HarvestDate,
	}

	c.JSON(http.StatusCreated, response)
}

// GetAllCrops handles GET /crops
func (h *CropHandler) GetAllCrops(c *gin.Context) {
	crops, err := h.cropUseCase.GetAllCrops()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]dto.CropDTO, len(crops))
	for i, crop := range crops {
		response[i] = dto.CropDTO{
			ID:          crop.ID,
			Name:        crop.Name,
			PlantedArea: crop.PlantedArea,
			FarmID:      crop.FarmID,
			PlantedDate: crop.PlantedDate,
			HarvestDate: crop.HarvestDate,
		}
	}

	c.JSON(http.StatusOK, response)
}

// GetCropByID handles GET /crops/:id
func (h *CropHandler) GetCropByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	crop, err := h.cropUseCase.GetCropByID(id)
	if err != nil {
		if err == usecases.ErrCropNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "crop not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.CropDTO{
		ID:          crop.ID,
		Name:        crop.Name,
		PlantedArea: crop.PlantedArea,
		FarmID:      crop.FarmID,
		PlantedDate: crop.PlantedDate,
		HarvestDate: crop.HarvestDate,
	}

	c.JSON(http.StatusOK, response)
}

// GetCropsByFarmID handles GET /farms/:id/crops
func (h *CropHandler) GetCropsByFarmID(c *gin.Context) {
	farmID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid farm id"})
		return
	}

	crops, err := h.cropUseCase.GetCropsByFarmID(farmID)
	if err != nil {
		if err == usecases.ErrFarmNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "farm not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]dto.CropDTO, len(crops))
	for i, crop := range crops {
		response[i] = dto.CropDTO{
			ID:          crop.ID,
			Name:        crop.Name,
			PlantedArea: crop.PlantedArea,
			FarmID:      crop.FarmID,
			PlantedDate: crop.PlantedDate,
			HarvestDate: crop.HarvestDate,
		}
	}

	c.JSON(http.StatusOK, response)
}

// AddFertilizerToCrop handles POST /crops/:cropId/fertilizers/:fertilizerId
func (h *CropHandler) AddFertilizerToCrop(c *gin.Context) {
	cropID, err := strconv.ParseInt(c.Param("cropId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid crop id"})
		return
	}

	fertilizerID, err := strconv.ParseInt(c.Param("fertilizerId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid fertilizer id"})
		return
	}

	if err := h.cropUseCase.AddFertilizerToCrop(cropID, fertilizerID); err != nil {
		if err == usecases.ErrCropNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "crop not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Fertilizante associado à plantação com sucesso"})
}

// GetFertilizersByCropID handles GET /crops/:cropId/fertilizers
func (h *CropHandler) GetFertilizersByCropID(c *gin.Context) {
	cropID, err := strconv.ParseInt(c.Param("cropId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid crop id"})
		return
	}

	fertilizers, err := h.cropUseCase.GetFertilizersByCropID(cropID)
	if err != nil {
		if err == usecases.ErrCropNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "crop not found"})
			return
		}
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