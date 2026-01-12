package entities

import "time"

// Role represents user roles in the system
type Role string

const (
	RoleUser    Role = "ROLE_USER"
	RoleManager Role = "ROLE_MANAGER"
	RoleAdmin   Role = "ROLE_ADMIN"
)

// Person represents a person/user entity
type Person struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"` // Password is never serialized to JSON
	Role      Role      `json:"role" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// TableName overrides the default table name
func (Person) TableName() string {
	return "person"
}

// GetAuthority returns the role name for authorization
func (p *Person) GetAuthority() string {
	return string(p.Role)
}
