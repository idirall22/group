package provider

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/idirall22/utilities"
	_ "github.com/lib/pq"
)

var (
	provider  *PostgresProvider
	database  *sql.DB
	tableName = "groups"
	groupNum  = 5
	query     = fmt.Sprintf(`
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

	provider = &PostgresProvider{DB: db, TableName: tableName}

	t.Run("New", testNew)
	t.Run("Get", testGet)
	t.Run("List", testList)
	t.Run("update", testUpdate)
	t.Run("delete", testDelete)
	t.Run("join", testJoin)
	// t.Run("leave", testLeave)
}

// Test New
func testNew(t *testing.T) {

	_, err := provider.New(context.Background(), 1, "super groupe")
	if err != nil {
		t.Error("Error should be nil but got:", err)
	}
}

// Test Get
func testGet(t *testing.T) {
	_, err := provider.Get(context.Background(), 1, 1)

	if err != nil {
		t.Error("Error should be nil but got:", err)
		return
	}
}

// Test List
func testList(t *testing.T) {
	_, err := provider.List(context.Background(), groupNum, 0)

	if err != nil {
		t.Error("Error should be nil but got:", err)
		return
	}
}

// Test update
func testUpdate(t *testing.T) {
	err := provider.Update(context.Background(), 1, 1, "updated goupe name")

	if err != nil {
		t.Error("Error should be nil but got:", err)
		return
	}
}

// Test delete
func testDelete(t *testing.T) {
	err := provider.Delete(context.Background(), 1, 1)
	if err != nil {
		t.Error("Error should be nil but got:", err)
		return
	}
}

// Test join a group
func testJoin(t *testing.T) {
	if err := provider.Join(context.Background(), 1, 1); err != nil {
		t.Error("Error should be nil but got:", err)
		return
	}

}

// Test leave a group
func testLeave(t *testing.T) {
	if err := provider.Leave(context.Background(), 2, 1); err != nil {
		t.Error("Error should be nil but got:", err)
		return
	}

}
