package cmd

import (
	c "github.com/andersnormal/penny/pkg/config"

	"github.com/spf13/cobra"
)

func addFlags(cmd *cobra.Command, cfg *c.Config) {
	// set the config file
	cmd.Flags().StringVar(&cfg.File, "config", "", "config file (default is $HOME/.autobot.yaml")
	// disable log
	cmd.Flags().BoolVar(&cfg.Logger.Enabled, "log", c.DefaultLogger, "enable log")
	// set log format
	cmd.Flags().StringVar(&cfg.Logger.Format, "log-format", c.DefaultLoggerFormat, "log format (default is 'text')")
	// set log level
	cmd.Flags().StringVar(&cfg.Logger.Level, "log-level", c.DefaultLoggerLevel, "log level (default is 'warn'")
}
