package backend

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func (backend *Backend) initializeDB() error {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		return err
	}

	backend.db = db

	return nil
}

func (backend *Backend) closeDB() error {
	err := backend.db.Close()
	if err != nil {
		return err
	}

	return nil
}

func migrate(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS migrations (id INTEGER PRIMARY KEY, name VARCHAR(255) NOT NULL, batch INTEGER DEFAULT 0)")
	if err != nil {
		return err
	}

	var batch int
	err = db.QueryRow("SELECT MAX(batch) FROM migrations").Scan(&batch)
	if err != nil {
		batch = 0
	}

	batch++

	files, err := os.ReadDir("./migrations")
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		rowCount := 0
		err = db.QueryRow("SELECT COUNT(*) FROM migrations WHERE name = ?", file.Name()).Scan(&rowCount)
		if err != nil {
			return err
		}

		if rowCount > 0 {
			continue
		}

		fileContent, err := os.ReadFile("./migrations/" + file.Name())
		if err != nil {
			return err
		}

		_, err = db.Exec(string(fileContent))
		if err != nil {
			return err
		}

		_, err = db.Exec("INSERT INTO migrations (name, batch) VALUES (?, ?)", file.Name(), batch)
		if err != nil {
			return err
		}
	}

	return nil
}
