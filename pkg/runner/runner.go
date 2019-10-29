package runner

import (
	"context"
	"os"
	"os/exec"
)

// Runner ...
type Runner interface {
	Exec(context.Context, []string, ...string) error
}

type runner struct{}

// New ...
func New() Runner {
	r := new(runner)

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

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

func newRunner() *runner {
	r := new(runner)

	return r
}
