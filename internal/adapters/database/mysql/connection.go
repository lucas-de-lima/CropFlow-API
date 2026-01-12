package mysql

import (
	"fmt"

	"github.com/cropflow/api/config"
	"github.com/cropflow/api/internal/domain/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewMySQLConnection creates a new MySQL database connection
func NewMySQLConnection(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

// RunMigrations runs database migrations
func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&entities.Farm{},
		&entities.Crop{},
		&entities.Fertilizer{},
		&entities.Person{},
	)
}
