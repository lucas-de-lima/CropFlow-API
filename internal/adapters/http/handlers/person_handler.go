package handlers

import (
	"net/http"
	"strconv"

	"github.com/cropflow/api/internal/adapters/http/dto"
	"github.com/cropflow/api/internal/domain/entities"
	"github.com/cropflow/api/internal/usecases"
	"github.com/gin-gonic/gin"
)

// PersonHandler handles person HTTP requests
type PersonHandler struct {
	personUseCase *usecases.PersonUseCase
}

// NewPersonHandler creates a new person handler
func NewPersonHandler(personUseCase *usecases.PersonUseCase) *PersonHandler {
	return &PersonHandler{
		personUseCase: personUseCase,
	}
}

// CreatePerson handles POST /persons
func (h *PersonHandler) CreatePerson(c *gin.Context) {
	var body dto.PersonBodyDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person := &entities.Person{
		Username: body.Username,
		Password: body.Password,
		Role:     body.Role,
	}

	if err := h.personUseCase.CreatePerson(person); err != nil {
		if err == usecases.ErrUsernameAlreadyUsed {
			c.JSON(http.StatusConflict, gin.H{"error": "username already in use"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.PersonDTO{
		ID:       person.ID,
		Username: person.Username,
		Role:     person.Role,
	}

	c.JSON(http.StatusCreated, response)
}

// GetAllPersons handles GET /persons
func (h *PersonHandler) GetAllPersons(c *gin.Context) {
	persons, err := h.personUseCase.GetAllPersons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]dto.PersonDTO, len(persons))
	for i, person := range persons {
		response[i] = dto.PersonDTO{
			ID:       person.ID,
			Username: person.Username,
			Role:     person.Role,
		}
	}

	c.JSON(http.StatusOK, response)
}

// GetPersonByID handles GET /persons/:id
func (h *PersonHandler) GetPersonByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	person, err := h.personUseCase.GetPersonByID(id)
	if err != nil {
		if err == usecases.ErrPersonNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "person not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.PersonDTO{
		ID:       person.ID,
		Username: person.Username,
		Role:     person.Role,
	}

	c.JSON(http.StatusOK, response)
}
