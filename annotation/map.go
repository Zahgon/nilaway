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

package annotation

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"regexp"

	"go.uber.org/nilaway/util/analysishelper"
)

// Map is an abstraction that concrete annotation maps must implement to be checked against.
type Map interface {
	CheckFieldAnn(*types.Var) (Val, bool)
	CheckFuncParamAnn(*types.Func, int) (Val, bool)
	CheckFuncRetAnn(*types.Func, int) (Val, bool)
	CheckFuncRecvAnn(*types.Func) (Val, bool)
	CheckDeepTypeAnn(*types.TypeName) (Val, bool)
	CheckGlobalVarAnn(*types.Var) (Val, bool)
	CheckFuncCallSiteParamAnn(*CallSiteParamAnnotationKey) (Val, bool)
	CheckFuncCallSiteRetAnn(*CallSiteRetAnnotationKey) (Val, bool)
}

// Val is a possible value of an Annotation
type Val struct {
	IsNilable        bool
	IsDeepNilable    bool
	IsNilableSet     bool
	IsDeepNilableSet bool
}

// EmptyVal indicates an annotation value that is fully nonnil but not "set"
var EmptyVal = Val{
	IsNilable:        false,
	IsDeepNilable:    false,
	IsNilableSet:     false,
	IsDeepNilableSet: false,
}

// makeNilable inspects a Val to see if its nilability has already been set.
// If it has, then makeNilable is a noop, otherwise, it returns a copy of the passed
// Val with the nilability set to true.
// The parameter isFinalVal indicates whether later calls to these utility functions
// can update the `Val` further - uses cases are commented on to discuss why
// this is or is not desirable in given cases.
func (a Val) makeNilable(isFinalVal bool) Val { _ = "STUB: not implemented"; return *new(Val) }

// makeDeepNilable inspects a Val to see if its deep nilability has already been set.
// If it has, then makeDeepNilable is a noop, otherwise, it returns a copy of the passed
// Val with the deep nilability set to true.
// The parameter isFinalVal indicates whether later calls to these utility functions
// can update the `Val` further - uses cases are commented on to discuss why
// this is or is not desirable in given cases.
func (a Val) makeDeepNilable(isFinalVal bool) Val { _ = "STUB: not implemented"; return *new(Val) }

// makeNonNil inspects a Val to see if its nilability has already been set.
// If it has, then makeNonNil is a noop, otherwise, it returns a copy of the passed
// Val with the nilability set to false.
// The parameter isFinalVal indicates whether later calls to these utility functions
// can update the `Val` further - uses cases are commented on to discuss why
// this is or is not desirable in given cases.
func (a Val) makeNonNil(isFinalVal bool) Val { _ = "STUB: not implemented"; return *new(Val) }

// makeDeepNonNil inspects a Val to see if its deep nilability has already been set.
// If it has, then makeDeepNonNil is a noop, otherwise, it returns a copy of the passed
// Val with the deep nilability set to false.
// The parameter isFinalVal indicates whether later calls to these utility functions
// can update the `Val` further - uses cases are commented on to discuss why
// this is or is not desirable in given cases.
func (a Val) makeDeepNonNil(isFinalVal bool) Val { _ = "STUB: not implemented"; return *new(Val) }

// A ObservedMap represents a completed set of annotations read from a file or set of files,
// it can be checked against an assertionTree using RootAssertionNode.ReportErrors
//
// The maps are keyed by *ast.Idents because such an object is unique at each site in the code
// it is used; canonically, declarations are identified with the identifier used at the site
// of the declaration
//
// TODO: handle annotations for anonymous functions too
type ObservedMap struct {
	// this maps fields by the identifier declaring them to their Annotation type
	fieldAnnMap map[*types.Var]Val

	// this maps functions by the identifier declaring them to a slice with the
	// annotations of its params
	funcParamAnnMap map[*types.Func][]Val

	// this maps functions by the identifier declaring them to a slice with the
	// annotations of its results
	funcRetAnnMap map[*types.Func][]Val

	// this maps functions by the identifier declaring them to the annotations of their receivers
	funcRecvAnnMap map[*types.Func]Val

	// this maps named types to annotations describing their deep nilability
	deepTypeAnnMap map[*types.TypeName]Val

	// this maps declarations of global variables to their annotations
	globalVarsAnnMap map[*types.Var]Val

	// funcCallSiteParamAnnMap maps a function call site to a slice with the annotations of its
	// duplicated params at the call site.
	funcCallSiteParamAnnMap map[CallSite][]ArgLocAndVal

	// funcCallSiteRetAnnMap maps a function call site to a slice with the annotations of its
	// duplicated returns at the call site.
	funcCallSiteRetAnnMap map[CallSite][]Val
}

// CallSite uniquely identifies a function call. It contains the called function object and the
// code location of the call expression.
type CallSite struct {
	Fun      *types.Func
	Location token.Position
}

// ArgLocAndVal pairs the code location of the argument expression and the annotation value.
type ArgLocAndVal struct {
	Location token.Position
	Val      Val
}

// Range calls the passed function `op` on each annotation site in this map. If `setSitesOnly`
// is true, then it only calls `op` only on the sites with is<Deep?>NilableSet true.
func (m *ObservedMap) Range(op func(key Key, isDeep bool, val bool), setSitesOnly bool) {
	_ = "STUB: not implemented"
	return
}

/* isDeep */

/* isDeep */

// the location inside the callSite is the location of the call expression, we want
// the location of every argument expression

// defaults for anonymous functions and structs (ones for which definitions just can't be found
// aren't even looked up for now)
var (
	nonAnnotatedDefault = EmptyVal
)

const nilableKeyword = "nilable"
const nonNilKeyword = "nonnil"

