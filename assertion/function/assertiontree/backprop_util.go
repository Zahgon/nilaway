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

package assertiontree

import (
	"go/ast"
	"go/types"

	"go.uber.org/nilaway/annotation"
	"go.uber.org/nilaway/util/analysishelper"
	"golang.org/x/tools/go/cfg"
)

// If `block` is a conditional branch (e.g. an if statement), return the expression on which it
// branches, otherwise return nil.
// nilable(result 0)
func getConditional(block *cfg.Block) ast.Expr {
	_ = "STUB: not implemented"
	// TODO: Nilness check for `block` is currently needed due to a FP, but should not be needed
	//
	//	after  is implemented.
	return *new(ast.Expr)
}

// If `block` is the precursor to a range statement, return the expression being ranged over,
// otherwise return nil.
//
// This function matches on two cases:
//   - a block terminating with an `*ast.AssignStmt` whose singular rhs is an `*ast.UnaryExpr` with operation
//     `token.RANGE` - i.e. `[x | x, y] = range z`
//   - a block terminating with an `*ast.UnaryExpr` directly, whose operation is `token.RANGE`
//
// Both of these forms are inserted during the pass in assertion.markRangeStatements when that pass
// determines it has found the output of CFG-parsing an `*ast.RangeStmt`. These two functions can thus
// be seen as a direct input/output pair.
//
// nilable(result 0)
func getRangeExpr(block *cfg.Block) ast.Expr { _ = "STUB: not implemented"; return *new(ast.Expr) }

// if we are in the former case described above, strip the assignment and focus only on
// its rhs

// check that the last node, or the rhs of the last node if it is an assignment, is a range expression

// we've matched the given block with one of the two desired cases, so we return
// the expression being ranged over

// a preprocessPair bundles a function `trueBranchFunc` that modifies a *RootAssertionNode from the true
// branch of a conditional with a function `falseBranchFunc` that modifies the false branch
type preprocessPair struct {
	trueBranchFunc  RootFunc
	falseBranchFunc RootFunc
}

const knownNilableErrFunc = "sometimesErrs"

// exprCallsKnownNilableErrFunc checks if expression calls a function that we know to be nilable without
// needing to consult annotations.
//
// The best mechanism for this would be to somehow expose a fixed library function that serves this
// purpose, but for now, we simply check that it has the special name "sometimesErrs" set above through
// the constant `knownNilableErrFunc`
func exprCallsKnownNilableErrFunc(expr ast.Expr) bool { _ = "STUB: not implemented"; return false }

// no ident - anonymous function

// For a return statement - make sure all returned results are computable by generating the
// appropriate assertions, and consume each as the respective return number of that function
// this indicates the "normal" case of backprop across return statements, and is called
// from backpropAcrossReturn when the interesting cases like a one-to-many return are eliminated.
// Much of the logic in this function deals with error-returning functions
// in particular, this function is responsible for splitting returns into the cases:
// 1: Normal Return - all results yield consume triggers eventually enforcing their annotated/inferred nilability
// 2: Error Return - consume triggers are created based on the error contract. i.e., based on the nilabiity status of the error return expression
// 3. Ok return - consume triggers are created based on the nilability status of the boolean (`ok`) return expression
func computeAndConsumeResults(rootNode *RootAssertionNode, node *ast.ReturnStmt) error {
	_ = "STUB: not implemented"
	// no matter what case the consumption of these returns ends up as - each must be computed
	return nil
}

// check if this is a named return case -- where an empty *ast.ReturnStmt shows up even in functions that have
// a nonzero number of results

// flatten return variables in the signature

// if the function has named error return variable, then handle specially using the error handling logic
/* isNamedReturn */

/* isNamedReturn */

// below is the normal handling for named return variables

// default handling if retVariable is not a blank identifier (e.g., i *int)

/* isNamedReturn */

// special handling if retVariable is a blank identifier (e.g., _ *int)

