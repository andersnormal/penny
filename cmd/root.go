package cmd

import (
	"os"

	"github.com/andersnormal/penny/pkg/config"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfg *config.Config
var build string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "penny",
	Short: "Penny is the nanny for your container commands",
	Long: `
		Penny wraps your command and executes it in the environment you configure.
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
	RootCmd.SilenceErrors = true
	RootCmd.SilenceUsage = true

	// initialize cobra
	cobra.OnInitialize(initConfig)

	// adding flags
	addFlags(RootCmd, cfg)

	// set the default format, which is basically text
	log.SetFormatter(&log.TextFormatter{})
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// allow to read in from environment
	viper.SetEnvPrefix("penny")
	viper.AutomaticEnv() // read in environment variables that match

	// unmarshal to config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf(errors.Wrap(err, "cannot unmarshal config").Error())
	}

	// config logger
	logConfig(cfg)
}

func logConfig(cfg *config.Config) {
	// reset log format
	if cfg.LogFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	if cfg.Verbose {
		cfg.LogLevel = "info"
	}

	// set the configured log level
	if level, err := log.ParseLevel(cfg.LogLevel); err == nil {
		log.SetLevel(level)
	}
}
