package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	hasDebugging string
)

// Initialize function will be called when every command gets called.
func init() {
	// Get our environment variables which will used to configure our application and save across all the sub-commands.
	rootCmd.PersistentFlags().StringVar(&hasDebugging, "hasDebugging", os.Getenv("LOGFARM_HAS_DEBUGGING"), "Indicate whether debugging is true or false.")

	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if hasDebugging == "true" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)

		// The following line of code adds a pretty output to the console.
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().
			Timestamp(). // Add timestamp to every call.
			Caller().    // IMPORTANT: Add line numbers to every call.
			Logger()
	} else {
		// The following line of code adds a pretty output to the console.
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().
			Timestamp(). // Add timestamp to every call.
			Logger()
	}

}

var rootCmd = &cobra.Command{
	Use:   "logfarm",
	Short: "Syslog server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Do nothing.
	},
}

// Execute is the main entry into the application from the command line terminal.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
