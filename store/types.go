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
	"sync"

	"github.com/aws/aws-sdk-go/service/ssm"
)

// ParameterStore represents the interface to the parameter store of SSM which needs to be implemented.
type ParameterStore interface {
	Parameters() ([]*ssm.Parameter, error)
}

// SSMStore contains the data, and clients to access the paramter store of SSM
type SSMStore struct {
	recursive      *bool
	withDecryption *bool
	parameters     []*ssm.Parameter
	ssm            *ssm.SSM
	ssmPath        *string

	ctx context.Context
	mux sync.Mutex
}
