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

package diagnostic

import (
	"go/token"

	"go.uber.org/nilaway/util/analysishelper"
)

type conflict struct {
	// position is the package-independent position where the conflict should be reported.
	position token.Position
	// flow stores nil flow from source to dereference point
	flow nilFlow
	// similarConflicts stores other conflicts that are similar to this one.
	similarConflicts []*conflict
}

func (c *conflict) String() string {
	_ = "STUB: not implemented"
	// build string for similar conflicts (i.e., conflicts with the same nil path)
	return ""
}

func (c *conflict) addSimilarConflict(conflict conflict) { _ = "STUB: not implemented"; return }

// groupConflicts groups conflicts with the same nil path together and update conflicts list.
func groupConflicts(allConflicts []conflict, pass *analysishelper.EnhancedPass) []conflict {
	_ = "STUB: not implemented"
	return nil
}

// key: nil path string, value: index in `allConflicts`
// indices of conflicts to be ignored from `allConflicts`, since they are grouped with other conflicts

// Handle the case of single assertion conflict separately

// This is the case of single assertion conflict. Use producer position and repr from the non-nil path as
// the key, if present, else use the producer and consumer repr as a heuristic key to group conflicts.

// The heuristic of using producer and consumer repr as key may not work perfectly, especially when the
// error messages in two different functions are exactly the same. Consider the following example:
// ```
// 	func f1() {
//		mp := make(map[int]*int)
//		_ = *mp[0] // error message: "deep read from local variable `mp` lacking guarding; dereferenced"
// 	}
//
// 	func f2() {
//		mp := make(map[int]*int)
//		_ = *mp[0] // error message: "deep read from local variable `mp` lacking guarding; dereferenced"
// 	}
// ```
// Here, the two error messages are exactly the same, but they should not be grouped together as they are
// from different functions. To handle such cases, we prepend the enclosing function name to the key.

// Check if the file is in scope and the conflict position is in the same file

// Check if the conflict position falls within the function's position range. If so, update the key to
// include the function name, and end the traversal.

// Grouping condition satisfied. Add new conflict to `similarConflicts` in `existingConflict`, and update groupedConflicts map

// update groupedConflicts list with grouped groupedConflicts
