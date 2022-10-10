package cmd

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/bartmika/logfarm/internal/config"
	"github.com/bartmika/logfarm/internal/domain/record"
	"github.com/bartmika/logfarm/internal/interfaceadapters"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print on screen all the data",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("Starting...")

		// Load up all the environment variables.
		appConf := config.AppConfig()

		// Load up all the interface adapters.
		adapters, err := interfaceadapters.NewServices(appConf)
		if err != nil {
			log.Fatal().
				Err(err).
				Msgf("Cannot start interfaceadapters %s", err)
		}

		ctx := context.Background()

		log.Info().Msg("Fetching records...")

		f := &record.RecordFilter{
			SortOrder:                   "ASC",
			SortField:                   "timestamp",
			TimestampGreaterThenOrEqual: time.Date(2022, 1, 1, 1, 30, 45, 100, time.Local),
			TimestampLessThenOrEqual:    time.Now(),
		}

		records, err := adapters.RecordRepo.ListByFilter(ctx, f)
		if err != nil {
			log.Fatal().
				Err(err).
				Msgf("Cannot start list %s", err)
		}

		log.Info().
			Int("length", len(records)).
			Msg("Finished fetching")

		for _, record := range records {
			log.Info().
				Str("Client", record.Client).
				Str("Content", record.Content).
				Int("Facility", record.Facility).
				Str("Hostname", record.Hostname).
				Int("Priority", record.Priority).
				Int("Severity", record.Severity).
				Str("Tag", record.Tag).
				Time("Timestamp", record.Timestamp).
				Str("TLSPeer", record.TLSPeer).
				Msg("Viewing record")
		}
	},
}
