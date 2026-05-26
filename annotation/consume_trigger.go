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

package annotation

import (
	"go/ast"
	"go/token"
	"go/types"

	"go.uber.org/nilaway/guard"
	"go.uber.org/nilaway/util/orderedmap"
)

// A ConsumingAnnotationTrigger indicated a possible reason that a nil flow to this site would indicate
// an error
//
// All ConsumingAnnotationTriggers must embed one of the following 3 structs:
// -TriggerIfNonnil
// -TriggerIfDeepNonnil
// -ConsumeTriggerTautology
type ConsumingAnnotationTrigger interface {
	// CheckConsume can be called to determined whether this trigger should be triggered
	// given a particular Annotation map
	// for example - an `ArgPass` trigger triggers iff the corresponding function arg has
	// nonNil type
	CheckConsume(Map) bool
	Prestring() Prestring

	// Kind returns the kind of the trigger.
	Kind() TriggerKind

	// UnderlyingSite returns the underlying site this trigger's nilability depends on. If the
	// trigger always or never fires, the site is nil.
	UnderlyingSite() Key

	customPos() (token.Pos, bool)

	// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
	equals(ConsumingAnnotationTrigger) bool

	// Copy returns a deep copy of this ConsumingAnnotationTrigger
	Copy() ConsumingAnnotationTrigger

	// AddAssignment adds an assignment to the trigger for tracking and printing informative error message.
	// NilAway's `backpropAcrossOneToOneAssignment()` lifts consumer triggers from the RHS of an assignment to the LHS.
	// This implies loss of information about the assignment. This method is used to track such assignments and print
	// a more informative error message.
	AddAssignment(Assignment)

	// NeedsGuard returns true if the trigger needs to be guarded, for example, by a nil check or an ok form.
	NeedsGuard() bool

	// SetNeedsGuard sets the underlying Guard-Neediness of this ConsumerTrigger, if present.
	// Default setting for ConsumerTriggers is that they need a guard. Override this method to set the need for a guard to false.
	SetNeedsGuard(bool)
}

// Prestring is an interface used to encode objects that have compact on-the-wire encodings
// (via gob) but can still be expanded into verbose string representations on demand using
// type information. These are key for compact encoding of InferredAnnotationMaps
type Prestring interface {
	String() string
}

// Assignment is a struct that represents an assignment to an expression
type Assignment struct {
	LHSExprStr string
	RHSExprStr string
	Position   token.Position
}

func (a *Assignment) String() string { _ = "STUB: not implemented"; return "" }

// assignmentFlow is a struct that represents a flow of assignments.
// Note that we implement a copy method for this struct, since we want to deep copy the assignments map when we copy
// ConsumerTriggers. However, we don't implement an `equals` method for this struct, since it would incur a performance
// penalty in situations where multiple nilable flows reach a dereference site by creating more full triggers and possibly
// more rounds through backpropagation fix point. Consider the following example:
//
//	func f(m map[int]*int) {
//	  var v *int
//	  var ok1, ok2 bool
//	  if cond {
//	    v, ok1 = m[0] // nilable flow 1, ok1 is false
//	  } else {
//	    v, ok2 = m[1] // nilable flow 2, ok2 is false
//	  }
//	  _, _ = ok1, ok2
//	  _ = *v // nil panic!
//	}
//
// Here `v` can be potentiall nilable from two flows: ok1 or ok2 is false. We would like to print only one error message
// for this situation with one representative flow printed in the error message. However, with an `equals` method, we would
// report multiple error messages, one for each flow, by creating multiple full triggers, thereby affecting performance.
type assignmentFlow struct {
	// We use ordered map for `assignments` to maintain the order of assignments in the flow, and also to avoid
	// duplicates that can get introduced due to fix point convergence in backpropagation.
	assignments *orderedmap.OrderedMap[Assignment, bool]
}

func (a *assignmentFlow) addEntry(entry Assignment) { _ = "STUB: not implemented"; return }

func (a *assignmentFlow) copy() assignmentFlow {
	_ = "STUB: not implemented"
	return *new(assignmentFlow)
}

func (a *assignmentFlow) String() string { _ = "STUB: not implemented"; return "" }

// backprop algorithm populates assignment entries in backward order. Reverse entries to get forward order of
// assignments, and store in `strs` slice.

// build the informative print string tracking the assignments

// TriggerIfNonNil is triggered if the contained Annotation is non-nil
type TriggerIfNonNil struct {
	Ann              Key
	IsGuardNotNeeded bool // ConsumeTriggers need guards by default, when applicable. Set this to true when guards are not needed.
	assignmentFlow
}

// Kind returns Conditional.
func (*TriggerIfNonNil) Kind() TriggerKind {
	_ = "STUB: not implemented"

	// UnderlyingSite the underlying site this trigger's nilability depends on.
	return *new(TriggerKind)
}

func (t *TriggerIfNonNil) UnderlyingSite() Key {
	_ = "STUB: not implemented"

	// CheckConsume returns true if the underlying annotation is present in the passed map and nonnil
	return *new(Key)
}

func (t *TriggerIfNonNil) CheckConsume(annMap Map) bool { _ = "STUB: not implemented"; return false }

