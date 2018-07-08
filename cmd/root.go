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

package cmd

import (
	"fmt"
	"os"

	config "github.com/andersnormal/penny/cfg"
	"github.com/andersnormal/penny/run"
	"github.com/andersnormal/penny/version"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	cfg     = config.Config
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "penny",
	Short: "Executes a subprocess with environment variables from SSM",
	RunE:  runE,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Init Cobra config
	cobra.OnInitialize(initConfig)

	// silence on the root cmd
	RootCmd.SilenceErrors = true
	RootCmd.SilenceUsage = true

	// add run command
	RootCmd.AddCommand(run.Cmd)

	// add version command
	RootCmd.AddCommand(version.Cmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.penny.yaml)")

	// Set the verbosity
	RootCmd.PersistentFlags().BoolVarP(&cfg.Verbose, "verbose", "v", cfg.Verbose, "enable verbose logging")

	// Bind to read in
	viper.BindPFlag("verbose", RootCmd.PersistentFlags().Lookup("verbose"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// set specific prefix
	viper.SetEnvPrefix("PENNY")

	// read in files
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".peny" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".penny")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// runE is running the root command of Penny
func runE(cmd *cobra.Command, args []string) error {
	var err error

	return err // noop

	// save for latet
	// for {
	// 	select {
	// 	case <-ctx.Done():
	// 		err = ctx.Err()
	// 		return err
	// 	default:
	// 	}
	// }
}
