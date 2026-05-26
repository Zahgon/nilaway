//  Copyright (c) 2024 Uber Technologies, Inc.
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

package preprocess

import (
	"go/ast"

	"golang.org/x/tools/go/cfg"
)

// CFG performs several passes on the CFG of utility to our analysis and returns a shallow copy of
// the modified CFG, with the original CFG untouched.
//
// Specifically, it performs the following modifications to the CFG:
//
// Canonicalize conditionals:
// - replace `if !cond {T} {F}` with `if cond {F} {T}` (swap successors)
// - replace `if cond1 && cond2 {T} {F}` with `if cond1 {if cond2 {T} else {F}}{F}` (nesting)
// - replace `if cond1 || cond2 {T} {F}` with `if cond1 {T} else {if cond2 {T} else {F}}` (nesting)
//
// Canonicalize nil comparisons:
// It also performs the following useful transformation:
// - replace `if x != nil {T} {F}` with `if x == nil {F} {T}` (swap successors)
// - replace `nil == x {T} {F}` with `if x == nil {T} {F}` (swap comparison order)
//
// Canonicalize explicit boolean comparisons:
// - replace `if x == true {T} {F}` with `if x {T} {F}`
// - replace `if x == false {T} {F}` with `if !x {T} {F}`
func (p *Preprocessor) CFG(graph *cfg.CFG, funcDecl *ast.FuncDecl) *cfg.CFG {
	_ = "STUB: not implemented"
	// The ASTs and CFGs are shared across all analyzers in the nogo framework, so we should never
	// modify them directly. Here, we make a copy of the graph (and all blocks in it) and modify
	// the copied graph instead.
	return nil
}

// Important: add all new blocks to the end, don't try to "move around" any existing blocks
// because they're all referenced by index!

// Create a failure block at the end of the blocks list to be used for trusted functions.

// Perform a series of CFG transformations here (for hooks and canonicalization). The order of
// these transformations matters due to canonicalization. Some transformations may expect the
// CFG to be in canonical form, and some transformations may change the CFG structure in a way
// that it needs to be re-canonicalized.

// split blocks do not require the CFG to be in canonical form, and it may modify the CFG
// structure in a way that it needs to be re-canonicalized. Here, we cleverly bundles the two
// operations together such that we only need to run canonicalization once.

// Replacing conditionals in the CFG requires the CFG to be in canonical form (such that it
// does not have to handle "trustedFunc() && trustedFunc()"), and it will canonicalize the
// modified block by itself.

// Next, we need to re-insert information that is lost during CFG build for *ast.RangeStmt
// and *ast.SwitchStmt by iterating through all blocks. This requires knowing the links between
// the nodes contained within a block to their parents (*ast.RangeStmt or *ast.SwitchStmt nodes).
// So, here establish the link and then do the work.

// Please check the docstring of the following call to see why this is needed.
// TODO: remove this once anonymous function support handles it naturally.

// copyGraph makes a semi-deep copy of the CFG and returns the copied graph. Note that only the
// graph itself is copied, i.e., the blocks and their edges (via block.Succs). The referenced AST
// nodes are _not_ copied (meaning we still should not modify the underlying AST nodes), but the
// slice storing the AST nodes (i.e., cfg.Block.Nodes) in each block is shallow-copied for modifications.
func copyGraph(graph *cfg.CFG) *cfg.CFG {
	_ = "STUB: not implemented"
	// For some large graphs, a recursion-based approach will exceed the runtime stack size limit
	// and a stack-based approach will have many allocations / de-allocations. For best performance
	// (both in time and space), we run two iterations, one simply copying the blocks without
	// copying the edges (Succs), and another that copies the edges.
	return nil
}

// Keep track of the mapping between original block ptr -> copied block ptr.

// Shallow copy the slice that stores the AST nodes.

// Copy the block and put it in the new graph.

// Store the mapping.

// Now, we iterate through the blocks again to fill in the edges (block.Succs). All blocks
// must have already been copied.

func (p *Preprocessor) restructureOnNoReturnCall(block *cfg.Block) {
	_ = "STUB: not implemented"
	return
}

// The rest of the nodes are now unreachable.
// There will be no successor block.

// splitBlockOnTrustedFuncs splits the CFG block into two parts upon seeing a trusted function
// from the hook framework (e.g., "require.Nil(t, arg)" to "if arg == nil { <all code after> }".
// This does not expect the CFG to be in canonical form, and it may change the CFG structure in a
// way that it needs to be re-canonicalized.
func (p *Preprocessor) splitBlockOnTrustedFuncs(graph *cfg.CFG, thisBlock, failureBlock *cfg.Block) {
	_ = "STUB: not implemented"
	return
}

// replaceConditional calls the hook functions and replaces the conditional expressions in the CFG
// with the returned equivalent expression for analysis.
//
// This function expects the CFG to be in canonical form to fully function (otherwise it may miss
// cases like "trustedFunc() && trustedFunc()").
//
// It also calls canonicalizeConditional to canonicalize the transformed block such that the CFG
// is still canonical.
func (p *Preprocessor) replaceConditional(graph *cfg.CFG, block *cfg.Block) {
	_ = "STUB: not implemented"
	// We only replace conditionals on branching blocks.
	return
}

// Last node is a call expression for `if foo() { ... }` case.

// Otherwise, we check if it is `if ok := foo(); ok { ... }` case.
// Note that this would fail for the following case:
//
// ok := foo()
// if dummy {
//   if ok {
//     ...
//   }
// }
//
// (The example above is canonicalized -- `if dummy && ok {...}` is equivalent, and is
// probably more common in practice).
//
// Here we will not find the declaration of `ok` in the block. Ideally we should really find
// the declaration node of `ok` instead of simply checking the last node in the block (possibly
// with the help of SSA).
// TODO: implement that.

