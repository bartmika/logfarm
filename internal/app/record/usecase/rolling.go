package usecase

import (
	"context"
	"time"

	"github.com/bartmika/logfarm/internal/domain/record"
	"github.com/rs/zerolog/log"
)

func (uc recordUsecase) RollOldestRecords(ctx context.Context) (err error) {

	// Generate the beginning of time.
	startAt := uc.Time.Date(2000, 1, 1, 1, 1, 1, 1, time.Local)

	// Generate the expiry day where any date older then this date is considered old.
	finishAt := uc.Time.Now()
	finishAt = finishAt.AddDate(0, -1*uc.MaxDayAge, 0)

	// Generate a filter that will fetch all the records from the beginning of
	// time to the expiry date.
	f := &record.RecordFilter{
		SortOrder:                   "ASC",
		SortField:                   "timestamp",
		TimestampGreaterThenOrEqual: startAt,
		TimestampLessThenOrEqual:    finishAt,
	}

	records, err := uc.RecordRepo.ListByFilter(ctx, f)
	if err != nil {
		return err
	}

	for _, record := range records {
		if err := uc.RecordRepo.DeleteByID(ctx, record.ID); err != nil {
			return err
		}
		log.Info().
			Str("ID", record.ID).
			Time("Timestamp", record.Timestamp).
			Str("Function", "RollOldestRecords").
			Msg("Deleted record")
	}

	log.Debug().
		Int("Deletion Count", len(records)).
		Str("Function", "RollOldestRecords").
		Msg("Finished deleting old records")
	return nil
}
