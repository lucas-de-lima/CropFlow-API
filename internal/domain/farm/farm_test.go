package farm_test

import (
	"testing"
	"time"

	"github.com/cropflow/api/internal/domain/farm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewFarm(t *testing.T) {
	t.Run("should create farm with valid data", func(t *testing.T) {
		// Arrange
		name := "Fazenda Boa Vista"
		size := 100.5

		// Act
		f, err := farm.NewFarm(name, size)

		// Assert
		require.NoError(t, err)
		assert.NotNil(t, f)
		assert.Equal(t, name, f.Name())
		assert.Equal(t, size, f.Size().Value())
		assert.NotZero(t, f.CreatedAt())
		assert.NotZero(t, f.UpdatedAt())
		assert.Equal(t, f.CreatedAt(), f.UpdatedAt())
	})

	t.Run("should return error when name is empty", func(t *testing.T) {
		// Arrange
		name := ""
		size := 100.5

		// Act
		f, err := farm.NewFarm(name, size)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, f)
		assert.Equal(t, farm.ErrInvalidFarmName, err)
	})

	t.Run("should return error when size is zero", func(t *testing.T) {
		// Arrange
		name := "Fazenda Teste"
		size := 0.0

		// Act
		f, err := farm.NewFarm(name, size)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, f)
		assert.Equal(t, farm.ErrInvalidFarmSize, err)
	})

	t.Run("should return error when size is negative", func(t *testing.T) {
		// Arrange
		name := "Fazenda Teste"
		size := -10.0

		// Act
		f, err := farm.NewFarm(name, size)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, f)
		assert.Equal(t, farm.ErrInvalidFarmSize, err)
	})
}

func TestFarm_ChangeName(t *testing.T) {
	t.Run("should change farm name successfully", func(t *testing.T) {
		// Arrange
		f, _ := farm.NewFarm("Fazenda Original", 100.0)
		originalUpdatedAt := f.UpdatedAt()
		time.Sleep(1 * time.Millisecond) // Ensure time difference
		newName := "Fazenda Nova"

		// Act
		err := f.ChangeName(newName)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, newName, f.Name())
		assert.True(t, f.UpdatedAt().After(originalUpdatedAt))
	})

	t.Run("should return error when new name is empty", func(t *testing.T) {
		// Arrange
		f, _ := farm.NewFarm("Fazenda Original", 100.0)

		// Act
		err := f.ChangeName("")

		// Assert
		assert.Error(t, err)
		assert.Equal(t, farm.ErrInvalidFarmName, err)
		assert.Equal(t, "Fazenda Original", f.Name()) // Name should not change
	})
}

func TestFarm_ChangeSize(t *testing.T) {
	t.Run("should change farm size successfully", func(t *testing.T) {
		// Arrange
		f, _ := farm.NewFarm("Fazenda Teste", 100.0)
		originalUpdatedAt := f.UpdatedAt()
		time.Sleep(1 * time.Millisecond)
		newSize := 200.5

		// Act
		err := f.ChangeSize(newSize)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, newSize, f.Size().Value())
		assert.True(t, f.UpdatedAt().After(originalUpdatedAt))
	})

	t.Run("should return error when new size is zero", func(t *testing.T) {
		// Arrange
		f, _ := farm.NewFarm("Fazenda Teste", 100.0)

		// Act
		err := f.ChangeSize(0)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, farm.ErrInvalidFarmSize, err)
		assert.Equal(t, 100.0, f.Size().Value())
	})

	t.Run("should return error when new size is negative", func(t *testing.T) {
		// Arrange
		f, _ := farm.NewFarm("Fazenda Teste", 100.0)

		// Act
		err := f.ChangeSize(-50.0)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, farm.ErrInvalidFarmSize, err)
		assert.Equal(t, 100.0, f.Size().Value())
	})
}

func TestFarm_IsValid(t *testing.T) {
	t.Run("should return true for valid farm", func(t *testing.T) {
		// Arrange
		f, _ := farm.NewFarm("Fazenda Válida", 100.0)

		// Act
		isValid := f.IsValid()

		// Assert
		assert.True(t, isValid)
	})
}

func TestRestore(t *testing.T) {
	t.Run("should restore farm from persistence", func(t *testing.T) {
		// Arrange
		id := int64(1)
		name := "Fazenda Restaurada"
		sizeValue := 150.75
		size, _ := farm.NewSize(sizeValue)
		createdAt := time.Now().Add(-24 * time.Hour)
		updatedAt := time.Now()

		// Act
		f := farm.Restore(id, name, size, createdAt, updatedAt)

		// Assert
		assert.NotNil(t, f)
		assert.Equal(t, id, f.ID())
		assert.Equal(t, name, f.Name())
		assert.Equal(t, sizeValue, f.Size().Value())
		assert.Equal(t, createdAt, f.CreatedAt())
		assert.Equal(t, updatedAt, f.UpdatedAt())
	})
}

func TestFarm_SetID(t *testing.T) {
	t.Run("should set farm ID", func(t *testing.T) {
		// Arrange
		f, _ := farm.NewFarm("Fazenda Teste", 100.0)
		assert.Zero(t, f.ID())

		// Act
		f.SetID(42)

		// Assert
		assert.Equal(t, int64(42), f.ID())
	})
}
