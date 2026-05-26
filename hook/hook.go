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

// Package hook implements a hook framework for NilAway where it hooks into different parts to
// provide additional context for certain function calls. This is useful for well-known standard
// or 3rd party libraries where we can encode certain knowledge about them (e.g.,
// `assert.Nil(t, x)` implies `x == nil`) and use that to provide better analysis.
package hook

import (
	"go/ast"
	"go/token"
	"regexp"

	"go.uber.org/nilaway/util/analysishelper"
)

// funcKind indicates the kind of the trusted function:
// (1) _method: it is a method of a struct;
// (2) _func: it is a top-level function of a package.
type funcKind uint8

const (
	_method funcKind = iota
	_func
)

// trustedFuncSig defines the signature of a function that we "trust" to have a certain effect on its arguments, for example.
type trustedFuncSig struct {
	kind           funcKind
	enclosingRegex *regexp.Regexp
	funcNameRegex  *regexp.Regexp
}

// match checks if a given call expression matches with a trusted function's signature. Namely,
// it performs a strict matching for the function / method name and a user-defined regex match for
// the enclosing package or struct path.
func (t *trustedFuncSig) match(pass *analysishelper.EnhancedPass, call *ast.CallExpr) bool {
	_ = "STUB: not implemented"
	return false
}

// Match fully qualified path of the call expression with the expected path specified in `t`
// if function, match enclosing "<pkg path>". E.g., for `assert.Error(err)`, path = github.com/stretchr/testify/assert
// if method, match with "<pkg path>.<struct name>". E.g., for `u.Require().Error(err)`, path = github.com/stretchr/testify/require.Assertions

// return early if the kind of `t` and `funcObj` don't match. Both should be functions (or methods) for the match to be performed
// `recv != nil` implies `funcObj` is a method, while `recv == nil` means it is a function

// add struct name to the path

// we should likely never hit this case, but is only added for extra safety since
// `util.TypeAsDeeplyNamed` can return nil

// newNilBinaryExpr creates a new binary expression "expr op nil".
func newNilBinaryExpr(expr ast.Expr, op token.Token) *ast.BinaryExpr {
	_ = "STUB: not implemented"
	return nil
}
