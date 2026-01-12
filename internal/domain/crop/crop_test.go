package crop_test

import (
	"testing"
	"time"

	"github.com/cropflow/api/internal/domain/crop"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCrop(t *testing.T) {
	t.Run("should create crop with valid data", func(t *testing.T) {
		// Arrange
		name := "Milho"
		plantedArea := 50.5
		farmID := int64(1)
		plantedDate := time.Now()
		harvestDate := plantedDate.AddDate(0, 6, 0)

		// Act
		c, err := crop.NewCrop(name, plantedArea, farmID, &plantedDate, &harvestDate)

		// Assert
		require.NoError(t, err)
		assert.NotNil(t, c)
		assert.Equal(t, name, c.Name())
		assert.Equal(t, plantedArea, c.PlantedArea())
		assert.Equal(t, farmID, c.FarmID())
		assert.Equal(t, &plantedDate, c.PlantedDate())
		assert.Equal(t, &harvestDate, c.HarvestDate())
		assert.NotZero(t, c.CreatedAt())
		assert.NotZero(t, c.UpdatedAt())
	})

	t.Run("should return error when name is empty", func(t *testing.T) {
		// Arrange
		name := ""
		plantedArea := 50.5
		farmID := int64(1)
		plantedDate := time.Now()
		harvestDate := plantedDate.AddDate(0, 6, 0)

		// Act
		c, err := crop.NewCrop(name, plantedArea, farmID, &plantedDate, &harvestDate)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, c)
		assert.Equal(t, crop.ErrInvalidCropName, err)
	})

	t.Run("should return error when planted area is zero", func(t *testing.T) {
		// Arrange
		name := "Milho"
		plantedArea := 0.0
		farmID := int64(1)
		plantedDate := time.Now()
		harvestDate := plantedDate.AddDate(0, 6, 0)

		// Act
		c, err := crop.NewCrop(name, plantedArea, farmID, &plantedDate, &harvestDate)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, c)
		assert.Equal(t, crop.ErrInvalidPlantedArea, err)
	})

	t.Run("should return error when planted area is negative", func(t *testing.T) {
		// Arrange
		name := "Milho"
		plantedArea := -10.0
		farmID := int64(1)
		plantedDate := time.Now()
		harvestDate := plantedDate.AddDate(0, 6, 0)

		// Act
		c, err := crop.NewCrop(name, plantedArea, farmID, &plantedDate, &harvestDate)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, c)
		assert.Equal(t, crop.ErrInvalidPlantedArea, err)
	})

	t.Run("should return error when farm ID is zero", func(t *testing.T) {
		// Arrange
		name := "Milho"
		plantedArea := 50.5
		farmID := int64(0)
		plantedDate := time.Now()
		harvestDate := plantedDate.AddDate(0, 6, 0)

		// Act
		c, err := crop.NewCrop(name, plantedArea, farmID, &plantedDate, &harvestDate)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, c)
		assert.Equal(t, crop.ErrInvalidFarmID, err)
	})

	t.Run("should return error when harvest date is before planted date", func(t *testing.T) {
		// Arrange
		name := "Milho"
		plantedArea := 50.5
		farmID := int64(1)
		plantedDate := time.Now()
		harvestDate := plantedDate.AddDate(0, -1, 0) // 1 month before

		// Act
		c, err := crop.NewCrop(name, plantedArea, farmID, &plantedDate, &harvestDate)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, c)
		assert.Equal(t, crop.ErrInvalidHarvestDate, err)
	})
}

func TestCrop_ChangeName(t *testing.T) {
	t.Run("should change crop name successfully", func(t *testing.T) {
		// Arrange
		plantedDate := time.Now()
		harvestDate := plantedDate.AddDate(0, 6, 0)
		c, _ := crop.NewCrop("Milho", 50.0, 1, &plantedDate, &harvestDate)
		originalUpdatedAt := c.UpdatedAt()
		time.Sleep(1 * time.Millisecond)
		newName := "Soja"

		// Act
		err := c.ChangeName(newName)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, newName, c.Name())
		assert.True(t, c.UpdatedAt().After(originalUpdatedAt))
	})

	t.Run("should return error when new name is empty", func(t *testing.T) {
		// Arrange
		plantedDate := time.Now()
		harvestDate := plantedDate.AddDate(0, 6, 0)
		c, _ := crop.NewCrop("Milho", 50.0, 1, &plantedDate, &harvestDate)

		// Act
		err := c.ChangeName("")

		// Assert
		assert.Error(t, err)
		assert.Equal(t, crop.ErrInvalidCropName, err)
		assert.Equal(t, "Milho", c.Name())
	})
}

