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

// Package functioncontracts implements a sub-analyzer to analyze function contracts in a package,
// i.e., parsing specified function contracts written as special comments before function
// declarations, or automatically inferring function contracts from the function body.
package functioncontracts

import (
	"go/types"
	"reflect"

	"go.uber.org/nilaway/config"
	"go.uber.org/nilaway/util/analysishelper"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

const _doc = "Read the contracts of each function in this package, returning the results."

// Analyzer here is the analyzer than reads function contracts. It returns the map generated from
// reading the function contracts in the source code.
var Analyzer = &analysis.Analyzer{
	Name:       "nilaway_function_contracts_analyzer",
	Doc:        _doc,
	Run:        analysishelper.WrapRun(run),
	ResultType: reflect.TypeOf((*analysishelper.Result[Map])(nil)),
	FactTypes:  []analysis.Fact{new(Contracts)},
	Requires:   []*analysis.Analyzer{config.Analyzer, buildssa.Analyzer},
}

// Contracts represents the list of contracts for a function.
type Contracts []Contract

// AFact enables use of the facts passing mechanism in Go's analysis framework.
func (*Contracts) AFact() {
	_ = "STUB: not implemented"

	// Map stores the mappings from *types.Func to associated function contracts.
	return
}

type Map map[*types.Func]Contracts

func run(p *analysis.Pass) (Map, error) { _ = "STUB: not implemented"; return *new(Map), nil }

// Collect contracts from the current package.

// The fact mechanism only allows exporting pointer types. However, internally we are using
// `Contract` as a value type because it is an underlying slice type (such that making it a
// pointer type will make the rest of the logic more complicated). Therefore, we strictly
// only convert it from/to a pointer type _here_ during the fact import/exports. Everywhere
// else in NilAway (this sub-analyzer, as well as the other analyzers) we treat `Contract`
// simply as a value type.

// Import contracts from upstream packages and merge it with the local contract map.

// The existing contracts are imported from upstream packages about upstream functions,
// therefore there should not be any conflicts with contracts collected from the current package.

// Now, export the contracts for the _exported_ functions in the current package only.

// Check if the function is (1) exported by name (i.e., starts with a capital letter), (2)
// it is directly inside the package scope (such that it is really visible in downstream
// packages).

// fn.Scope() -> the scope of the function body.

// fn.Scope().Parent() -> the scope of the file.

// fn.Scope().Parent().Parent() -> the scope of the package.

// functionResult is the struct that is received from the channel for each function.
type functionResult struct {
	funcObj   *types.Func
	contracts Contracts
	err       error
}

// collectFunctionContracts collects all the function contracts and returns a map that associates
// every function with its contracts if it has any. We prefer to parse handwritten contracts from
// the comments at the top of each function. Only when there are no handwritten contracts there,
// do we try to automatically infer contracts.
func collectFunctionContracts(pass *analysishelper.EnhancedPass) (Map, error) {
	_ = "STUB: not implemented"
	return *new(Map), nil
}

// Collect ssa for every function.

// should be guaranteed to be non-nil; otherwise it would have paniced in the library
// https://cs.opensource.google/go/x/tools/+/refs/tags/v0.12.0:go/analysis/passes/buildssa/buildssa.go;l=99

// Set up variables for synchronization and communication.

// Ignore any non-function declaration
// TODO: If we want to support contracts for anonymous functions (function
//  literals) in the future, then we need to handle more types here.

// First, we try to parse the contracts from the comments at the top of the function.
// If there are any, we do not need to infer contracts for this function.

// If we reach here, it means that there are no handwritten contracts for this
// function. We need to infer contracts for this function.

// We definitely want to ignore any function without any parameters or return
// values since they cannot have any contracts.

// TODO: However, we want to analyze for multiple param/return in the future; for
//  now we consider contract(nonnil->nonnil) only.

// TODO: If the function has only one parameter and the parameter is variadic, then
//  it may happen that no argument is passed when calling the function. Such cases
//  are not handled well when duplicating full triggers from contracted functions,
//  so we don't infer contract(nonnil->nonnil) for such a function although we can
//  already.

// For some reason, we cannot find the ssa for this function. We ignore this
// function.

// For external functions whose function bodies are defined outside Go (e.g.,
// assembly), we do not actually have Go source code for them, and there will be no
// blocks (see the documentation of the ssa package). Therefore, we ignore such functions.

// Infer contracts for a function that does not have any contracts specified.

// As a last resort, convert the panics into errors and return.

// Spawn another goroutine that will close the channel when all analyses are done. This makes
// sure the channel receive logic in the main thread (below) can properly terminate.

// Collect inferred contracts from the channel.
