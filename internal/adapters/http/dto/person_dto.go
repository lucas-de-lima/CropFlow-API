package dto

import "github.com/cropflow/api/internal/domain/entities"

// PersonBodyDTO represents the request body for person creation
type PersonBodyDTO struct {
	Username string         `json:"username" binding:"required"`
	Password string         `json:"password" binding:"required"`
	Role     entities.Role  `json:"role" binding:"required"`
}

// PersonDTO represents the response for person data
type PersonDTO struct {
	ID       int64         `json:"id"`
	Username string        `json:"username"`
	Role     entities.Role `json:"role"`
}

// LoginBodyDTO represents the login request body
type LoginBodyDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// TokenDTO represents the token response
type TokenDTO struct {
	Token string `json:"token"`
}

// ResponseDTO represents a generic response
type ResponseDTO struct {
	Message string `json:"message"`
}
