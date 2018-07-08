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

package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Cmd exports the run command
var Cmd *cobra.Command

// Version exports the version number of the current build
var Version = "0.0.1"

// exports command by default
func init() {
	Cmd = &cobra.Command{
		Use:   "version",
		Short: "Prints the version number",
		Run:   run,
	}
}

func run(cmd *cobra.Command, args []string) {
	fmt.Printf("Version: %s", Version)
}
