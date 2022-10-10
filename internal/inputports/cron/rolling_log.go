package cron

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (cronServer *Server) ProcessRollOldestRecordsTask() {
	log.Debug().Str("func", "ProcessRollOldestRecordsTask").Str("service", "cron").Msg("Executing.")

	ctx := context.Background()
	if err := cronServer.AppServices.RecordUsecase.RollOldestRecords(ctx); err != nil {
		log.Error().Err(err).Caller().Msg("Failed performing rolling on records.")
	}
}
