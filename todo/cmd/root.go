package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var dbFile string
var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple command line todo list",
	Long:  ``, // TODO: fill in
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.WithField("error", err).Fatal("Failed to execute root command")
	}
}

func init() {
	cobra.OnInitialize(initLogging)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&dbFile, "database", "", "database file (default is $HOME/.todo.db)")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Enable verbose output")
}

func initLogging() {
	if verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}
