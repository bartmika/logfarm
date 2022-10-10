package interfaceadapters

import (
	"github.com/rs/zerolog/log"

	record_r "github.com/bartmika/logfarm/internal/app/record/repository"
	"github.com/bartmika/logfarm/internal/config"
	record_d "github.com/bartmika/logfarm/internal/domain/record"
	"github.com/bartmika/logfarm/internal/interfaceadapters/migrations"
	"github.com/bartmika/logfarm/internal/interfaceadapters/storage/sqlite"
)

// Services contains the exposed services of interface adapters
type Services struct {
	RecordRepo record_d.Repository
}

// NewServices Instantiates the interface adapter services
func NewServices(appConf *config.Conf) (*Services, error) {
	// Step 2: Connect to database.
	db, err := sqlite.ConnectDB(appConf.DB.FilePath)
	if err != nil {
		return nil, err
	}

	// Step 2: Perform our automatic database migrations (if enabled)
	if appConf.DB.HasAutoMigrations {
		if err := migrations.RunOnDB(db); err != nil {
			return nil, err
		}
	} else {
		log.Warn().Msg("No migrations occured - you must do this manually.")
	}

	return &Services{
		RecordRepo: record_r.NewRecordRepoImpl(db),
	}, nil
}
