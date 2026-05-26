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

package structfield

import (
	"go/ast"
	"go/types"

	"go.uber.org/nilaway/annotation"
	"go.uber.org/nilaway/util/analysishelper"
)

// fieldUse is an int type to indicate how a field is used in the function
type fieldUse uint8

// The use of a field can be represented as `assigned` (e.g., s.f = v) or `accessed` (e.g., g(s.f))
const (
	Assigned fieldUse = 1 << iota // 0b01
	Accessed                      // 0b10
)

// relevantFieldsMap is a type to store the assigned/accessed fields of a struct in a function represented by the ParamAnnotationKey
type relevantFieldsMap map[annotation.ParamAnnotationKey]map[string]fieldUse

// FieldContext stores field information (i.e., assignment and/or access) collected by parsing a function
type FieldContext struct {
	fieldMap relevantFieldsMap
}

// IsFieldUsedInFunc returns true if the passed `fieldName` of struct at index `param` is found to be direct used in the function `funcDecl` for assignment or access
func (f *FieldContext) IsFieldUsedInFunc(funcDecl *types.Func, param int, fieldName string, expectedUse fieldUse) bool {
	_ = "STUB: not implemented"
	return false
}

// addEntry adds a new entry in the map for a field that was assigned or accessed
func (f *FieldContext) addEntry(funcDecl *types.Func, param int, fieldName string, use fieldUse) {
	_ = "STUB: not implemented"
	return
}

// TODO: this check should not be required after . Currently it is needed since NilAway reports a FP on `f.fieldMap[p][fieldName] = ...`

// helps to encode if the field is used as both, assigned and accessed

// processFunc parses the function body for collecting field uses (i.e., assignments and accesses) of a given struct passed as a parameter to the function
func (f *FieldContext) processFunc(funcDecl *ast.FuncDecl, pass *analysishelper.EnhancedPass) {
	_ = "STUB: not implemented"
	// get all assigned and accessed selector expressions from `funcDecl` of the form `s.f`
	return
}

// if fldRefUse.x is a receiver or parameter, then add field name (fldRefUse.field.Name) to the map

// match with params

// match with receiver

// fieldRefUse stores the selector expression `x.field`, where `x` and `field` are of the type *ast.Ident and `use` indicates how `x.field` was used in the function, assigned or accessed
type fieldRefUse struct {
	x     *ast.Ident
	field *ast.Ident
	use   fieldUse
}

func newIdentSelExpr(x *ast.Ident, sel *ast.Ident, use fieldUse) *fieldRefUse {
	_ = "STUB: not implemented"
	return nil
}

// collectIdentSelExpr walks over the AST of the function `f` to collect all selector expressions of the form `s.f`, where `s` is an
// identifier (expressions like `g().f` will be ignored). Each selector expression is marked as `assigned` or `accessed`,
// depending on how it was used within the function. (`assigned` expressions come from LHS of `ast.AssignStmt`, while all other selector
// expressions are grouped as `accessed`.)
func collectIdentSelExpr(f *ast.FuncDecl) []*fieldRefUse { _ = "STUB: not implemented"; return nil }

// assignedSelExprVisited keeps a track of the assigned selector expressions (LHS of ast.AssignStmt) that have been
// visited to prevent those selector expressions from being processed again as "accessed" `case *ast.SelectorExpr`.
// This can happen since ast.Inspect() visits each node in the AST in a DFS fashion

// Simple assignments (e.g., s.f = v) as well as compound assignments in `if` (e.g., if s.f = v; <cond> {}), `for`, `switch`,
// and `typeswitch` are covered by the AssignStmt processing here
