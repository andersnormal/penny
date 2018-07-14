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

package cfg

import (
	"syscall"

	log "github.com/sirupsen/logrus"
)

// Config exports a reusable config
var Config *CmdConfig

const (
	// DefaultLogLevel is the default logging level.
	DefaultLogLevel = log.WarnLevel

	// DefaultReloadSignal is the default signal for reload.
	DefaultReloadSignal = syscall.SIGHUP

	// DefaultKillSignal is the default signal for termination.
	DefaultKillSignal = syscall.SIGINT

	// DefaultVerbose is the default verbosity.
	DefaultVerbose = false

	// DefaultRecursive is the default for recursive lookup.
	DefaultRecursive = true

	// DefaultTimeout is the default time to configure the runtime
	DefaultTimeout = 60

	// DefaultPrefix for the environment variables
	DefaultPrefix = "SSM"

	// DefaultOverwrite for environment variables
	DefaultOverwrite = false

	// DefaultForce for environment setup
	DefaultForce = false

	// DefaultSSMWithDecryption for SSM decryption
	DefaultSSMWithDecryption = true

	// DefaultSSMRegion is the default AWS region
	DefaultSSMRegion = "eu-west-1"
)

func init() {
	ssm := &SSMConfig{
		Region:         DefaultSSMRegion,
		WithDecryption: DefaultSSMWithDecryption,
	}

	Config = &CmdConfig{
		Verbose:      DefaultVerbose,
		LogLevel:     DefaultLogLevel,
		ReloadSignal: DefaultReloadSignal,
		KillSignal:   DefaultKillSignal,
		Recursive:    DefaultRecursive,
		Prefix:       DefaultPrefix,
		Timeout:      DefaultTimeout,
		Overwrite:    DefaultOverwrite,
		Force:        DefaultForce,
		SSM:          ssm,
	}
}

// Bool returns a pointer to a boolean
func Bool(f bool) *bool {
	return &f
}

// String returns a pointer to a string
func String(f string) *string {
	return &f
}
