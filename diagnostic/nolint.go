//	Copyright (c) 2025 Uber Technologies, Inc.
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

package diagnostic

import (
	"reflect"

	"go.uber.org/nilaway/util/analysishelper"
	"golang.org/x/tools/go/analysis"
)

// NoLintAnalyzer is an analyzer that reads all NilAway's nolint comments. This is needed since NilAway
// is able to report upstream violations when analyzing downstream packages. Currently, most
// drivers are not able to respect this (they only respect nolint comments on current packages).
// Therefore, inside NilAway we parse the nolint comments, export them as facts, and then do the
// filtering ourselves in [diagnostic.Engine].
var NoLintAnalyzer = &analysis.Analyzer{
	Name:       "nilaway_nolint_analyzer",
	Doc:        "Read NilAway's nolint comments and export them as facts for NilAway's diagnostic engine.",
	Run:        analysishelper.WrapRun(run),
	FactTypes:  []analysis.Fact{new(NoLint)},
	Requires:   []*analysis.Analyzer{},
	ResultType: reflect.TypeOf((*analysishelper.Result[[]Range])(nil)),
}

// NoLint is a fact that stores the ranges of "//nolint:nilaway" comments for cross-package nolint
// suppression support.
type NoLint struct {
	// Ranges lists the ranges of the nolint scopes in the package.
	Ranges []Range
}

// AFact makes NoLint satisfy the analysis.Fact interface such that it can be exported as a fact.
func (*NoLint) AFact() {
	_ = "STUB: not implemented"

	// Range is a minimal struct that stores the filename and the start and end lines of a nolint scopes.
	return
}

type Range struct {
	// Filename is the filename of the file where the nolint comment is located.
	Filename string
	// From and To are the start and end lines of the nolint scope.
	From, To int
}

func run(p *analysis.Pass) ([]Range, error) { _ = "STUB: not implemented"; return nil, nil }

// CommentMap will correctly associate comments to the largest node group
// applicable. This handles inline comments that might trail a large
// assignment and will apply the comment to the entire assignment.

// Import all nolint ranges from upstream.

// Export local nolint ranges (if available) for downstream uses.

// nolintContainsNilAway checks if the particular comment is a nolint comment for NilAway suppression.
func nolintContainsNilAway(text string) bool {
	_ = "STUB: not implemented"
	// This implementation is adapted from
	// https://github.com/bazel-contrib/rules_go/blob/eb13b736d9568044427f23359329155e67071948/go/tools/builders/nolint.go#L21
	// under Apache 2.0 license.
	return false
}

// strip explanation comments
