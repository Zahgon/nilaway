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

// Package assertiontree contains the node definitions for the assertion tree, as well as the main
// backpropagation algorithm.
package assertiontree

import (
	"context"
	"go/ast"

	"go.uber.org/nilaway/annotation"
	"go.uber.org/nilaway/util/analysishelper"
	"golang.org/x/tools/go/cfg"
)

// backpropAcrossBlock iterates over all nodes in the CFG block in _reverse_ order, writes logs,
// and delegates the handling of each node to backpropAcrossNode.
func backpropAcrossBlock(rootNode *RootAssertionNode, block *cfg.Block) error {
	_ = "STUB: not implemented"
	// Iterate over all blocks in _reverse_ order
	return nil
}

// If any error occurs when back-propagating a node, we wrap the error with more
// information such as file name and positions for easier debugging.

// backpropAcrossNode is the main driver function for the backpropagation of each node of
// different types. For some complicated cases, it further delegates the handling to other
// finer-grained backpropX functions for better code clarity.
func backpropAcrossNode(rootNode *RootAssertionNode, node ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// These nodes represent declarations such as `var x, y : int = 4, 3`

// The following cases are not interesting to our nilness analysis, or are currently
// unsupported, so we do nothing for them.

// TODO: figure out what source code generates these cases - it's not obvious
// TODO: handle defers

// backpropAcrossSend handles backpropagation for send statements. It is designed to be called from
// backpropAcrossNode as a special handler.
func backpropAcrossSend(rootNode *RootAssertionNode, node *ast.SendStmt) error {
	_ = "STUB: not implemented"
	// Note that for channel sends, we have:
	// (1) A send to a nil channel blocks forever;
	// (2) A send to a closed channel panics.
	// (1) falls out of scope for NilAway and hence we do not create a consumer here for the
	// channel variable. For (2), since we do not track the state of the channels, we currently
	// cannot support it.
	// TODO: rethink our strategy of handling channels (#192).
	return nil
}

// backpropAcrossReturn handles backpropagation for return statements. It is designed to be called
// from backpropAcrossNode as a special handler.
func backpropAcrossReturn(rootNode *RootAssertionNode, node *ast.ReturnStmt) error {
	_ = "STUB: not implemented"
	// we have to handle the case that a multiply-returning function is being returned, and split
	// the productions appropriate instead of just calling computeAndConsumeResults directly in that case
	return nil
}

// definitely not multiply returning here

// this is a call to a variable with function type -
// we could try to search these for their annotations (and should) but not yet

// In this case - an anonymous function is called and returned, for now I don't
// know what to do here, so we just compute
// TODO - handle this case (and similar case in ParseExprAsProducer)

// Since functions are assumed to be without side effects, we don't know that
// `fident` is actually definitely, non-nil here, tracked as

// this is the case we were looking for!
// we've identified that a multiply-returning function is being returned

// this nil check reflects programmer logic

// since we don't individually track the returns of a multiply returning function,
// we form full triggers for each return whose type doesn't bar nilness

// since the value is being returned directly, only its shallow nilability
// matters (but deep would matter if we were enforcing correct variance)

// if an error returning function returns directly as the result of
// another error returning function, then its results can safely be
// interpreted as guarded

// This is a duplicate trigger for tracking "always safe" paths. The analysis of these triggers
// will be processed at the inference stage.

// is a function call but not a multiply returning one - just compute!

// backpropAcrossAssignment handles backpropagation for assignments, including special cases such
// as type switches and range statements. Moreover, it also handles special contracts such as map
// reads, channel reads, and type assertions by adding appropriate guards to them.
// Specifically, there are three phases in our process here:
// Phase 1. move any assertions on trackable LHS expressions to the RHS;
// Phase 2. mark any assignments into fields as consuming the assigned value by that fields' annotation;
// Phase 3. mark all LHS and RHS as computed.
// nonnil(lhs, rhs)
func backpropAcrossAssignment(rootNode *RootAssertionNode, lhs, rhs []ast.Expr) error {
	_ = "STUB: not implemented"
	// For phase 1 and 2, we will first handle a few special assignments (with early return), then
	// if none of the special cases are hit, which means it is a normal assignment, we further
	// delegate the process to other functions depending on if it is many-to-one or one-to-one
	// assignment for better code clarity.
	// In all cases, we should do phase 3 before returning, so here we defer a function for phase 3.
	return nil
}

// Phase 3
// Now that we've back-propagated across the assignment itself, make sure we can compute
// all of the lhs and rhs.

// Phase 1 and 2
// First handle a few special cases, e.g., type switches, type assertions, range statements,
// and some cases for "ok" contracts, all of which will have a rhs with length 1.

// Here we first strip the parentheses of the rhs to reveal the underlying nodes.

// Type switch `x := y.(type)`, which needs special handling because TypesInfo.Defs
// can't find an object for the lhs.
// Note that the key distinction between a "type switch" and a "type assertion" in the
// AST is whether the `Type` field of the AST node is nil.

// lhs must have one element, which is *ast.Ident.

// Range statement of the form `for x := range y`, which is not overly complex to
// handle but does involve distinct semantics.

// This is the case of creating a pointer and assigning it to a variable, e.g., `x := &y`,
// where y is a non-pointer type (e.g., y := S{}).
// Here, the pointer is always nonnil, so we can just add a ProduceTriggerNever.

// Now we handle special cases for "ok" contracts, the lhs must have length of 2, the first
// being the processed variable and second being the `ok` boolean. Specifically, we
// currently handle the following cases in NilAway:
// 1. Map read: `v, ok := m[k]`
// 2. Channel receive: `v, ok := <-ch`
// TODO: 3. Type assertion: `v, ok := y.(*type)`

// Map read

// Channel read
// There is a slight difference in the handling of map reads and channel receives
// that is driven by Go's behavior. Reading from a nil map returns `nil`, but
// reading from a nil channel gives a deadlock error. Hence, for maps guarding is
// essentially required if the map is determined to be nilable, however, for a
// channel guarding logic is enforced only if it is in an `ok` form. That is why
// the NeedsGuard of ChanRecv is set to false in all other cases, but this one,
// where we know that it is an `ok` receive case.

// Add produce trigger for channel receive on the expression `v` here itself,
// since we want to set guarding = true.

// set the guard on channel receive since it is an ok form

// We do not need to "backpropAcrossOneToOneAssignment" since we explicitly
// added a produce trigger above.

// Type assertion

// TODO: properly handle type assertions' "OK" contract

// Generic function call
// TODO: currently we do not handle generic rich-check-effect function calls, so as a temporary solution,
//  we suppress reporting of errors. Note that this helps suppress false positives, but it also means that
//  we don't report true positives either. We should fix this in the future when we add support for generics.

// If the above code did not catch any special cases, it means this assignment is a normal
// assignment, and we further delegate the handling to other functions.

// backpropAcrossRange handles range expression (e.g., "for i, v := range lst"), it is designed to
// be called from backpropAcrossAssignment as a finer-grained handler for special assignment cases.
func backpropAcrossRange(rootNode *RootAssertionNode, lhs []ast.Expr, rhs ast.Expr) error {
	_ = "STUB: not implemented"
	// produceAsIndex(i) marks the ith lhs expression as a range index, producing it as non-nil
	// because it necessarily has basic type (int or char)
	return nil
}

// if nonempty, produce the index as definitely non-nil

// produceAsDeepRHS(i) marks the ith lhs expression as flowing deeply from the rhs

// if nonempty, produce the ranging value from the deep nilability of the rhs
// we can't track the rhs of ranges since we would need to discover non-nil assignments
// to an unbounded number of indices to conclude anything other than the annotation-based
// deep nilability of rhs

// we remove the guard on any deep types read from a range because reading
// them through a range guarantees they exist, removing the need for an ok check

// produceNonNil marks the ith lhs expression as nonnil due to limitations of NilAway.

// Go 1.23 introduced [range-over-func] language feature, where the `range` statement can
// now take the following types:
//
// 1. `func(func() bool)`
// 2. `func(func(K) bool)`
// 3. `func(func(K, V) bool)`
//
// We currently do not handle these types yet, so here we assume that they are deeply non-nil
// (by adding nonnil producers to both K and V if given).
//
// Note that the `iter` package provides `iter.Seq` and `iter.Seq2` generic types for 2 and 3
// specifically. Therefore, we need to `.Underlying()` on the rhsType to find the underlying
// func type for simplicity.
//
// [range-over-func]: https://tip.golang.org/doc/go1.23
// TODO: handle that (#287).

// This block breaks down the cases for the `range` statement being analyzed,
// starting by switching on how many left-hand operands there are

// If we have two left hand operands, the first is always int-valued
// This checks if we are ranging over a string
// If we are ranging over a string, then the second lhs operand is also non-nil

// If we are not ranging over a string, then we cannot assume basic type

// If we are ranging over a map slice or string with only a single lhs operand, then that
// operand will be int-valued.

// Iterating over a channel with only a single lhs operand will still result in deeply
// produced lhs values.

// Here the range is over basic types, such as integers (e.g., "for i := range 10").
// We do not need to do anything here, as the basic types are always presumed to be non-nil.

// We could be ranging over a generic slice (where rhsType is a *types.TypeParam) but
// we do not handle generics yet. Here we just assume generic slices are all deeply
// nonnil - we do not need to do anything here.
// TODO: handle that.

// backpropAcrossTypeSwitch handles type switches (e.g., "switch v := a.(*type)"), it is designed
// to be called from backpropAcrossAssignment as a finer-grained handler for special assignment
// cases. The main reason that this case has to be handled separately is that it introduces a
// "symbolic" variable to track the result of the type switch. This variable does not have a
// canonical instance of `types.Var` declaring it, in fact there is NO instance of `types.Var`
// associated with the declaration site as there usually would be through TypesInfo.Defs, and
// TypesInfo. Uses will give a fresh `types.Var` at every usage site. This is why we have to
// inspect the assertion tree for any variables that match the symbolic type switch variable
// without being able to compare the identity of `types.Var` instances as we usually do.
// nonnil(lhs, rhs)
func backpropAcrossTypeSwitch(rootNode *RootAssertionNode, lhs *ast.Ident, rhs ast.Expr) error {
	_ = "STUB: not implemented"
	// First, make a copy of the children array to iterate over, as we will mutate it.
	return nil
}

// For each variable in the assertion tree, check if it's equal to the symbolic variable
// being instantiated by this type switch, and, if so, assign to it.
// we (annonyingly) have to handle this as a separate case and continue after the first child
// is found because there is no canonical instance of types.Object for symbolic type switch
// variables

// even though we can't do the matching through a types.Object instance, we conclude
// that the *types.Var in the assertion nodes matches the lhs of this assignment

// this nil check reflects programmer logic

// rhs is trackable, so move assertions as we would in the vanilla assignment case

// assignment complete

// lhsVal expression will never be nil here because rhsVal will never be nil

// no current assertion matches the type switch lhs variable here, so it's a no-op

// backpropAcrossOneToOneAssignment handles normal one-to-one assignment (e.g, "var a *int = b", or
// "var a, b, c *int = d, e, f"), it is designed to be called from backpropAcrossAssignment as a
// finer-grained handler for one-to-one normal assignments.
// nonnil(lhs, rhs)
func backpropAcrossOneToOneAssignment(rootNode *RootAssertionNode, lhs, rhs []ast.Expr) error {
	_ = "STUB: not implemented"

	// precompute the parses of all LHS expressions - we'll need them
	return nil
}

// Phase 1

// We now have n expressions on the RHS and n on the LHS
// For each of these n pairs, we have 3 cases:
// A) LHS is not trackable - nothing to be done in Phase 1
// B) LHS is trackable but RHS is not - mark LHS as produced by RHS value
// C) LHS and RHS are both trackable - move assertions from LHS to RHS
//
// Notably, the two phases of 3 - remove assertions from LHS with LiftFromPath and
// add assertions to RHS with LandAtPat - must be done in parallel so that swaps behave
// correctly, e.g.: x, y = y, x (see the test multipleassignment.go)

// Before we can even start though, we have to deal with another problem: shadowing.
// If x.f is non-nil, then after either of the following assignments y.f will be non-nil:
//
// y.f, y = nil, x
// y, y.f = x, nil
//
// This is because the assignment y.f = nil is always to the *old* y, whereas the other
// assignment y = x makes sure that subsequently we touch the *new* y which is x
//
// To handle this, we have to do a preliminary pass to figure out which expressions will
// be assigned to over the course of this entire multiple assignment, and not propagate
// assertions for any member assignments that will be shadowed. The two ways we conclude
// that a member assignment M to E will be shadowed are:
// i) there exists another member assignment to a strict prefix expression of E
// ii) there exists another member assignment to E to the right of M
// We test for both cases, putting the results in the array `shadowMask` below -
// which is true at index i iff member assignment i is shadowed.

// TODO: compute this more efficiently using a tree

// This struct is declared to assist with deferring the second phase of the assignments

// To process all case C's in parallel, we accumulate a list of the first phases,
// the "lifts", and where they will land once all accumulated

// Split cases A, B, C from above

// If lpath == nil we're in case A so we do nothing

// Both lhsVal and rhsVal are trackable! we're in case C

// If rhs is a function call that is tracked then we just add field producers before detaching
// the assertion nodes

// Length of rproducers must be 1 since assignment is one-one

// TODO: below check for `lhsNode != nil` should not be needed when NilAway supports Ok form for
//  used-defined functions (tracked issue #77)

// Add assignment entries to the consumers of lhsNode for informative printing of errors

// If the lhsVal path is not only trackable but tracked, we add it as
// a deferred landing

// We're in case B

// lhsVal expression will never be nil here because rhsVal will never be nil

// beforeTriggersLastIndex is used to find the newly added triggers on the next line

// Update consumers of newly added triggers with assignment entries for informative printing of errors
// TODO: the below check `len(rootNode.triggers) == 0` should not be needed, however, it is added to
//  satisfy NilAway's analysis

// Now we actually land each of the nodes lifted above, this guarantees parallelism

// Phase 2

// Check whether a consumption trigger needs to be added for a field assignment here

// backpropAcrossManyToOneAssignment handles normal many-to-one assignment (e.g, "a, b := foo()"),
// it is designed to be called from backpropAcrossAssignment as a finer-grained handler for
// many-to-one normal assignments.
// nonnil(lhs, rhs)
func backpropAcrossManyToOneAssignment(rootNode *RootAssertionNode, lhs, rhs []ast.Expr) error {
	_ = "STUB: not implemented"
	// Single rhsVal value assigned to multiple lhs values - the only option here is that it is
	// a function return.
	return nil
}

// Eliminates checking of the `_` instances in the lhs of a multiple assignment

// Phase 1

// beforeTriggersLastIndex is used to find the newly added triggers on the next line

// Update consumers of newly added triggers with assignment entries for informative printing of errors

// Phase 2

// Update consumeTrigger with assignment entries for informative printing of errors

// lhsVal is a field read, so this is a field assignment
// since multiple return functions aren't trackable, this is a completed trigger
// as long as the type of the expression being assigned doesn't bar nilness

// We are assigning directly into the field, so we only care about shallow,
// but we would have to check deep if we were checking dep nilability variance

// Update consumeTrigger with assignment entries for informative printing of errors

// computePostOrder computes the postorder of depth first search tree (DFST) of the live blocks of the CFG.
// The backpropagation algorithm converges faster if the CFG blocks are traversed in postorder, compared to the random
// order (Check [Kam, Ullman 76']).
func computePostOrder(blocks []*cfg.Block) []int { _ = "STUB: not implemented"; return nil }

// Start traversal from entry block (index 0): https://pkg.go.dev/golang.org/x/tools/go/cfg#CFG

// BackpropAcrossFunc is the main driver of the backpropagation, it takes a function declaration
// with accompanying CFG, and back-propagates a tree of assertions across it to generate, at entry
// to the function, the set of assertions that must hold to avoid possible nil flow errors.
func BackpropAcrossFunc(
	ctx context.Context,
	pass *analysishelper.EnhancedPass,
	decl *ast.FuncDecl,
	functionContext FunctionContext,
	graph *cfg.CFG,
) ([]annotation.FullTrigger, int, int, error) {
	_ = "STUB: not implemented"
	// We transform the CFG to have it reflect the implicit control flow that happens
	// inside short-circuiting boolean expressions.
	return nil, 0, 0, nil
}

// Generate rick check effects.

// The assertion nodes for each block and an array of bools to indicate whether each block is
// updated in this round or not.
// DANGER: anytime a pointer is copied from currAssertions to nextAssertions, it MUST be
// treated as immutable - if any modification (e.g. merging or backprop) is going to occur you
// must perform a deep copy with CopyNode

// The assertion nodes for the entry block, we will use it as an indication for stabilization.
// We consider the backpropagation stable if # of stable rounds > # of live blocks + tolerance.

// Initialize the process by creating the assertion nodes for the return block.

// There is no need to process non-live blocks; their assertions will simply be nil.

// No need to re-process the assertion node for the current block if it does not have
// successors, or they were not updated in current or last round.

// Normally we should copy the assertion node, but here it is ok to simply pass it
// to the next round since we are not modifying it.

// Before doing actual back propagation for the current block, we need to first prepare
// the assertion node for its successors: (1) deep copy for modifications, and (2)
// apply preprocessor (insert nil checks) corresponding to the branch condition
// (if it is a branch block). In the meantime, we filter out any successors if there
// are no assertion nodes associated with it.

// No need to preprocess if there is no assertion node for the successor.

// If the successor was updated this round
// deep copy the node for modifications.

// If the successor was updated last round
// deep copy the node for modifications.

// Apply preprocessor (for branches) if there is any.

// No assertion nodes attached with any successors, this should never happen since we
// will only reach here if any of the successors were updated in the current or last round.

// Merge the branch successors if they are both available.

// Now, the final processed node is in succs[0], we can back-propagate across it.

// Monotonize updates updatedThisRound to reflect whether the assertions changed at a given index.

// ProcessEntry is expensive, and ideally we would only call it once after the fixed point
// of the backprop has terminated. We call it every time because waiting for all of the assertion
// trees (i.e. the assertion tree for each block) to stabilize sometimes never occurs. As an alternative,
// we wait for only the entry to stabilize because that's the one we care about anyways, provided its been
// stable for at least as many rounds as there are blocks because that means the information
// from each block has had a chance to reach entry. But waiting for the entire tree at entry block
// to stabilize because sometimes consume triggers generated in a loop that generates them
// on slightly different trackable expressions every time are paired with produce triggers
// similarly generated in a loop, (see infiniteAssertions test in loopflow.go), so the assertion
// tree will continue to grow unboundedly even though this growth produced no new full triggers
// i.e. no new errors. To get around this, we generate the full triggers every round and
// track stabilization of those not stabilization of the root node.
// TODO: implement that

// Move variables from this round to last round and create new ones for next round.
// For best performance, we reuse the slices by simply swapping them and clearing the
// slices for next rounds.

// Return the generated full triggers at the entry block; we're done!
