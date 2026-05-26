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

// Package main implements the integration test framework for checking cross-package inference with
// different analyzer drivers. It compares the diagnostics reported by running NilAway separately
// and the diagnostics specified in the comments of the `testdata/integration` project.
// See `testdata/integration/README.md` for more details.
package main

import (
	"fmt"
	"os"
	"regexp"
)

// Position represents a line position in a file.
type Position struct {
	Filename string
	Line     int
}

// Driver is the analyzer driver interface that runs NilAway on the test project.
type Driver interface {
	// Run runs NilAway on the test project specified by dir and returns the diagnostics reported
	// by NilAway (in a map from Position to the diagnostic message).
	Run(dir string) (map[Position]string, error)
}

// CollectGroundTruths collects the ground truths from the test project specified by the "//want"
// comments in the test code (see `testdata/integration` for more details).
func CollectGroundTruths(dir string) (map[Position]*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	// First load all packages in the directory.
	return nil, nil
}

// Traverse all comment nodes and collect corresponding comments with "want" strings.

// CompareDiagnostics compares the ground truths with the collected diagnostics and returns a
// joined error containing the mismatched/missing/unexpected diagnostics (or nil if none).
func CompareDiagnostics(truth map[Position]*regexp.Regexp, collected map[Position]string) (err error) {
	_ = "STUB: not implemented"
	// Keep track of the positions that we have seen.
	return nil
}

// Check for missing diagnostics.

// Run runs the integration test.
func Run() error {
	_ = "STUB: not implemented"
	// Make sure we are at the root of the git repository.
	return nil
}

// Set up the root directory for the integration test project.

// Collect ground truths first.

func main() {
	if err := Run(); err != nil {
		fmt.Printf("FAILED: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("PASSED")
}