// We're returning a multiply returning function, but one that couldn't be parsed in
// backpropAcrossReturn (likely due to being anonymous)
// there is no consumption we can compute here, so abort

/* isNamedReturn */

/* isNamedReturn */

// we've excluded all abnormal cases - here, just really consume each result as a return value

/* isNamedReturn */

// isErrorReturnNil returns true if the error return is guaranteed to be nil, false otherwise
func isErrorReturnNil(rootNode *RootAssertionNode, errRet ast.Expr) bool {
	_ = "STUB: not implemented"
	return false
}

// error return is the literal nil

// check for false cases where error return value may be nil

// error result is a blank named return ("_ error"), so it's always nil

// error value is the return of a known nilable function

// isErrorReturnNonnil returns true if the error return is guaranteed to be nonnil, false otherwise
func isErrorReturnNonnil(rootNode *RootAssertionNode, errRet ast.Expr) bool {
	_ = "STUB: not implemented"
	return false
}

// handleErrorReturns handles the special case for error returning functions (n-th result of type `error` which guards at least one of the first n-1 non-error results).
// It generates consumers by applying the error contract:
// (1) if error return value = nil, create consumers for the non-error returns
// (2) if error return value = non-nil, create consumer for error return
// (3) if error return value = unknown, create consumers for all returns (error and non-error), and defer applying of the error contract when the nilability status is known, such as at `ProcessEntry`
//
// Note that `results` should be explicitly passed since `retStmt` of a named return will contain no results
func handleErrorReturns(rootNode *RootAssertionNode, retStmt *ast.ReturnStmt, results []ast.Expr, isNamedReturn bool) bool {
	_ = "STUB: not implemented"
	return false
}

// n-th expression
// n-1 expressions

// default tracking to support potential "always safe" cases

// check if the error return is at all guarding any nilable returns, such as pointers, maps, and slices

// if error is the only return expression in the statement, then create a consumer for it, else create consumers for the non-error return expressions

// create general return consume triggers for all n-1 (non-error) return expressions

// TODO: handle struct init in the context of error return in a better way in a follow up diff

// create consume trigger for only the error return

// the nilability of error return is unknown, hence create special consume triggers for all returns

// TODO: handle struct init in the context of error return in a better way in a follow up diff

// handleBooleanReturns handles the special case for boolean (`ok`) returning functions (n-th result of type `bool`
// which guards at least one of the first n-1 non-bool results). Similar to the handling of error returning functions,
// for boolean returns, we generate consumers by applying the following boolean contract:
// (1) if boolean return value = true, create consumers for the non-boolean returns
// TODO: currently we support only explicit boolean returns (i.e., `return r0, r1, ..., {true|false}`). We should also support implicit boolean returns, i.e., `return` or `return <expr>` in the future.
//
// handleBooleanReturns returns true if the above contract is satisfied and consumers are created, false otherwise
func handleBooleanReturns(rootNode *RootAssertionNode, retStmt *ast.ReturnStmt, results []ast.Expr, isNamedReturn bool) bool {
	_ = "STUB: not implemented"
	// FuncIsOkReturning checks that the length of the results defined for the current function is at least 2, and that
	// the last return type is a boolean, the value of which can be determined at compile time (e.g., return true)
	return false
}

// n-th expression
// n-1 expressions

// check if the return statement is of the currently supported explicit boolean return form (`return ..., {true|false}`)

// default tracking to support potential "always safe" cases

// If return is "true", then track its n-1 returns. Create return consume triggers for all n-1 return expressions.
// If return is "false", then do nothing, since we don't track boolean values.

// createConsumerForErrorReturn creates a consumer for the error return enforcing it to be non-nil
func createConsumerForErrorReturn(rootNode *RootAssertionNode, errRetExpr ast.Expr, errRetIndex int, retStmt *ast.ReturnStmt, isNamedReturn bool) {
	_ = "STUB: not implemented"
	return
}

