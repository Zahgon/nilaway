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

// Package function implements a sub-analyzer to create full triggers for each function declaration.
package function

import (
	"context"
	"go/ast"
	"go/types"
	"reflect"

	"go.uber.org/nilaway/annotation"
	"go.uber.org/nilaway/assertion/anonymousfunc"
	"go.uber.org/nilaway/assertion/function/assertiontree"
	"go.uber.org/nilaway/assertion/function/functioncontracts"
	"go.uber.org/nilaway/assertion/structfield"
	"go.uber.org/nilaway/config"
	"go.uber.org/nilaway/util/analysishelper"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/ctrlflow"
	"golang.org/x/tools/go/cfg"
)

const _doc = "Build the trees of assertions for each function in this package, propagating them to " +
	"entry and then matching them with possible sources of production to create a list of triggers " +
	"that can then be matched against a set of annotations to generate nil flow errors"

// Analyzer here is the analyzer than generates assertions and passes them onto the accumulator to
// be matched against annotations
var Analyzer = &analysis.Analyzer{
	Name:       "nilaway_function_analyzer",
	Doc:        _doc,
	Run:        analysishelper.WrapRun(run),
	ResultType: reflect.TypeOf((*analysishelper.Result[[]annotation.FullTrigger])(nil)),
	Requires: []*analysis.Analyzer{
		config.Analyzer,
		ctrlflow.Analyzer,
		structfield.Analyzer,
		anonymousfunc.Analyzer,
		functioncontracts.Analyzer,
	},
}

// This limit is in place to prevent the expensive assertions analyzer from being run on
// overly-sized functions. The limit is based on the number of CFG blocks, which provides
// a better measure of function complexity than token/byte count.
const _maxFuncSizeInCFGBlocks = 500

// functionResult is the struct that stores the results for analyzing a function declaration.
type functionResult struct {
	// triggers is the slice of triggers generated from analyzing a particular function.
	triggers []annotation.FullTrigger
	// err stores any error occurred during the analysis.
	err error
	// index is the index of the function declaration in the package. This is particularly
	// important since currently we have hidden coupling in NilAway that requires the generated
	// triggers be placed in order of their declarations. Here, the index will ensure that we can
	// place the triggers in their original order, even though the analyses of function
	// declarations can be parallelized.
	// TODO: remove this.
	index int
	// funcDecl is the function declaration itself.
	funcDecl *ast.FuncDecl
}

