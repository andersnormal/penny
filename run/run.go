// Copyright 2018 Sebastian Döll
// Copyright 2018 Axel Springer SE
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package run

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	config "github.com/andersnormal/penny/cfg"
)

// Cmd exports the run command
var Cmd *cobra.Command

// config
var cfg = config.Config

// New creates a new command line interface.
func New() *Run {
	return &Run{}
}

// exports command by default
func init() {
	Cmd = &cobra.Command{
		Use:   "run",
		Short: "Runs command in configured environment",
		RunE:  runE,
	}

	// Path
	Cmd.Flags().StringVarP(&cfg.Path, "path", "p", cfg.Path, "path in the parameter store")

	// Enable SSM
	Cmd.Flags().BoolVar(&cfg.SSM.Enable, "ssm", cfg.SSM.Enable, "enable SSM Parameter Store")

	// Recursive lookup
	Cmd.Flags().BoolVarP(&cfg.Recursive, "recursive", "r", cfg.Recursive, "recursive lookup")

	// With decryption
	Cmd.Flags().BoolVarP(&cfg.SSM.WithDecryption, "ssm-decrypt", "d", cfg.SSM.WithDecryption, "disable decryption")

	// AWS Region
	Cmd.Flags().StringVar(&cfg.SSM.Region, "ssm-region", cfg.SSM.Region, "AWS Region")

	// Timeout
	Cmd.Flags().DurationVarP(&cfg.Timeout, "timeout", "t", cfg.Timeout, "timeout of the config (in seconds)")

	// SSM Prefix
	Cmd.Flags().StringVar(&cfg.Prefix, "prefix", cfg.Prefix, "prefix for the environment variables")

	// Overwrite existing envs
	Cmd.Flags().BoolVar(&cfg.Overwrite, "overwrite", cfg.Overwrite, "overwrite existing environment variables")

	// Force execution
	Cmd.Flags().BoolVar(&cfg.Force, "force", cfg.Force, "force run the command")

	// bind to read in
	viper.BindPFlag("path", Cmd.Flags().Lookup("path"))

	// parse arbitrary args at the end
	Cmd.Args = cobra.ArbitraryArgs
}

func runE(cmd *cobra.Command, args []string) error {
	// new Run
	// var run = new(Run)

	// // create new ctx
	// ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout*time.Second)
	// defer cancel()

	// new AWS Session
	// session := session.Must(session.NewSession(&aws.Config{
	// 	Region: aws.String(cfg.SSM.Region),
	// }))
	// ssmSvc := ssm.New(session)

	// // set path
	// cfg.Path = viper.GetString("path")

	// // create a new SSM store and SSM environment
	// ssmStore, err := store.Must(ctx, ssmSvc)
	// if !cfg.Force && err != nil {
	// 	return err
	// }

	viper.Debug()

	// // configure run
	// run.store = ssmStore
	// run.args = args

	// do simple execute, should be more complex later
	// run.Exec()

	return nil // noop
}

// Exec is setting up the environment with the configured store
// func (e *Run) Exec() error {
// 	var err error
// 	var execCmd string
// 	var execArgs []string

// 	if len(e.args) > 0 {
// 		execCmd = e.args[0]
// 	}

// 	if len(e.args) > 1 {
// 		execArgs = append(execArgs, e.args[1:]...)
// 	}

// 	cmd := exec.Command(execCmd, execArgs...)
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr

// 	env, err := e.Env()
// 	if !cfg.Force && err != nil {
// 		return err
// 	}
// 	cmd.Env = append(os.Environ(), env...)

// 	// todo: listen for syscalls

// 	// exec
// 	err = cmd.Start()
// 	if err != nil {
// 		return err
// 	}

// 	err = cmd.Wait()

// 	return err // on error
// }

// // Env returns an environment
// func (e *Run) Env() ([]string, error) {
// 	var err error
// 	var env []string

// 	// setup env
// 	for _, parameter := range e.store.Parameters() {
// 		name := strings.TrimPrefix(aws.StringValue(parameter.Name), cfg.Path)
// 		parts := strings.Split(name, "/")
// 		parts = format(notEmpty(parts))

// 		// prefix
// 		parts = append([]string{cfg.Prefix}, parts...)

// 		key := strings.Join(parts, "_")
// 		if _, ok := os.LookupEnv(key); ok && !cfg.Overwrite {
// 			return env, errors.New("could not")
// 		}

// 		env = append(env, fmt.Sprintf("%s=%s", key, aws.StringValue(parameter.Value)))
// 	}

// 	return env, err
// }

func format(s []string) []string {
	var r []string
	for _, str := range s {
		r = append(r, strings.ToUpper(str))
	}
	return r
}

func notEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
