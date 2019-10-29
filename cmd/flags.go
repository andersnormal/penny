package cmd

import (
	c "github.com/andersnormal/penny/pkg/config"

	"github.com/spf13/cobra"
)

func addFlags(cmd *cobra.Command, cfg *c.Config) {
	// set the config file
	cmd.Flags().StringVar(&cfg.File, "config", "", "config file (default is $HOME/.autobot.yaml")
	// enable verbose output
	cmd.Flags().BoolVar(&cfg.Verbose, "verbose", c.DefaultVerbose, "enable verbose output")
	// set log format
	cmd.Flags().StringVar(&cfg.LogFormat, "log-format", c.DefaultLogFormat, "log format (default is 'text')")
	// set log level
	cmd.Flags().StringVar(&cfg.LogLevel, "log-level", c.DefaultLogLevel, "log level (default is 'warn'")
}