var annotationKeyword = fmt.Sprintf("(%s|%s)", nilableKeyword, nonNilKeyword)

const sep = ","
const identRegexStr = "[a-zA-Z][a-zA-Z0-9]*"

const paramTemplateStr = "param %s"

var paramRegexStr = fmt.Sprintf(paramTemplateStr, "[0-9]+")

const resultTemplateStr = "result %s"

var resultRegexStr = fmt.Sprintf(resultTemplateStr, "[0-9]+")

var tokenRegexStr = fmt.Sprintf("((%s)|(%s)|(%s))",
	identRegexStr, paramRegexStr, resultRegexStr)

func paramStr(i int) string { _ = "STUB: not implemented"; return "" }

func resultStr(i int) string { _ = "STUB: not implemented"; return "" }

var deepIdentRegexStr = fmt.Sprintf("((\\*%s)|(%s\\[\\])|(<-%s)|%s)",
	tokenRegexStr, tokenRegexStr, tokenRegexStr, tokenRegexStr)
var seqRegexStr = fmt.Sprintf("%s\\((\\s*%s\\s*(%s\\s*%s\\s*)*)\\)",
	annotationKeyword, deepIdentRegexStr, sep, deepIdentRegexStr)
var seqRegex = regexp.MustCompile(seqRegexStr)

type nilabilitySet map[string]Val

// from a CommentGroup return a nilabilitySet of which identifiers are known annotated nilable
func nilabilityFromCommentGroup(group *ast.CommentGroup) nilabilitySet {
	_ = "STUB: not implemented"
	return *

	// in each of the following utility functions, isFinalVal=true because literally read annotations
	// are considered final
	new(nilabilitySet)
}

// TypeIsDefaultNilable takes a type and returns true iff we assume default nilability for that
// type - in contrast to the remaining cases, in which we assume default non-nil.
func TypeIsDefaultNilable(t types.Type) bool { _ = "STUB: not implemented"; return false }

// Builtin error type should be nilable by default.

// Slice, map, and chan should also be nilable by default (after unwrapping the named type).

// TypeIsDeepDefaultNilable takes an `ast.Expr` that evaluates to a type, and returns true iff
// we assume default deep nilability for that type - in contrast to the remaining cases, in which
// we assume default deep non-nil.
func TypeIsDeepDefaultNilable(t types.Type) bool { _ = "STUB: not implemented"; return false }

// the array case is handled different from others, since an array is not default nilable,
// but can be default deeply nilable based on its containing type

// recurse if multi-dimensional array until containing type is reached

// assign deep nilability based on the element type

// checkNilability for a nilabilitySet checks to see if a string is mapped to an Annotation by that
// set. If it is, then that Annotation is returned. If not, then `nonNil` is returned.
// the type of the Annotation site is also passed, and it can possibly serve to mark a site
// as `nilable` when its Annotation doesn't indicate so.
func (set nilabilitySet) checkNilability(name string, t types.Type) Val {
	_ = "STUB: not implemented"
	return *new(Val)
}

// in each of the following cases, isFinalVal=false because defaults are not considered final

func newObservedMap(pass *analysishelper.EnhancedPass, files []*ast.File) *ObservedMap {
	_ = "STUB: not implemented"
	return nil
}

// TODO - only store annotations for fields/vars/parameters of types that do not bar nilness

// for a function declaration, accumulate its parameters from an *ast.Fieldlist object
// listing them, look them up in the docstring, and return an equally long list of
// annotationVals

// this is included for nil-safety

// case of anonymous field - on which we do not permit annotations
// non-named fields

// for each named field, check the docstring for its Annotation and append that

// in the case that our argument is variadic (hence has a type expression
// of the form `...T`, we treat the arguments as having type `T` not type
// `T[]`

// store the mapping from the function object to the ast node.

// this is used for any declaration besides a function
// here, we specifically look for declarations of struct types

// this set will contain the nilability annotations read from the appropriate
// docstring (this takes into account the syntax option to group declarations -
// in which a single keyword may be used to declare a group)

// this reads declarations like type A struct {}

// this reads declarations like type (A struct{}, B struct{})

// we've found a declaration using the `var` or `const` keyword

// narrow down to the case of a `var` declaration - i.e., a global var

// we've found a declaration for a `type`

// readDeepNilability is called on type declarations of maps, slices, and
// pointers to see if their contained values are nilable

// iterate over the methods of this interface

// this is the common case - a simply declared method

// this is the case of inheritance - i.e. a method with another
// method named within it, in this case the identifiers will
// correctly resolve field references to the super-interface
// so no work needs to be done

// unrecognized

// type alias - do nothing
// type alias - do nothing
// function type - do nothing (for now)

// TODO - treat channel types as deeply nilable at the typedef level

// TODO - handle generics

// do nothing - we don't care about these for annotations' sake

// Parse inline annotations at call sites.

// Store a mapping between single comment's line number to its text.

// Now, find all *ast.CallExpr nodes and find their comment and extract annotations.
// Comment nodes are floating in GO asts. https://github.com/golang/go/issues/20744
// Thus, we require the comments for annotations are written in the same line as the
// function call expression, so we can match them by line numbers.

// No annotations for this CallExpr node, but we still need to traverse further since
// there could be comments for nested CallExpr nodes.

// if ident is nil, keep searching for nested CallExpr nodes.

// not a function, keep searching for nested CallExpr nodes.

// empty set, no annotation, keep searching for nested CallExpr nodes.

// keep searching for nested CallExpr nodes.

func getLineFromPos(pos token.Pos, pass *analysishelper.EnhancedPass) int {
	_ = "STUB: not implemented"
	return 0
}
