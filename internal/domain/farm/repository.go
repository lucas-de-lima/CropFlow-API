package farm

import "context"

// Repository defines the interface for farm persistence (Port)
type Repository interface {
	Save(ctx context.Context, farm *Farm) error
	FindByID(ctx context.Context, id int64) (*Farm, error)
	FindAll(ctx context.Context) ([]*Farm, error)
	Delete(ctx context.Context, id int64) error
	ExistsByID(ctx context.Context, id int64) (bool, error)
}
