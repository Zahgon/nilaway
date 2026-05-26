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

	"go.uber.org/nilaway/assertion/function/producer"
)

// ParseExprAsProducer takes an expression, and determines whether it is `trackable` - i.e. if it is a
// linear sequence of variable reads, field reads, indexes by `stable` expressions, and function calls
// with `stable` arguments. An expression is `stable` if our static analysis assume that multiple
// syntactic occurrences of it will always yield the same value - i.e. they are assumed to be constant.
//
// This function and the cases in which it returns a sequence of nodes serve as our internal
// definition of `trackable`, and similarly the function isStable below serves as our internal
// definition of `stable`.
//
// # Of its two return values, shallowSeq and producer, only one will be non-nil
//
// In the case that expr is trackable, shallowSeq will be non-nil, and contain the AssertionNodes
// without pointers between them that characterize the give expression.
//
// In the case that expr is not trackable, shallowSeq will be nil. If expr is known to be non-nil
// (e.g. a non-nil constant) then producer will be nil too, but otherwise it will be a slice of
// produceTriggers encapsulating the conditions under which expr could be nil. The slice will have
// length 1 for every expr except multiply returning functions, for which it will have length equal
// to the number of returns of that function.
//
// The function also takes a flag doNotTrack which, if set to true, always treats the expr
// as non-trackable and gives its producer trigger or nil if it's not a nilable expression.
//
// ParseExprAsProducer will panic if passed the empty expression `_`
//
// nilable(shallowSeq)
// nilable(producers)
//
// TODO: split this up into smaller functions with more granular documentation
func (r *RootAssertionNode) ParseExprAsProducer(expr ast.Expr, doNotTrack bool) (
	shallowSeq TrackableExpr, producers []producer.ParsedProducer) {
	_ = "STUB: not implemented"
	return *new(TrackableExpr), nil
}

// we assume none of these types of identifiers return nil
// TODO: refine this handling of constants

// in the case of a totally unrecognized identifier - we assume nilability

// by process of elimination, it's a variable, so track it!

// this function represents the case in which we have identified that the value of the
// expression being parsed flows from a _deep read_ to the expression `deepExpr`.
// this function is only to be used in cases when we have determined that the parsed
// expression is not trackable.

// the already parsed prefix to `expr` - all we care about is whether nil
// the expression we identified is being deeply read for this parse
// the overall expression being parsed - used to construct `annotation.ProduceTrigger`s
// the, possibly already set, parse of `deepExpr`
// in general - our goal is to obtain the parse of `deepExpr` - then lift its deep producer to
// the shallow producer of a new `ParsedProducer`, and populate the new deep producer by a default
// based on type name if applicable

// this is so that if the first time we parsed this we determined it was trackable,
// then re re-parse to obtain as a non-tracked producer
// an example case is that the receiver was trackable in an index expression,
// but the index was non-literal

// there is no possible source for a doubly deep nilability annotation except
// the named type of the expression

// if we reach here - that should mean that expr.X is not deeply nilable, so we know this
// read cannot produce nil

// if we've reduced to a package-qualified identifier like pkg.A, just interpret it
// as a bare identifier

// we assume builtins aren't nilable
// functions are definitely not nilable

// treat as non-trackable

// trackable access to a field

// non-trackable access to a field - just return a produce trigger for that field

// we delay this check until we're sure we have to make it, as it could be expensive

// the cases of a function and method call are different enough here that it would be useless
// to try to subsume this switch with funcIdentFromCallExpr

// direct function call
// Handle specially if the function is an anonymous function.

// TODO: this is a temporary fix to handle the case of anonymous functions.
//  Remove this once we have have enabled the anonymous function support.

// Check if it is a type alias for a function type.
// e.g., type MyFunc func() (*int, error)
// func foo(f MyErrRetFunc) {
// 		x, err := f()
// 		if err != nil {
// 			return
// 		}
// 		_ = *x
// }
// TODO: this is only a temporary fix to suppress false positives caused by type aliases.
//  Remove this once we have implemented complete support for type aliases.

// Check if the method is a function value, e.g., `f := func() {}` and then `f()`.
// TODO: this is a temporary fix to suppress false positives caused by function values.
//  Remove this once we have have implemented the function value support.

// The following block implements the basic support for append function where it has
// only two arguments and the first argument is the same as the lhs of assignment.
// Since in Go it is allowed to have only one argument in the append method, we need
// to have a check to make sure that len(expr.Args) > 1

// TODO: handle the correlation of return type of append with its first argument .
// TODO: iterate over the arguments of the append call if it has more than two args

