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

package inference

import (
	"go/types"

	"go.uber.org/nilaway/annotation"
	"go.uber.org/nilaway/util/analysishelper"
	"go.uber.org/nilaway/util/orderedmap"
)

// An InferredMap is the state accumulated by multi-package inference. It's
// field `Mapping` maps a set of known annotation sites to InferredAnnotationVals - which can
// be either a fixed bool value along with explanation for why it was fixed - an DeterminedVal
// - or an UndeterminedVal indicating that site's place in the known implication graph
// between underconstrained sites. The set of sites mapped to UndeterminedBoolVals is guaranteed
// to be closed under following `Implicant`s and `Implicate`s pointers.
//
// Additionally, a field upstreamMapping is stored indicating a stable copy of the information
// gleaned from upstream packages. Both mapping and upstreamMapping are initially populated
// with the same informations, but observation functions (observeSiteExplanation and observeImplication)
// add information only to Mapping. On export, iterations combined with calls to
// inferredValDiff on shared keys is used to ensure that only
// information present in `Mapping` but not `UpstreamMapping` is exported.
type InferredMap struct {
	primitive       *primitivizer
	upstreamMapping map[primitiveSite]InferredVal
	mapping         *orderedmap.OrderedMap[primitiveSite, InferredVal]
}

// newInferredMap returns a new, empty InferredMap.
func newInferredMap(primitive *primitivizer) *InferredMap { _ = "STUB: not implemented"; return nil }

// AFact allows InferredAnnotationMaps to be imported and exported via the Facts mechanism.
func (*InferredMap) AFact() {
	_ = "STUB: not implemented"

	// Load returns the value stored in the map for an annotation site, or nil if no value is present.
	// The ok result indicates whether value was found in the map.
	return
}

func (i *InferredMap) Load(site primitiveSite) (value InferredVal, ok bool) {
	_ = "STUB: not implemented"
	return *new(InferredVal), false
}

// StoreDetermined sets the inferred value for an annotation site.
func (i *InferredMap) StoreDetermined(site primitiveSite, value ExplainedBool) {
	_ = "STUB: not implemented"
	return
}

// StoreImplication stores an implication edge between the `from` and `to` annotation sites in the
// graph with the assertion for error reporting.
func (i *InferredMap) StoreImplication(from primitiveSite, to primitiveSite, assertion primitiveFullTrigger) {
	_ = "STUB: not implemented"
	// First create UndeterminedVal in the map if it does not exist yet.
	return
}

// Len returns the number of annotation sites currently stored in the map.
func (i *InferredMap) Len() int { _ = "STUB: not implemented"; return 0 }

// OrderedRange calls f sequentially for each annotation site and inferred value present in the map
// in insertion order. If f returns false, range stops the iteration.
func (i *InferredMap) OrderedRange(f func(primitiveSite, InferredVal) bool) {
	_ = "STUB: not implemented"
	return
}

// Export only encodes new information not already present in the upstream maps, and it does not
// encode all (in the go sense; i.e. capitalized) annotation sites (See chooseSitesToExport).
// This ensures that only _incremental_ information is exported by this package and plays a _vital_
// role in minimizing build output.
func (i *InferredMap) Export(pass *analysishelper.EnhancedPass) { _ = "STUB: not implemented"; return }

// If we are testing, we encode and decode the inferred map to ensure that the gob encoding
// works correctly (i.e., there are no un-registered types to Gob encoding).
// This is a little hacky given that this should belong to the test logic instead of production
// logic. However, our current analyzer architecture (i.e., the accumulation.Analyzer generates
// the diagnostics and exports the facts, where the top-level nilaway.Analyzer only does the
// reporting) prevents us from accessing the facts of accumulation.Analyzer in the test logic,
// since facts are assumed to be "private" to an analysis. In the meantime, we do not want to
// merge the accumulation.Analyzer and the top-level nilaway.Analyzer since the `analysistest`
// framework will then require us to write "want" strings for facts as well.
// We also encode/decode the entire map instead of the incremental map to have as much coverage
// as possible.

