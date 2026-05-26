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

// Package asthelper implements utility functions for AST.
package asthelper

import (
	"go/ast"
	"go/token"
	"go/types"
	"io"
)

// DocContains returns true if the comment group contains the given string.
func DocContains(file *ast.File, s string) bool { _ = "STUB: not implemented"; return false }

// The comment group here contains all comments in the file. However, we should only check
// the comments before the package name (e.g., `package Foo`) line.

// PrintExpr converts AST expression to string, and shortens long expressions if isShortenExpr is true
func PrintExpr(e ast.Expr, fset *token.FileSet, isShortenExpr bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// traverse over the AST expression's subtree and shorten long expressions
// (e.g., s.foo(longVarName, anotherLongVarName, someOtherLongVarName) --> s.foo(...))

func printExpr(writer io.Writer, fset *token.FileSet, e ast.Expr) (err error) {
	_ = "STUB: not implemented"
	// _shortenExprLen is the maximum length of an expression to be printed in full. The value is set to 3 to account for
	// the length of the ellipsis ("..."), which is used to shorten long expressions.
	return nil
}

// fullExpr returns true if the expression is short enough (<= _shortenExprLen) to be printed in full

// ExtractLHSRHS extracts the left-hand side and right-hand side of an assignment statement or a variable declaration
func ExtractLHSRHS(node ast.Node) (lhs, rhs []ast.Expr) { _ = "STUB: not implemented"; return nil, nil }

// IsLiteral returns true if `expr` is a literal that matches with one of the given literal values (e.g., "nil", "true", "false)
func IsLiteral(expr ast.Expr, literals ...string) bool { _ = "STUB: not implemented"; return false }

// CallExprFromExpr returns the call expression from the given expression. It recursively
// traverses the expression tree to find the call expression. If the expression is not a
// call expression, it returns nil.
func CallExprFromExpr(expr ast.Expr) *ast.CallExpr { _ = "STUB: not implemented"; return nil }

// GetSelectorExprHeadIdent gets the head of the chained selector expression if it is an ident. Returns nil otherwise
func GetSelectorExprHeadIdent(selExpr *ast.SelectorExpr) *ast.Ident {
	_ = "STUB: not implemented"
	return nil
}

// IsFieldSelectorChain returns true if the expr is chain of idents. e.g, x.y.z
// It returns for false for expressions such as x.y().z
func IsFieldSelectorChain(expr ast.Expr) bool { _ = "STUB: not implemented"; return false }

// IsEmptyExpr checks if an expression is the empty identifier
func IsEmptyExpr(expr ast.Expr) bool { _ = "STUB: not implemented"; return false }

// GetFieldVal returns the assigned value for the field at index. compElts holds the  elements of the composite literal expression
// for struct initialization
func GetFieldVal(compElts []ast.Expr, fieldName string, numFields int, index int) ast.Expr {
	_ = "STUB: not implemented"
	return *new(ast.Expr)
}

// In this case the initialization is serial e.g. a = &A{p, q}

// FuncIdentFromCallExpr return a function identified from a call expression, nil otherwise
// nilable(result 0)
func FuncIdentFromCallExpr(expr *ast.CallExpr) *ast.Ident { _ = "STUB: not implemented"; return nil }

// case of anonymous function

// GetFunctionParamNode returns the ast param node matching the variable searchParam
func GetFunctionParamNode(funcDecl *ast.FuncDecl, searchParam *types.Var) ast.Expr {
	_ = "STUB: not implemented"
	return *new(ast.Expr)
}
