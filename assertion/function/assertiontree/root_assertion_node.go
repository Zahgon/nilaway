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
	"go/token"
	"go/types"

	"go.uber.org/nilaway/annotation"
	"go.uber.org/nilaway/guard"
	"go.uber.org/nilaway/util/analysishelper"
)

// RootAssertionNode is the object that will be directly handled by the propagation algorithm,
// their only children should be VarAssertionNodes and FuncAssertionNodes
//
// the triggers field keeps track of productions and consumptions that have been directly matched
// its consumeTriggers field should be kept empty
//
// //nilable(funcObj)
type RootAssertionNode struct {
	assertionNodeCommon
	triggers []annotation.FullTrigger

	// funcObj does not have to be set. When set, it indicates the object corresponding to this function
	funcObj *types.Func

	// exprNonceMap maps expressions to nonces created to track their contracts
	exprNonceMap guard.ExprNonceMap

	// functionContext holds the context of the function during backpropagation. The state includes
	// map objects that are created at initialization, and configurations that are passed through function analyzer.
	functionContext FunctionContext
}

// LocationOf returns the location of the given expression.
func (r *RootAssertionNode) LocationOf(expr ast.Expr) token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

// HasContract returns if the given function has any contracts.
func (r *RootAssertionNode) HasContract(funcObj *types.Func) bool {
	_ = "STUB: not implemented"
	return false
}

// MinimalString for a RootAssertionNode returns a minimal string representation of that root node
func (r *RootAssertionNode) MinimalString() string { _ = "STUB: not implemented"; return "" }

// AddNewTriggers adds the given new triggers to the existing set of triggers of this node
func (r *RootAssertionNode) AddNewTriggers(newTrigger ...annotation.FullTrigger) {
	_ = "STUB: not implemented"
	return
}

// FuncDecl returns the underlying function declaration of this node
func (r *RootAssertionNode) FuncDecl() *ast.FuncDecl { _ = "STUB: not implemented"; return nil }

// Pass the overarching analysis pass
func (r *RootAssertionNode) Pass() *analysishelper.EnhancedPass {
	_ = "STUB: not implemented"
	return nil
}

// FuncNameIdent returns the function name identifier node
func (r *RootAssertionNode) FuncNameIdent() *ast.Ident { _ = "STUB: not implemented"; return nil }

// DefaultTrigger is not well defined for root nodes
func (r *RootAssertionNode) DefaultTrigger() annotation.ProducingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(annotation.ProducingAnnotationTrigger)
}

// BuildExpr is not well defined for root nodes
func (r *RootAssertionNode) BuildExpr(_ ast.Expr) ast.Expr {
	_ = "STUB: not implemented"
	return *new(ast.Expr)
}

// Root for a RootAssertionNode is the identity function
func (r *RootAssertionNode) Root() *RootAssertionNode {
	_ = "STUB: not implemented"

	// Size for a RootAssertionNode also includes the full triggers
	return nil
}

func (r *RootAssertionNode) Size() int { _ = "STUB: not implemented"; return 0 }

// FuncObj returns the underlying function declaration of this node as a types.Func
func (r *RootAssertionNode) FuncObj() *types.Func { _ = "STUB: not implemented"; return nil }

// GetNonce returns the nonce associated with the passed expression, if one exists. the boolean
// return indicates whether a nonce was found
func (r *RootAssertionNode) GetNonce(expr ast.Expr) (guard.Nonce, bool) {
	_ = "STUB: not implemented"
	return *new(guard.Nonce), false
}

// GetTriggers returns the full triggers accumulated at this root node
func (r *RootAssertionNode) GetTriggers() []annotation.FullTrigger {
	_ = "STUB: not implemented"

	// GetDeclaringIdent finds the identifier that serves as the declaration of the passed object
	return nil
}

func (r *RootAssertionNode) GetDeclaringIdent(obj types.Object) *ast.Ident {
	_ = "STUB: not implemented"
	return nil
}

// In case the declaration is package.ident

// create a fake object just to allow lookups

// ObjectOf is the same as [types.Info.ObjectOf], but if an identifier cannot be looked up (e.g.,
// it is an artificial identifier we created to aid the analysis), we look up the internal backup
// map instead. ObjectOf returns nil if and only if both attempts fail.
func (r *RootAssertionNode) ObjectOf(ident *ast.Ident) types.Object {
	_ = "STUB: not implemented"
	return *new(types.Object)
}

