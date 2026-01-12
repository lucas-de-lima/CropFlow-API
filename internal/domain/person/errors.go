package person

import "errors"

var (
	ErrPersonNotFound        = errors.New("person not found")
	ErrInvalidCredentials    = errors.New("invalid username or password")
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrInvalidUsername       = errors.New("invalid username: must be at least 3 characters")
	ErrInvalidPassword       = errors.New("invalid password: must be at least 8 characters")
	ErrInvalidRole           = errors.New("invalid role: must be USER, MANAGER, or ADMIN")
	ErrSamePassword          = errors.New("new password must be different from old password")
	ErrSameRole              = errors.New("person already has this role")
)
