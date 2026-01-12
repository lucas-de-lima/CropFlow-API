package person

import (
	"time"
)

// Person represents a person aggregate root
type Person struct {
	id        int64
	username  string
	password  Password
	role      Role
	createdAt time.Time
	updatedAt time.Time
}

// NewPerson creates a new Person with validation (Factory Method)
func NewPerson(username, plainPassword, roleStr string) (*Person, error) {
	if username == "" {
		return nil, ErrInvalidUsername
	}

	if len(username) < 3 {
		return nil, ErrInvalidUsername
	}

	password, err := NewPassword(plainPassword)
	if err != nil {
		return nil, err
	}

	role, err := NewRole(roleStr)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &Person{
		username:  username,
		password:  password,
		role:      role,
		createdAt: now,
		updatedAt: now,
	}, nil
}

// Restore reconstructs a Person from persistence (used by repository)
func Restore(id int64, username, passwordHash string, role Role, createdAt, updatedAt time.Time) *Person {
	return &Person{
		id:        id,
		username:  username,
		password:  NewPasswordFromHash(passwordHash),
		role:      role,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

// Getters (encapsulation)
func (p *Person) ID() int64 {
	return p.id
}

func (p *Person) Username() string {
	return p.username
}

func (p *Person) Password() Password {
	return p.password
}

func (p *Person) Role() Role {
	return p.role
}

func (p *Person) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Person) UpdatedAt() time.Time {
	return p.updatedAt
}

// SetID is used by repository after insertion
func (p *Person) SetID(id int64) {
	p.id = id
}

// Business Methods

// Authenticate validates the provided password
func (p *Person) Authenticate(plainPassword string) error {
	if err := p.password.Compare(plainPassword); err != nil {
		return ErrInvalidCredentials
	}
	return nil
}

// ChangePassword changes the person's password
func (p *Person) ChangePassword(oldPassword, newPassword string) error {
	// Verify old password
	if err := p.password.Compare(oldPassword); err != nil {
		return ErrInvalidCredentials
	}

	// Check if new password is different
	if err := p.password.Compare(newPassword); err == nil {
		return ErrSamePassword
	}

	// Create new password
	newPass, err := NewPassword(newPassword)
	if err != nil {
		return err
	}

	p.password = newPass
	p.updatedAt = time.Now()
	return nil
}

// PromoteToRole changes the person's role
func (p *Person) PromoteToRole(newRoleStr string) error {
	newRole, err := NewRole(newRoleStr)
	if err != nil {
		return err
	}

	if p.role == newRole {
		return ErrSameRole
	}

	p.role = newRole
	p.updatedAt = time.Now()
	return nil
}

// CanAccessResource checks if this person can access a specific resource
func (p *Person) CanAccessResource(resource string) bool {
	return p.role.CanAccess(resource)
}

// HasRole checks if person has a specific role
func (p *Person) HasRole(role Role) bool {
	return p.role == role
}

// HasMinimumRole checks if person has minimum required role
func (p *Person) HasMinimumRole(minimumRole Role) bool {
	return p.role.HasPermission(minimumRole)
}

// IsValid checks if the person is in a valid state
func (p *Person) IsValid() bool {
	return p.username != "" && p.password.IsValid() && p.role != ""
}
