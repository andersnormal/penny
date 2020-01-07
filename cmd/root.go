package cmd

import (
	"os"

	"github.com/andersnormal/penny/pkg/config"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var cfg *config.Config
var build string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "penny",
	Short: "Penny is the nanny for your container commands",
	Long: `
		Penny wraps your command and executes in a prepopulated environment.
  	`,
	PreRunE: preRunE,
	RunE:    runE,
	Version: build,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func preRunE(cmd *cobra.Command, args []string) error {
	return nil
}

func init() {
	// init config
	cfg = config.New()

	// silence on the root cmd
	RootCmd.SilenceErrors = false
	RootCmd.SilenceUsage = true

	// adding flags
	addFlags(RootCmd, cfg)

	// set the default format, which is basically text
	log.SetFormatter(&log.TextFormatter{})
}
