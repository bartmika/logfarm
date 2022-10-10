package repository

import (
	"database/sql"
)

type RecordRepoImpl struct {
	db *sql.DB
}

func NewRecordRepoImpl(db *sql.DB) *RecordRepoImpl {
	return &RecordRepoImpl{
		db: db,
	}
}
