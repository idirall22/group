package provider

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "diskshar_test"
)

const (
	groupNum = 5
)

var provider *PostgresProvider

func cleanDB(db *sql.DB) error {
	query := fmt.Sprintf(`
		DROP TABLE IF EXISTS groups;

		CREATE TABLE IF NOT EXISTS groups(
		    id SERIAL PRIMARY KEY,
			name VARCHAR NOT NULL,
			admin_id INTEGER NOT NULL,
			user_ids INTEGER[] DEFAULT '{}',
		    created_at TIMESTAMP with TIME ZONE DEFAULT now(),
		    deleted_at TIMESTAMP DEFAULT NULL
		);
		`)

	_, err := db.Exec(query)

	if err != nil {
		return err
	}
	return nil
}

func closeDB(db *sql.DB) {
	db.Close()
}

func connectDB() error {

	dbInfos := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbInfos)
	if err != nil {
		return err
	}

	provider = &PostgresProvider{DB: db, TableName: "groups"}

	err = cleanDB(db)
	if err != nil {
		return err
	}

	return nil
}

func TestGlobal(t *testing.T) {
	if err := connectDB(); err != nil {
		log.Fatal(err)
		return
	}
	defer closeDB(provider.DB)

	t.Run("New", testNew)
	// t.Run("List", testList)
	// t.Run("update", testUpdate)
	// t.Run("delete", testDelete)
}

// Test New
func testNew(t *testing.T) {

	_, err := provider.New(context.Background(), 1, "super groupe")
	if err != nil {
		t.Error("Error should be nil but got:", err)
	}
}

// // Test List
// func testList(t *testing.T) {
// 	comments, err := provider.List(context.Background(), 1, commentNum, 0)
//
// 	if err != nil {
// 		t.Error("Error should be nil but got:", err)
// 		return
// 	}
//
// 	if len(comments) != commentNum {
// 		t.Errorf("Error comments slice length should be %d But got %d",
// 			commentNum, len(comments))
// 	}
// }
//
// // Test update
// func testUpdate(t *testing.T) {
// 	_, err := provider.Update(context.Background(), 1, 1, "updated message")
//
// 	if err != nil {
// 		t.Error("Error should be nil but got:", err)
// 		return
// 	}
// }
//
// // Test delete
// func testDelete(t *testing.T) {
// 	err := provider.Delete(context.Background(), 1, 1)
// 	if err != nil {
// 		t.Error("Error should be nil but got:", err)
// 		return
// 	}
// }