// check if ident points to an anonymous function literal

// funcArgsFromCallExpr returns the set of arguments that are passed to the method at the call site. If the method
// is an anonymous function, it expands the argument set with the closure variables collected for that function
func (r *RootAssertionNode) funcArgsFromCallExpr(expr *ast.CallExpr) []ast.Expr {
	_ = "STUB: not implemented"
	return nil
}

// if the declaration of the ident points to a function literal node,
// then update fun with the function literal node

// Equal returns true iff a is the same path as b
// nilable(a, b)
func (r *RootAssertionNode) Equal(a, b TrackableExpr) bool { _ = "STUB: not implemented"; return false }

// IsPrefix returns true iff a is a prefix of b
func (r *RootAssertionNode) IsPrefix(a, b TrackableExpr) bool {
	_ = "STUB: not implemented"
	return false
}

// IsStrictPrefix returns true iff a is a prefix of b and a does not equal b
func (r *RootAssertionNode) IsStrictPrefix(a, b TrackableExpr) bool {
	_ = "STUB: not implemented"
	return false
}

func newRootAssertionNode(exprNonceMap guard.ExprNonceMap, functionContext FunctionContext) *RootAssertionNode {
	_ = "STUB: not implemented"
	return nil
}

// using information from self (pass and funcDecl only) - turn a path into a new assertion tree starting
// at a new root. Except for that new root, all nodes are preserved so they can still be accessed as before
// the call. The new root is returned
func (r *RootAssertionNode) linkPath(path TrackableExpr) *RootAssertionNode {
	_ = "STUB: not implemented"
	return nil
}

// use this currNode to build a linear tree to merge into r

// AddConsumption takes the knowledge that consumer.expr will be consumed at a site characterized by the trigger
// consumer.annotation, and incorporate it into the assertion tree self
func (r *RootAssertionNode) AddConsumption(consumer *annotation.ConsumeTrigger) {
	_ = "STUB: not implemented"

	// we check if the type of the expression `expr` prevents it from ever being nil in the first place
	return
}

// expr cannot be nil, so do nothing

// expr is not trackable

// Here we can infer that the expression is non-nil by definition. Instead of ignoring creation of a trigger,
// particularly for always safe tracking, we create a trigger with ProduceTriggerNever.

// expr can be nil - complete the trigger and add to root

// we are consuming the expression directly - so only its shallow nilability counts

// we're adding a fresh node to the assertion tree to represent this consumption!

// merge it in - increasing the set of consumeTrigges as far as the path already exists
// in the assertion tree and extending the tree beyond that
// TODO - possibly avoid merging in a whole new path
// ^^^^ But I suspect gains would be marginal or non-existant - same logic either way

// This function takes an expression, represented as a path of AssertionNodes returned from ParseExprAsProducer,
// and searches for it in the assertion tree self
//
// if nodePtr != nil, it was found, and whichChild indicates which child it is of its parent.
// if nodePtr == nil, the expression was not trackable, or it was trackable but not present
//
// nilable(path, nodePtr)
func (r *RootAssertionNode) lookupPath(path TrackableExpr) (nodePtr AssertionNode, whichChild int) {
	_ = "STUB: not implemented"

	// expr is not trackable - return nil
	return *new(AssertionNode), 0
}

// lookup that path in r
// this tracks our lookup
// this tracks which child number we took to reach that lookup - useful for removing

// path does not exist in r, so even though the expr is trackable no assertions we're tracking care about it

// AddProduction takes the knowledge that producer.expr will have a value produced by the trigger producer.annotation,
// and incorporates it into the assertion tree rootNode
func (r *RootAssertionNode) AddProduction(producer *annotation.ProduceTrigger, deeperProducer ...*annotation.ProduceTrigger) {
	_ = "STUB: not implemented"
	return
}

// we don't care if this expression has a value produced because it's not tracked

// If we've reached here, that means currNode points to a subtree of r matching producer.expr
// since a value has now been produced for producer.expr, we can remove it from the assertion tree.
// Note that it's safe to remove the entire subtree under current node, since productions to paths accessible
// from the current expression and happening before the current line will have no effect on those paths'
// values going forward.
// e.g. in `x.f = nonNilVal(); x = foo(); x.f.g()` the production at `x = foo()` also has the effect of
// invalidating the previous assignment to `x.f`.

