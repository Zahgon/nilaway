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

// Package diagnostic hosts the diagnostic engine, which is responsible for collecting the
// conflicts from annotation-based checks (no-infer mode) and/or inference (full-infer mode) and
// generating user-friendly diagnostics from those conflicts.
package diagnostic

import (
	"go/token"

	"go.uber.org/nilaway/annotation"
	"go.uber.org/nilaway/inference"
	"go.uber.org/nilaway/util/analysishelper"
	"golang.org/x/tools/go/analysis"
)

// fileInfo bundles the token.File object and auxiliary information about it, e.g., whether it is
// a fake file (i.e., imported from archive), for uses in primitivizer.
type fileInfo struct {
	file   *token.File
	isFake bool
}

// Engine is the main engine for generating diagnostics from conflicts.
type Engine struct {
	pass      *analysishelper.EnhancedPass
	conflicts []conflict
	// files maps the file name (modulo the possible build-system prefix) to the token.File object
	// for faster lookup when converting correct upstream position back to local token.Pos for
	// reporting purposes.
	files map[string]fileInfo
}

// NewEngine creates a new diagnostic engine.
func NewEngine(pass *analysishelper.EnhancedPass) *Engine {
	_ = "STUB: not implemented"
	// Iterate all files within the Fset (which includes upstream and current-package files), and
	// store the mapping between its file name (modulo the possible build-system prefix) and the
	// token.File object. This is needed for converting correct upstream position back to local
	// incorrect token.Pos for error reporting purposes. Also see
	// [inference.primitivizer.toPosition] for more detailed explanations.
	return nil
}

// For files that are not in the execroot (e.g., stdlib files start with "$GOROOT", and
// upstream files that do not have the build-system prefix), it simply returns the original.

// The file will be fake (conceptually "\n" * 65535) if it is imported from archive. So we
// check if there are any gaps between the line starts to determine if the file is fake.

// Diagnostics generates diagnostics from the internally-stored conflicts. The grouping parameter
// controls whether the conflicts with the same nil flow -- the part in the complete nil flow going
// from a nilable source point to the conflict point -- are grouped together (under the first
// diagnostic) for concise reporting. The returned slice of diagnostics are sorted by file names
// and then offsets in the file.
func (e *Engine) Diagnostics(grouping bool) []analysis.Diagnostic {
	_ = "STUB: not implemented"
	// First sort the conflicts by position such that similar conflicts are grouped under the
	// first diagnostic.
	return nil
}

// Group conflicts with the same nil path together for concise reporting.

// Build diagnostics from conflicts. Apply cross-package nolint suppressions here as well.

// AddSingleAssertionConflict adds a new single assertion conflict to the engine.
func (e *Engine) AddSingleAssertionConflict(trigger annotation.FullTrigger) {
	_ = "STUB: not implemented"
	return
}

// Try to trim the build system prefix (i.e., the current working directory) if present.

// AddOverconstraintConflict adds a new overconstraint conflict to the engine.
func (e *Engine) AddOverconstraintConflict(nilReason, nonnilReason inference.ExplainedBool) {
	_ = "STUB: not implemented"

	// Build nil path by traversing the inference graph from `nilReason` part of the overconstraint failure.
	// (Note that this traversal gives us a backward path from point of conflict to the source of nilability. Hence, we
	// must take this into consideration while printing the flow, which is currently being handled in `addNilPathNode()`.)
	return
}

// We have two cases here:
// 1. No annotation present (i.e., full inference): we have producer and consumer explanations available; use them directly
// 2: Annotation present (i.e., no inference): we construct the reason from the annotation string

// Build nonnil path by traversing the inference graph from `nonnilReason` part of the overconstraint failure.
// (Note that this traversal is forward from the point of conflict to dereference. Hence, we don't need to make
// any special considerations while printing the flow.)
// Different from building the nil path above, here we also want to deduce the position where the error should be reported,
// i.e., the point of dereference where the nil panic would occur. In NilAway's context this is the last node
// in the non-nil path. Therefore, we keep updating `c.pos` until we reach the end of the non-nil path.

// Similar to above, we have two cases here:
// 1. No annotation present (i.e., full inference): we have producer and consumer explanations available; use them directly
// 2: Annotation present (i.e., no inference): we construct the reason from the annotation string

// _fakeFileMaxLines is the maximum number of lines that the archive importer will add to a (fake)
// file when it imports a package. See [the importer code] for more details. We use this to create
// more fake files when necessary (see [primitivizer.sitePos]).
// [the importer code]: https://cs.opensource.google/go/x/tools/+/master:internal/gcimporter/bimport.go;l=34;bpv=0;bpt=1
const _fakeFileMaxLines = 64 * 1024

// toPos converts the token.Position back to a token.Pos that is relative to local Fset for
// reporting purposes _only_. Note that the input position could be obtained from facts or
// inference, so the position might not exist in the local Fset. In such cases, we pad the local
// Fset for correct reporting.
func (e *Engine) toPos(position token.Position) token.Pos {
	_ = "STUB: not implemented"
	return *new(token.Pos)
}

// For incremental build systems like bazel, the pass.Fset contains only the files in
// current and _directly_ imported packages (see [gcexportdata] for more details). However,
// analyzer facts are imported transitively from all imported packages, and NilAway is able
// to operate across all those packages. As a result, if NilAway ever needs to report an
// error on a file from a transitively imported package, we need to create a fake file in
// the file set.
// [gcexportdata]: https://pkg.go.dev/golang.org/x/tools/go/gcexportdata

// Set up fake lines for the fake file.

// If the file is fake (imported from archive), it may not contain fake lines for unexported
// objects (as an "optimization", see [importer code]). However, NilAway may report errors
// on unexported objects due to multi-package inference. In such cases, we pad the file with
// more fake lines.
// [importer code]: https://cs.opensource.google/go/x/tools/+/refs/tags/v0.12.0:internal/gcimporter/bimport.go;l=36-69;drc=ad74ff6345e3663a8f1a4ba5c6e85d54a6fd5615

// We are adding offsets to fake lines here, and offset == fake line number - 1. So we
// can start from the current max line number to the desired line number - 1.

// For fake files, we can only report accurate line number but not column number.

// For non-fake files, the position is accurate.

// involvesTestFile returns true if the conflict's report position or any node position in the
// nil flow originates from a test file (i.e., a file ending with "_test.go").
func involvesTestFile(c conflict) bool { _ = "STUB: not implemented"; return false }
