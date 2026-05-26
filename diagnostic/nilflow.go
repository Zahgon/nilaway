//	Copyright (c) 2023 Uber Technologies, Inc.
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

package diagnostic

import (
	"go/token"

	"go.uber.org/nilaway/annotation"
)

type nilFlow struct {
	nilPath    []node // stores nil path of the flow from nilable source to conflict point
	nonnilPath []node // stores non-nil path of the flow from conflict point to dereference point
}

// addNilPathNode adds a new node to the nil path.
func (n *nilFlow) addNilPathNode(p annotation.Prestring, c annotation.Prestring) {
	_ = "STUB: not implemented"
	return

	// Note that in the implication graph, we traverse backwards from the point of conflict to the source of nilability.
	// Therefore, they are added in reverse order from what the program flow would look like. To account for this we
	// prepend the new node to nilPath because we want to print the program flow in its correct (forward) order.
	// TODO: instead of prepending here, we can reverse the nilPath slice while printing.
}

// addNonNilPathNode adds a new node to the non-nil path
func (n *nilFlow) addNonNilPathNode(p annotation.Prestring, c annotation.Prestring) {
	_ = "STUB: not implemented"
	return
}

// String converts a nilFlow to a string representation, where each entry is the flow of the form: `<pos>: <reason>`
func (n *nilFlow) String() string { _ = "STUB: not implemented"; return "" }

// involvesTestFile returns true if any node position in the nil or non-nil path originates from
// a test file (i.e., a file ending with "_test.go").
func (n *nilFlow) involvesTestFile() bool { _ = "STUB: not implemented"; return false }

type node struct {
	producerPosition token.Position
	consumerPosition token.Position
	producerRepr     string
	consumerRepr     string
}

// newNode creates a new node object from the given producer and consumer Prestrings.
// LocatedPrestring contains accurate information about the position and the reason why NilAway deemed that position
// to be nilable. We use it if available, else we use the raw string representation available from the Prestring.
func newNode(p annotation.Prestring, c annotation.Prestring) node {
	_ = "STUB: not implemented"

	// get producer representation string
	return *new(node)
}

// get consumer representation string

func (n *node) String() string { _ = "STUB: not implemented"; return "" }

func pathString(nodes []node) string { _ = "STUB: not implemented"; return "" }
