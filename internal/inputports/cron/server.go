package cron

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"

	"github.com/bartmika/logfarm/internal/app"
	"github.com/bartmika/logfarm/internal/config"
)

// Server Represents the http server running for this service
type Server struct {
	Scheduler   *gocron.Scheduler
	AppServices app.Services
}

// NewServer HTTP Server constructor
func NewServer(appConf *config.Conf, appServices app.Services) *Server {

	scheduler := gocron.NewScheduler(time.UTC)

	cronServer := &Server{
		AppServices: appServices,
		Scheduler:   scheduler,
	}
	return cronServer
}

// ListenAndServe Starts listening for requests
func (cronServer *Server) ListenAndServe() {
	cronServer.Scheduler.Cron("*/1 * * * *").Do(cronServer.ProcessPingTask) // Every minute.
	cronServer.Scheduler.Cron("*/1 * * * *").Do(cronServer.ProcessRollOldestRecordsTask)

	log.Info().Str("func", "ListenAndServe").Str("service", "cron").Msg("Service started.")
	defer log.Info().Str("func", "ListenAndServe").Str("service", "cron").Msg("Service stopped.")

	// Starts the scheduler and blocks current execution path.
	cronServer.Scheduler.StartBlocking()
}

func (cronServer *Server) ProcessPingTask() {
	log.Debug().Str("func", "ProcessPingTask").Str("service", "cron").Msg("ping")
}
