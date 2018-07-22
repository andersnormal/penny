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
	"sync"

	"github.com/andersnormal/penny/provider"
)

// Runner is the interface to a runtime environment
type Runner interface {
	// Setup should setup the runtime environment
	Exec() error
}

// Run is a runtime environment for SSM
type Run struct {
	sync.Mutex

	args   []string
	kvPair []*provider.KVPair
}