// triggerProductions takes a node (assumed to be attached to its parent) and matches any of its
// consumeTriggers with the given produceTrigger, as well as matching any more deeply found consumeTriggers
// with the default non-tracked produceTriggers of their consuming expressions. Direct children of the
// node being produced also have the option to be matches with a single optionally passed `deeperProducer`,
// used for assignments by values with known deep nilness properties.
func (r *RootAssertionNode) triggerProductions(node AssertionNode, producer *annotation.ProduceTrigger, deeperProducer ...*annotation.ProduceTrigger) {
	_ = "STUB: not implemented"

	// first we check if we were passed a deeper producer. If so, we use it to produce any \
	// indexAssertionNode children of the currNode
	return
}

// TODO: consider allowing multiple levels of deeper producers to be passed -
//   but very incompatible with current annotations approach so not yet

// for any consumeTriggers as the indexed expr, directly match them with this produceTrigger

// 	now we search for any deeper consumeTriggers in our indexed subtree, trying to match them
// with default producers as we go. These default producers are constructed using the methods
// BuildExpr and DefaultTrigger of assertion nodes. The former allows us to build up an expression
// to use to symbolize the production, and the latter allows us to point out the particular
// annotation that will yield the production of the found consumeTrigger

// GuardMatchBehavior as a type represents the set of possible effects of obtaining a guard match.
type GuardMatchBehavior = int

const (
	// ContinueTracking is a GuardMatchBehavior indicating that the field
	// GuardMatched should be set to true and the ConsumeTrigger that was matched should
	// otherwise be left in the assertion tree to flow through the function
	ContinueTracking GuardMatchBehavior = iota

	// ProduceAsNonnil is a GuardMatchBehavior indicating that the ConsumeTrigger
	// that was matched should be treated as nonnil-produced at this point, using the
	// trigger OkReadReflCheck
	ProduceAsNonnil
)

// AddGuardMatch takes an expression, and sees if that expression is mapped to a nonce
// indicating a RichCheckEffect that has been propagated from the concrete site of a check to the
// earlier site whose nilability semantics depend on that check.
// If it is mapped to a nonce, it sees if that expression is also present in the assertion
// tree with a consume trigger guarded by that nonce. This indicates that the flow we were
// looking for - for example, from `v` in `v, ok := m[k]` to `if ok {needsNonnil(v)}` - exists.
// The function takes a `GuardMatchBehavior` indicating what to do if the guard is found,
// for now, either continue tracking its expression or produce it as nonnil.
//
// To elaborate further, here is a complete rundown of the guarding mechanism.
//
// During preprocessing (preprocess_blocks.go) some statements are identified as producing a `RichCheckEffect`
// - a contract indicating that certain conditionals later in the program should have an effect on the
// semantics of that earlier statement. As an example, if `v, ok := m[k]` is encountered, then regardless
// of the deep nilability of `m`, `v` will be nilable. However, if `ok` is checked later in the program,
// it will be exactly as nilable as `m` is deeply nilable. This non-local reliance is propagated in the
// form of a RichCheckEffect that takes a GuardNonce uniquely generated corresponding to the AST node `v`
// at that site, and indicates that any time the expression `ok` is checked to be true, `v` should have
// that GuardNonce added to the set `Guards` of all of its `ConsumeTrigger`s. This indicates that those
// consumptions occur in a context "guarded" by that check. These `Guards` sets are intersected at control
// flow points (see `MergeConsumeTriggerSlices`), to ensure that the presence of a guard on a consumer
// really does indicate that it only occurs in a context in which the appropriate check has been made.
//
// This intersecting guard propagation then ensures that by the time any `ConsumeTrigger`s reach the
// statement that was dependent on the associated nonce, they will contain the information of whether
// they are properly guarded by that nonce. For example, in the below code snippet, line 1 will
// associate a nonce with `v`, to be applied when `ok` is checked. That contract will be propagated
// to the check on line 3 by a RichCheckEffect, so when backpropagation occurs across that positive branch
// of the check, it will see that `v` has two ConsumeTriggers, one generated by line 4 and one generated
// by line 7, and apply the nonce guard to both. However, on unifying the two branches, it will see that
// the ConsumeTrigger generated on line 7 is present on both sides, so it will intersect the Guards sets
// on each side and erase the nonce. Two ConsumeTriggers will then reach line 1, one from line 4 and
// one from line 7, but only the one from line 4 will have the appropriate nonce in its Guards set.
//
// ```
//
//	1 v, ok := m[k]
//	2
//	3 if ok {
//	4    consume(v)
//	5 }
//	6
//	7 consume(v)
//
// ```
//
// The role of this function, AddGuardMatch, is to look at an expression, take all the ConsumeTriggers
// for that expression in the current assertion tree, and set GuardMatched to true for them if they
// have the appropriate nonce in their Guards set. In the above example, this function would be called
// when backpropagating across line 1 with `v` for expr. The appropriate nonce would be found, and this
// function would see that it is present for 4's ConsumeTrigger but not 7's. Thus 4's would get
// GuardMatched set to true and 7's would not. If both of these ConsumeTriggers flowed to the beginning
// of the program, then they would get matched with a default ProduceTrigger as a deep read of `m`, which
// `checkGuardOnFullTrigger` would invalidate unless paired with a ConsumeTrigger with GuardMatched = true.
// GuardMatched for a ConsumeTrigger takes a conjunction over all paths from that production site to
// that ConsumeTrigger, so it is true iff the trigger has had every guard in its Guards set required
// every time it has passed through a contract-generating statement on any path.
//
// This description characterizes the `ContinueTracking` behavior. A simpler alternative, `ProduceAsNonnil`,
// indicates that if the appropriate nonce is found in a ConsumeTrigger's Guards set, the ConsumeTrigger
// should be matched immediately with a ProduceTrigger indicating nonnil production. This behavior is
// appropriate, for example, for the map itself in a read `v, ok := m[k]` - where consumptions of `m`
// guarded by a check `ok == true` are guaranteed to be produced as nonnil
func (r *RootAssertionNode) AddGuardMatch(expr ast.Expr, behavior GuardMatchBehavior) {
	_ = "STUB: not implemented"
	return
}

