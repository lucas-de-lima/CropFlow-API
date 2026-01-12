package main

import (
	"log"
	"os"

	"github.com/cropflow/api/config"
	"github.com/cropflow/api/internal/adapters/database/mysql"
	"github.com/cropflow/api/internal/adapters/http/handlers"
	"github.com/cropflow/api/internal/adapters/http/routes"
	"github.com/cropflow/api/internal/infrastructure/security"
	"github.com/cropflow/api/internal/usecases"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize configuration
	cfg := config.NewConfig()

	// Initialize database
	db, err := mysql.NewMySQLConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	if err := mysql.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize repositories
	farmRepo := mysql.NewFarmRepository(db)
	cropRepo := mysql.NewCropRepository(db)
	fertilizerRepo := mysql.NewFertilizerRepository(db)
	personRepo := mysql.NewPersonRepository(db)

	// Initialize security services
	jwtService := security.NewJWTService(cfg.JWTSecret, cfg.JWTIssuer)
	passwordService := security.NewPasswordService()

	// Initialize use cases
	farmUseCase := usecases.NewFarmUseCase(farmRepo)
	cropUseCase := usecases.NewCropUseCase(cropRepo, farmRepo, fertilizerRepo)
	fertilizerUseCase := usecases.NewFertilizerUseCase(fertilizerRepo)
	personUseCase := usecases.NewPersonUseCase(personRepo, passwordService)
	authUseCase := usecases.NewAuthUseCase(personRepo, passwordService, jwtService)

	// Initialize handlers
	farmHandler := handlers.NewFarmHandler(farmUseCase)
	cropHandler := handlers.NewCropHandler(cropUseCase)
	fertilizerHandler := handlers.NewFertilizerHandler(fertilizerUseCase)
	personHandler := handlers.NewPersonHandler(personUseCase)
	authHandler := handlers.NewAuthHandler(authUseCase)

	// Setup router
	router := gin.Default()
	routes.SetupRoutes(router, farmHandler, cropHandler, fertilizerHandler, personHandler, authHandler, jwtService)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