// customPos has the below default implementation for TriggerIfNonNil, in which case ConsumeTrigger.Pos() will return a default value.
// To return non-default position values, this method should be overridden appropriately.
func (*TriggerIfNonNil) customPos() (token.Pos, bool) {
	_ = "STUB: not implemented"
	return *

	// NeedsGuard is the default implementation for TriggerIfNonNil. To return non-default value, this method should be overridden.
	new(token.Pos), false
}

func (t *TriggerIfNonNil) NeedsGuard() bool { _ = "STUB: not implemented"; return false }

// SetNeedsGuard sets the underlying Guard-Neediness of this ConsumerTrigger
func (t *TriggerIfNonNil) SetNeedsGuard(b bool) { _ = "STUB: not implemented"; return }

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (t *TriggerIfNonNil) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (t *TriggerIfNonNil) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// AddAssignment adds an assignment to the trigger.
func (t *TriggerIfNonNil) AddAssignment(e Assignment) {
	_ = "STUB: not implemented"

	// Prestring returns this Prestring as a Prestring
	return
}

func (t *TriggerIfNonNil) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// TriggerIfNonNilPrestring is a Prestring storing the needed information to compactly encode a TriggerIfNonNil
type TriggerIfNonNilPrestring struct {
	AssignmentStr string
}

func (t TriggerIfNonNilPrestring) String() string { _ = "STUB: not implemented"; return "" }

// TriggerIfDeepNonNil is triggered if the contained Annotation is deeply non-nil
type TriggerIfDeepNonNil struct {
	Ann              Key
	IsGuardNotNeeded bool // ConsumeTriggers need guards by default, when applicable. Set this to true when guards are not needed.
	assignmentFlow
}

// Kind returns DeepConditional.
func (*TriggerIfDeepNonNil) Kind() TriggerKind {
	_ = "STUB: not implemented"
	return *

	// UnderlyingSite the underlying site this trigger's nilability depends on.
	new(TriggerKind)
}

func (t *TriggerIfDeepNonNil) UnderlyingSite() Key {
	_ = "STUB: not implemented"

	// CheckConsume returns true if the underlying annotation is present in the passed map and deeply nonnil
	return *new(Key)
}

func (t *TriggerIfDeepNonNil) CheckConsume(annMap Map) bool {
	_ = "STUB: not implemented"
	return false
}

// customPos has the below default implementation for TriggerIfDeepNonNil, in which case ConsumeTrigger.Pos() will return a default value.
// To return non-default position values, this method should be overridden appropriately.
func (*TriggerIfDeepNonNil) customPos() (token.Pos, bool) {
	_ = "STUB: not implemented"
	return *

	// NeedsGuard default implementation for TriggerIfDeepNonNil. To return non-default value, this method should be overridden.
	new(token.Pos), false
}

func (t *TriggerIfDeepNonNil) NeedsGuard() bool { _ = "STUB: not implemented"; return false }

// SetNeedsGuard sets the underlying Guard-Neediness of this ConsumerTrigger
func (t *TriggerIfDeepNonNil) SetNeedsGuard(b bool) { _ = "STUB: not implemented"; return }

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (t *TriggerIfDeepNonNil) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (t *TriggerIfDeepNonNil) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// AddAssignment adds an assignment to the trigger.
func (t *TriggerIfDeepNonNil) AddAssignment(e Assignment) {
	_ = "STUB: not implemented"

	// Prestring returns this Prestring as a Prestring
	return
}

func (t *TriggerIfDeepNonNil) Prestring() Prestring {
	_ = "STUB: not implemented"
	return *new(Prestring)
}

// TriggerIfDeepNonNilPrestring is a Prestring storing the needed information to compactly encode a TriggerIfDeepNonNil
type TriggerIfDeepNonNilPrestring struct {
	AssignmentStr string
}

func (t TriggerIfDeepNonNilPrestring) String() string { _ = "STUB: not implemented"; return "" }

// ConsumeTriggerTautology is used at consumption sites were consuming nil is always an error
type ConsumeTriggerTautology struct {
	IsGuardNotNeeded bool // ConsumeTriggers need guards by default, when applicable. Set this to true when guards are not needed.
	assignmentFlow
}

// Kind returns Always.
func (*ConsumeTriggerTautology) Kind() TriggerKind {
	_ = "STUB: not implemented"

	// UnderlyingSite always returns nil.
	return *new(TriggerKind)
}

func (*ConsumeTriggerTautology) UnderlyingSite() Key {
	_ = "STUB: not implemented"

	// CheckConsume returns true
	return *new(Key)
}

func (*ConsumeTriggerTautology) CheckConsume(Map) bool {
	_ = "STUB: not implemented"

	// customPos has the below default implementation for ConsumeTriggerTautology, in which case ConsumeTrigger.Pos() will return a default value.
	// To return non-default position values, this method should be overridden appropriately.
	return false
}

func (*ConsumeTriggerTautology) customPos() (token.Pos, bool) {
	_ = "STUB: not implemented"
	return *

	// NeedsGuard default implementation for ConsumeTriggerTautology. To return non-default value, this method should be overridden.
	new(token.Pos), false
}

func (c *ConsumeTriggerTautology) NeedsGuard() bool { _ = "STUB: not implemented"; return false }

