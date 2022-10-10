package usecase

import (
	"context"

	"github.com/bartmika/logfarm/internal/config"
	ls_d "github.com/bartmika/logfarm/internal/domain/record"
	record_d "github.com/bartmika/logfarm/internal/domain/record"
	"github.com/bartmika/logfarm/internal/pkg/time"
	"github.com/bartmika/logfarm/internal/pkg/uuid"
	"gopkg.in/mcuadros/go-syslog.v2/format"
)

// Usecase Provides interface for the record use cases.
type Usecase interface {
	// InsertLogParts Function will take the `syslog` to parse it and save it into our app.
	InsertLogParts(ctx context.Context, logParts format.LogParts) (record *ls_d.Record, err error)

	// RollingRecords Function will delete all records past the `max_day_age`.
	RollOldestRecords(ctx context.Context) (err error)
}

type recordUsecase struct {
	MaxDayAge  int
	UUID       uuid.Provider
	Time       time.Provider
	RecordRepo record_d.Repository
}

// NewRecordUsecase Constructor function for the `RecordUsecase` implementation.
func NewRecordUsecase(appConf *config.Conf, uuidProvider uuid.Provider, timeProvider time.Provider, recordRepo record_d.Repository) *recordUsecase {
	return &recordUsecase{
		MaxDayAge:  appConf.Setting.MaxDayAge,
		UUID:       uuidProvider,
		Time:       timeProvider,
		RecordRepo: recordRepo,
	}
}
