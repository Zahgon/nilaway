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
	"go.uber.org/nilaway/assertion/structfield"
)

// addProductionsForAssignmentFields adds production for each produce trigger in fieldProducers.
// fieldProducers contain all the field producers due to the rhs of the assignment.
// lhsVal is the assigned lhs expression
// If the assignment is lhsVal = f(), then the expression for the field producer is `lhsVal.field` for
// the corresponding field
func (r *RootAssertionNode) addProductionsForAssignmentFields(fieldProducers []*annotation.ProduceTrigger, lhsVal ast.Expr) {
	_ = "STUB: not implemented"
	return
}

// addConsumptionsForFieldsOfReturns adds consumptions for each field of retNum-th return of a function.
// This creates consumptions in one of the following cases
// 1. if the return expression is chain of field accesses
// For example, an ident `x` is also a chain of selector expression, thus it will consume all `x.field` for fields
// that have nilable type
// 2. If the expression gives field producers, we create full triggers for those producers with the
// consumers of field return keys
func (r *RootAssertionNode) addConsumptionsForFieldsOfReturns(retExpr ast.Expr, retNum int) {
	_ = "STUB: not implemented"
	return

	// fdecl Type() is always a *Signature
}

// For field selection chains we add the consumptions for fields by creating artificial selector expression

// We do not create field triggers for types that are not nilable

// create an artificial selector expression ast node

// Also add escape consumer

// For the expressions that return field producers we create full triggers

// field producer is nil for fields that have non-nilable type

// Also add escape consumer

// getFieldProducersForFuncReturns creates and returns producers for nilable-typed fields of the struct at return index retNum.
// The method returns nil if the return at index retNum is not a struct. calledFuncDecl is the declaration of the function that
// is called.
func (r *RootAssertionNode) getFieldProducersForFuncReturns(calledFuncDecl *types.Func, retNum int) []*annotation.ProduceTrigger {
	_ = "STUB: not implemented"
	// calledFuncDecl Type() is always *Signature
	return nil
}

// We do not create field triggers for types that are not nilable

// addProductionsForParamFields adds productions for fields of params and receivers. It is called while processing entry
// block during backprop
func (r *RootAssertionNode) addProductionsForParamFields(node AssertionNode, builtExpr ast.Expr) {
	_ = "STUB: not implemented"
	return
}

// we need to make copy of the child nodes as they might get deleted when the productions are matched.
// Each such child node represents a consumption on an argument-field pair s.f, and as we generate the
// corresponding production in `addProductionsForParamFieldNode` (and call `r.AddProduction`
// within that function), the two will get matched into a full trigger, which removes the consumption trigger
// from its current position in the assertion tree.

// We do not add production for types that are not nilable

// addProductionForVarFieldNode adds productions for fields of a struct defined as a local variable represented by varAssertionNode
func (r *RootAssertionNode) addProductionForVarFieldNode(varNode *varAssertionNode, varAstExpr ast.Expr) {
	_ = "STUB: not implemented"
	return
}

// addProductionsForParamFieldNode adds productions for fields of params and receiver at the entry node for a specific
// fldAssertionNode
func (r *RootAssertionNode) addProductionsForParamFieldNode(selExpr *ast.SelectorExpr, node *fldAssertionNode) {
	_ = "STUB: not implemented"
	return
}

// If funcDecl is a method then add production for fields of receivers

// addConsumptionsForArgAndReceiverFields adds consumptions for fields of receivers and arguments of a function call. It takes
// the call expression and adds consumptions for each param and receiver by calling addConsumptionsForArgFields and
// addConsumptionsForReceiverFields respectively
func (r *RootAssertionNode) addConsumptionsForArgAndReceiverFields(call *ast.CallExpr, funcIdent *ast.Ident) {
	_ = "STUB: not implemented"
	return
}

// In case we are dealing with a method call or a function call in another package (and not dot-imported)

// addConsumptionsForReceiverFields adds consumptions for receiver fields at function call
func (r *RootAssertionNode) addConsumptionsForReceiverFields(call *ast.CallExpr, fieldContext *structfield.FieldContext) {
	_ = "STUB: not implemented"
	return
}

// In case we are dealing with a method call add consumers for fields of receiver

// addConsumptionsForReceiverFields adds consumptions for param fields at function call
func (r *RootAssertionNode) addConsumptionsForArgFields(call *ast.CallExpr, funcName *ast.Ident, fieldContext *structfield.FieldContext) {
	_ = "STUB: not implemented"
	return
}