func TestCrop_ChangePlantedArea(t *testing.T) {
	t.Run("should change planted area successfully", func(t *testing.T) {
		// Arrange
		plantedDate := time.Now()
		harvestDate := plantedDate.AddDate(0, 6, 0)
		c, _ := crop.NewCrop("Milho", 50.0, 1, &plantedDate, &harvestDate)
		originalUpdatedAt := c.UpdatedAt()
		time.Sleep(1 * time.Millisecond)
		newArea := 75.5

		// Act
		err := c.ChangePlantedArea(newArea)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, newArea, c.PlantedArea())
		assert.True(t, c.UpdatedAt().After(originalUpdatedAt))
	})

	t.Run("should return error when new area is zero", func(t *testing.T) {
		// Arrange
		plantedDate := time.Now()
		harvestDate := plantedDate.AddDate(0, 6, 0)
		c, _ := crop.NewCrop("Milho", 50.0, 1, &plantedDate, &harvestDate)

		// Act
		err := c.ChangePlantedArea(0)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, crop.ErrInvalidPlantedArea, err)
		assert.Equal(t, 50.0, c.PlantedArea())
	})
}

func TestCrop_ChangeHarvestDate(t *testing.T) {
	t.Run("should change harvest date successfully", func(t *testing.T) {
		// Arrange
		plantedDate := time.Now()
		harvestDate := plantedDate.AddDate(0, 6, 0)
		c, _ := crop.NewCrop("Milho", 50.0, 1, &plantedDate, &harvestDate)
		originalUpdatedAt := c.UpdatedAt()
		time.Sleep(1 * time.Millisecond)
		newHarvestDate := plantedDate.AddDate(0, 8, 0)

		// Act
		err := c.ChangeHarvestDate(&newHarvestDate)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, &newHarvestDate, c.HarvestDate())
		assert.True(t, c.UpdatedAt().After(originalUpdatedAt))
	})

	t.Run("should return error when new harvest date is before planted date", func(t *testing.T) {
		// Arrange
		plantedDate := time.Now()
		harvestDate := plantedDate.AddDate(0, 6, 0)
		c, _ := crop.NewCrop("Milho", 50.0, 1, &plantedDate, &harvestDate)
		newHarvestDate := plantedDate.AddDate(0, -1, 0)

		// Act
		err := c.ChangeHarvestDate(&newHarvestDate)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, crop.ErrInvalidHarvestDate, err)
		assert.Equal(t, &harvestDate, c.HarvestDate())
	})
}

func TestCrop_IsValid(t *testing.T) {
	t.Run("should return true for valid crop", func(t *testing.T) {
		// Arrange
		plantedDate := time.Now()
		harvestDate := plantedDate.AddDate(0, 6, 0)
		c, _ := crop.NewCrop("Milho", 50.0, 1, &plantedDate, &harvestDate)

		// Act
		isValid := c.IsValid()

		// Assert
		assert.True(t, isValid)
	})
}

func TestRestore(t *testing.T) {
	t.Run("should restore crop from persistence", func(t *testing.T) {
		// Arrange
		id := int64(1)
		name := "Soja Restaurada"
		plantedArea := 100.5
		farmID := int64(5)
		plantedDate := time.Now().Add(-30 * 24 * time.Hour)
		harvestDate := time.Now().Add(150 * 24 * time.Hour)
		createdAt := time.Now().Add(-25 * 24 * time.Hour)
		updatedAt := time.Now()

		// Act
		c := crop.Restore(id, name, plantedArea, farmID, plantedDate, harvestDate, createdAt, updatedAt)

		// Assert
		assert.NotNil(t, c)
		assert.Equal(t, id, c.ID())
		assert.Equal(t, name, c.Name())
		assert.Equal(t, plantedArea, c.PlantedArea())
		assert.Equal(t, farmID, c.FarmID())
		assert.Equal(t, &plantedDate, c.PlantedDate())
		assert.Equal(t, &harvestDate, c.HarvestDate())
		assert.Equal(t, createdAt, c.CreatedAt())
		assert.Equal(t, updatedAt, c.UpdatedAt())
	})
}

func TestCrop_SetID(t *testing.T) {
	t.Run("should set crop ID", func(t *testing.T) {
		// Arrange
		plantedDate := time.Now()
		harvestDate := plantedDate.AddDate(0, 6, 0)
		c, _ := crop.NewCrop("Milho", 50.0, 1, &plantedDate, &harvestDate)
		assert.Zero(t, c.ID())

		// Act
		c.SetID(42)

		// Assert
		assert.Equal(t, int64(42), c.ID())
	})
}
