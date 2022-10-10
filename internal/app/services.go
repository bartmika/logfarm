package app

import (
	record_usecase "github.com/bartmika/logfarm/internal/app/record/usecase"
	"github.com/bartmika/logfarm/internal/config"
	"github.com/bartmika/logfarm/internal/interfaceadapters"
	"github.com/bartmika/logfarm/internal/pkg/time"
	"github.com/bartmika/logfarm/internal/pkg/uuid"
)

//Services contains all exposed services of the application layer
type Services struct {
	RecordUsecase record_usecase.Usecase
}

// NewAppServices Bootstraps Application Layer dependencies
func NewAppServices(appConf *config.Conf, uuidProvider uuid.Provider, timeProvider time.Provider, adapters *interfaceadapters.Services) Services {
	return Services{
		RecordUsecase: record_usecase.NewRecordUsecase(appConf, uuidProvider, timeProvider, adapters.RecordRepo),
	}
}