// we don't care if this expression could become guarded because it's not tracked

func (r *RootAssertionNode) consumeIndexExpr(expr ast.Expr) { _ = "STUB: not implemented"; return }

// AddComputation takes the knowledge that the expression expr has to be computed to generate any necessary assertions to
// ensure that the access is safe. This will take the form of nested calls to AddConsumption
//
// basic semantics: any ast node with an ast.Expr field recurs into that field
func (r *RootAssertionNode) AddComputation(expr ast.Expr) { _ = "STUB: not implemented"; return }

// We seek to recur through the AST to look for any sites at which an expression
// must be non-nil we ignore any expressions that provide types not values since
// assignments and branching can't happen within expressions in Go, the order in
// which we recur doesn't matter

// Process the binary expression `X op Y` in reverse, i.e., add consumers for Y first and then X

// Consider the example of the binary expression in: `x != nil && x.f != nil && x.f.g == 1`.
// If the binary expr is a short-circuiting `&&`, recursively iterate through every sub-expression in a
// right-to-left manner to check if any of the previous expressions contain appropriate negative nil checks that
// can mark the subsequent dereference of that expression as safe. For example, `x.f != nil` can mark the field
// access `x.f.g` as safe. Similarly, `x != nil` makes `x.f` safe. The expressions are marked safe by adding a
// producer right away to match with a consumer for that expression.
//
// An AST binary expression has two parts: X and Y. We recursively iterate through Y first and then X to achieve
// the right-to-left processing described above. We use `AddNilCheck()` to check if the expression is an
// atomic nil check or len check, i.e., not compounded with other expressions, and get the function pointer for
// the appropriate action to be taken. For `x != nil`, `AddNilCheck()` returns a function pointer for adding a
// nil check producer for the true branch, while a noop for the false branch, and vice versa for `x == nil`.
// In this case, with a `&&` short-circuiting operator, we only need to care about the true branch since the Y
// expression won't be executed if the X expression is false.
//
// Considering the above example,
// round 1: expr.Y: x.f.g == 1, and expr.X: x != nil && x.f != nil =>  AddNilCheck() returns noop since Y is not a nil check and X is non-atomic.
// round 2: expr.Y: x.f != nil, and expr.X: x != nil => AddNilCheck() returns successfully for both X and Y, where Y marks x.f.g as safe and X marks x.f as safe
//
// A similar approach is followed for the `||` operator, where we only need to care about the false branch since the
// Y expression won't be executed if the X expression is true.

