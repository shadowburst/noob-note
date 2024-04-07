package backend

import (
	"database/sql"
)

// Backend struct
type Backend struct {
	db *sql.DB
}

// NewBackend creates a new Backend application struct
func NewBackend() *Backend {
	return &Backend{}
}

func (backend *Backend) Start() error {
	err := backend.initializeDB()
	if err != nil {
		return err
	}

	return nil
}

func (backend *Backend) Stop() error {
	err := backend.closeDB()
	if err != nil {
		return err
	}

	return nil
}
