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

package inference

import (
	"fmt"
	"go/token"
)

// An ExplainedBool is a boolean value, wrapped by a "reason" that we came to the conclusion it should
// have that value. ExplainedBools are used as labels on annotation sites once their state (true for
// nilable, false for nonnil) is established, with the explanations serving primarily to generate
// error messages in the case that conflicting labels are required for a given site. The currently used
// explanations are:
// - <Val>BecauseShallowConstraint: Applied to site X when X was half of an assertion where the other half was fixed as a definite site of nil production or nonnil consumption
// - <Val>BecauseDeepConstraint: Applied to site X when X was half of an assertion where the other half was fixed, but through a deeper chain of assertions
// - <Val>BecauseAnnotation: Applied to site X when a syntactic annotation was discovered on X
type ExplainedBool interface {
	fmt.Stringer

	Val() bool
	Position() token.Position
	TriggerReprs() (producer fmt.Stringer, consumer fmt.Stringer)
	DeeperReason() ExplainedBool
}

// ExplainedTrue is a common embedding in all instances of ExplainedBool that wrap the value `true`
type ExplainedTrue struct{}

// Val for an ExplainedTrue always returns `true` (this is the point of an ExplainedTrue)
func (ExplainedTrue) Val() bool {
	_ = "STUB: not implemented"

	// ExplainedFalse is a common embedding in all instances of ExplainedBool that wrap the value `false`
	return false
}

type ExplainedFalse struct{}

// Val for an ExplainedFalse always returns `false` (this is the point of an ExplainedFalse)
func (ExplainedFalse) Val() bool {
	_ = "STUB: not implemented"

	// TrueBecauseShallowConstraint is used as the label for site Y when an assertion of the form
	// `nilable X -> nilable Y` is discovered and the trigger for `nilable X` always fires (i.e. yields
	// nilable) - for example because it is the literal nil or an unguarded map read. In all cases, this
	// constrains the site Y to be nilable, so we label Y with `ExplainedTrue` as a
	// `TrueBecauseShallowConstraint`, wrapped along with the assertion that we discovered to yield the
	// truth.
	return false
}

type TrueBecauseShallowConstraint struct {
	ExplainedTrue
	ExternalAssertion primitiveFullTrigger
}

func (t TrueBecauseShallowConstraint) String() string { _ = "STUB: not implemented"; return "" }

// Position is the position of underlying site.
func (t TrueBecauseShallowConstraint) Position() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

// TriggerReprs returns the compact representation structs for the producer and consumer.
func (t TrueBecauseShallowConstraint) TriggerReprs() (fmt.Stringer, fmt.Stringer) {
	_ = "STUB: not implemented"
	return *new(fmt.Stringer), *new(fmt.Stringer)
}

// DeeperReason returns another ExplainedBool that marks the deeper reason of this constraint.
// It is only nonnil for deep constraints.
func (t TrueBecauseShallowConstraint) DeeperReason() ExplainedBool {
	_ = "STUB: not implemented"

	// FalseBecauseShallowConstraint is used as the label for site X when an assertion of the form
	// `nilable X -> nilable Y` is discovered and the trigger for `nilable Y` always fires (i.e. yields
	// nonnil) - for example because it is the dereferenced as a pointer or passed to a field access. In
	// all cases, this constrains the site X to be nonnil, so we label X with `ExplainedFalse` as a
	// `FalseBecauseShallowConstraint`, wrapped along with the assertion that we discovered to yield the
	// falsehood.
	return *new(ExplainedBool)
}

type FalseBecauseShallowConstraint struct {
	ExplainedFalse
	ExternalAssertion primitiveFullTrigger
}

func (f FalseBecauseShallowConstraint) String() string { _ = "STUB: not implemented"; return "" }

// Position is the position of underlying site.
func (f FalseBecauseShallowConstraint) Position() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

// TriggerReprs returns the compact representation structs for the producer and consumer.
func (f FalseBecauseShallowConstraint) TriggerReprs() (fmt.Stringer, fmt.Stringer) {
	_ = "STUB: not implemented"
	return *new(fmt.Stringer), *new(fmt.Stringer)
}

// DeeperReason returns another ExplainedBool that marks the deeper reason of this constraint.
// It is only nonnil for deep constraints.
func (f FalseBecauseShallowConstraint) DeeperReason() ExplainedBool {
	_ = "STUB: not implemented"

	// TrueBecauseDeepConstraint is used as the label for a site Y when an assertion of the form
	// `nilable X -> nilable Y` is discovered along with some reason for X to be nilable, besides it
	// necessarily being so because it always fires. This reason could be any ExplainedTrue - such as
	// `TrueBecauseAnnotation`, `TrueBecauseShallowConstraint`, or another `TrueBecauseDeepConstraint`.
	return *new(ExplainedBool)
}

