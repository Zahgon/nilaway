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
	"go/types"
)

// DeepNilabilityAsNamedType tries to interpret the named type as a typedef of a map or slice,
// returning the deep nilability annotation of that typedef if found. Otherwise, it returns
// ProduceTriggerNever to indicate that we assume in the default case the type is NOT deeply nilable
func DeepNilabilityAsNamedType(typ types.Type) ProducingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ProducingAnnotationTrigger)
}

// Calling Underlying on [types.Named] will always return the unnamed type, so we
// do not have to recursively "unwrap" the [types.Named].
// See [https://github.com/golang/example/tree/master/gotypes#named-types].

// DeepNilabilityOfFuncRet inspects a function return for deep nilability annotation
func DeepNilabilityOfFuncRet(fn *types.Func, retNum int) ProducingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ProducingAnnotationTrigger)
}

// DeepNilabilityOfFld inspects a struct field for deep nilability annotation
func DeepNilabilityOfFld(fld *types.Var) ProducingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ProducingAnnotationTrigger)
}

// in this case, the deep nilability of the field comes from its declaring annotations

// DeepNilabilityOfVar inspects a variable for deep nilability annotation
func DeepNilabilityOfVar(fdecl *types.Func, v *types.Var) ProducingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ProducingAnnotationTrigger)
}

// in each of the following cases, the deep nilability of the variable comes from its
// declaring annotations

// in this case, the deep nilability of the variable is dependent only on its possible guarding

// otherwise, the deep nilability of this variable is either that of its named type,
// or not deeply nilable - a logical split captured in the method DeepNilabilityAsNamedType

// VarIsGlobal returns true iff `v` is a global variable
// this check is performed by looking up the package of the variable,
// then the declaring scope of that package,
// then checking that declaring scope to see if it maps the name of `v` to
// the passed `*types.Var` instance of `v`
func VarIsGlobal(v *types.Var) bool { _ = "STUB: not implemented"; return false }

// VarIsParam returns true iff `v` is a parameter of `fdecl`
func VarIsParam(fdecl *types.Func, v *types.Var) bool { _ = "STUB: not implemented"; return false }

// could avoid this iteration by hashing once, but param lists are short enough that
// it feels like a premature optimization

// VarIsVariadicParam returns true iff `v` is a variadic parameter of `fdecl`
func VarIsVariadicParam(fdecl *types.Func, v *types.Var) bool {
	_ = "STUB: not implemented"
	return false
}

// VarIsRecv returns true iff `v` is the receiver of `fdecl`
func VarIsRecv(fdecl *types.Func, v *types.Var) bool { _ = "STUB: not implemented"; return false }

// ParamAsProducer inspects a variable, which must be a parameter to the passed function, and returns
// a produce trigger for its value as annotated at its function's declaration. The interesting case
// is when the parameter is variadic - and then the annotation is interpreted as referring to the
// elements of the variadic slice not the slice itself
func ParamAsProducer(fdecl *types.Func, param *types.Var) ProducingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ProducingAnnotationTrigger)
}

// paramAsDeepProducer inspects a variable, which must be a parameter to the passed function, and returns
// a produce trigger for its deep value as annotated at its function's declaration.
//
// As above, the interesting case is when the parameter is variadic
func paramAsDeepProducer(fdecl *types.Func, param *types.Var) ProducingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ProducingAnnotationTrigger)
}

// NewRetFldAnnKey returns the RetFieldAnnotationKey for return at retNum and the field fieldDecl.
// This function is called from multiple places where funcDecl could be the declaration of the function being analyzed
// (when looking at a return statement) or a function called from the function being analyzed
// (when looking at a function call statement, in which case funcDecl is the callee).
func NewRetFldAnnKey(funcDecl *types.Func, retNum int, fieldDecl *types.Var) *RetFieldAnnotationKey {
	_ = "STUB: not implemented"
	return nil
}

// NewParamFldAnnKey returns ParamFieldAnnotationKey for a field fieldDecl of param at index of the function funcObj
func NewParamFldAnnKey(funcObj *types.Func, index int, fieldDecl *types.Var) *ParamFieldAnnotationKey {
	_ = "STUB: not implemented"
	return nil
}

// NewEscapeFldAnnKey returns a new EscapeFieldAnnotationKey for field fieldObj
func NewEscapeFldAnnKey(fieldObj *types.Var) *EscapeFieldAnnotationKey {
	_ = "STUB: not implemented"
	return nil
}
