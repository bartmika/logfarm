package cmd

import (
	"log/syslog"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

// Ex:
// go run main.go send --message="This is a test message"

var (
	sendMessage string
	// changePassPassword string
)

func init() {
	sendCmd.Flags().StringVarP(&sendMessage, "message", "m", "", "The log message content.")
	sendCmd.MarkFlagRequired("message")
	rootCmd.AddCommand(sendCmd)
}

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a log via syslog using `rs/zerolog`",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logwriter, _ := syslog.Dial("udp", "localhost:514", syslog.LOG_DEBUG|syslog.LOG_ERR|syslog.LOG_INFO, "logfarm")

		// UNIX Time is faster and smaller than most timestamps
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

		log := zerolog.New(logwriter).With().
			Str("cmd", "send"). // Add extra context items.
			Timestamp().        // Add timestamp to every call.
			Caller().           // Add line numbers to every call.
			Logger()

		log.Info().Msg(sendMessage) // The content message to send.

		// DEVELOPERS NOTE:
		// EXAMPLE CONTENT OUTPUT:
		// {"level":"info","command":"send","caller":"/Users/bmika/go/src/github.com/bartmika/logfarm/cmd/send.go:44","message":"This is a test message"}
	},
}
