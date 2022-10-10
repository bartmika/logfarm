package syslog

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"gopkg.in/mcuadros/go-syslog.v2"

	"github.com/bartmika/logfarm/internal/app"
	"github.com/bartmika/logfarm/internal/config"
)

// Server Represents the http server running for this service
type Server struct {
	Address     string
	AppServices app.Services
}

// NewServer HTTP Server constructor
func NewServer(appConf *config.Conf, appServices app.Services) *Server {
	syslogServer := &Server{
		Address:     fmt.Sprintf("%s:%s", appConf.Server.IP, appConf.Server.Port),
		AppServices: appServices,
	}
	return syslogServer
}

// ListenAndServe Starts listening for requests
func (syslogServer *Server) ListenAndServe() {
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.Automatic) // https://github.com/mcuadros/go-syslog/issues/29
	server.SetHandler(handler)

	server.ListenUDP(syslogServer.Address)
	server.Boot()

	go func(channel syslog.LogPartsChannel) {
		ctx := context.Background()

		for logParts := range channel {

			// DEVELOPERS NOTE:
			// Take the raw data and submit it into our app to handle.
			_, err := syslogServer.AppServices.RecordUsecase.InsertLogParts(ctx, logParts)
			if err != nil {
				log.Err(err).Caller().Msg("Failed executing insertion of log parts")
			}
		}
	}(channel)

	// For debugging purposes only.
	log.Info().Str("func", "ListenAndServe").Str("service", "syslog").Msg("Service started.")
	defer log.Info().Str("func", "ListenAndServe").Str("service", "syslog").Msg("Service stopped.")

	// Block the current thread.
	server.Wait()
}
