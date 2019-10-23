package group

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	pr "github.com/idirall22/group/providers/postgres"
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

var testService *Service
var database *sql.DB

func cleanDB(db *sql.DB) error {
	query := fmt.Sprintf(`
		DROP TABLE IF EXISTS groups;

		CREATE TABLE IF NOT EXISTS groups(
		    id SERIAL PRIMARY KEY,
			name VARCHAR NOT NULL,
			admin_id INTEGER NOT NULL,
			users_ids INTEGER[] default '{}'::int[],
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

	provider := &pr.PostgresProvider{DB: db, TableName: "groups"}
	testService = &Service{provider: provider}

	err = cleanDB(db)
	if err != nil {
		return err
	}

	database = db
	return nil
}

func TestGlobal(t *testing.T) {
	if err := connectDB(); err != nil {
		log.Fatal(err)
		return
	}
	defer closeDB(database)

	t.Run("add group", testAddGroup)
}
