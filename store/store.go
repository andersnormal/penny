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

package store

import (
	"context"

	config "github.com/andersnormal/penny/cfg"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var cfg = config.Config

// New returns a new Store with the fetched parameters, and service id
func New() *SSMStore {
	return &SSMStore{}
}

// Must returns a new Store with parameters configured
func Must(ctx context.Context, ssm *ssm.SSM) (*SSMStore, error) {
	var err error
	var store = New()

	store.ctx = ctx
	store.ssm = ssm
	store.ssmPath = config.String(cfg.SSMPath)
	store.recursive = config.Bool(cfg.Recursive)
	store.withDecryption = config.Bool(cfg.WithDecryption)

	_, err = store.getParameters(nil) // harvest err

	return store, err
}

// Parameters returns the fetches parameters
func (s *SSMStore) Parameters() (parameters []*ssm.Parameter) {
	return s.parameters
}

// getParameters is fetching the parameters beloging to a service id from the parameter store
func (s *SSMStore) getParameters(nextToken *string) (parameters []*ssm.Parameter, err error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	// input to the SSM to get parameters by path
	input := &ssm.GetParametersByPathInput{
		Path:           s.ssmPath,
		Recursive:      s.recursive,
		WithDecryption: s.withDecryption,
	}

	if nextToken != nil {
		input.NextToken = nextToken
	}

	output, err := s.ssm.GetParametersByPathWithContext(s.ctx, input)
	if err != nil {
		return parameters, err
	}
	s.parameters = append(s.parameters, output.Parameters...)

	if nextToken != nil {
		s.getParameters(nextToken)
	}

	return s.parameters, err
}
