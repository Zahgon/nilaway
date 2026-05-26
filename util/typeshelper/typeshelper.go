//  Copyright (c) 2025 Uber Technologies, Inc.
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

// Package typeshelper implements utility functions for the go/types package.
package typeshelper

import (
	"go/types"
)

// ErrorType is the type of the builtin "error" interface.
var ErrorType = types.Universe.Lookup("error").Type()

// ErrorInterface is the underlying type of the builtin "error" interface.
var ErrorInterface = ErrorType.Underlying().(*types.Interface)

// BoolType is the type of the builtin "bool" interface.
var BoolType = types.Universe.Lookup("bool").Type()

// BuiltinLen is the builtin "len" function object.
var BuiltinLen = types.Universe.Lookup("len")

// BuiltinAppend is the builtin "append" function object.
var BuiltinAppend = types.Universe.Lookup("append")

// BuiltinNew is the builtin "new" function object.
var BuiltinNew = types.Universe.Lookup("new")

// IsDeep checks if a type is an expression that admits deep nilability, such as maps, slices, arrays, etc.
// Only consider pointers to deep types (e.g., `var x *[]int`) as deep type,
// not pointers to basic types (e.g., `var x *int`) or struct types (e.g., `var x *S`)
func IsDeep(t types.Type) bool { _ = "STUB: not implemented"; return false }

// IsSlice returns true if `t` is of slice type
func IsSlice(t types.Type) bool { _ = "STUB: not implemented"; return false }

// IsDeeplyArray returns true if `t` is of array type, including
// transitively through Named types
func IsDeeplyArray(t types.Type) bool { _ = "STUB: not implemented"; return false }

// IsDeeplySlice returns true if `t` is of slice type, including
// transitively through Named types
func IsDeeplySlice(t types.Type) bool { _ = "STUB: not implemented"; return false }

// IsDeeplyMap returns true if `t` is of map type, including
// transitively through Named types
func IsDeeplyMap(t types.Type) bool { _ = "STUB: not implemented"; return false }

// IsDeeplyPtr returns true if `t` is of pointer type, including
// transitively through Named types
func IsDeeplyPtr(t types.Type) bool { _ = "STUB: not implemented"; return false }

// IsDeeplyChan returns true if `t` is of channel type, including
// transitively through Named types
func IsDeeplyChan(t types.Type) bool { _ = "STUB: not implemented"; return false }

// AsDeeplyStruct returns underlying struct type if the type is struct type or a pointer to a struct type
// returns nil otherwise
func AsDeeplyStruct(typ types.Type) *types.Struct { _ = "STUB: not implemented"; return nil }

// IsDeeplyInterface returns true if `t` is of struct type, including
// transitively through Named types
func IsDeeplyInterface(t types.Type) bool { _ = "STUB: not implemented"; return false }

// IsPointer checks whether the type `t` is an explicit or implicit pointer type, which could also be of deep type.
// Examples of explicit pointer types are `*int`, `*S`, etc.
// Examples of implicit pointer types are `[]int`, `map[string]*S`, `chan int`, etc.
func IsPointer(t types.Type) bool { _ = "STUB: not implemented"; return false }

// UnwrapPtr unwraps a pointer type and returns the element type. For all other types it returns
// the type unmodified.
func UnwrapPtr(t types.Type) types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

// PartiallyQualifiedFuncName returns the name of the passed function, with the name of its receiver
// if defined
func PartiallyQualifiedFuncName(f *types.Func) string { _ = "STUB: not implemented"; return "" }

// FuncNumResults looks at a function declaration and returns the number of results of that function
func FuncNumResults(decl *types.Func) int { _ = "STUB: not implemented"; return 0 }

// ImplementsError checks if the given object implements the error interface. It also covers the case of
// interfaces that embed the error interface.
func ImplementsError(t types.Type) bool { _ = "STUB: not implemented"; return false }

// FuncIsErrReturning encodes the conditions that a function is deemed "error-returning".
// This guards its results to require an `err` check before use as nonnil.
// A function is deemed "error-returning" iff it has a single result of type `error`, and that
// result is the last in the list of results.
func FuncIsErrReturning(sig *types.Signature) bool { _ = "STUB: not implemented"; return false }

// FuncIsOkReturning encodes the conditions that a function is deemed "ok-returning".
// This guards its results to require an `ok` check before use as nonnil.
// A function is deemed "ok-returning" iff it has a single result of type `bool`, and that
// result is the last in the list of results.
func FuncIsOkReturning(sig *types.Signature) bool { _ = "STUB: not implemented"; return false }

// GetParamObjFromIndex get the variable corresponding to the parameter from the function functionType
func GetParamObjFromIndex(functionType *types.Func, argIdx int) *types.Var {
	_ = "STUB: not implemented"
	return nil
}

// In this case the argument is given to a variadic function and the object is last element of the param signature

// IsIterType returns true if the underlying type is an iterator func:
//
// func(func() bool)
// func(func(K) bool)
// func(func(K, V) bool)
//
// See more at https://tip.golang.org/doc/go1.23.
func IsIterType(t types.Type) bool {
	_ = "STUB: not implemented"
	// Ensure it is a function signature.
	return false
}

// Ensure it has exactly one parameter (the yield func).

// Ensure the single parameter is a function type (the yield func).

// Ensure the yield func takes fewer than 2 arguments and returns exactly one boolean value.

// Final check: ensure the return type of the yield func is a boolean.

// GetFuncSignature returns the signature of a function or an anonymous function.
func GetFuncSignature(t types.Type) *types.Signature { _ = "STUB: not implemented"; return nil }

// If the alias is a named function pointer, we extract its signature.
// Example: `type MyFunc func() (*int, error)`

// TypeBarsNilness returns false iff the type `t` is inhabited by nil.
func TypeBarsNilness(t types.Type) bool { _ = "STUB: not implemented"; return false }

// function-types are not inhabited by nil

// all basic types except UntypedNil are not inhabited by nil