// The returned expression may be a binary expression, so we need to canonicalize the CFG again
// after such replacement.

// canonicalizeConditional canonicalizes the conditional CFG structures to make it easier to reason
// about control flows later. For example, it rewrites
// `if !cond {T} {F}` to `if cond {F} {T}` (swap successors), and rewrites
// `if cond1 && cond2 {T} {F}` to `if cond1 {if cond2 {T} else {F}}{F}` (nesting).
func (p *Preprocessor) canonicalizeConditional(graph *cfg.CFG, thisBlock *cfg.Block) {
	_ = "STUB: not implemented"
	// We only restructure non-empty branching blocks.
	return
}

// type *cfg.Block
// type *cfg.Block

// A few helper functions to make the code more readable.
// The conditional expr is the last node in the block.

// if a parenexpr, strip and restart - this is done with recursion to account for ((((x)))) case

// recur within parens

// swap successors - i.e. swap true and false branches

// recur within NOT

// Logical AND and Logical OR actually require the exact same short circuiting behavior
// except for whether the true or false branch leads to the short circuiting. This split
// is captured by the following switch, and, as can be observed, all other logic is the
// same

// Standardize binary expressions to be of the form `expr OP literal` by swapping `x` and `y`, if `x` is a literal.
// For example, standardizes `nil == v` to the `v == nil` form

// Swap X and Y

// The NEQ and EQL cases here rewrite the ASTs to ensure all _nil comparisons_ are
// standardized to the form of `x == nil` (i.e., "variable" == "literal nil"). For example:
// (1) x != nil -> x == nil, and swapping the true and false branches in the CFG.

// Similarly, we also rewrite the ASTs for explicit _boolean comparisons_. For example,
// (1) `ok == true` and `ok != false` -> `ok`
// (2) `ok == false` and `ok != true` -> `!ok`, and swapping the true and false branches in the CFG.

// Note that we _should not_ directly modify the AST nodes, since they are shared across
// other nogo analyzers. Instead, whenever a rewrite is needed we create a new AST node
// and replace the original node pointer with the clone in the block.Nodes slice instead.

// Rewrite when operand `y` is a literal `nil`.

// Copy the AST Node first.

// As discussed, we change the operator to EQL here.

// Replace the condition, and swap the branches since we modified a NEQ conditional
// to a EQL one.

// For explicit boolean NEQ checks, we replace the AST nodes for `ok != true` and `ok != false`
// (also, `true != ok` and `false != ok`) with `ok` and `!ok` form for the true and false cases, respectively.

// replaces `ok != false` with `ok`

// replaces `ok != true` with `!ok`
// recur to swap true and false branches for the unary expr `!ok`

// For explicit boolean EQL checks, we replace the AST nodes for `ok == true` and `ok == false`
// (also, `true == ok` and `false == ok`) with `ok` and `!ok` form for the true and false cases, respectively.

// replaces `ok == true` with `ok`

// replaces `ok == false` with `!ok`
// recur to swap true and false branches for the unary expr `!ok`

// collectChildren establishes the links between the range / switch statement nodes and their child
// nodes. This is specifically designed for our preprocess function: when we rewrite the CFG to
// re-insert the lost information, we need to know if a block in CFG belongs to a certain range
// statement or switch statement AST node for retrieving lost information.
func collectChildren(funcDecl *ast.FuncDecl) (map[ast.Node]*ast.RangeStmt, map[ast.Node]*ast.SwitchStmt) {
	_ = "STUB: not implemented"
	return nil, nil
}

// markRangeStatements rewrites a cfg to reflect ranging loops - the assignments in a `for... range y {}`
// loop are by default erased in the CFG pass, so we match on the structure of all blocks in the CFG
// and their AST nodes to rediscover and reinsert these assignments or, in the case of a `for range`
// loop with no assignments - we insert a fresh *ast.UnaryExpr simply indicating that this is a range
func markRangeStatements(graph *cfg.CFG, rangeChildren map[ast.Node]*ast.RangeStmt) {
	_ = "STUB: not implemented"
	return
}

// we have a `range` statement! now time to figure out which one

// we have a `for range expr {}` loop

// we have a `for x := range expr {}` loop

// we have a `for x, y := range expr {}` loop

// markSwitchStatements restructures a cfg to reflect switch statements.
//
// In particular, `switch x { case y0 : e0 case y1 : e1 ... }` will be parsed by the CFG into:
//
// Block0: Nodes: x, y0, Succs: Block1, Block2
// Block1: e0
// Block2: Nodes: y1, Succs: Block3, Block4
// Block3: e1
// Block4: Nodes: y2, Succs: Block5, Block6
//
// Which we transform into:
//
// Block0: Nodes: x == y0, Succs: Block1, Block2
// Block1: e0
// Block2: Nodes: x == y1, Succs: Block3, Block4
// Block3: e1
// Block4: Nodes: x == y2, Succs: Block5, Block6
//
// This will allow the existing logic for reading conditionals from the CFG to handle `switch` statements,
// which simply checks that the block ends with a Binary check like `x == y` and has two successors.
//
// invariant - consecutive cases of a switch statement have block numbers whose ordering
// reflects the syntactic ordering of the cases - if a case were to have a lower block number
// than its initial switch statement this would be broken
func markSwitchStatements(graph *cfg.CFG, switchChildren map[ast.Node]*ast.SwitchStmt) {
	_ = "STUB: not implemented"
	return
}

// we've found a switch statement!

// use the position of the case expression

// use the position of the case expression