// this returns a function that adds a consume trigger for the i-th argument to
// an annotated function call. One case we handle specially is that of a
// multiply-returning function passed directly to a multiple param function, say
// `foo(bar())`. In that case, we eagerly generate full triggers matching the
// producer for the i-th result of `bar()` to a consumer for the i-th parameter
// of `foo()`. In that case, since adding the consumer is already handled by the
// call to `consumeArgTrigger` itself, the returned `func(i, expr)` becomes a
// no-op. In all other cases, the function returned by `consumeArgTrigger` will
// add a consumption on the annotation of the i-th parameter of `fdecl` and the
// expression `expr` to the root node.

// Check if it is a rich check effect function call. If yes, this needs special handling.

// TODO: this is a temporary suppression which will be removed once we add support for
//  handling such cases precisely.

// is a pass of a multiply returning function to another function

// the argument is consumed directly - it's deep nilability
// doesn't matter (but it would if we were checking correct
// variance of nilability types)

// we have already handled

// is a pass of a function to another function, but not multiply returning

// in this case - the identifier for the argument function did not have
// a declaration available, so don't try to consume it

// application is anonymous - no annotations
// TODO implement
// unfortunately, we can't compute the appropriate consumption here

// this is an unpacking of a variadic argument: i.e. the call `foo(_, _, a...)`

// Creates a new param site with location information at every call site
// for a function with contracts. The param site is unique at every call
// site, even with the same function called.

// If arg is a deep type, we add a full trigger for it to track its deep nilability.
// ```
// E.g., func bar(s []*int) {
//   foo(s) // <-- track shallow and deep nilability of `s` here
// }
// ```

// since this is an implicit tracking of the deep nilability of arg, we don't need to
// check for its guarding

// here we have found a call to a function whose declaration we have access to,
// so we can mark its arguments as consumed

// Add Productions for struct field params

// Add Consumptions for struct field params

// here we have found either a builtin function like make or new,
// or a typecast like int(x) - in either case (at least for now), do nothing to try
// to consume the arguments

// when we reach this point, consumeArg will be set to a no-op exactly if we don't know
// how to process consumption of this function's arguments (e.g. anonymous funcs) or if
// we already have, namely through the multiple consumption case above

// if arguments are to a known-annotated function, consume with its annotations

// check if this is just qualifying a package:

// A selector expression (`X.Sel`, where X is an expression and Sel is a selector) can be handled in the following two ways:
// - (1) Allow the expression X to be nilable by creating a TriggerIfNonNil consumer for it. This is a special case,
//       with so far the only known case being of method invocations for supporting nilable receivers. Our support
//       is currently limited to enabling this analysis only if the below criteria is satisfied.
//       - Check 1: selector expression is a method invocation (e.g., `s.foo()`)
//       - Check 2: receiver is a pointer receiver (e.g., `func (s *S) foo()` or `func (*S) foo()`). Go automatically
//			dereferences a value (non-pointer) receiver when a method is called on a pointer to the type. This means that
//			this is not a candidate for analyzing nilable receiver, instead we should check for nilablilty of the
//			receiver at the call site itself.
//       - In-scope flow:
//       	- Check 3: the invoked method is in scope
//       	- Check 4: the invoking expression (caller) is of a non-interface type (e.g., struct or named). (We are
//       		restricting support only for non-interfaces due to the challenges of secret nil for interfaces.)
//       - Out-of-scope flow:
//          - Check 5: consider the criteria satisfied to support optimistic default
//
// - (2) Don't allow the expression X to be nilable by creating a FldAccess (ConsumeTriggerTautology) consumer for it.
//       This is default behavior which gets triggered if the above special case is not satisfied.

// Check 1:  selector expression is a method invocation

// Check 2: receiver is an explicit or implicit pointer receiver

// Check 3: invoked method is in scope
// Here, `t` can only be of type interface, struct, or named, of which we only support for struct and named types.
// Check 4: invoking expression (caller) is of a non-interface type (e.g., struct or named)

// We are in the special case of supporting nilable receivers! Can be nilable depending on declaration annotation/inferred nilability.

// Check 5: invoked method is out of scope
// We are setting an optimistic default here for methods out of scope, specifically to avoid
// false positives being reported for methods in generated code. It means that such external
// methods are assumed to be safely handling nil receivers

