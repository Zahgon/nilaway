//  Copyright (c) 2025 Uber Technologies, Inc.
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

package main

// GolangCILintDriver implements Driver for running NilAway via golangci-lint.
type GolangCILintDriver struct{}

// Run runs NilAway via golangci-lint on the test project and returns the diagnostics.
func (d *GolangCILintDriver) Run(dir string) (diagnostics map[Position]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create a temporary directory to host the custom gcl binary.

// Read GCLVersion from .golangci.version

// Instantiate the template with the version and write it to the temp dir for building the
// custom gcl binary.

// Install the bootstrap gcl and build the custom binary.

// Run the custom-gcl to collect NilAway diagnostics.

// golangci-lint exits with status 1 when it finds issues, which is expected.

func parseGolangCILintOutput(output []byte, dir string) (map[Position]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Only include diagnostics in files under the test directory; golangci-lint surfaces
// diagnostics from dependencies (e.g. stdlib, modules in the cache) that the standalone
// driver does not, but those are not part of the test contract.
