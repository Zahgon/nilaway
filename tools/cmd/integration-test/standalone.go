//  Copyright (c) 2023 Uber Technologies, Inc.
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

// StandaloneDriver implements Driver for running NilAway as a standalone binary.
type StandaloneDriver struct{}

// Run runs NilAway as a standalone binary on the test project and returns the diagnostics.
func (d *StandaloneDriver) Run(dir string) (map[Position]string, error) {
	_ = "STUB: not implemented"
	// Build NilAway first.
	return nil, nil
}

// Run the NilAway binary on the integration test project, with redirects to an internal buffer.

// Disable group error messages to make the output accurate for comparisons.

// Parse the diagnostics.

// pkg name -> "nilaway" -> list of diagnostics.

// Convert diagnostic output from NilAway to canonical form.