// We are in the default case -- it's a field/method access! Must be non-nil.

// similar to index case

// zero slicing contains b[:0] b[0:0] b[0:] b[:] b[:0:0] b[0:0:0], which are safe even when b is
// nil, so we do not create consumer triggers for those slicing.

// For all the other slicing, the slice must be nonnil, so we create a consumer
// trigger.

// pointer load! definitely must be non-nil

// doesn't need to be non-nil, but really should be

// Note if expr.Op == token.ARROW it represents a channel receive (<-X), and we have:
// (1) A receive from a nil channel blocks forever;
// (2) A receive from a closed channel returns the zero value immediately.
// (1) falls out of scope of NilAway, and we have a lot of valid Go code that receives
// from nil channels (e.g., select statements with nilable channels). So we do not create
// consumer for the channel variable here. For (2), since we currently do not track the
// state of channels, we currently cannot support it either.
// TODO: rethink our strategy of handling channels (#192).

// TODO: analyze the bodies of anonymous functions

// TODO - once debugger is working - fill in cases here
// if we don't recognize the node - do nothing

// getFuncIdent returns the function identified from a call expression. If the function
// is an anonymous function, it will return the fake function declaration created in the
// function analyzer
func getFuncIdent(expr *ast.CallExpr, fc *FunctionContext) *ast.Ident {
	_ = "STUB: not implemented"
	return nil
}

// if ident is nil, check if the expr represents a FuncLit node

// check if the declaration the ident points to a function literal node

// getFuncLitFromAssignment if the declaration of the ident is an assignment
// statement and Rhs of the assignment is a call expression which represents an
// anonymous function, returns the ident of the fake function declaration created
// for that. Otherwise, return nil.
func getFuncLitFromAssignment(ident *ast.Ident) *ast.FuncLit { _ = "STUB: not implemented"; return nil }

// TODO get the correct ident for many to one assignments

// LiftFromPath takes a `path` of assertion nodes, and searches for it in the assertion tree rooted
// at `rootNode`. If found, it removes that tree and returns its root as `node`, with `ok` = true.
// If not found, it returns `node`, `ok` = nil, false
//
// This is used as the first half of an assignment between trackable expressions. The two halves are
// kept separate to allow them to be separated into two parallel phases in the case of multiple
// assignments, but for illustrative purposes, here is how a self-contained single assignment method
// would look:
//
// ```
//
//	func (rootNode *RootAssertionNode) AddAssignment(dstpath, srcpath TrackableExpr) {
//		node, ok := rootNode.LiftFromPath(dstpath)
//		if ok {
//			rootNode.LandAtPath(srcpath, node)
//		}
//	}
//
// ```
func (r *RootAssertionNode) LiftFromPath(path TrackableExpr) (AssertionNode, bool) {
	_ = "STUB: not implemented"
	return *new(AssertionNode), false
}

// LandAtPath takes a `path` of assertion nodes, and another target `node`, and places that target
// into the assertion tree rooted at `rootNode` at the location specified by `path`. It fails only
// if `path` is nil.
//
// This is used as the second half of an assignment between trackable expressions. For information on
// why this is done, and an example of how to complete an entire assignment, see `LiftFromPath`'s
// documentation.
func (r *RootAssertionNode) LandAtPath(path TrackableExpr, node AssertionNode) {
	_ = "STUB: not implemented"
	return
}

// To restrict the assertion tree from growing unboundedly, we add node.children to `newNode` iff
// they are not equal to `newNode` itself.

// RootFunc is a function type taking a RootAssertionNode pointer as a parameter
type RootFunc = func(*RootAssertionNode)

// ProcessEntry is called when an assertion tree is known to have reached the entry to its function
// It takes any remaining assertions (consumeTriggers) and conclusively resolves them
// (see for len(self.Children()) > 0) condition by:
// - producing all parameters to the function from their appropriate annotations (paramAnnotationKey)
// - producing all non-parameter variables as definitely nil (noVarAssign)
// - producing all remaining function assertions according to their annotation (retAnnotationKey)
func (r *RootAssertionNode) ProcessEntry() { _ = "STUB: not implemented"; return }

// process field Assertion nodes of function parameters

// filter triggers for error return handling -- intra-procedural

