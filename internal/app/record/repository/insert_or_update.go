package repository

import (
	"context"

	domain "github.com/bartmika/logfarm/internal/domain/record"
)

func (r *RecordRepoImpl) InsertOrUpdateByID(ctx context.Context, m *domain.Record) error {
	if m.ID == "" {
		return r.Insert(ctx, m)
	}

	doesExist, err := r.CheckIfExistsByID(ctx, m.ID)
	if err != nil {
		return err
	}

	if doesExist == false {
		return r.Insert(ctx, m)
	}
	return r.UpdateByID(ctx, m)
}
