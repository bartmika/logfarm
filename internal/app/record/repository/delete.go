package repository

import (
	"context"
	"database/sql"
	"time"
)

func (s *RecordRepoImpl) DeleteByID(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `DELETE FROM records WHERE id = $1;`
	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		if err == sql.ErrNoRows { // CASE 1 OF 2: Cannot find record with that ID.
			return nil
		}
		// CASE 2 OF 2: All other errors.
		return err
	}
	return nil
}