// First create a new map containing only the sites and their inferred values that we would
// like to export.

// We do not need to encode the primitivizer since it is just a helper for the analysis of
// the current package.
/* primitive */

// GobEncode encodes the inferred map via gob encoding.
func (i *InferredMap) GobEncode() (b []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

// Close the s2 writer before getting the bytes such that we have complete information.

// GobDecode decodes the InferredMap from buffer.
func (i *InferredMap) GobDecode(input []byte) error { _ = "STUB: not implemented"; return nil }

// chooseSitesToExport returns the set of AnnotationSites mapped by this InferredMap that are both
// reachable from and that reach an Exported (in the go sense; i.e. capitalized) site. We define
// reachability  here to be reflexive, and we choose this definition so that the returned set is
// convex -guaranteeing that we never forget a semantically meaningful implication - yet minimal -
// containing no site that could be forgotten without sacrificing soundness
func (i *InferredMap) chooseSitesToExport() map[primitiveSite]bool {
	_ = "STUB: not implemented"
	return nil
}

// Mark the current site as to be exported.

// For UndeterminedVal, we visit the implicants and implicates recursively and mark
// them as to be exported as well.

// The following method implementations make InferredMap satisfy the annotation.Map
// interface, so that triggers can be checked against it.

// CheckFieldAnn checks this InferredMap for a concrete mapping of the field key provided
func (i *InferredMap) CheckFieldAnn(fld *types.Var) (annotation.Val, bool) {
	_ = "STUB: not implemented"
	return *new(annotation.Val), false
}

// CheckFuncParamAnn checks this InferredMap for a concrete mapping of the param key provided
func (i *InferredMap) CheckFuncParamAnn(fdecl *types.Func, num int) (annotation.Val, bool) {
	_ = "STUB: not implemented"
	return *new(annotation.Val), false
}

// CheckFuncRetAnn checks this InferredMap for a concrete mapping of the return key provided
func (i *InferredMap) CheckFuncRetAnn(fdecl *types.Func, num int) (annotation.Val, bool) {
	_ = "STUB: not implemented"
	return *new(annotation.Val), false
}

// CheckFuncRecvAnn checks this InferredMap for a concrete mapping of the receiver key provided
func (i *InferredMap) CheckFuncRecvAnn(fdecl *types.Func) (annotation.Val, bool) {
	_ = "STUB: not implemented"
	return *new(annotation.Val), false
}

// CheckDeepTypeAnn checks this InferredMap for a concrete mapping of the type name key provideed
func (i *InferredMap) CheckDeepTypeAnn(name *types.TypeName) (annotation.Val, bool) {
	_ = "STUB: not implemented"
	return *new(annotation.Val), false
}

// CheckGlobalVarAnn checks this InferredMap for a concrete mapping of the global variable key provided
func (i *InferredMap) CheckGlobalVarAnn(v *types.Var) (annotation.Val, bool) {
	_ = "STUB: not implemented"
	return *new(annotation.Val), false
}

// CheckFuncCallSiteParamAnn checks this InferredMap for a concrete mapping of the call site param
// key provided.
func (i *InferredMap) CheckFuncCallSiteParamAnn(key *annotation.CallSiteParamAnnotationKey) (annotation.Val, bool) {
	_ = "STUB: not implemented"
	return *new(annotation.Val), false
}

// CheckFuncCallSiteRetAnn checks this InferredMap for a concrete mapping of the call site return
// key provided.
func (i *InferredMap) CheckFuncCallSiteRetAnn(key *annotation.CallSiteRetAnnotationKey) (annotation.Val, bool) {
	_ = "STUB: not implemented"
	return *new(annotation.Val), false
}

func (i *InferredMap) checkAnnotationKey(key annotation.Key) (annotation.Val, bool) {
	_ = "STUB: not implemented"
	return *new(annotation.Val), false
}
