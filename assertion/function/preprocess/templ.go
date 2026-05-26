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

package preprocess

import (
	"go/ast"

	"golang.org/x/tools/go/cfg"
)

const (
	_templPkgPath        = "github.com/a-h/templ"
	_templRuntimePkgPath = "github.com/a-h/templ/runtime"
)

// inlineTemplComponentFuncLit "inlines" the function literal that is used to create a templ component
// that is the first argument to templruntime.GeneratedTemplate in the return statement of
// a templ component function.
//
// A typical generated templ component function looks like this:
//
//	func MyComponent(...) templ.Component {
//	    return templruntime.GeneratedTemplate(func(templruntime.GeneratedComponentInput) error {
//	        // ... component logic ...
//	        return nil
//	    })
//	}
//
// Currently, NilAway is not able to analyze such functions since it involves passing a function
// literal to a function call (which is eventually invoked by the templ runtime). To aid NilAway's
// analysis for now, we "inline" the function literal by replacing the CFG of the function with the
// CFG of the function literal. Moreover, we replace the inner return statements inside the
// function literal with the real return statement, this helps NilAway to understand the return
// value is actually non-nil.
//
// Note that this is a temporary workaround until NilAway has better support for function literals
// in general.
// TODO: remove this once anonymous function support handles it naturally.
func (p *Preprocessor) inlineTemplComponentFuncLit(graph *cfg.CFG, funcDecl *ast.FuncDecl) {
	_ = "STUB: not implemented"
	return
}

// If the function is not a templ component function, we don't need to do anything.

// Now, we "inline" the function literal by replacing the CFG of the function with the CFG of
// the function literal.

// Replace the inner return statements inside the function literal with the real return
// statement, this helps NilAway to understand the return value is non-nil.

func (p *Preprocessor) extractTemplComponentFuncLit(funcDecl *ast.FuncDecl) (*ast.FuncLit, *ast.ReturnStmt) {
	_ = "STUB: not implemented"
	// Check if the function returns a single result of type `templ.Component`.
	return nil, nil
}

// Check if the function contains only a single return statement that calls
// `templruntime.GeneratedTemplate(func() { ... })`.

// Check if the first argument is a function literal.
