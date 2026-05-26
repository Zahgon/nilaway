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

package functioncontracts

import (
	"golang.org/x/tools/go/ssa"
)

// _maxNumTablesPerBlock is the maximum number of nilnessTables that we will keep for each block.
// If the number of nilnessTables exceeds this number, we will stop analyzing the function and
// infer nothing, out of performance consideration.
const _maxNumTablesPerBlock = 1024

// inferContracts infers function contracts for a function if it has no contracts written. It
// returns a list of inferred contracts, which may be empty if no contract is inferred but is never
// nil.
func inferContracts(fn *ssa.Function) Contracts { _ = "STUB: not implemented"; return *new(Contracts) }

// TODO: Consider *ssa.Panic
// No need of an expensive dataflow analysis if we can derive contracts from the return
// instructions directly.

// Add the entry block to the queue.
// TODO: visit fn.Recover.

// seen[i] means we have visited block i at least once.

// Visit every block in the queue.

// Create a nilnessTableSet for this block if it doesn't exist.

// Propagate nilnessTables from predecessors. Union every predecessor's nilnessTableSet
// into this block's nilnessTableSet.

// Map each predecessor to the nilnessTables propagated from it.

// Skip any predecessor that has not been visited yet.

// Copy all the existing values.

// Learn new nilness from the branch if any.

// conflict found when extending this table, we should drop this table.

// Transfer nilness inside the block

// There is no table set up for this predecessor because the predecessor is
// skipped in the previous for loop, since it is not visited yet. We skip
// it here as well.

// Do not save the nilness if it is unknown. There are two cases:
// 1. the phi value cannot have nil as a valid value, e.g. it is an
// int.
// 2. the phi value can have nil as a valid value, but we do not know
// the nilness of the value.

// Append the nilness of the phi value and the related values' nilness
// expanded from the nilness of the phi value.

// TODO: function call, field addr, store, etc.

// Update nilnessTableSetByBB for this block.

// Only save the table if it is not empty.

// No need to visit the successors if the nilness table set of this block is not updated.

// TODO: nicely handle exponential explosion of tables.

// Too many tables, we should give up inferring contracts for this function.

// Add successors to queue since the nilness table set of this block has been updated.

// learnNilness learns nilness for the block succ, extended from one nilnessTable table of its
// predecessor pred. The function creates a new nilnessTable that stores **only newly learned**
// nilness for succ block and a bool flag that indicates whether a conflict is seen when going to
// succ block from pred block. If there is no a conflict, the bool flag is true, otherwise it is
// false. Note, when the bool flag is false, the returned nilnessTable can be empty because we do
// not learn nilness for al kinds of conditions. For now we support only what is supported in
// the function branch.
func learnNilness(succ *ssa.BasicBlock, pred *ssa.BasicBlock, table nilnessTable) (nilnessTable, bool) {
	_ = "STUB: not implemented"
	return *
	// learned nilnessTable
	new(nilnessTable), false
}

// TODO: for now we learn nilness from only these two condition: ? == ? and ? != ?

// TODO: only binOp == nil is sufficient, but nilaway would not be happy with it.

// TODO: each successor should learn a different nilness in terms of the variable
//  involved in the condition.

// Both operands are known of nilness.
// Determine whether the branch is reachable.

// conflict

// Only one operand is known of nilness.

// learn the nilness of Y

// learn the nilness of X

// deriveContracts checks nilness of parameter and return values at every exit block to infer
// contracts.
func deriveContracts(
	retInstrs []*ssa.Return,
	fn *ssa.Function,
	nilnessTableSetByBB map[*ssa.BasicBlock]nilnessTableSet,
) Contracts {
	_ = "STUB: not implemented"
	// TODO: verify other or multiple param/return contracts in the future; for now we consider
	//  contract(nonnil->nonnil) only.
	return *new(Contracts)
}

// We try to find a counterexample to nonnil->nonnil. If we find one, we immediately break out
// of the loop and return the contract(nonnil->nonnil) does not hold. Otherwise, we will
// move on to post-check before we can conclude the contaract(nonnil->nonnil) holds.

// b ends with a return

// All the possibilities:
// nonnil->nonnil     // OK
// unknown->nonnil    // OK
// nil->nonnil        // OK
// nonnil->unknown    // counterexample to contract(nonnil->nonnil)
// unknown->unknown   // two cases
//   unknown->unknown // counterexample in general
//   x->x             // OK if two unknown are the same value
// nil->unknown       // OK
// nonnil->nil        // counterexample to contract(nonnil->nonnil)
// unknown->nil       // counterexample to contract(nonnil->nonnil)
// nil->nil           // OK
//

// nil param never leads to a counterexample to contract(nonnil->nonnil)

// pNil == isnonnil or unknown, rNil can be anything, i.e. isnonnil, unknown, isnil.

// Absolutely OK if rNil == isnonnil
// The only OK case otherwise
// Those cases are not counterexamples to contract(nonnil->nonnil)

// All the remaining cases are counterexamples to contract(nonnil->nonnil)

// Post-check: we deny contract(nonnil->nonnil) for the following cases:
//
// 1. If the parameter is nil at every path, the contract nonnil->nonnil is trivially true, but
// we suppress inferring such a useless contract to avoid the following overhead, such as
// trigger duplication. However, I feel it is not possible that all paths have the parameter as
// isnil, since the parameter always starts with unknown, and we do only branching nil and
// nonnil.
//
// 2. It is essentially _->nonnil that holds for this function, so it is not necessary to
// infer nonnil->nonnil.

