package group

import (
	"database/sql"
	"fmt"
	"testing"

	pr "github.com/idirall22/group/providers/postgres"
	"github.com/idirall22/utilities"
	_ "github.com/lib/pq"
)

var (
	testService *Service
	database    *sql.DB
	testToken   string
	tableName   = "groups"
	query       = fmt.Sprintf(`
	DROP TABLE IF EXISTS %s;

	CREATE TABLE IF NOT EXISTS %s(
		id SERIAL PRIMARY KEY,
		name VARCHAR NOT NULL,
		admin_id INTEGER NOT NULL,
		users_ids INTEGER[] default '{}'::int[],
		created_at TIMESTAMP with TIME ZONE DEFAULT now(),
		deleted_at TIMESTAMP DEFAULT NULL
	);
	`, tableName, tableName)
)

// TestGlobal run tests
func TestGlobal(t *testing.T) {

	db, err := utilities.ConnectDataBaseTest()

	if err != nil {
		t.Error(err)
		return
	}

	err = utilities.BuildDataBase(db, query)

	if err != nil {
		t.Error(err)
		return
	}

	defer utilities.CloseDataBaseTest(db)

	provider := &pr.PostgresProvider{DB: db, TableName: tableName}
	testService = &Service{provider: provider}

	testToken = utilities.LoginUser(db)

	t.Run("add group", testAddGroupHandler)
	t.Run("get group", testGetGroupHandler)
	t.Run("list group", testListGroupHandler)
	t.Run("update group", testUpdateGroupHandler)
	t.Run("delete group", testDeleteGroupHandler)
	t.Run("join group", testJoinGroupHandler)
	t.Run("leave group", testLeaveGroupHandler)
}
