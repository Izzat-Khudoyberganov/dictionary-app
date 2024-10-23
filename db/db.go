package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite", "example.db")

	if err != nil {
		panic("Database could not connect: " + err.Error())
	}

	DB = db

	err = createTables()
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}
}

func createTables() error {
	createAdminTable := `
		CREATE TABLE IF NOT EXISTS admin (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			login TEXT NOT NULL, 
			password TEXT NOT NULL
		)
	`

	_, err := DB.Exec(createAdminTable)
	if err != nil {
		panic("Could not create admin table")
	}

	createDictionaryTable := `
		CREATE TABLE IF NOT EXISTS dictionary (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			word TEXT NOT NULL,
			translate TEXT NOT NULL,
			description TEXT NOT NULL,
			example TEXT NOT NULL
		)
	`

	_, err = DB.Exec(createDictionaryTable)
	if err != nil {
		panic("Could not create tests table")
	}

	createPhrasaTable := `
		CREATE TABLE IF NOT EXISTS phrasa (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			word TEXT NOT NULL,
			translate TEXT NOT NULL
		)
	`

	_, err = DB.Exec(createPhrasaTable)
	if err != nil {
		panic("Could not create phrasa table")
	}

	return err
}
