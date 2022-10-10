package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/bartmika/logfarm/internal/app"
	"github.com/bartmika/logfarm/internal/config"
	"github.com/bartmika/logfarm/internal/inputports"
	"github.com/bartmika/logfarm/internal/interfaceadapters"
	"github.com/bartmika/logfarm/internal/pkg/time"
	"github.com/bartmika/logfarm/internal/pkg/uuid"
)

func init() {
	rootCmd.AddCommand(logservCmd)
}

var logservCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the logfarm server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Load up all the environment variables.
		appConf := config.AppConfig()

		uuidProvider := uuid.NewUUIDProvider()
		timeProvider := time.NewTimeProvider()

		// Load up all the interface adapters.
		adapters, err := interfaceadapters.NewServices(appConf)
		if err != nil {
			log.Fatal().
				Err(err).
				Msgf("Cannot start interfaceadapters %s", err)
		}

		// Load up all the app services.
		appServices := app.NewAppServices(appConf, uuidProvider, timeProvider, adapters)

		// Load up our HTTP server and connect it with the rest of our application.
		inputPortsServices := inputports.NewServices(appConf, appServices)

		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		go inputPortsServices.CronServer.ListenAndServe()
		go inputPortsServices.SyslogServer.ListenAndServe()

		addr := fmt.Sprintf("%s:%s", appConf.Server.IP, appConf.Server.Port)

		log.Info().Msgf("Server started listening on UDP via %s", addr)

		// Run the main loop blocking code.
		<-done

		stopMainRuntimeLoop(inputPortsServices)
	},
}

func stopMainRuntimeLoop(services inputports.Services) {
	log.Info().Msg("Starting graceful shutdown now...")

	// DEVELOPERS NOTE:
	// Write your closing code here.
	// . . .

	log.Info().Msg("Graceful shutdown finished.")
	log.Info().Msg("Server Exited")
}
