package person

import "errors"

// Role represents a user role value object
type Role string

const (
	RoleUser    Role = "USER"
	RoleManager Role = "MANAGER"
	RoleAdmin   Role = "ADMIN"
)

// NewRole creates a new Role with validation
func NewRole(value string) (Role, error) {
	role := Role(value)
	switch role {
	case RoleUser, RoleManager, RoleAdmin:
		return role, nil
	default:
		return "", errors.New("invalid role: must be USER, MANAGER, or ADMIN")
	}
}

// String returns the string representation of the role
func (r Role) String() string {
	return string(r)
}

// HasPermission checks if this role has permission level equal or higher than required
func (r Role) HasPermission(required Role) bool {
	hierarchy := map[Role]int{
		RoleUser:    1,
		RoleManager: 2,
		RoleAdmin:   3,
	}
	return hierarchy[r] >= hierarchy[required]
}

// CanAccess checks if this role can access a specific resource
func (r Role) CanAccess(resource string) bool {
	permissions := map[string][]Role{
		"farms":       {RoleUser, RoleManager, RoleAdmin},
		"crops":       {RoleManager, RoleAdmin},
		"fertilizers": {RoleAdmin},
	}

	allowedRoles, exists := permissions[resource]
	if !exists {
		return false
	}

	for _, allowed := range allowedRoles {
		if r == allowed {
			return true
		}
	}
	return false
}

// IsUser checks if role is USER
func (r Role) IsUser() bool {
	return r == RoleUser
}

// IsManager checks if role is MANAGER
func (r Role) IsManager() bool {
	return r == RoleManager
}

// IsAdmin checks if role is ADMIN
func (r Role) IsAdmin() bool {
	return r == RoleAdmin
}
