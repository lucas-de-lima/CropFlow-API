package routes

import (
	"strings"

	"github.com/cropflow/api/internal/adapters/http/handlers"
	"github.com/cropflow/api/internal/infrastructure/security"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all routes for the application
func SetupRoutes(
	router *gin.Engine,
	farmHandler *handlers.FarmHandler,
	cropHandler *handlers.CropHandler,
	fertilizerHandler *handlers.FertilizerHandler,
	personHandler *handlers.PersonHandler,
	authHandler *handlers.AuthHandler,
	jwtService *security.JWTService,
) {
	// Public routes
	router.POST("/persons", personHandler.CreatePerson)
	router.POST("/auth/login", authHandler.Login)

	// Farm routes
	router.POST("/farms", farmHandler.CreateFarm)
	router.GET("/farms", AuthMiddleware(jwtService, "ROLE_USER", "ROLE_MANAGER", "ROLE_ADMIN"), farmHandler.GetAllFarms)
	router.GET("/farms/:id", farmHandler.GetFarmByID)
	
	// Farm-Crop relationship routes
	router.POST("/farms/:id/crops", cropHandler.CreateCrop)
	router.GET("/farms/:id/crops", cropHandler.GetCropsByFarmID)

	// Crop routes
	router.GET("/crops", AuthMiddleware(jwtService, "ROLE_MANAGER", "ROLE_ADMIN"), cropHandler.GetAllCrops)
	router.GET("/crops/:id", cropHandler.GetCropByID)

	// Fertilizer routes
	router.POST("/fertilizers", fertilizerHandler.CreateFertilizer)
	router.GET("/fertilizers", AuthMiddleware(jwtService, "ROLE_ADMIN"), fertilizerHandler.GetAllFertilizers)
	router.GET("/fertilizers/:id", fertilizerHandler.GetFertilizerByID)
	
	// Crop-Fertilizer relationship routes (using different base path to avoid conflicts)
	router.POST("/crop/:cropId/fertilizer/:fertilizerId", cropHandler.AddFertilizerToCrop)
	router.GET("/crop/:cropId/fertilizers", cropHandler.GetFertilizersByCropID)
}

// AuthMiddleware validates JWT token and checks user roles
func AuthMiddleware(jwtService *security.JWTService, allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(401, gin.H{"error": "invalid authorization header format"})
			c.Abort()
			return
		}

		// Validate token
		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// Check if user has required role
		hasRole := false
		for _, role := range allowedRoles {
			if claims.Role == role {
				hasRole = true
				break
			}
		}

		if !hasRole {
			c.JSON(403, gin.H{"error": "insufficient permissions"})
			c.Abort()
			return
		}

		// Set user info in context
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}