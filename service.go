package group

import (
	"database/sql"

	pr "github.com/idirall22/group/providers/postgres"
)

// Service structure
type Service struct {
	provider Provider
}

// StartService start a group service
func StartService(db *sql.DB, tableName string) *Service {
	return &Service{provider: &pr.PostgresProvider{DB: db, TableName: tableName}}
}
