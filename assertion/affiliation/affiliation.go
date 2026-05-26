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

// Package affiliation implements the affliation analyzer that tries to find the concrete
// implementation of an interface and create full triggers for them.
package affiliation

import (
	"go/types"

	"go.uber.org/nilaway/annotation"
	"go.uber.org/nilaway/config"
	"go.uber.org/nilaway/util/analysishelper"
	"go.uber.org/nilaway/util/orderedmap"
)

// Affiliation is used to track the association between an interface and its concrete implementations in the form of a map,
// where the key is a function where the affiliation was witnessed and value is an array of full triggers computed for
// the affiliation key
type Affiliation struct {
	conf     *config.Config
	triggers []annotation.FullTrigger
}

// Pair is a struct to store struct-interface affiliation pairs
type Pair struct {
	ImplementedID string
	DeclaredID    string
}

// Cache stores the mapping between interfaces and their implementations that have been analyzed. This
// information can be used by downstream packages to avoid re-analysis of the same affiliations
type Cache struct {
	// Content declares a map of a concrete implementation (e.g., struct) and its implemented interfaces
	// in the form of their fully qualified paths (key: <structFQ>#<interfaceFQ>, value: true/false)
	Content *orderedmap.OrderedMap[Pair, bool]
}

// AFact enables use of the facts passing mechanism in Go's analysis framework
func (*Cache) AFact() {
	_ = "STUB: not implemented"

	// extractAffiliations processes all affiliations (e.g., interface and its implementing struct) and returns map documenting
	// the affiliations
	return
}

func (a *Affiliation) extractAffiliations(pass *analysishelper.EnhancedPass) {
	_ = "STUB: not implemented"
	return
}

// store entries passed from upstream packages
// store new entries witnessed in this current package

// populate upstreamCache by importing entries passed from upstream packages

// export upstreamCache from this package by adding new entries (if any)

// computeTriggersForCastingSites analyzes all explicit and implicit sites of casts in the AST. For example, explicit casts,
// variable assignments, variable declaration and initialization, method returns, and method parameters.
func (a *Affiliation) computeTriggersForCastingSites(pass *analysishelper.EnhancedPass, upstreamCache, currentCache *orderedmap.OrderedMap[Pair, bool]) {
	_ = "STUB: not implemented"
	return
}

// identify sites of explicit or implicit casts

// special case of n-to-1 assignment from a function with multiple returns: e.g., i1, i2 = foo(), where foo() return s1, s2
// note that other n-to-1 assignments (e.g. v, ok := m[k]) are handled by the loop below, since only the first LHS element is
// being directly assigned to in a way we care about

// e.g., var i I, var s *S, i = s, or more generally, i1, i2, i3 = s1, s2, s3

// e.g., var i I = &S{}

// e.g., func foo(i I), foo(&S{})

// receiver param of method declaration
// caller param

// slice is declared to be of interface type, and append function is used to add struct

// e.g., v, ok := i.(*S)

// function signature states interface return, but the actual return is a struct
// e.g., m(x *A) I { return x }

// A slice (or array) declared of type interface, and initialized with a struct
// e.g., _ = []I{&S{}}
// TODO: currently, nested composite literal for ArrayType is not supported (e.g., _ = [][]I{{&A1{}}}).
//  Tracked in issue #46.

// Key, value, or both of a map declared of type interface, and initialized with a struct
// e.g., _ = map[int]I{0: &S{}}

// A struct field (embedded or explicit) declared of type interface, and initialized with a struct
// e.g., var i I = S{t:&T{}}, where `type S struct { t J }`. (Here I and J are interfaces,
// and S and T are structs implementing them, respectively.)
// Similarly, embedding is also supported. E.g., var i I = &S{&T{}}, where `type S struct { J }`.

// In this case the initialization is key-value based. E.g. s = &S{t: &T{}}

// In this case the initialization is serial. E.g. s = &S{&T{}}

// TODO: Nilability analysis support for anonymous functions is currently not
//       implemented (tracked in issue #52), so here we completely skip
//       the affiliation analysis for them.

// computeTriggersForTypes finds corresponding concrete implementation and their declared methods and populates them in a map
func (a *Affiliation) computeTriggersForTypes(lhsType types.Type, rhsType types.Type, upstreamCache, currentCache *orderedmap.OrderedMap[Pair, bool]) []annotation.FullTrigger {
	_ = "STUB: not implemented"
	return nil
}

// Don't process if the affiliation is already analyzed in upstream packages' upstreamCache or
// the current package's upstreamCache.

// Add unvisited entry.

// for each method declared in the interface, find its corresponding concrete implementation

func getFullyQualifiedName(t types.Type) string { _ = "STUB: not implemented"; return "" }

// interface has no exported field/method that can be used to get its fully qualified path directly. However,
// its declared methods (*types.Func) have such exported methods. Therefore, the below logic extracts the
// interface's fully qualified path from its method's FullName()

// funcName.FullName() returns a string of the form "(/path/to/interface).funcName". The below code strips
// off the method name and parentheses to get only "/path/to/interface"

func computeAfflitiationCacheKey(interfaceObj *types.Interface, concreteObj *types.Named) Pair {
	_ = "STUB: not implemented"
	return *new(Pair)
}

// createFunctionTriggers verifies the nilability annotations of the concrete implementation of a method
// against its interface declaration for covariant return types and contravariant parameter types
func createFunctionTriggers(implementingMethod *types.Func, interfaceMethod *types.Func) []annotation.FullTrigger {
	_ = "STUB: not implemented"
	return nil
}

// check for covariance in return types

// check for contravariance in parameter types
