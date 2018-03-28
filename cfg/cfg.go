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

	"github.com/aws/aws-sdk-go/aws/session"
	log "github.com/sirupsen/logrus"
)

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

	// DefaultWithDecryption is the default for decryption.
	DefaultWithDecryption = true
)

// New is returning a new config with default paramaeters
func New() *Config {
	return &Config{
		Verbose:        DefaultVerbose,
		LogLevel:       DefaultLogLevel,
		ReloadSignal:   DefaultReloadSignal,
		KillSignal:     DefaultKillSignal,
		WithDecryption: DefaultWithDecryption,
		Recursive:      DefaultRecursive,
	}
}

// Must returns a new Config with required parameters
func Must(session *session.Session) *Config {
	c := New()
	c.Session = session

	return c
}

// Bool returns a pointer to a boolean
func Bool(f bool) *bool {
	return &f
}

// String returns a pointer to a string
func String(f string) *string {
	return &f
}