// We are in the case of built-in functions. The below block particularly checks for the case of the
// built-in `new` function for struct initialization handling. The `new` function returns a pointer to
// the passed type (e.g., new(S) returns *S), which is same as creating a struct using composite
// literal `&S{}`. We are interested in handling this case since all fields of the struct `S` would be
// uninitialized with a `new(S)`.
// TODO: below logic won't be required once we standardize the calls by replacing `new(S)` with `&S{}`
//  in the preprocessing phase after  is implemented.

// for builtin funcs (e.g. new, make), we assume their return is never nil
// similarly, we assume type casts (e.g. `int(x)`) never return nil
// anonymous functions will also fall into this case

// non-builtin funcs

// function call has non-literal args, so is not literal, use its return annotation
// alternatively, doNotTrack was set

// method call
// Check if the method is a function value, e.g., `f := func() {}` and then `f()`.
// TODO: this is a temporary fix to handle the case of function values.
//  Remove this once we have have implemented the function value support.

// we assume builtins and type casts don't return nil

// receiver is not trackable, use its return annotation

// function call has non-literal args, so is not literal, use its return annotation

// TODO: This case can possibly be combined with the case of *ast.Ident above.

// TODO: this is a temporary fix to handle the case of anonymous functions.
//  Remove this once we have have enabled the anonymous function support.

// non-builtin funcs

// function call has non-literal args, so is not literal, use its return annotation
// alternatively, doNotTrack was set

// this could result from calling a function returned anonymously from another function, such as f(4)(3), and
// although theoretically we should track that, we're going to leave it as an unhandled edge case for now
// TODO: consider handling this case (and similar case in backPropAcrossReturn)

// X part of the expression is trackable. Now we need to check if the index is stable or trackable.
// If the index is not stable, it is still considered trackable if it falls into any of these categories:
// - Index is a variable (e.g., `m[i]`)
// - Index is a built-in function (e.g., `m[len(m)-1]`)
// - Index is a field selector chain (e.g., `m[g.h.i]`)
// TODO: above non-literal indices should only be considered trackable if no reassignment is found between
//  accesses. For example, `i := 0; if m[i] != nil { i = 10; return *m[i] }` should not be considered trackable
//  as the index `i` is reassigned between accesses. Towards, we plan to add another analyzer pass based on
//  SSA to determine if the index is reassigned between accesses.

// iterate over the arguments of the call expression

// receiver is trackable and index is stable or trackable, so return an augmented path

// index is non-trackable, so the expression is not trackable, just return nilable for index without check

// reciever is non-trackable, just return nilable for index without check

// For slice expressions `b[_:0:_]`, the result is always an empty (nilable in
// NilAway's eyes) slice. (`_` can be anything including empty.)

// We should create a nilable producer.

// For slice expressions `b[0:]` and `b[:]`, the result's nilability depends on the
// nilability of the original slice. Note that you cannot give empty High in 3-index
// slices.

// TODO: for now we directly return the trackable expression of the original slice. We
// should instead properly create a trackable expression for the slice expression. See
//  for more details.

// Return the trackable expression of the original slice

// For all other cases, the result must be a nonnil slice.

// Returning nil to indicate the slice expression results in a nonnil slice.

// TODO - if `recv` is trackable, then track expression instead, as in the index case

// we've found a receive expression

// we treat a struct object pointer (e.g., &A{}) and struct object (e.g., A{}) identically for creating field producers

// simply parse the underlying expression

// TODO: right now this default case assumes that unhandled expressions are non-nil, consider changing this

// getFuncReturnProducers returns a list of producers that are triggered at the call expression
func (r *RootAssertionNode) getFuncReturnProducers(ident *ast.Ident, expr *ast.CallExpr) []producer.ParsedProducer {
	_ = "STUB: not implemented"
	return nil
}

// Creates a new return site with location information at every call site for a
// function with contracts. The return site is unique at every call site, even with the
// same function called.

// for an error-returning function, all but the last result are guarded
// TODO: add an annotation that allows more results to escape from guarding
// such as "error-nonnil" or "always-nonnil"

// parseStructCreateExprAsProducer parses composite expressions used to initialize a struct e.g. A{f1: v1, f2: v2}
func (r *RootAssertionNode) parseStructCreateExprAsProducer(expr ast.Expr, fieldInitializations []ast.Expr) producer.ParsedProducer {
	_ = "STUB: not implemented"
	return *new(producer.ParsedProducer)
}

// we do not create producers for fields that are not nilable

// extract the value assigned to the field in the composite

// this means the field is not assigned any value, thus unassigned field should be produced

// do not track. Get producer for expression `fieldVal` assigned to the field

// since we only track field producers at depth one, we ignore deep producers from the field

// If the field producer is nil, that means it is not a nilable expression
