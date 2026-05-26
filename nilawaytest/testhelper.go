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

// Package nilawaytest implements utility functions for tests.
package nilawaytest

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// FindExpectedValues inspects test files and gathers expected values' comment strings
func FindExpectedValues(pass *analysis.Pass, expectedPrefix string) map[ast.Node][]string {
	_ = "STUB: not implemented"
	return nil
}

// Store a mapping between single comment's line number to its text.

// Now, find all nodes of interest, such as *ast.FuncLit and *ast.FuncDecl, and find their comment.

// It is ok to not leave annotations for a node - it simply does not use
// any closure variables. We still need to traverse further since there could be
// comments for nested func lit nodes.

// Trim the trailing slashes and extra spaces and extract the set of expected values.

// If no expected values are written after the `expectedPrefix`, we simply ignore it.
