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

package analysishelper

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

// EnhancedPass is a drop-in replacement for `*analysis.Pass` that provides additional helper methods
// to make it easier to work with the analysis pass.
type EnhancedPass struct {
	*analysis.Pass
}

// NewEnhancedPass creates a new EnhancedPass from the given *analysis.Pass.
func NewEnhancedPass(pass *analysis.Pass) *EnhancedPass { _ = "STUB: not implemented"; return nil }

// Panic panics with the given message and additional position information.
func (p *EnhancedPass) Panic(msg string, pos token.Pos) { _ = "STUB: not implemented"; return }

// IsZero returns if the given expression is evaluated to integer zero at compile time. For
// example, zero literal, zero const or binary expression that evaluates to zero, e.g., 1 - 1
// should all return true. Note the function will return false for zero string `"0"`.
func (p *EnhancedPass) IsZero(expr ast.Expr) bool { _ = "STUB: not implemented"; return false }

// ConstInt returns the constant integer value of the given expression if it is a constant. The
// boolean return value indicates whether the expression is a constant integer or not.
func (p *EnhancedPass) ConstInt(expr ast.Expr) (int64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// IsNil checks if the given expression evaluates to untyped nil at compile time. It also treats
// the identifier `nil` as nil too to support cases where we have inserted a fake identifier.
func (p *EnhancedPass) IsNil(expr ast.Expr) bool { _ = "STUB: not implemented"; return false }

// HumanReadablePosition modifies the Position's filename to be more human-friendly (truncated or relative to cwd).
func (p *EnhancedPass) HumanReadablePosition(position token.Position) token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

// PosToLocation converts a token.Pos as a real code location, of token.Position.
func (p *EnhancedPass) PosToLocation(pos token.Pos) token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

// ExprBarsNilness returns if the expression can never be nil for the simple reason that nil does
// not inhabit its type.
func (p *EnhancedPass) ExprBarsNilness(expr ast.Expr) bool { _ = "STUB: not implemented"; return false }

// `p.TypesInfo.TypeOf` only checks Types, Uses, and Defs maps in TypesInfo. However, we may
// miss types for some expressions. For example, `f` in `s.f` can only be found in
// `p.TypesInfo.Selections` map (see the comments of p.TypesInfo.Types for more details).
// Be conservative for those cases for now.
// TODO:  to investigate and find more cases.

// IsSliceAppendCall checks if `node` represents the builtin append(slice []Type, elems ...Type) []Type
// call on a slice.
// The function checks 2 things,
// 1) Name of the called function is "builtin append"
// 2) The first argument to the function is a slice
func (p *EnhancedPass) IsSliceAppendCall(node *ast.CallExpr) (*types.Slice, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// ExprIsAuthentic aims to return true iff the passed expression is an AST node
// found in the source program of this pass - not one that we created as an intermediate value.
// There is no fully sound way to do this - but returning whether it is present in the `Types` map
// map is a good approximation.
// Right now, this is used only to decide whether to print the location of the producer expression
// in a full trigger.
func (p *EnhancedPass) ExprIsAuthentic(expr ast.Expr) bool { _ = "STUB: not implemented"; return false }