func run(p *analysis.Pass) ([]annotation.FullTrigger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Construct experimental features. By default, enable all features on NilAway itself.

//nolint:revive
// TODO: enable struct initialization flag (tracked in Issue #23).
// TODO: enable anonymous function flag.

// Create a fake ident map for the fake func decl nodes to be shared for all function contexts.

// Set up variables for synchronization and communication.

// We use this to keep track of the index of the function declaration we are analyzing.
// TODO: remove this once  is done.

// Skip if a file is marked to be ignored, or it is not in scope of our analysis.

// Collect all function declarations and function literals if anonymous function support
// is enabled.

// We need a stable order of triggers for inference. However, the
// fake func decl nodes generated from the anonymous function analyzer are stored in
// a map. Hence, here we traverse the file and append the fake func decl nodes in
// depth-first order.

// Retrieve the auxiliary information about a function to be analyzed, since it is
// slightly different to do so for function declarations and function literals.

// Skip if function declaration has an empty body.

// Skip if the function is too large based on CFG complexity.
// Use CFG block count as a more accurate measure of function complexity
// than token/byte count, which can be misleading due to comments and formatting.

// Now, analyze the function declarations concurrently.

// Spawn another goroutine that will close the channel when all analyses are done. This makes
// sure the channel receive logic in the main thread (below) can properly terminate.

// Now we collect the results for each function analysis. Note that due to hidden couplings in
// NilAway, the order of the triggers must align with the order of the function declarations (
// as if the analyses were done serially). So we first store the result triggers in order,
// then flatten the slice.
// TODO: remove this extra logic once  is done.

// Duplicate triggers in contracted functions in the callers of the function

// Flatten the triggers

// duplicateFullTriggersFromContractedFunctionsToCallers duplicates all the full triggers that have
// FuncParam producer or UseAsReturn consumer or both, from the contracted functions to the callers
// of all the contracted functions. This is necessary because we have created new
// producers/consumers for argument pass or result return at every call site of the contracted
// function. In order to connect such producers/consumers back to the contracted functions, we
// create new full triggers that duplicates the original full triggers of the contracted functions
// but uses the new producers/consumers instead.
func duplicateFullTriggersFromContractedFunctionsToCallers(
	pass *analysishelper.EnhancedPass,
	funcContracts functioncontracts.Map,
	funcTriggers [][]annotation.FullTrigger,
	funcResults map[*types.Func]*functionResult,
) {
	_ = "STUB: not implemented"

	// Find all the calls to contracted functions
	// callsByCtrtFunc is a mapping: contracted function -> caller -> all the call expressions
	return
}

// TODO: Ideally, we should do
//
// if _, ok := callsByCtrtFunc[ctrFunc]; !ok {
//  callsByCtrtFunc[ctrFunc] = map[*types.Func][]*ast.CallExpr{}
// }
// callsByCtrtFunc[ctrFunc][funcObj] = append(callsByCtrtFunc[ctrFunc][funcObj], call)
//
// However, NilAway complains that callsByCtrtFunc[ctrFunc] can be nil. Thus, we
// introduce an intermediate variable v and the following instead.

// For every contracted function, duplicate some of its full triggers (that involves param or
// return) into all the callers

// The contracted function is imported from upstream, and the local package analysis
// does not involve it.

// If the full trigger has a FuncParam producer or a UseAsReturn consumer, then create
// a duplicated (possibly controlled) full trigger from it and add the created full
// trigger to every caller.

// No need to duplicate the full trigger

// Duplicate the full trigger in every caller

// Store the duplicated full trigger

// Update funcTriggers with duplicated triggers

// Should not happen since we would not have created the duplicated triggers if the
// contracted function is not involved in the analysis of local package.

// duplicateFullTrigger creates a (possibly controlled) full trigger from the given full trigger
// with FuncParam producer or UseAsReturn consumer or both.
// Precondition: isParamProducer or isReturnConsumer is true; also they can be both true.
func duplicateFullTrigger(
	trigger annotation.FullTrigger,
	callee *types.Func,
	callExpr *ast.CallExpr,
	pass *analysishelper.EnhancedPass,
	isParamProducer bool,
	isReturnConsumer bool,
) annotation.FullTrigger {
	_ = "STUB: not implemented"
	// TODO: what if we have more than one parameter, planned in future revisions
	return *new(annotation.FullTrigger)
}

// Create the duplicated full trigger
// TODO: we just copy the pointer for producer and consumer because I don't see a problem when
//  two full triggers share a producer or consumer. We do deep duplication for the param or
//  return related producer/consumer, i.e., FuncParam, FuncReturn, ArgPass, UseAsReturn, and I
//  don't see other conditional producer/consumer that can be shared between two call sites.
//  If we did see such cases in the future, we would want to add a deep copy function for every
//  ProduceTrigger or ConsumeTrigger type and make the deep copy here. In this case, we would
//  also want to see if it is OK to share the underlying site of the producer/consumer in
//  inference engine because we would not want to see a conflict at this site due to different
//  call sites.

// Set up the site that controls the controlled full trigger to be created

// findCallsToContractedFunctions finds all the calls to the contracted functions in the given
// function, and returns a map from every called contracted function to the call expressions that
// call it.
func findCallsToContractedFunctions(
	funcNode *ast.FuncDecl,
	pass *analysishelper.EnhancedPass,
	functionContracts functioncontracts.Map,
) map[*types.Func][]*ast.CallExpr {
	_ = "STUB: not implemented"
	return nil
}

// TODO: for now we find the functions with only a single contract nonnil -> nonnil. If we
//  want to support multiple contracts or contracts with multiple/other values not only we
//  should update here, but we should also make changes to other parts of duplicating
//  triggers.

// hasOnlyNonNilToNonNilContract returns whether the given function has only one contract that is
// nonnil->nonnil.
func hasOnlyNonNilToNonNilContract(funcContracts functioncontracts.Map, funcObj *types.Func) bool {
	_ = "STUB: not implemented"
	return false
}

// analyzeFunc analyzes a given function declaration and emit generated triggers, or an error if
// something went wrong during the analysis. It is mainly a wrapper function for
// assertiontree.BackpropAcrossFunc with synchronization and communication support for concurrency.
// The actual result will be sent via the channel.
func analyzeFunc(
	ctx context.Context,
	pass *analysishelper.EnhancedPass,
	funcDecl *ast.FuncDecl,
	funcContext assertiontree.FunctionContext,
	graph *cfg.CFG,
	index int,
	funcChan chan functionResult,
) {
	_ = "STUB: not implemented"
	// As a last resort, convert the panics into errors and return.
	return
}

// Do the actual backpropagation.

// If any error occurs in back-propagating the function, we wrap the error with more information.
