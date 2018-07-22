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
	"time"

	log "github.com/sirupsen/logrus"
)

// CmdConfig contains a current config of Penny
type CmdConfig struct {
	// Verbose toggles the verbosity
	Verbose bool

	// LogLevel is the level with with to log for this config
	LogLevel log.Level

	// ReloadSignal
	ReloadSignal syscall.Signal

	// KillSignal
	KillSignal syscall.Signal

	// Recursive lookup
	Recursive bool

	// Timeout of the runtime
	Timeout time.Duration

	// Prefix for the environment variable in runtime to use
	Prefix string

	// Overwrite existing environment variables
	Overwrite bool

	// Force execution
	Force bool

	// Path specifies the path in a provider
	Path string

	// SSM exposes an AWS Parameter Store Config
	SSM *SSMConfig
}

// SSMConfig is the config for an SSM Parameter Store in AWS
type SSMConfig struct {
	// Enable is enabling the SSM provider
	Enable bool

	// Region specifies the AWS Region
	Region string

	// WithDecryption specifies to allow to KMS encrypted values
	WithDecryption bool

	// Token specifies an AWS Access Token
	Token string
}
