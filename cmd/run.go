package cmd

import (
	"context"

	"github.com/andersnormal/penny/pkg/runner"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type root struct {
	logger *log.Entry
}

func runE(c *cobra.Command, args []string) error {
	// create a new root
	root := new(root)

	// init logger
	root.logger = log.WithFields(log.Fields{
		"verbose": cfg.Verbose,
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := runner.New(cfg)

	if err := r.Exec(ctx, []string{}, args...); err != nil {
		return err
	}

	// noop
	return nil
}