// addConsumptionsForArgFieldsAtIndex adds consumptions for fields of argument at index argIdx of a function call. Thus, it captures the
// nilability flow to the function through the fields of param. in case, argIdx is equal to ReceiverParamIndex then the consumptions
// are created for the receiver
func (r *RootAssertionNode) addConsumptionsForArgFieldsAtIndex(arg ast.Expr, funcObj *types.Func, argIdx int, fieldContext *structfield.FieldContext) {
	_ = "STUB: not implemented"
	return
}

// For field selection chains we add the consumptions for fields by creating artificial selector expression

// only create this trigger if the field was found to be accessed in the function given by `funcObj`

// Also add escape consumer

// For the argument expressions that produce field producers we create full triggers
// e.g., f(&A{...}) or f(g())

// TODO: We only handle first producer of the expressions

// for fields that have non-nilable type we don't do anything

// add escape trigger
// TODO: Do not call this for list of special functions

// addProductionForFuncCallArgAndReceiverFields is called while consuming a function call. Productions for fields of params
// and receivers are added to track the effect the function call can have on the fields.
func (r *RootAssertionNode) addProductionForFuncCallArgAndReceiverFields(call *ast.CallExpr, funcIdent *ast.Ident) {
	_ = "STUB: not implemented"
	return
}

// In case we are dealing with a method call or call to other package

// addProductionForFuncCallReceiverFields adds productions for fields of receivers are added to track the
// effect the function call can have on the fields.
func (r *RootAssertionNode) addProductionForFuncCallReceiverFields(call *ast.CallExpr, fieldContext *structfield.FieldContext) {
	_ = "STUB: not implemented"
	return
}

// addProductionForFuncCallArgFields adds productions for fields of params are added to track the
// effect the function call can have on the fields.
func (r *RootAssertionNode) addProductionForFuncCallArgFields(funcName *ast.Ident, call *ast.CallExpr, fieldContext *structfield.FieldContext) {
	_ = "STUB: not implemented"
	return
}

// addProductionForFuncCallArgFieldsAtIndex is called from addProductionForFuncCallArgFields for individual
// argument arg. Productions are added for fields of arguments of functions to track the side effect
// on the fields due to the function call.
func (r *RootAssertionNode) addProductionForFuncCallArgFieldsAtIndex(arg ast.Expr, methodType *types.Func, argIdx int, fieldContext *structfield.FieldContext) {
	_ = "STUB: not implemented"
	return
}

// addConsumptionsForFieldsOfParams is called at the return statement of the function during backprop. This consumption captures the
// possible side effect on the fields of params that the function call can have
func (r *RootAssertionNode) addConsumptionsForFieldsOfParams() { _ = "STUB: not implemented"; return }

// `funcSig.Recv() != nil` actually implies `r.FuncDecl().Recv != nil` since they are referring
// to the same function represented in two different systems (i.e., language objects and ASTs).
// However, NilAway does not know this correlation, hence this redundant check.

// The [language specs] require that the receiver must be a single non-variadic parameter.
// However, they are represented as a regular parameter _list_ in the AST. Here, we can
// safely use the first element of the list.
// [language specs]: https://go.dev/ref/spec#Method_declarations

// The length of receivers can only be 0 (unnamed receiver) or 1 (named receiver).
// We only need to handle the named case if it is not an empty (`_`) receiver.

// addConsumptionsForFieldsOfParam is called by addConsumptionsForFieldsOfParams for parameter at index paramIdx
func (r *RootAssertionNode) addConsumptionsForFieldsOfParam(param *types.Var, paramNode ast.Expr, paramIdx int, fieldContext *structfield.FieldContext) {
	_ = "STUB: not implemented"
	return
}

// getParamFieldKey returns param field annotation key and selector expression that selects the field of expression arg
func (r *RootAssertionNode) getParamFieldKey(arg ast.Expr, methodType *types.Func, argIdx int, structType *types.Struct, fieldID int) (annotation.Key, *ast.SelectorExpr) {
	_ = "STUB: not implemented"
	return *new(annotation.Key), nil
}

// addEscapeFullTrigger adds escape full trigger for the field with fieldIdx
func (r *RootAssertionNode) addEscapeFullTrigger(expr ast.Expr, structType *types.Struct, fieldIdx int, fieldProducer *annotation.ProduceTrigger) {
	_ = "STUB: not implemented"
	return
}

// getSelectorExpr gets the declaring ident for field fieldDecl, and returns the selector expression
func (r *RootAssertionNode) getSelectorExpr(fieldDecl *types.Var, fieldOf ast.Expr) *ast.SelectorExpr {
	_ = "STUB: not implemented"
	return nil
}
