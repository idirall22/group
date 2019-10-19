package group

import (
	"context"

	"github.com/idirall22/group/models"
)

// Provider interface
type Provider interface {
	New(ctx context.Context, userID int64, name string) (int64, error)
	Get(ctx context.Context, id int64, name string) (*models.Group, error)
	List(ctx context.Context, offset, limit int) ([]*models.Group, error)
	Update(ctx context.Context, id int64, name string) (int64, error)
	Delete(ctx context.Context, id, userID int64) error
	Join(ctx context.Context, userID, groupID int64) (int64, error)
	Leave(ctx context.Context, userID, groupID int64) (int64, error)
}
