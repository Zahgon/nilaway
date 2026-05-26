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

// Package main implements the golden tests for NilAway to ensure that the errors reported on the
// stdlib are equal between the base branch and the test branch for preventing functionality
// regressions during development.
package main

import (
	"errors"
	"flag"
	"io"
	"log"
	"os"
	"os/exec"
)

// Diagnostic is the diagnostic reported by NilAway.
type Diagnostic struct {
	// Posn is the position string of the diagnostic.
	Posn string `json:"posn"`
	// Message is the message reported by NilAway.
	Message string `json:"message"`
}

// BranchResult stores the information about a branch, and the diagnostics reported on that branch.
type BranchResult struct {
	// Name is the friendly name of the branch (if available and not "HEAD", otherwise it is equal
	// to its ShortSHA).
	Name string
	// ShortSHA is the short SHA of the branch.
	ShortSHA string
	// Result is the set of diagnostics NilAway reported on the branch.
	Result map[Diagnostic]bool
}

// Run runs the golden tests on the base branch and the test branch and writes the summary and
// diff to the writer.
func Run(writer io.Writer, baseBranch, testBranch string) error {
	_ = "STUB: not implemented"
	// First verify that the git repository is clean.
	return nil
}

// Then verify that we are at the root of the git project.

// Get the current branch name and switch back to it after the golden test.

// If test branch is not specified, use the current branch.

// Initialize the base and test branch SHAs.

// Now the golden test starts. From here on, we should use the `branches` variable to refer to
// the base and test branches.

// Run the built NilAway binary on the stdlib and parse the diagnostics.

// Inherit env vars such that users can control the resource usages via GOMEMLIMIT, GOGC
// etc. env vars.

// ParseDiagnostics parses the diagnostics from the raw JSON output of NilAway and returns the
// set of diagnostics.
func ParseDiagnostics(reader io.Reader) (map[Diagnostic]bool, error) {
	_ = "STUB: not implemented"
	// Package name -> "nilaway" -> slice of diagnostics.
	return nil, nil
}

// WriteDiff writes the summary and the diff (if the base and test are different) between the base
// and test diagnostics to the writer. If the writer is os.Stdout, it will write the diff in color.
func WriteDiff(writer io.Writer, branches [2]*BranchResult) {
	_ = "STUB: not implemented"
	// Compute the diagnostic differences between base and test branches.
	return
}

// Write the summary lines first.

// Optionally write the direction of the change (if present).

// Now write the statistics of the diagnostics in each branch.

// Early return if there is no diff to write.

// If the writer is os.Stdout, we will write the diff in color.

// Add Posn to the first line and prefix to each line for diff formatting.

// Diff computes the diff between the first and second diagnostics and returns the diff in
// alphabetical order.
func Diff(first, second map[Diagnostic]bool) []Diagnostic { _ = "STUB: not implemented"; return nil }

// Sort the diff such that we have stable ordering for the same runs.

// MustFprint is a helper function that takes the result of the family of Fprint functions and
// panics if the error is nonnil.
func MustFprint(_ int, err error) { _ = "STUB: not implemented"; return }

func main() {
	fset := flag.NewFlagSet("golden-test", flag.ExitOnError)
	baseBranch := fset.String("base-branch", "main", "the base branch to compare against")
	testBranch := fset.String("test-branch", "", "the test branch to run golden tests (default current branch)")
	resultFile := fset.String("result-file", "", "the file to write the diff to, default stdout")
	if err := fset.Parse(os.Args[1:]); err != nil {
		log.Printf("failed to parse flags: %v\n", err)
		flag.PrintDefaults()
		os.Exit(1)
	}

	writer := os.Stdout
	if *resultFile != "" {
		w, err := os.OpenFile(*resultFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatalf("failed to open file %q: %v", *resultFile, err)
		}
		writer = w
	}

	if err := Run(writer, *baseBranch, *testBranch); err != nil {
		log.Printf("failed to run golden test: %v", err)
		var e *exec.ExitError
		if errors.As(err, &e) {
			log.Printf("failed command output: %v", e.Stderr)
		}
		os.Exit(1)
	}
}
