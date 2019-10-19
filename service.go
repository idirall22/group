package group

import (
	"database/sql"

	pr "github.com/idirall22/group/providers/postgres"
)

// Service structure
type Service struct {
	provider Provider
}

// NewService create a group service
func NewService(db *sql.DB, tableName string) *Service {
	return &Service{provider: &pr.PostgresProvider{DB: db, TableName: tableName}}
}
