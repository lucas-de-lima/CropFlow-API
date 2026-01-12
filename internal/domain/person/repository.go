package person

import "context"

// Repository defines the interface for person persistence (Port)
type Repository interface {
	Save(ctx context.Context, person *Person) error
	FindByID(ctx context.Context, id int64) (*Person, error)
	FindByUsername(ctx context.Context, username string) (*Person, error)
	FindAll(ctx context.Context) ([]*Person, error)
	Delete(ctx context.Context, id int64) error
	UsernameExists(ctx context.Context, username string) (bool, error)
}