// SetNeedsGuard sets the underlying Guard-Neediness of this ConsumerTrigger
func (c *ConsumeTriggerTautology) SetNeedsGuard(b bool) { _ = "STUB: not implemented"; return }

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (c *ConsumeTriggerTautology) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (c *ConsumeTriggerTautology) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// AddAssignment adds an assignment to the trigger.
func (c *ConsumeTriggerTautology) AddAssignment(e Assignment) {
	_ = "STUB: not implemented"

	// Prestring returns this Prestring as a Prestring
	return
}

func (c *ConsumeTriggerTautology) Prestring() Prestring {
	_ = "STUB: not implemented"
	return *new(Prestring)
}

// ConsumeTriggerTautologyPrestring is a Prestring storing the needed information to compactly encode a ConsumeTriggerTautology
type ConsumeTriggerTautologyPrestring struct {
	AssignmentStr string
}

func (c ConsumeTriggerTautologyPrestring) String() string { _ = "STUB: not implemented"; return "" }

// PtrLoad is when a value flows to a point where it is loaded as a pointer
type PtrLoad struct {
	*ConsumeTriggerTautology
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (p *PtrLoad) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (p *PtrLoad) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this PtrLoad as a Prestring
func (p *PtrLoad) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// PtrLoadPrestring is a Prestring storing the needed information to compactly encode a PtrLoad
type PtrLoadPrestring struct {
	AssignmentStr string
}

func (p PtrLoadPrestring) String() string { _ = "STUB: not implemented"; return "" }

// MapAccess is when a map value flows to a point where it is indexed, and thus must be non-nil
//
// note: this trigger is produced only if config.ErrorOnNilableMapRead == true
type MapAccess struct {
	*ConsumeTriggerTautology
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (i *MapAccess) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (i *MapAccess) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this MapAccess as a Prestring
func (i *MapAccess) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// MapAccessPrestring is a Prestring storing the needed information to compactly encode a MapAccess
type MapAccessPrestring struct {
	AssignmentStr string
}

func (i MapAccessPrestring) String() string { _ = "STUB: not implemented"; return "" }

// MapWrittenTo is when a map value flows to a point where one of its indices is written to, and thus
// must be non-nil
type MapWrittenTo struct {
	*ConsumeTriggerTautology
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (m *MapWrittenTo) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (m *MapWrittenTo) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this MapWrittenTo as a Prestring
func (m *MapWrittenTo) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// MapWrittenToPrestring is a Prestring storing the needed information to compactly encode a MapWrittenTo
type MapWrittenToPrestring struct {
	AssignmentStr string
}

func (m MapWrittenToPrestring) String() string { _ = "STUB: not implemented"; return "" }

// SliceAccess is when a slice value flows to a point where it is sliced, and thus must be non-nil
type SliceAccess struct {
	*ConsumeTriggerTautology
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (s *SliceAccess) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (s *SliceAccess) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this SliceAccess as a Prestring
func (s *SliceAccess) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// SliceAccessPrestring is a Prestring storing the needed information to compactly encode a SliceAccess
type SliceAccessPrestring struct {
	AssignmentStr string
}

func (s SliceAccessPrestring) String() string { _ = "STUB: not implemented"; return "" }

// FldAccess is when a value flows to a point where a field of it is accessed, and so it must be non-nil
type FldAccess struct {
	*ConsumeTriggerTautology

	Sel types.Object
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (f *FldAccess) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (f *FldAccess) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this FldAccess as a Prestring
func (f *FldAccess) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// FldAccessPrestring is a Prestring storing the needed information to compactly encode a FldAccess
type FldAccessPrestring struct {
	FieldName     string
	MethodName    string
	AssignmentStr string
}

func (f FldAccessPrestring) String() string { _ = "STUB: not implemented"; return "" }

// UseAsErrorResult is when a value flows to the error result of a function, where it is expected to be non-nil
type UseAsErrorResult struct {
	*TriggerIfNonNil

	RetStmt       *ast.ReturnStmt
	IsNamedReturn bool
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (u *UseAsErrorResult) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (u *UseAsErrorResult) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this UseAsErrorResult as a Prestring
func (u *UseAsErrorResult) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// UseAsErrorResultPrestring is a Prestring storing the needed information to compactly encode a UseAsErrorResult
type UseAsErrorResultPrestring struct {
	Pos              int
	ReturningFuncStr string
	IsNamedReturn    bool
	RetName          string
	AssignmentStr    string
}

func (u UseAsErrorResultPrestring) String() string { _ = "STUB: not implemented"; return "" }

// overriding position value to point to the raw return statement, which is the source of the potential error
func (u *UseAsErrorResult) customPos() (token.Pos, bool) {
	_ = "STUB: not implemented"
	return *new(token.Pos), false
}

// FldAssign is when a value flows to a point where it is assigned into a field
type FldAssign struct {
	*TriggerIfNonNil
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (f *FldAssign) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (f *FldAssign) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this FldAssign as a Prestring
func (f *FldAssign) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// FldAssignPrestring is a Prestring storing the needed information to compactly encode a FldAssign
type FldAssignPrestring struct {
	FieldName     string
	AssignmentStr string
}

func (f FldAssignPrestring) String() string { _ = "STUB: not implemented"; return "" }

// ArgFldPass is when a struct field value (A.f) flows to a point where it is passed to a function with a param of
// the same struct type (A)
type ArgFldPass struct {
	*TriggerIfNonNil
	IsPassed bool
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (f *ArgFldPass) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (f *ArgFldPass) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this ArgFldPass as a Prestring
func (f *ArgFldPass) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// ArgFldPassPrestring is a Prestring storing the needed information to compactly encode a ArgFldPass
type ArgFldPassPrestring struct {
	FieldName     string
	FuncName      string
	ParamNum      int
	RecvName      string
	IsPassed      bool
	AssignmentStr string
}

func (f ArgFldPassPrestring) String() string { _ = "STUB: not implemented"; return "" }

// GlobalVarAssign is when a value flows to a point where it is assigned into a global variable
type GlobalVarAssign struct {
	*TriggerIfNonNil
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (g *GlobalVarAssign) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (g *GlobalVarAssign) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this GlobalVarAssign as a Prestring
func (g *GlobalVarAssign) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// GlobalVarAssignPrestring is a Prestring storing the needed information to compactly encode a GlobalVarAssign
type GlobalVarAssignPrestring struct {
	VarName       string
	AssignmentStr string
}

func (g GlobalVarAssignPrestring) String() string { _ = "STUB: not implemented"; return "" }

// ArgPass is when a value flows to a point where it is passed as an argument to a function. This
// consumer trigger can be used on top of two different sites: ParamAnnotationKey &
// CallSiteParamAnnotationKey. ParamAnnotationKey is the parameter site in the function
// declaration; CallSiteParamAnnotationKey is the argument site in the call expression.
// CallSiteParamAnnotationKey is specifically used for functions with contracts since we need to
// duplicate the sites for context sensitivity.
type ArgPass struct {
	*TriggerIfNonNil
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (a *ArgPass) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (a *ArgPass) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this ArgPass as a Prestring
func (a *ArgPass) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// ArgPassPrestring is a Prestring storing the needed information to compactly encode a ArgPass
type ArgPassPrestring struct {
	ParamName string
	FuncName  string
	// Location points to the code location of the argument pass at the call site for a ArgPass
	// enclosing CallSiteParamAnnotationKey; Location is empty for a ArgPass enclosing ParamAnnotationKey.
	Location      string
	AssignmentStr string
}

func (a ArgPassPrestring) String() string { _ = "STUB: not implemented"; return "" }

// ArgPassDeep is when a value deeply flows to a point where it is passed as an argument to a function
type ArgPassDeep struct {
	*TriggerIfDeepNonNil
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (a *ArgPassDeep) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (a *ArgPassDeep) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this ArgPassDeep as a Prestring
func (a *ArgPassDeep) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// ArgPassDeepPrestring is a Prestring storing the needed information to compactly encode a ArgPassDeep
type ArgPassDeepPrestring struct {
	ParamName string
	FuncName  string
	// Location points to the code location of the argument pass at the call site for a ArgPass
	// enclosing CallSiteParamAnnotationKey; Location is empty for a ArgPass enclosing ParamAnnotationKey.
	Location      string
	AssignmentStr string
}

func (a ArgPassDeepPrestring) String() string { _ = "STUB: not implemented"; return "" }

// RecvPass is when a receiver value flows to a point where it is used to invoke a method.
// E.g., `s.foo()`, here `s` is a receiver and forms the RecvPass Consumer
type RecvPass struct {
	*TriggerIfNonNil
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (a *RecvPass) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (a *RecvPass) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this RecvPass as a Prestring
func (a *RecvPass) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// RecvPassPrestring is a Prestring storing the needed information to compactly encode a RecvPass
type RecvPassPrestring struct {
	FuncName      string
	AssignmentStr string
}

func (a RecvPassPrestring) String() string { _ = "STUB: not implemented"; return "" }

// InterfaceResultFromImplementation is when a result is determined to flow from a concrete method to an interface method via implementation
type InterfaceResultFromImplementation struct {
	*TriggerIfNonNil
	AffiliationPair
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (i *InterfaceResultFromImplementation) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (i *InterfaceResultFromImplementation) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this InterfaceResultFromImplementation as a Prestring
func (i *InterfaceResultFromImplementation) Prestring() Prestring {
	_ = "STUB: not implemented"
	return *new(Prestring)
}

// InterfaceResultFromImplementationPrestring is a Prestring storing the needed information to compactly encode a InterfaceResultFromImplementation
type InterfaceResultFromImplementationPrestring struct {
	RetNum        int
	IntName       string
	ImplName      string
	AssignmentStr string
}

func (i InterfaceResultFromImplementationPrestring) String() string {
	_ = "STUB: not implemented"
	return ""
}

// MethodParamFromInterface is when a param flows from an interface method to a concrete method via implementation
type MethodParamFromInterface struct {
	*TriggerIfNonNil
	AffiliationPair
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (m *MethodParamFromInterface) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (m *MethodParamFromInterface) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this MethodParamFromInterface as a Prestring
func (m *MethodParamFromInterface) Prestring() Prestring {
	_ = "STUB: not implemented"
	return *new(Prestring)
}

// MethodParamFromInterfacePrestring is a Prestring storing the needed information to compactly encode a MethodParamFromInterface
type MethodParamFromInterfacePrestring struct {
	ParamName     string
	ImplName      string
	IntName       string
	AssignmentStr string
}

func (m MethodParamFromInterfacePrestring) String() string { _ = "STUB: not implemented"; return "" }

// DuplicateReturnConsumer duplicates a given consume trigger, assuming the given consumer trigger
// is for a UseAsReturn annotation.
func DuplicateReturnConsumer(t *ConsumeTrigger, location token.Position) *ConsumeTrigger {
	_ = "STUB: not implemented"
	return nil
}

// TODO: probably, we might not need a deep copy all the time

// UseAsReturn is when a value flows to a point where it is returned from a function.
// This consumer trigger can be used on top of two different sites: RetAnnotationKey &
// CallSiteRetAnnotationKey. RetAnnotationKey is the parameter site in the function declaration;
// CallSiteRetAnnotationKey is the argument site in the call expression. CallSiteRetAnnotationKey is specifically
// used for functions with contracts since we need to duplicate the sites for context sensitivity.
type UseAsReturn struct {
	*TriggerIfNonNil
	IsNamedReturn        bool
	IsTrackingAlwaysSafe bool
	RetStmt              *ast.ReturnStmt
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (u *UseAsReturn) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (u *UseAsReturn) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this UseAsReturn as a Prestring
func (u *UseAsReturn) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// UseAsReturnPrestring is a Prestring storing the needed information to compactly encode a UseAsReturn
type UseAsReturnPrestring struct {
	FuncName      string
	RetNum        int
	IsNamedReturn bool
	RetName       string
	// Location is empty for a UseAsReturn enclosing RetAnnotationKey. Location points to the
	// location of the result at the call site for a UseAsReturn enclosing
	// CallSiteRetAnnotationKey.
	Location      string
	AssignmentStr string
}

func (u UseAsReturnPrestring) String() string { _ = "STUB: not implemented"; return "" }

// overriding position value to point to the raw return statement, which is the source of the potential error
func (u *UseAsReturn) customPos() (token.Pos, bool) {
	_ = "STUB: not implemented"
	return *new(token.Pos), false
}

// UseAsReturnDeep is when a deep value flows to a point where it is returned from a function.
type UseAsReturnDeep struct {
	*TriggerIfDeepNonNil
	IsNamedReturn bool
	RetStmt       *ast.ReturnStmt
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (u *UseAsReturnDeep) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (u *UseAsReturnDeep) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this UseAsReturn as a Prestring
func (u *UseAsReturnDeep) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// UseAsReturnDeepPrestring is a Prestring storing the needed information to compactly encode a UseAsReturnDeep
type UseAsReturnDeepPrestring struct {
	FuncName      string
	RetNum        int
	RetName       string
	AssignmentStr string
}

func (u UseAsReturnDeepPrestring) String() string { _ = "STUB: not implemented"; return "" }

// overriding position value to point to the raw return statement, which is the source of the potential error
func (u UseAsReturnDeep) customPos() (token.Pos, bool) {
	_ = "STUB: not implemented"
	return *new(token.Pos), false
}

// UseAsFldOfReturn is when a struct field value (A.f) flows to a point where it is returned from a function with the
// return expression of the same struct type (A)
type UseAsFldOfReturn struct {
	*TriggerIfNonNil
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (u *UseAsFldOfReturn) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (u *UseAsFldOfReturn) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this UseAsFldOfReturn as a Prestring
func (u *UseAsFldOfReturn) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// UseAsFldOfReturnPrestring is a Prestring storing the needed information to compactly encode a UseAsFldOfReturn
type UseAsFldOfReturnPrestring struct {
	FuncName      string
	FieldName     string
	RetNum        int
	AssignmentStr string
}

func (u UseAsFldOfReturnPrestring) String() string { _ = "STUB: not implemented"; return "" }

// GetRetFldConsumer returns the UseAsFldOfReturn consume trigger with given retKey and expr
func GetRetFldConsumer(retKey Key, expr ast.Expr) *ConsumeTrigger {
	_ = "STUB: not implemented"
	return nil
}

// GetEscapeFldConsumer returns the FldEscape consume trigger with given escKey and selExpr
func GetEscapeFldConsumer(escKey Key, selExpr ast.Expr) *ConsumeTrigger {
	_ = "STUB: not implemented"
	return nil
}

// GetParamFldConsumer returns the ArgFldPass consume trigger with given paramKey and expr
func GetParamFldConsumer(paramKey Key, expr ast.Expr) *ConsumeTrigger {
	_ = "STUB: not implemented"
	return nil
}

// SliceAssign is when a value flows to a point where it is assigned into a slice
type SliceAssign struct {
	*TriggerIfDeepNonNil
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (f *SliceAssign) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (f *SliceAssign) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this SliceAssign as a Prestring
func (f *SliceAssign) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// SliceAssignPrestring is a Prestring storing the needed information to compactly encode a SliceAssign
type SliceAssignPrestring struct {
	TypeName      string
	AssignmentStr string
}

func (f SliceAssignPrestring) String() string { _ = "STUB: not implemented"; return "" }

// ArrayAssign is when a value flows to a point where it is assigned into an array
type ArrayAssign struct {
	*TriggerIfDeepNonNil
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (a *ArrayAssign) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (a *ArrayAssign) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this ArrayAssign as a Prestring
func (a *ArrayAssign) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// ArrayAssignPrestring is a Prestring storing the needed information to compactly encode a SliceAssign
type ArrayAssignPrestring struct {
	TypeName      string
	AssignmentStr string
}

func (a ArrayAssignPrestring) String() string { _ = "STUB: not implemented"; return "" }

// PtrAssign is when a value flows to a point where it is assigned into a pointer
type PtrAssign struct {
	*TriggerIfDeepNonNil
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (f *PtrAssign) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (f *PtrAssign) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this PtrAssign as a Prestring
func (f *PtrAssign) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// PtrAssignPrestring is a Prestring storing the needed information to compactly encode a PtrAssign
type PtrAssignPrestring struct {
	TypeName      string
	AssignmentStr string
}

func (f PtrAssignPrestring) String() string { _ = "STUB: not implemented"; return "" }

// MapAssign is when a value flows to a point where it is assigned into an annotated map
type MapAssign struct {
	*TriggerIfDeepNonNil
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (f *MapAssign) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (f *MapAssign) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this MapAssign as a Prestring
func (f *MapAssign) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// MapAssignPrestring is a Prestring storing the needed information to compactly encode a MapAssign
type MapAssignPrestring struct {
	TypeName      string
	AssignmentStr string
}

func (f MapAssignPrestring) String() string { _ = "STUB: not implemented"; return "" }

// DeepAssignPrimitive is when a value flows to a point where it is assigned
// deeply into an unnannotated object
type DeepAssignPrimitive struct {
	*ConsumeTriggerTautology
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (d *DeepAssignPrimitive) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (d *DeepAssignPrimitive) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this Prestring as a Prestring
func (d *DeepAssignPrimitive) Prestring() Prestring {
	_ = "STUB: not implemented"
	return *new(Prestring)
}

// DeepAssignPrimitivePrestring is a Prestring storing the needed information to compactly encode a DeepAssignPrimitive
type DeepAssignPrimitivePrestring struct {
	AssignmentStr string
}

func (d DeepAssignPrimitivePrestring) String() string { _ = "STUB: not implemented"; return "" }

// ParamAssignDeep is when a value flows to a point where it is assigned deeply into a function parameter
type ParamAssignDeep struct {
	*TriggerIfDeepNonNil
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (p *ParamAssignDeep) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (p *ParamAssignDeep) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this ParamAssignDeep as a Prestring
func (p *ParamAssignDeep) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// ParamAssignDeepPrestring is a Prestring storing the needed information to compactly encode a ParamAssignDeep
type ParamAssignDeepPrestring struct {
	ParamName     string
	AssignmentStr string
}

func (p ParamAssignDeepPrestring) String() string { _ = "STUB: not implemented"; return "" }

// FuncRetAssignDeep is when a value flows to a point where it is assigned deeply into a function return
type FuncRetAssignDeep struct {
	*TriggerIfDeepNonNil
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (f *FuncRetAssignDeep) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (f *FuncRetAssignDeep) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this FuncRetAssignDeep as a Prestring
func (f *FuncRetAssignDeep) Prestring() Prestring {
	_ = "STUB: not implemented"
	return *new(Prestring)
}

// FuncRetAssignDeepPrestring is a Prestring storing the needed information to compactly encode a FuncRetAssignDeep
type FuncRetAssignDeepPrestring struct {
	FuncName      string
	RetNum        int
	AssignmentStr string
}

func (f FuncRetAssignDeepPrestring) String() string { _ = "STUB: not implemented"; return "" }

// VariadicParamAssignDeep is when a value flows to a point where it is assigned deeply into a variadic
// function parameter
type VariadicParamAssignDeep struct {
	*TriggerIfNonNil
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (v *VariadicParamAssignDeep) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (v *VariadicParamAssignDeep) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this VariadicParamAssignDeep as a Prestring
func (v *VariadicParamAssignDeep) Prestring() Prestring {
	_ = "STUB: not implemented"
	return *new(Prestring)
}

// VariadicParamAssignDeepPrestring is a Prestring storing the needed information to compactly encode a VariadicParamAssignDeep
type VariadicParamAssignDeepPrestring struct {
	ParamName     string
	AssignmentStr string
}

func (v VariadicParamAssignDeepPrestring) String() string { _ = "STUB: not implemented"; return "" }

// FieldAssignDeep is when a value flows to a point where it is assigned deeply into a field
type FieldAssignDeep struct {
	*TriggerIfDeepNonNil
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (f *FieldAssignDeep) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (f *FieldAssignDeep) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this FieldAssignDeep as a Prestring
func (f *FieldAssignDeep) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// FieldAssignDeepPrestring is a Prestring storing the needed information to compactly encode a FieldAssignDeep
type FieldAssignDeepPrestring struct {
	FldName       string
	AssignmentStr string
}

func (f FieldAssignDeepPrestring) String() string { _ = "STUB: not implemented"; return "" }

// GlobalVarAssignDeep is when a value flows to a point where it is assigned deeply into a global variable
type GlobalVarAssignDeep struct {
	*TriggerIfDeepNonNil
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (g *GlobalVarAssignDeep) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (g *GlobalVarAssignDeep) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this GlobalVarAssignDeep as a Prestring
func (g *GlobalVarAssignDeep) Prestring() Prestring {
	_ = "STUB: not implemented"
	return *new(Prestring)
}

// GlobalVarAssignDeepPrestring is a Prestring storing the needed information to compactly encode a GlobalVarAssignDeep
type GlobalVarAssignDeepPrestring struct {
	VarName       string
	AssignmentStr string
}

func (g GlobalVarAssignDeepPrestring) String() string { _ = "STUB: not implemented"; return "" }

// LocalVarAssignDeep is when a value flows to a point where it is assigned deeply into a local variable of deeply nonnil type
type LocalVarAssignDeep struct {
	*TriggerIfDeepNonNil
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (l *LocalVarAssignDeep) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (l *LocalVarAssignDeep) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this LocalVarAssignDeep as a Prestring
func (l *LocalVarAssignDeep) Prestring() Prestring {
	_ = "STUB: not implemented"
	return *new(Prestring)
}

// LocalVarAssignDeepPrestring is a Prestring storing the needed information to compactly encode a LocalVarAssignDeep
type LocalVarAssignDeepPrestring struct {
	VarName       string
	AssignmentStr string
}

func (l LocalVarAssignDeepPrestring) String() string { _ = "STUB: not implemented"; return "" }

// ChanSend is when a value flows to a point where it is sent to a channel
type ChanSend struct {
	*TriggerIfDeepNonNil
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (c *ChanSend) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (c *ChanSend) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this ChanSend as a Prestring
func (c *ChanSend) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// ChanSendPrestring is a Prestring storing the needed information to compactly encode a ChanSend
type ChanSendPrestring struct {
	TypeName      string
	AssignmentStr string
}

func (c ChanSendPrestring) String() string { _ = "STUB: not implemented"; return "" }

// FldEscape is when a nilable value flows through a field of a struct that escapes.
// The consumer is added for the fields at sites of escape.
// There are 2 cases, that we currently consider as escaping:
// 1. If a struct is returned from the function where the field has nilable value,
// e.g, If aptr is pointer in struct A, then  `return &A{}` causes the field aptr to escape
// 2. If a struct is parameter of a function and the field is not initialized
// e.g., if we have fun(&A{}) then the field aptr is considered escaped
// TODO: Add struct assignment as another possible cause of field escape
type FldEscape struct {
	*TriggerIfNonNil
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (f *FldEscape) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (f *FldEscape) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this FldEscape as a Prestring
func (f *FldEscape) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// FldEscapePrestring is a Prestring storing the needed information to compactly encode a FldEscape
type FldEscapePrestring struct {
	FieldName     string
	AssignmentStr string
}

func (f FldEscapePrestring) String() string { _ = "STUB: not implemented"; return "" }

// UseAsNonErrorRetDependentOnErrorRetNilability is when a value flows to a point where it is returned from an error returning function
type UseAsNonErrorRetDependentOnErrorRetNilability struct {
	*TriggerIfNonNil

	IsNamedReturn bool
	RetStmt       *ast.ReturnStmt
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (u *UseAsNonErrorRetDependentOnErrorRetNilability) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (u *UseAsNonErrorRetDependentOnErrorRetNilability) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this UseAsNonErrorRetDependentOnErrorRetNilability as a Prestring
func (u *UseAsNonErrorRetDependentOnErrorRetNilability) Prestring() Prestring {
	_ = "STUB: not implemented"
	return *new(Prestring)
}

// UseAsNonErrorRetDependentOnErrorRetNilabilityPrestring is a Prestring storing the needed information to compactly encode a UseAsNonErrorRetDependentOnErrorRetNilability
type UseAsNonErrorRetDependentOnErrorRetNilabilityPrestring struct {
	FuncName      string
	RetNum        int
	RetName       string
	ErrRetNum     int
	IsNamedReturn bool
	AssignmentStr string
}

func (u UseAsNonErrorRetDependentOnErrorRetNilabilityPrestring) String() string {
	_ = "STUB: not implemented"
	return ""
}

// overriding position value to point to the raw return statement, which is the source of the potential error
func (u *UseAsNonErrorRetDependentOnErrorRetNilability) customPos() (token.Pos, bool) {
	_ = "STUB: not implemented"
	return *new(token.Pos), false
}

// UseAsErrorRetWithNilabilityUnknown is when a value flows to a point where it is returned from an error returning function
type UseAsErrorRetWithNilabilityUnknown struct {
	*TriggerIfNonNil

	IsNamedReturn bool
	RetStmt       *ast.ReturnStmt
}

// equals returns true if the passed ConsumingAnnotationTrigger is equal to this one
func (u *UseAsErrorRetWithNilabilityUnknown) equals(other ConsumingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Copy returns a deep copy of this ConsumingAnnotationTrigger
func (u *UseAsErrorRetWithNilabilityUnknown) Copy() ConsumingAnnotationTrigger {
	_ = "STUB: not implemented"
	return *new(ConsumingAnnotationTrigger)
}

// Prestring returns this UseAsErrorRetWithNilabilityUnknown as a Prestring
func (u *UseAsErrorRetWithNilabilityUnknown) Prestring() Prestring {
	_ = "STUB: not implemented"
	return *new(Prestring)
}

// UseAsErrorRetWithNilabilityUnknownPrestring is a Prestring storing the needed information to compactly encode a UseAsErrorRetWithNilabilityUnknown
type UseAsErrorRetWithNilabilityUnknownPrestring struct {
	FuncName      string
	RetNum        int
	IsNamedReturn bool
	RetName       string
	AssignmentStr string
}

func (u UseAsErrorRetWithNilabilityUnknownPrestring) String() string {
	_ = "STUB: not implemented"
	return ""
}

// overriding position value to point to the raw return statement, which is the source of the potential error
func (u *UseAsErrorRetWithNilabilityUnknown) customPos() (token.Pos, bool) {
	_ = "STUB: not implemented"
	return *new(token.Pos), false
}

// don't modify the ConsumeTrigger and ProduceTrigger objects after construction! Pointers
// to them are duplicated

// A ConsumeTrigger represents a point at which a value is consumed that may be required to be
// non-nil by some Annotation (ConsumingAnnotationTrigger). If Parent is not a RootAssertionNode,
// then that AssertionNode represents the expression that will flow into this consumption point.
// If Parent is a RootAssertionNode, then it will be paired with a ProduceTrigger
//
// Expr should be the expression being consumed, not the expression doing the consumption.
// For example, if the field access x.f requires x to be non-nil, then x should be the
// expression embedded in the ConsumeTrigger not x.f.
//
// The set Guards indicates whether this consumption takes places in a context in which
// it is known to be _guarded_ by one or more conditional checks that refine its behavior.
// This is not _all_ conditional checks this is a very small subset of them.
// Consume triggers become guarded via backpropagation across a check that
// `propagateRichChecks` identified with a `RichCheckEffect`. This pass will
// embed a call to `ConsumeTriggerSliceAsGuarded` that will modify all consume
// triggers for the value targeted by the check as guarded by the guard nonces of the
// flowed `RichCheckEffect`.
//
// Like a nil check, guarding is used to indicate information
// refinement local to one branch. The presence of a guard is overwritten by the absence of a guard
// on a given ConsumeTrigger - see MergeConsumeTriggerSlices. Beyond RichCheckEffects,
// Guards consume triggers can be introduced by other sites that are known to
// obey compatible semantics - such as passing the results of one error-returning function
// directly to a return of another.
//
// ConsumeTriggers arise at consumption sites that may guarded by a meaningful conditional check,
// adding that guard as a unique nonce to the set Guards of the trigger. The guard is added when the
// trigger is propagated across the check, so that when it reaches the statement that relies on the
// guard, the statement can see that the check was performed around the site of the consumption. This
// allows the statement to switch to more permissive semantics.
//
// GuardMatched is a boolean used to indicate that this ConsumeTrigger, by the current point in
// backpropagation, passed through a conditional that granted it a guard, and that that guard was
// determined to match the guard expected by a statement such as `v, ok := m[k]`. Since there could have
// been multiple paths in the CFG between the current point in backpropagation and the site at which the
// trigger arose, GuardMatched is true only if a guard arose and was matched along every path. This
// allows the trigger to maintain its more permissive semantics in later stages of backpropagation.
//
// For some productions, such as reading an index of a map, there is no way for them to generate
// nonnil without such a guarding along every path to their point of consumption, so if GuardMatched
// is not true then they will be replaced (by `checkGuardOnFullTrigger`) with an always-produce-nil
// producer. More explanation of this mechanism is provided in the documentation for
// `RootAssertionNode.AddGuardMatch`
//
// nonnil(Guards)
type ConsumeTrigger struct {
	Annotation   ConsumingAnnotationTrigger
	Expr         ast.Expr
	Guards       guard.NonceSet
	GuardMatched bool
}

// equals compares two ConsumeTrigger pointers for equality
func (c *ConsumeTrigger) equals(c2 *ConsumeTrigger) bool { _ = "STUB: not implemented"; return false }

// Copy returns a deep copy of the ConsumeTrigger
func (c *ConsumeTrigger) Copy() *ConsumeTrigger { _ = "STUB: not implemented"; return nil }

// Pos returns the source position (e.g., line) of the consumer's expression. In special cases, such as named return, it
// returns the position of the stored return AST node
func (c *ConsumeTrigger) Pos() token.Pos { _ = "STUB: not implemented"; return *new(token.Pos) }

// MergeConsumeTriggerSlices merges two slices of `ConsumeTrigger`s
// its semantics are slightly unexpected only in its treatment of guarding:
// it intersects guard sets
func MergeConsumeTriggerSlices(left, right []*ConsumeTrigger) []*ConsumeTrigger {
	_ = "STUB: not implemented"
	return nil
}

// intersect guard sets - if a guard isn't present in both branches it can't
// be considered present before the branch

// ConsumeTriggerSliceAsGuarded takes a slice of consume triggers,
// and returns a new slice identical except that each trigger is guarded
func ConsumeTriggerSliceAsGuarded(slice []*ConsumeTrigger, guards ...guard.Nonce) []*ConsumeTrigger {
	_ = "STUB: not implemented"
	return nil
}

// ConsumeTriggerSlicesEq returns true if the two passed slices of ConsumeTrigger contain the same elements
// precondition: no duplications
func ConsumeTriggerSlicesEq(left, right []*ConsumeTrigger) bool {
	_ = "STUB: not implemented"
	return false
}
