package runner

import (
	"context"
	"os"
	"os/exec"

	"github.com/andersnormal/penny/pkg/config"

	log "github.com/sirupsen/logrus"
)

// Runner ...
type Runner interface {
	Exec(context.Context, []string, ...string) error
}

type runner struct {
	cfg *config.Config
}

// New ...
func New(cfg *config.Config) Runner {
	r := newRunner(cfg)

	return r
}

// Exec ...
func (r *runner) Exec(ctx context.Context, env []string, args ...string) error {
	var execCmd string
	var execArgs []string

	if len(args) > 0 {
		execCmd = args[0]
	}

	if len(args) > 1 {
		execArgs = append(execArgs, args[1:]...)
	}

	cmd := exec.CommandContext(ctx, execCmd, execArgs...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), env...)

	if r.cfg.Logger.Enabled {
		r.logConfig(cmd)
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

func (r *runner) logConfig(cmd *exec.Cmd) {
	ll := log.New()
	ll.SetFormatter(&log.TextFormatter{})

	// reset log format
	if r.cfg.Logger.Format == "json" {
		ll.SetFormatter(&log.JSONFormatter{})
	}

	// set the configured log level
	if level, err := log.ParseLevel(r.cfg.Logger.Level); err == nil {
		ll.SetLevel(level)
	}

	cmd.Stdout = ll.WriterLevel(log.InfoLevel)
	cmd.Stderr = ll.WriterLevel(log.ErrorLevel)
}

func newRunner(cfg *config.Config) *runner {
	r := new(runner)
	r.cfg = cfg

	return r
}
