package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type taskDB struct {
	db         *sql.DB
	dataDir    string
	dbFileName string
}

// SetupDb creates the database and tables if they don't exists
func SetupDb(path string) error {
	dbPath := path + "/taskman.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// check for existing table
	query := `SELECT name FROM sqlite_master WHERE type='table' AND name='tasks';`
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return nil
	}

	if err := createTables(db); err != nil {
		return err
	}

	return nil
}

// create table if not exists
func createTables(db *sql.DB) error {
	ddlQuery := `CREATE TABLE 'tasks' (
		'id' INTEGER PRIMARY KEY AUTOINCREMENT, 
		'parent_id' VARCHAR(255) NULL, 
		'user_id' VARCHAR(64) NULL, 
		'status' VARCHAR(10) NOT NULL, 
		'title' VARCHAR(255) NOT NULL, 
		'detail' TEXT NULL, 
		'created_at' TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
		'created_by' INTEGER NOT NULL, 
		'modified_at' TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
		'modified_by' INTEGER NULL
);`
	_, err := db.Exec(ddlQuery)
	if err != nil {
		return err
	}
	return nil
}