// totalChoices > nilParamChoices >= 0 && totalChoices >= nonnilOrUnknownParamChoices > 0 &&
// nonnilRetChoices < totalChoices

// nonnil->nonnil is valid at all exit blocks

func getReturnInstrs(fn *ssa.Function) []*ssa.Return { _ = "STUB: not implemented"; return nil }

// len(x) still returns 0 when x is nil

// b is an exit block, either panic or return

// branch reports whether we can make a branch and learn nilness from the true and false successor.
// The function returns equal, not-equal successor and the binary operation iff
//
//	(1) b ends with an equality or inequality comparison;
//	(2) the operands can have nil as a valid value;
//
// Otherwise, the function returns all nil.
func branch(b *ssa.BasicBlock) (*ssa.BasicBlock, *ssa.BasicBlock, *ssa.BinOp) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Check only one operand is sufficient since the two operands must have the same type.

// not a binary comparison or the type cannot have nil as a value.

// isBuiltinAppendCall reports if the call is a call to builtin append.
func isBuiltinAppendCall(v *ssa.Call) bool {
	_ = "STUB: not implemented"
	// TODO: consider merge this and assertion.BuiltinAppend
	return false
}

type nilness int

func (n nilness) negate() nilness { _ = "STUB: not implemented"; return *new(nilness) }

const (
	isnonnil         = -1
	unknown  nilness = 0
	isnil            = 1
)

var nilnessStrings = []string{"non-nil", "unknown", "nil"}

func (n nilness) String() string { _ = "STUB: not implemented"; return "" }

type nilnessTable map[ssa.Value]nilness

// nilnessOf reports whether v is definitely nil, definitely not nil, or unknown given the nilness
// table.
// Adapted from org_golang_x_tools/go/analysis/passes/nilness/nilness.go
func (t nilnessTable) nilnessOf(v ssa.Value) nilness {
	_ = "STUB: not implemented"
	return *

	// unwrap ChangeInterface and Slice values recursively, to detect if underlying values have any
	// facts recorded or are otherwise known with regard to nilness.
	//
	// This work must be in addition to expanding facts about ChangeInterfaces during
	// inference/fact gathering because this covers cases where the nilness of a value is
	// intrinsic, rather than based on inferred facts, such as a zero value interface variable.
	// That said, this work alone would only inform us when facts are about underlying values,
	// rather than outer values, when the analysis is transitive in both directions.
	new(nilness)
}

// Get the length of underlying array pointer of slice

// We know that *(*[1]byte)(nil) is going to panic because of the conversion. So
// return unknown to the caller, prevent useless nil deference reporting due to *
// operator.

// Otherwise, the conversion will yield a non-nil pointer to array. Note that the
// instruction can still panic if array length greater than slice length. If the value
// is used by another instruction, that instruction can assume the panic did not happen
// when that instruction is reached.

// In case array length is zero, the conversion result depends on nilness of the slice.

// append(s, x) always returns a nonnil.

// append(s) depends on the nilability of s.

// Is value intrinsically nil or non-nil?

// nil or zero value of a pointer-like type

// non-pointer

// Search table for the value.

// TODO: When we see a pointer indirection, we can check if the indirection of the same pointer
//  has been saved in the table. If so, we can use the saved value.

// expandNilness takes a single known nilness and learn the set of nilness that can be known
// about it or any of its related values. Some operations, like ChangeInterface, have transitive
// nilness, such that if you know the underlying value is nil, you also know the value itself is
// nil, and vice versa. This operation allows callers to match on any of the related values in
// analyses, rather than just the one form of the value that happened to appear in a comparison.
//
// This work must be in addition to unwrapping values within nilnessOf because while this work
// helps give facts about transitively known values based on inferred facts, the recursive check
// within nilnessOf covers cases where nilness facts are intrinsic to the underlying value, such as
// a zero value interface variables.
//
// ChangeInterface is the only expansion currently supported, but others, like Slice, could be
// added. At this time, this tool does not check slice operations in a way this expansion could
// help, for example:
//
// var s []string
//
//	if s0 := s[:0]; s0 == nil {
//	  fmt.Println("s0 and s are both nil.")
//	}
//
// Adapted from org_golang_x_tools/go/analysis/passes/nilness/nilness.go
func (t nilnessTable) expandNilness(val ssa.Value, nn nilness) { _ = "STUB: not implemented"; return }

// Don't reprocess.

// append(s) always has the same nilness as s.

// copy returns a copy of the nilnessTableSet.
func (t nilnessTable) copy() nilnessTable { _ = "STUB: not implemented"; return *new(nilnessTable) }

// equals returns if two nilnessTableSet are equal.
func (t nilnessTable) equals(other nilnessTable) bool { _ = "STUB: not implemented"; return false }

// addAll adds all the entries from other to t.
func (t nilnessTable) addAll(other nilnessTable) { _ = "STUB: not implemented"; return }

type nilnessTableSet []nilnessTable

func add(s nilnessTableSet, t nilnessTable) (nilnessTableSet, bool) {
	_ = "STUB: not implemented"
	// TODO: an efficient check for whether t is already in s.
	return *new(nilnessTableSet), false
}

func newNilnessTableSet() nilnessTableSet { _ = "STUB: not implemented"; return *new(nilnessTableSet) }