// createGeneralReturnConsumers creates general return consumers for the non-return expressions in the return statement
func createGeneralReturnConsumers(rootNode *RootAssertionNode, results []ast.Expr, retStmt *ast.ReturnStmt, isNamedReturn bool) {
	_ = "STUB: not implemented"
	return

	// don't do anything if the expression is a blank identifier ("_")
}

// createReturnConsumersForAlwaysSafe creates return consumers for the non-return expressions in the return statement
// for tracking potential "always safe" cases
func createReturnConsumersForAlwaysSafe(rootNode *RootAssertionNode, nonErrResults []ast.Expr, retStmt *ast.ReturnStmt, isNamedReturn bool) {
	_ = "STUB: not implemented"
	return
}

// don't do anything if the expression is a blank identifier ("_")

// createSpecialConsumersForAllReturns conservatively creates specially designed consumers for all return expressions, error and non-error
func createSpecialConsumersForAllReturns(rootNode *RootAssertionNode, nonErrRetExpr []ast.Expr, errRetExpr ast.Expr, errRetIndex int, retStmt *ast.ReturnStmt, isNamedReturn bool) {
	_ = "STUB: not implemented"
	return
}

// don't do anything if the expression is a blank identifier ("_")

func typeIsString(t types.Type) bool { _ = "STUB: not implemented"; return false }

// some expressions consume their subexpressions specifically when assigned to - for now, we are
// aware only of map indices written to as having this behavior
// exprAsConsumedByAssignment recognizes these cases, and returns the corresponding consumeTrigger
// if one is found, otherwise returning `nil, false`
// nilable(result 0)
func exprAsConsumedByAssignment(rootNode *RootAssertionNode, expr ast.Node) *annotation.ConsumeTrigger {
	_ = "STUB: not implemented"
	return nil
}

// exprAsAssignmentConsumer is similar to parseExprAsProducer, but tries to parse the passed
// expression as a _consumer_ instead of as a _producer_. The simplest illustrative example of
// this is when a field read expression is passed as `expr` - meaning a field is being assigned
// to - such as `x.f = y`. This will result in an `annotation.FldAssign` being returned, which
// will serve to produce an error if that field is non-nil and a nilable value flows into it
// through the assignment triggering this call to exprAsAssignmentConsumer.
// other notable cases include passing a send expression here (which is why we take an `ast.Node`
// not `ast.Expr`, and various "deep" assignments such as to an index of an object
// nilable(result 0)
func exprAsAssignmentConsumer(rootNode *RootAssertionNode, expr ast.Node, exprRHS ast.Node) (annotation.ConsumingAnnotationTrigger, error) {
	_ = "STUB: not implemented"
	return *new(annotation.ConsumingAnnotationTrigger), nil
}

// we've found an assignment to a global

// we've found an assignment to a parameter with deep type - have to check its deep annotation!

// but first - if it's a variadic parameter then its "deep" annotation is really just
// its shallow annotation:

// we've concluded it's not a variadic parameter

// we've found an assignment to a global var with deep type - have to check its deep annotation!

// this is an assignment to an index of a field

// check if this is a call to a function by name

// Calling Underlying on [types.Named] will always return the unnamed type, so we
// do not have to recursively "unwrap" the [types.Named].
// See [https://github.com/golang/example/tree/master/gotypes#named-types].

// at this point - the value being deeply assigned to is of deep type but is not linked
// to an annotation site, for example, local variables.
// so we introspect on its type alone

// This block checks if the rhs of the assignment is the builtin append function for slices.

// If there is a deep assignment to a slice using append method

// If field access for a variable that is not a global var we rely on default field nilability based on
// escape analysis, and thus we do not create any triggers for field assignments.
// For global variables we still maintain the previous behaviour. Thus do not return anything.
// For a global variable g, `g.f = nil` would result in a const nil field assignment trigger.
// However, for other type of variables `p.f = nil` would result into an escape trigger only if the
// field escapes as per the definition of field escape in our analysis.

