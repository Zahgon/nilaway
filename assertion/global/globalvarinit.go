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

package global

import (
	"go/ast"

	"go.uber.org/nilaway/annotation"
	"go.uber.org/nilaway/util/analysishelper"
)

// analyzeValueSpec returns full triggers corresponding to the declaration
func analyzeValueSpec(pass *analysishelper.EnhancedPass, spec *ast.ValueSpec) []annotation.FullTrigger {
	_ = "STUB: not implemented"
	return nil
}

// Case: variables are not initialized
// All the variables in this case have same type

// Case: variables are initialized and the assignment is 1-1

// Case: variables are initialized using a multiple return function

// Returns a list of consumers corresponding to a global level variable declaration
func getGlobalConsumers(pass *analysishelper.EnhancedPass, valspec *ast.ValueSpec) []*annotation.ConsumeTrigger {
	_ = "STUB: not implemented"
	return nil
}

// Types that are not nilable are eliminated here

// Returns a producer in the cases: 1) func call 2) literal nil 3) another global var 4) struct field/method.
// In all other cases, it returns nil.
func getGlobalProducer(pass *analysishelper.EnhancedPass, valspec *ast.ValueSpec, lid int, rid int) *annotation.ProduceTrigger {
	_ = "STUB: not implemented"
	return nil
}

// We assume builtin functions do not return nil.

// Method call

// if rhs is literal nil

// if rhs is another global

// Struct field access

func getProducerForVar(pass *analysishelper.EnhancedPass, rhs *ast.Ident) *annotation.ProduceTrigger {
	_ = "STUB: not implemented"
	return nil
}

// If rhs is not a global variable (e.g., a constant), we ignore it.

func getProducerForField(pass *analysishelper.EnhancedPass, rhs *ast.Ident) *annotation.ProduceTrigger {
	_ = "STUB: not implemented"
	return nil
}

// If rhs is not a variable (e.g., a constant from an upstream package), we ignore it.

func getProducerForFuncCall(pass *analysishelper.EnhancedPass, methName *ast.Ident, lid int, rid int, rhs ast.Expr) *annotation.ProduceTrigger {
	_ = "STUB: not implemented"
	return nil
}

// We ignore if the method is anonymous

// We are interested in `lid-rid`-th return of the function
// In single return function this is `0` and in multiple return function it is `lid`

func getProducerForMethodCall(pass *analysishelper.EnhancedPass, methName *ast.Ident, lid int, rid int, rhs ast.Expr) *annotation.ProduceTrigger {
	_ = "STUB: not implemented"
	return nil
}

// We ignore if the method is anonymous

// We are interested in `lid-rid`-th return of the function
// In single return function this is `0` and in multiple return function it is `lid`
