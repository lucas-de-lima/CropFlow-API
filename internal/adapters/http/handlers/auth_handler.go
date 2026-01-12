package handlers

import (
	"net/http"

	"github.com/cropflow/api/internal/adapters/http/dto"
	"github.com/cropflow/api/internal/usecases"
	"github.com/gin-gonic/gin"
)

// AuthHandler handles authentication HTTP requests
type AuthHandler struct {
	authUseCase *usecases.AuthUseCase
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authUseCase *usecases.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

// Login handles POST /auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	var body dto.LoginBodyDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authUseCase.Login(body.Username, body.Password)
	if err != nil {
		if err == usecases.ErrInvalidCredentials {
			c.JSON(http.StatusForbidden, gin.H{"error": "invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.TokenDTO{
		Token: token,
	}

	c.JSON(http.StatusOK, response)
}
