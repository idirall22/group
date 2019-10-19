package provider

import (
	"context"
	"database/sql"

	"github.com/idirall22/group/models"
)

// PostgresProvider structure
type PostgresProvider struct {
	DB        *sql.DB
	TableName string
}

// New add a group
func (s *PostgresProvider) New(ctx context.Context, userID int64, name string) (*models.Group, error) {
	return nil, nil
}

// Get get a group
func (s *PostgresProvider) Get(ctx context.Context, id int64, name string) (*models.Group, error) {
	return nil, nil
}

// List get a list of groups
func (s *PostgresProvider) List(ctx context.Context, offset, limit int) ([]*models.Group, error) {
	return nil, nil
}

// Update update a group
func (s *PostgresProvider) Update(ctx context.Context, id int64, name string) (*models.Group, error) {
	return nil, nil
}

// Delete delete a group
func (s *PostgresProvider) Delete(ctx context.Context, id, userID int64) error {
	return nil
}

// Join join a group
func (s *PostgresProvider) Join(ctx context.Context, userID, groupID int64) (int64, error) {
	return 0, nil
}

// Leave leave a group
func (s *PostgresProvider) Leave(ctx context.Context, userID, groupID int64) (int64, error) {
	return 0, nil
}
