package sqlite

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

func ConnectDB(filePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", filePath)
	if err != nil {
		return nil, err
	}
	return db, nil
}
