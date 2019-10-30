package config

import (
	"syscall"
)

// Config contains a configuration for Autobot
type Config struct {
	// File is a config file provided
	File string
	// Verbose toggles the verbosity
	Verbose bool
	// ReloadSignal
	ReloadSignal syscall.Signal
	// TermSignal
	TermSignal syscall.Signal
	// KillSignal
	KillSignal syscall.Signal
	// Logger ...
	Logger *Logger
}

// Logger ...
type Logger struct {
	// Enabled ...
	Enabled bool
	// Level ...
	Level string
	// Format ...
	Format string
}

const (
	// DefaultLogger ...
	DefaultLogger = false
	// DefaultLoggerLevel is the default logging level.
	DefaultLoggerLevel = "warn"
	// DefaultLoggerFormat is the default format of the logger
	DefaultLoggerFormat = "text"
	// DefaultTermSignal is the signal to term the agent.
	DefaultTermSignal = syscall.SIGTERM
	// DefaultReloadSignal is the default signal for reload.
	DefaultReloadSignal = syscall.SIGHUP
	// DefaultKillSignal is the default signal for termination.
	DefaultKillSignal = syscall.SIGINT
	// DefaultVerbose is the default verbosity.
	DefaultVerbose = false
)

// New returns a new Config
func New() *Config {
	return &Config{
		Verbose: DefaultVerbose,
		Logger: &Logger{
			Enabled: DefaultLogger,
			Level:   DefaultLoggerLevel,
			Format:  DefaultLoggerFormat,
		},
		ReloadSignal: DefaultReloadSignal,
		TermSignal:   DefaultTermSignal,
		KillSignal:   DefaultKillSignal,
	}
}
