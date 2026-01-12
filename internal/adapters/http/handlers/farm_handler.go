package handlers

import (
	"net/http"
	"strconv"

	"github.com/cropflow/api/internal/adapters/http/dto"
	"github.com/cropflow/api/internal/domain/entities"
	"github.com/cropflow/api/internal/usecases"
	"github.com/gin-gonic/gin"
)

// FarmHandler handles farm HTTP requests
type FarmHandler struct {
	farmUseCase *usecases.FarmUseCase
}

// NewFarmHandler creates a new farm handler
func NewFarmHandler(farmUseCase *usecases.FarmUseCase) *FarmHandler {
	return &FarmHandler{
		farmUseCase: farmUseCase,
	}
}

// CreateFarm handles POST /farms
func (h *FarmHandler) CreateFarm(c *gin.Context) {
	var body dto.FarmBodyDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	farm := &entities.Farm{
		Name: body.Name,
		Size: body.Size,
	}

	if err := h.farmUseCase.CreateFarm(farm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.FarmDTO{
		ID:   farm.ID,
		Name: farm.Name,
		Size: farm.Size,
	}

	c.JSON(http.StatusCreated, response)
}

// GetAllFarms handles GET /farms
func (h *FarmHandler) GetAllFarms(c *gin.Context) {
	farms, err := h.farmUseCase.GetAllFarms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]dto.FarmDTO, len(farms))
	for i, farm := range farms {
		response[i] = dto.FarmDTO{
			ID:   farm.ID,
			Name: farm.Name,
			Size: farm.Size,
		}
	}

	c.JSON(http.StatusOK, response)
}

// GetFarmByID handles GET /farms/:id
func (h *FarmHandler) GetFarmByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	farm, err := h.farmUseCase.GetFarmByID(id)
	if err != nil {
		if err == usecases.ErrFarmNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "farm not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.FarmDTO{
		ID:   farm.ID,
		Name: farm.Name,
		Size: farm.Size,
	}

	c.JSON(http.StatusOK, response)
}