// no recognized source of deep nilability consumption

func composeRootFuncs(f1, f2 RootFunc) RootFunc { _ = "STUB: not implemented"; return *new(RootFunc) }

// This takes a cfg, and generates the information we need from it:
//  1. its set of blocks, but with a "return" block appended that's a successor of every block that returns
//     we need this as an index of where to start our backpropagation
//  2. for conditional branching blocks, add "pre-processing" to insert nil-checks corresponding to
//     their branch condition. If blocks[i] is a conditional, then preprocessing[i].trueBranchFunc will
//     be a function to insert the true result of the check and preprocessing[i].falseBranchFunc will
//     be a function to insert the false result
//
// The `richCheckBlocks` that it takes represents, for each block, which richCheckEffects hold at
// the end of that block
//
// postcondition - length of two return slices is equal
func blocksAndPreprocessingFromCFG(pass *analysishelper.EnhancedPass, graph *cfg.CFG, richCheckBlocks [][]RichCheckEffect) (
	[]*cfg.Block, []*preprocessPair) {
	_ = "STUB: not implemented"
	return nil, nil

	// add an empty "return" block
}

// add "return" block as a successor for:
// - all returning blocks
// - while loops (`for <EXPR> {}` or `for {}`)

// TODO: storing `blocks[i].Succs` in a local variable should not be needed. But NilAway complaints about slicing
//  of the field `blocks[i].Succs` in the if condition. This should be fixed.

// generate pre-processing

// blocks[i] is a conditional

// so add nil check productions to each successor
// this is where the assumption that True Name = Succs[0], False Name = Succs[1] shows up

// we've discovered that this is a nil check

// now check for RichCheckEffects triggered by this conditional

// blocks[i] is a precursor to a range loop

// this is the actual range loop node with two successors

// producing ranging expression as nonnil

// no-op

// nonnil(idents, result 0)
func toExprSlice(idents []*ast.Ident) []ast.Expr { _ = "STUB: not implemented"; return nil }

func exprAsDeepProducer(rootNode *RootAssertionNode, expr ast.Expr) annotation.ProducingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(annotation.ProducingAnnotationTrigger)
}

// the expr is not deeply nilable

// CheckGuardOnFullTrigger gives guarding its intended semantics:
// if a full trigger would be created with a guarded producer but
// not a guarded consumer, then the production as written in the
// trigger is ignored and replaced with an always-nilable-producing
// instance of annotation.GuardMissing
func CheckGuardOnFullTrigger(trigger annotation.FullTrigger) annotation.FullTrigger {
	_ = "STUB: not implemented"
	return *new(annotation.FullTrigger)
}

// addAssignmentToConsumer updates the consumer with assignment entries for informative printing of errors
func addAssignmentToConsumer(lhs, rhs ast.Expr, pass *analysishelper.EnhancedPass, consumer annotation.ConsumingAnnotationTrigger) error {
	_ = "STUB: not implemented"
	return nil
}

/* isShortenExpr */

/* isShortenExpr */

func addReturnConsumers(rootNode *RootAssertionNode, node *ast.ReturnStmt, expr ast.Expr, retKey *annotation.RetAnnotationKey, isNamedReturn bool) {
	_ = "STUB: not implemented"
	// add shallow consumer
	return
}

// If expr is a deep type, then we track its deep nilability as well.
// ```
// E.g., func foo(s []*int) []*int {
//   s[0] = nil
//   return s  // <-- track shallow and deep nilability of `s` here
// }
// ```

// since this is an implicit tracking of the deep nilability of expr, we don't need to
// check for its guarding.

// We add a full trigger here directly because if we add only a deep consumer here, then it gets added
// to the same assertion node in the assertion tree as for the shallow consumer above. This is a problem
// since a producer actually meant for the shallow consumer also incorrectly matches the deep consumer.
