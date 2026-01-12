package person_test

import (
	"testing"
	"time"

	"github.com/cropflow/api/internal/domain/person"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPerson(t *testing.T) {
	t.Run("should create person with valid data and hashed password", func(t *testing.T) {
		// Arrange
		username := "john_doe"
		plainPassword := "SecurePassword123"
		role := "ROLE_USER"

		// Act
		p, err := person.NewPerson(username, plainPassword, role)

		// Assert
		require.NoError(t, err)
		assert.NotNil(t, p)
		assert.Equal(t, username, p.Username())
		assert.Equal(t, role, p.Role().String())
		assert.NotEqual(t, plainPassword, p.Password().Hash()) // Password should be hashed
		assert.NotZero(t, p.CreatedAt())
		assert.NotZero(t, p.UpdatedAt())
	})

	t.Run("should return error when username is empty", func(t *testing.T) {
		// Arrange
		username := ""
		plainPassword := "SecurePassword123"
		role := "ROLE_USER"

		// Act
		p, err := person.NewPerson(username, plainPassword, role)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, p)
		assert.Equal(t, person.ErrInvalidUsername, err)
	})

	t.Run("should return error when password is empty", func(t *testing.T) {
		// Arrange
		username := "john_doe"
		plainPassword := ""
		role := "ROLE_USER"

		// Act
		p, err := person.NewPerson(username, plainPassword, role)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, p)
	})

	t.Run("should return error when role is invalid", func(t *testing.T) {
		// Arrange
		username := "john_doe"
		plainPassword := "SecurePassword123"
		role := "INVALID_ROLE"

		// Act
		p, err := person.NewPerson(username, plainPassword, role)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, p)
	})

	t.Run("should create person with ROLE_MANAGER", func(t *testing.T) {
		// Arrange
		username := "manager_user"
		plainPassword := "SecurePassword123"
		role := "ROLE_MANAGER"

		// Act
		p, err := person.NewPerson(username, plainPassword, role)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, "ROLE_MANAGER", p.Role().String())
	})

	t.Run("should create person with ROLE_ADMIN", func(t *testing.T) {
		// Arrange
		username := "admin_user"
		plainPassword := "SecurePassword123"
		role := "ROLE_ADMIN"

		// Act
		p, err := person.NewPerson(username, plainPassword, role)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, "ROLE_ADMIN", p.Role().String())
	})
}

func TestPerson_Authenticate(t *testing.T) {
	t.Run("should authenticate with correct password", func(t *testing.T) {
		// Arrange
		plainPassword := "SecurePassword123"
		p, _ := person.NewPerson("john_doe", plainPassword, "ROLE_USER")

		// Act
		err := p.Authenticate(plainPassword)

		// Assert
		assert.NoError(t, err)
	})

	t.Run("should fail authentication with incorrect password", func(t *testing.T) {
		// Arrange
		plainPassword := "SecurePassword123"
		p, _ := person.NewPerson("john_doe", plainPassword, "ROLE_USER")

		// Act
		err := p.Authenticate("WrongPassword")

		// Assert
		assert.Error(t, err)
		assert.Equal(t, person.ErrInvalidCredentials, err)
	})
}

func TestPerson_ChangePassword(t *testing.T) {
	t.Run("should change password successfully", func(t *testing.T) {
		// Arrange
		oldPassword := "OldPassword123"
		newPassword := "NewPassword123"
		p, _ := person.NewPerson("john_doe", oldPassword, "ROLE_USER")
		originalUpdatedAt := p.UpdatedAt()
		time.Sleep(1 * time.Millisecond)

		// Act
		err := p.ChangePassword(oldPassword, newPassword)

		// Assert
		require.NoError(t, err)
		assert.NoError(t, p.Authenticate(newPassword))
		assert.Error(t, p.Authenticate(oldPassword))
		assert.True(t, p.UpdatedAt().After(originalUpdatedAt))
	})

	t.Run("should return error when new password is empty", func(t *testing.T) {
		// Arrange
		oldPassword := "OldPassword123"
		p, _ := person.NewPerson("john_doe", oldPassword, "ROLE_USER")

		// Act
		err := p.ChangePassword(oldPassword, "")

		// Assert
		assert.Error(t, err)
	})

	t.Run("should return error when old password is incorrect", func(t *testing.T) {
		// Arrange
		oldPassword := "OldPassword123"
		p, _ := person.NewPerson("john_doe", oldPassword, "ROLE_USER")

		// Act
		err := p.ChangePassword("WrongPassword", "NewPassword123")

		// Assert
		assert.Error(t, err)
		assert.Equal(t, person.ErrInvalidCredentials, err)
	})
}

func TestPerson_PromoteToRole(t *testing.T) {
	t.Run("should promote to role successfully", func(t *testing.T) {
		// Arrange
		p, _ := person.NewPerson("john_doe", "Password123", "ROLE_USER")
		originalUpdatedAt := p.UpdatedAt()
		time.Sleep(1 * time.Millisecond)

		// Act
		err := p.PromoteToRole("ROLE_MANAGER")

		// Assert
		require.NoError(t, err)
		assert.Equal(t, "ROLE_MANAGER", p.Role().String())
		assert.True(t, p.UpdatedAt().After(originalUpdatedAt))
	})

	t.Run("should return error for invalid role", func(t *testing.T) {
		// Arrange
		p, _ := person.NewPerson("john_doe", "Password123", "ROLE_USER")

		// Act
		err := p.PromoteToRole("INVALID_ROLE")

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "ROLE_USER", p.Role().String()) // Role should not change
	})
}

func TestPerson_IsValid(t *testing.T) {
	t.Run("should return true for valid person", func(t *testing.T) {
		// Arrange
		p, _ := person.NewPerson("john_doe", "Password123", "ROLE_USER")

		// Act
		isValid := p.IsValid()

		// Assert
		assert.True(t, isValid)
	})
}

func TestRestore(t *testing.T) {
	t.Run("should restore person from persistence", func(t *testing.T) {
		// Arrange
		id := int64(1)
		username := "restored_user"
		passwordHash := "$2a$10$hashedpassword"
		roleStr := "ROLE_ADMIN"
		role, _ := person.NewRole(roleStr)
		createdAt := time.Now().Add(-24 * time.Hour)
		updatedAt := time.Now()

		// Act
		p := person.Restore(id, username, passwordHash, role, createdAt, updatedAt)

		// Assert
		assert.NotNil(t, p)
		assert.Equal(t, id, p.ID())
		assert.Equal(t, username, p.Username())
		assert.Equal(t, passwordHash, p.Password().Hash())
		assert.Equal(t, roleStr, p.Role().String())
		assert.Equal(t, createdAt, p.CreatedAt())
		assert.Equal(t, updatedAt, p.UpdatedAt())
	})
}

func TestPerson_SetID(t *testing.T) {
	t.Run("should set person ID", func(t *testing.T) {
		// Arrange
		p, _ := person.NewPerson("john_doe", "Password123", "ROLE_USER")
		assert.Zero(t, p.ID())

		// Act
		p.SetID(42)

		// Assert
		assert.Equal(t, int64(42), p.ID())
	})
}