type TrueBecauseDeepConstraint struct {
	ExplainedTrue
	InternalAssertion primitiveFullTrigger
	DeeperExplanation ExplainedBool
}

func (t TrueBecauseDeepConstraint) String() string { _ = "STUB: not implemented"; return "" }

// Position is the position of underlying site.
func (t TrueBecauseDeepConstraint) Position() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

// TriggerReprs returns the compact representation structs for the producer and consumer.
func (t TrueBecauseDeepConstraint) TriggerReprs() (fmt.Stringer, fmt.Stringer) {
	_ = "STUB: not implemented"
	return *new(fmt.Stringer), *new(fmt.Stringer)
}

// DeeperReason returns another ExplainedBool that marks the deeper reason of this constraint.
// It is only nonnil for deep constraints.
func (t TrueBecauseDeepConstraint) DeeperReason() ExplainedBool {
	_ = "STUB: not implemented"
	return *new(ExplainedBool)
}

// FalseBecauseDeepConstraint is used as the label for a site X when an assertion of the form
// `nilable X -> nilable Y` is discovered along with some reason for Y to be nonnil, besides it
// necessarily being so because it always fires. This reason could be any ExplainedFalse - such as
// `FalseBecauseAnnotation`, `FalseBecauseShallowConstraint`, or another `FalseBecauseDeepConstraint`.
type FalseBecauseDeepConstraint struct {
	ExplainedFalse
	InternalAssertion primitiveFullTrigger
	DeeperExplanation ExplainedBool
}

func (f FalseBecauseDeepConstraint) String() string { _ = "STUB: not implemented"; return "" }

// Position is the position of underlying site.
func (f FalseBecauseDeepConstraint) Position() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

// TriggerReprs returns the compact representation structs for the producer and consumer.
func (f FalseBecauseDeepConstraint) TriggerReprs() (fmt.Stringer, fmt.Stringer) {
	_ = "STUB: not implemented"
	return *new(fmt.Stringer), *new(fmt.Stringer)
}

// DeeperReason returns another ExplainedBool that marks the deeper reason of this constraint.
// It is only nonnil for deep constraints.
func (f FalseBecauseDeepConstraint) DeeperReason() ExplainedBool {
	_ = "STUB: not implemented"
	return *new(ExplainedBool)
}

// TrueBecauseAnnotation is used as the label for a site X on which a literal annotation "//nilable(x)"
// has been discovered - forcing that site to be nilable.
type TrueBecauseAnnotation struct {
	ExplainedTrue
	AnnotationPos token.Position
}

func (TrueBecauseAnnotation) String() string { _ = "STUB: not implemented"; return "" }

// Position is the position of underlying site.
func (t TrueBecauseAnnotation) Position() token.Position {
	_ = "STUB: not implemented"
	return *

	// TriggerReprs simply returns nil, nil since this constraint is the result of an annotation.
	new(token.Position)
}

func (TrueBecauseAnnotation) TriggerReprs() (fmt.Stringer, fmt.Stringer) {
	_ = "STUB: not implemented"

	// DeeperReason returns another ExplainedBool that marks the deeper reason of this constraint.
	// It is only nonnil for deep constraints.
	return *new(fmt.Stringer), *new(fmt.Stringer)
}

func (TrueBecauseAnnotation) DeeperReason() ExplainedBool {
	_ = "STUB: not implemented"

	// FalseBecauseAnnotation is used as the label for a site X on which a literal annotation "//nonnil(x)"
	// has been discovered - forcing that site to be nonnil.
	return *new(ExplainedBool)
}

type FalseBecauseAnnotation struct {
	ExplainedFalse
	AnnotationPos token.Position
}

func (FalseBecauseAnnotation) String() string { _ = "STUB: not implemented"; return "" }

// Position is the position of underlying site.
func (f FalseBecauseAnnotation) Position() token.Position {
	_ = "STUB: not implemented"
	return *

	// TriggerReprs simply returns nil, nil since this constraint is the result of an annotation.
	new(token.Position)
}

func (FalseBecauseAnnotation) TriggerReprs() (fmt.Stringer, fmt.Stringer) {
	_ = "STUB: not implemented"

	// DeeperReason returns another ExplainedBool that marks the deeper reason of this constraint.
	// It is only nonnil for deep constraints.
	return *new(fmt.Stringer), *new(fmt.Stringer)
}

func (f FalseBecauseAnnotation) DeeperReason() ExplainedBool {
	_ = "STUB: not implemented"
	return *new(ExplainedBool)
}