// performs a shallow comparison of two nodes - doesn't recur into their subtrees and doesn't look at triggers
// invariant on AssertionNodes is that this can never hold between any two of their distinct children
func (r *RootAssertionNode) shallowEqNodes(left, right AssertionNode) bool {
	_ = "STUB: not implemented"
	return false
}

// TODO: remove this when  is implemented and we can replace it with a real suppression

// compares full equality, used as fixed point condition for iteration
func (r *RootAssertionNode) eqNodes(left, right AssertionNode) bool {
	_ = "STUB: not implemented"
	return false
}

// nodes have different types!

// checks if a builtin - e.g. "new" or "make"
func (r *RootAssertionNode) isBuiltIn(ident *ast.Ident) bool {
	_ = "STUB: not implemented"
	return false
}

// builtInConversionFuncBasicType checks if it is a built-in conversion function call, such as `string(x)`.
// If yes returns the basic type object, otherwise nil.
func (r *RootAssertionNode) builtInConversionFuncBasicType(call *ast.CallExpr) (b *types.Basic) {
	_ = "STUB: not implemented"
	return nil
}

// checks if a constant - e.g. "true"
func (r *RootAssertionNode) isConst(ident *ast.Ident) bool { _ = "STUB: not implemented"; return false }

// checks if the literal value nil
func (r *RootAssertionNode) isNil(ident *ast.Ident) bool {
	_ = "STUB: not implemented"
	// sometimes we have to insert freshly created nil literal ast nodes, so we add this check to make
	// sure they're handled
	// it's sound because nil is a reserved name, so if an object is named nil it really has to be nil
	// the case handled below takes care of known compile time aliases of nil
	return false
}

// checks if this expression is an instance of types.Func
// this condition holds only for functions defined in the source - not builtins
func (r *RootAssertionNode) isFunc(ident *ast.Ident) bool { _ = "STUB: not implemented"; return false }

// checks if this expression is an instance of types.Var
func (r *RootAssertionNode) isVariable(ident *ast.Ident) bool {
	_ = "STUB: not implemented"
	return false
}

// checks if this is a package name
func (r *RootAssertionNode) isPkgName(expr ast.Expr) bool { _ = "STUB: not implemented"; return false }

// checks if this is a type name
func (r *RootAssertionNode) isTypeName(expr ast.Expr) bool { _ = "STUB: not implemented"; return false }

// checks if an expression is a type
func (r *RootAssertionNode) isType(expr ast.Expr) bool { _ = "STUB: not implemented"; return false }

// isZeroSlicing returns if the given slice expression is a special case that will not cause panic
// even when the slice itself is nil, i.e, one of [:0] [0:0] [0:] [:] [:0:0] [0:0:0]
func (r *RootAssertionNode) isZeroSlicing(expr *ast.SliceExpr) bool {
	_ = "STUB: not implemented"
	return false
}

// [:0] [0:0]
// [0:] [:]
// [:0:0] [0:0:0]

// This function defines whether an expression is `stable` - i.e. whether we assume it constant
// across multiple syntactic accesses. This obviously includes literal expressions closed under
// builtin logical and arithmetic expressions, but, by assumption, includes function calls and
// indexes by other `stable` expressions
func (r *RootAssertionNode) isStable(expr ast.Expr) bool { _ = "STUB: not implemented"; return false }

// There are three cases in which we admit an identifier is a stable:
// if it is a builtin name, if it is a function name, or if it is const.
// Package is considered a special case of ident to suppport selector expressions used to access stable
// expressions, such as constants declared in another package (e.g., pkg.Const)

// TODO: check for function names

// Between two stable expressions, check if we expect them to produce the same value
// precondition: isStable(left) && isStable(right), then checks if left and right are equal
func (r *RootAssertionNode) eqStable(left, right ast.Expr) bool {
	_ = "STUB: not implemented"
	return false
}

// Check if the call expr is a built-in conversion function, e.g., `string(x)`

// if the two identifiers are special values, just check them for string equality

// here, we have eliminated all of the cases in which
// non-variable identifiers can be equal, so if either side is a
// non-variable then the sides are not equal

// if they are variables, check them for declaration equality

// precondition: shallowEqNodes(left, right), then copies remaining data from RIGHT INTO LEFT
func (r *RootAssertionNode) mergeInto(left, right AssertionNode) { _ = "STUB: not implemented"; return }

// merge in consumers

// merge in children

// no existing matching child found, so add one
