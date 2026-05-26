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

package annotation

import (
	"go/ast"
	"go/token"
	"go/types"
)

// A ProducingAnnotationTrigger is a possible reason that a nil value might be produced
//
// All ProducingAnnotationTriggers must embed one of the following 4 structs:
// -TriggerIfNilable
// -TriggerIfDeepNilable
// -ProduceTriggerTautology
// -ProduceTriggerNever
//
// This is because there are interfaces, such as AdmitsPrimitive, that are implemented only for those
// structs, and to which a ProducingAnnotationTrigger must be able to be case
type ProducingAnnotationTrigger interface {
	// CheckProduce can be called to determined whether this trigger should be triggered
	// given a particular Annotation map
	// for example - a `FuncReturn` trigger triggers iff the corresponding function has
	// nilable return type
	CheckProduce(Map) bool

	// NeedsGuardMatch returns whether this production is contingent on being
	// paired with a guarded consumer.
	// In other words, this production is only given the freedom to produce
	// a non-nil value in the case that it is matched with a guarded consumer.
	// otherwise, it is replaced with annotation.GuardMissing
	NeedsGuardMatch() bool

	// SetNeedsGuard sets the underlying Guard-Neediness of this ProduceTrigger, if present
	// This should be very sparingly used, and only with utter conviction of correctness.
	// Default setting for ProduceTriggers is to not need a guard.
	SetNeedsGuard(bool)

	Prestring() Prestring

	// Kind returns the kind of the trigger.
	Kind() TriggerKind

	// UnderlyingSite returns the underlying site this trigger's nilability depends on. If the
	// trigger always or never fires, the site is nil.
	UnderlyingSite() Key

	// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
	equals(ProducingAnnotationTrigger) bool
}

// TriggerIfNilable is a general trigger indicating that the bad case occurs when a certain Annotation
// key is nilable
type TriggerIfNilable struct {
	Ann        Key
	NeedsGuard bool
}

// Prestring returns this Prestring as a Prestring
func (*TriggerIfNilable) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// TriggerIfNilablePrestring is a Prestring storing the needed information to compactly encode a TriggerIfNilable
type TriggerIfNilablePrestring struct{}

func (TriggerIfNilablePrestring) String() string { _ = "STUB: not implemented"; return "" }

// CheckProduce returns true if the underlying annotation is present in the passed map and nilable
func (t *TriggerIfNilable) CheckProduce(annMap Map) bool { _ = "STUB: not implemented"; return false }

// NeedsGuardMatch returns true if this trigger needs to be matched with a guarded consumer
func (t *TriggerIfNilable) NeedsGuardMatch() bool {
	_ = "STUB: not implemented"

	// SetNeedsGuard sets the underlying Guard-Neediness of this ProduceTrigger, if present
	return false
}

func (t *TriggerIfNilable) SetNeedsGuard(b bool) {
	_ = "STUB: not implemented"

	// Kind returns Conditional.
	return
}

func (t *TriggerIfNilable) Kind() TriggerKind {
	_ = "STUB: not implemented"

	// UnderlyingSite returns the underlying site this trigger's nilability depends on.
	return *new(TriggerKind)
}

func (t *TriggerIfNilable) UnderlyingSite() Key {
	_ = "STUB: not implemented"

	// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
	return *new(Key)
}

func (t *TriggerIfNilable) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// TriggerIfDeepNilable is a general trigger indicating the the bad case occurs when a certain Annotation
// key is deeply nilable
type TriggerIfDeepNilable struct {
	Ann        Key
	NeedsGuard bool
}

// Prestring returns this Prestring as a Prestring
func (*TriggerIfDeepNilable) Prestring() Prestring {
	_ = "STUB: not implemented"
	return *new(Prestring)
}

// TriggerIfDeepNilablePrestring is a Prestring storing the needed information to compactly encode a TriggerIfDeepNilable
type TriggerIfDeepNilablePrestring struct{}

func (TriggerIfDeepNilablePrestring) String() string { _ = "STUB: not implemented"; return "" }

// CheckProduce returns true if the underlying annotation is present in the passed map and deeply nilable
func (t *TriggerIfDeepNilable) CheckProduce(annMap Map) bool {
	_ = "STUB: not implemented"
	return false
}

// NeedsGuardMatch returns true if this trigger needs to be matched with a guarded consumer
func (t *TriggerIfDeepNilable) NeedsGuardMatch() bool {
	_ = "STUB: not implemented"

	// SetNeedsGuard sets the underlying Guard-Neediness of this ProduceTrigger, if present
	return false
}

func (t *TriggerIfDeepNilable) SetNeedsGuard(b bool) {
	_ = "STUB: not implemented"

	// Kind returns DeepConditional.
	return
}

func (t *TriggerIfDeepNilable) Kind() TriggerKind {
	_ = "STUB: not implemented"
	return *

	// UnderlyingSite returns the underlying site this trigger's nilability depends on.
	new(TriggerKind)
}

func (t *TriggerIfDeepNilable) UnderlyingSite() Key {
	_ = "STUB: not implemented"

	// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
	return *new(Key)
}

func (t *TriggerIfDeepNilable) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// ProduceTriggerTautology is used for trigger producers that will always result in nil
type ProduceTriggerTautology struct {
	NeedsGuard bool
}

// CheckProduce returns true
func (*ProduceTriggerTautology) CheckProduce(Map) bool {
	_ = "STUB: not implemented"

	// NeedsGuardMatch returns true if this trigger needs to be matched with a guarded consumer
	return false
}

func (p *ProduceTriggerTautology) NeedsGuardMatch() bool { _ = "STUB: not implemented"; return false }

// SetNeedsGuard sets the underlying Guard-Neediness of this ProduceTrigger, if present
func (p *ProduceTriggerTautology) SetNeedsGuard(b bool) {
	_ = "STUB: not implemented"

	// Prestring returns this Prestring as a Prestring
	return
}

func (*ProduceTriggerTautology) Prestring() Prestring {
	_ = "STUB: not implemented"
	return *new(Prestring)
}

// ProduceTriggerTautologyPrestring is a Prestring storing the needed information to compactly encode a ProduceTriggerTautology
type ProduceTriggerTautologyPrestring struct{}

func (ProduceTriggerTautologyPrestring) String() string { _ = "STUB: not implemented"; return "" }

// Kind returns Always.
func (*ProduceTriggerTautology) Kind() TriggerKind {
	_ = "STUB: not implemented"

	// UnderlyingSite always returns nil.
	return *new(TriggerKind)
}

func (*ProduceTriggerTautology) UnderlyingSite() Key {
	_ = "STUB: not implemented"

	// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
	return *new(Key)
}

func (p *ProduceTriggerTautology) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// ProduceTriggerNever is used for trigger producers that will never be nil
type ProduceTriggerNever struct {
	NeedsGuard bool
}

// Prestring returns this Prestring as a Prestring
func (*ProduceTriggerNever) Prestring() Prestring {
	_ = "STUB: not implemented"
	return *new(Prestring)
}

// ProduceTriggerNeverPrestring is a Prestring storing the needed information to compactly encode a ProduceTriggerNever
type ProduceTriggerNeverPrestring struct{}

func (ProduceTriggerNeverPrestring) String() string { _ = "STUB: not implemented"; return "" }

// CheckProduce returns true false
func (*ProduceTriggerNever) CheckProduce(Map) bool {
	_ = "STUB: not implemented"

	// NeedsGuardMatch returns true if this trigger needs to be matched with a guarded consumer
	return false
}

func (p *ProduceTriggerNever) NeedsGuardMatch() bool {
	_ = "STUB: not implemented"

	// SetNeedsGuard sets the underlying Guard-Neediness of this ProduceTrigger, if present
	return false
}

func (p *ProduceTriggerNever) SetNeedsGuard(b bool) {
	_ = "STUB: not implemented"

	// Kind returns Never.
	return
}

func (*ProduceTriggerNever) Kind() TriggerKind {
	_ = "STUB: not implemented"

	// UnderlyingSite always returns nil.
	return *new(TriggerKind)
}

func (*ProduceTriggerNever) UnderlyingSite() Key {
	_ = "STUB: not implemented"

	// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
	return *new(Key)
}

func (p *ProduceTriggerNever) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// note: each of the following two productions, ExprOkCheck, and RangeIndexAssignment, should be
// obselete now that we don't add consumptions for basic-typed expressions like ints and bools to
// begin with - TODO: verify that these productions are always no-ops and remove

// ExprOkCheck is used when a value is determined to flow from the second argument of a map or typecast
// operation that necessarily makes it boolean and thus non-nil
type ExprOkCheck struct {
	*ProduceTriggerNever
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (e *ExprOkCheck) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// RangeIndexAssignment is used when a value is determined to flow from the first argument of a
// range loop, and thus be an integer and non-nil
type RangeIndexAssignment struct {
	*ProduceTriggerNever
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (r *RangeIndexAssignment) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// PositiveNilCheck is used when a value is checked in a conditional to BE nil
type PositiveNilCheck struct {
	*ProduceTriggerTautology
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (p *PositiveNilCheck) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this Prestring as a Prestring
func (*PositiveNilCheck) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// PositiveNilCheckPrestring is a Prestring storing the needed information to compactly encode a PositiveNilCheck
type PositiveNilCheckPrestring struct{}

func (PositiveNilCheckPrestring) String() string { _ = "STUB: not implemented"; return "" }

// NegativeNilCheck is used when a value is checked in a conditional to NOT BE nil
type NegativeNilCheck struct {
	*ProduceTriggerNever
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (n *NegativeNilCheck) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this Prestring as a Prestring
func (*NegativeNilCheck) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// NegativeNilCheckPrestring is a Prestring storing the needed information to compactly encode a NegativeNilCheck
type NegativeNilCheckPrestring struct{}

func (NegativeNilCheckPrestring) String() string { _ = "STUB: not implemented"; return "" }

// OkReadReflCheck is used to produce nonnil for artifacts of successful `ok` forms (e.g., maps, channels, type casts).
// For example, a map value `m` that was read from in a `v, ok := m[k]` check followed by a positive check of `ok`, implies `m` is non-nil.
// This is valid because nil maps contain no keys.
type OkReadReflCheck struct {
	*ProduceTriggerNever
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (o *OkReadReflCheck) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// RangeOver is used when a value is ranged over - and thus nonnil in its range body
type RangeOver struct {
	*ProduceTriggerNever
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (r *RangeOver) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// ConstNil is when a value is determined to flow from a constant nil expression
type ConstNil struct {
	*ProduceTriggerTautology
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (c *ConstNil) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this Prestring as a Prestring
func (*ConstNil) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// ConstNilPrestring is a Prestring storing the needed information to compactly encode a ConstNil
type ConstNilPrestring struct{}

func (ConstNilPrestring) String() string { _ = "STUB: not implemented"; return "" }

// UnassignedFld is when a field of struct is not assigned at initialization
type UnassignedFld struct {
	*ProduceTriggerTautology
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (u *UnassignedFld) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this Prestring as a Prestring
func (*UnassignedFld) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// UnassignedFldPrestring is a Prestring storing the needed information to compactly encode a UnassignedFld
type UnassignedFldPrestring struct{}

func (UnassignedFldPrestring) String() string { _ = "STUB: not implemented"; return "" }

// NoVarAssign is when a value is determined to flow from a variable that wasn't assigned to
type NoVarAssign struct {
	*ProduceTriggerTautology
	VarObj *types.Var
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (n *NoVarAssign) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this Prestring as a Prestring
func (n *NoVarAssign) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// NoVarAssignPrestring is a Prestring storing the needed information to compactly encode a NoVarAssign
type NoVarAssignPrestring struct {
	VarName string
}

func (n NoVarAssignPrestring) String() string { _ = "STUB: not implemented"; return "" }

// BlankVarReturn is when a value is determined to flow from a blank variable ('_') to a return of the function
type BlankVarReturn struct {
	*ProduceTriggerTautology
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (b *BlankVarReturn) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this Prestring as a Prestring
func (*BlankVarReturn) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// BlankVarReturnPrestring is a Prestring storing the needed information to compactly encode a BlankVarReturn
type BlankVarReturnPrestring struct{}

func (BlankVarReturnPrestring) String() string { _ = "STUB: not implemented"; return "" }

// DuplicateParamProducer duplicates a given produce trigger, assuming the given produce trigger
// is of FuncParam.
func DuplicateParamProducer(t *ProduceTrigger, location token.Position) *ProduceTrigger {
	_ = "STUB: not implemented"
	return nil
}

// FuncParam is used when a value is determined to flow from a function parameter. This consumer
// trigger can be used on top of two different sites: ParamAnnotationKey &
// CallSiteParamAnnotationKey. ParamAnnotationKey is the parameter site in the function
// declaration; CallSiteParamAnnotationKey is the argument site in the call expression.
// CallSiteParamAnnotationKey is specifically used for functions with contracts since we need to
// duplicate the sites for context sensitivity.
type FuncParam struct {
	*TriggerIfNilable
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (f *FuncParam) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this FuncParam as a Prestring
func (f *FuncParam) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// FuncParamPrestring is a Prestring storing the needed information to compactly encode a FuncParam
type FuncParamPrestring struct {
	ParamName string
	FuncName  string
	// Location is empty for a FuncParam enclosing ParamAnnotationKey. Location points to the
	// location of the argument pass at the call site for a FuncParam enclosing CallSiteParamAnnotationKey.
	Location string
}

func (f FuncParamPrestring) String() string { _ = "STUB: not implemented"; return "" }

// MethodRecv is used when a value is determined to flow from a method receiver
type MethodRecv struct {
	*TriggerIfNilable
	VarDecl *types.Var
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (m *MethodRecv) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this MethodRecv as a Prestring
func (m *MethodRecv) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// MethodRecvPrestring is a Prestring storing the needed information to compactly encode a MethodRecv
type MethodRecvPrestring struct {
	RecvName string
}

func (m MethodRecvPrestring) String() string { _ = "STUB: not implemented"; return "" }

// MethodRecvDeep is used when a value is determined to flow deeply from a method receiver
type MethodRecvDeep struct {
	*TriggerIfDeepNilable
	VarDecl *types.Var
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (m *MethodRecvDeep) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this MethodRecv as a Prestring
func (m *MethodRecvDeep) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// MethodRecvDeepPrestring is a Prestring storing the needed information to compactly encode a MethodRecv
type MethodRecvDeepPrestring struct {
	RecvName string
}

func (m MethodRecvDeepPrestring) String() string { _ = "STUB: not implemented"; return "" }

// VariadicFuncParam is used when a value is determined to flow from a variadic function parameter,
// and thus always be nilable
type VariadicFuncParam struct {
	*ProduceTriggerTautology
	VarDecl *types.Var
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (v *VariadicFuncParam) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this Prestring as a Prestring
func (v *VariadicFuncParam) Prestring() Prestring {
	_ = "STUB: not implemented"
	return *new(Prestring)
}

// VariadicFuncParamPrestring is a Prestring storing the needed information to compactly encode a VariadicFuncParam
type VariadicFuncParamPrestring struct {
	ParamName string
}

func (v VariadicFuncParamPrestring) String() string { _ = "STUB: not implemented"; return "" }

// TrustedFuncNilable is used when a value is determined to be nilable by a trusted function call
type TrustedFuncNilable struct {
	*ProduceTriggerTautology
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (t *TrustedFuncNilable) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this Prestring as a Prestring
func (*TrustedFuncNilable) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// TrustedFuncNilablePrestring is a Prestring storing the needed information to compactly encode a TrustedFuncNilable
type TrustedFuncNilablePrestring struct{}

func (TrustedFuncNilablePrestring) String() string { _ = "STUB: not implemented"; return "" }

// TrustedFuncNonnil is used when a value is determined to be nonnil by a trusted function call
type TrustedFuncNonnil struct {
	*ProduceTriggerNever
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (t *TrustedFuncNonnil) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this Prestring as a Prestring
func (*TrustedFuncNonnil) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// TrustedFuncNonnilPrestring is a Prestring storing the needed information to compactly encode a TrustedFuncNonnil
type TrustedFuncNonnilPrestring struct{}

func (TrustedFuncNonnilPrestring) String() string { _ = "STUB: not implemented"; return "" }

// FldRead is used when a value is determined to flow from a read to a field
type FldRead struct {
	*TriggerIfNilable
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (f *FldRead) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this FldRead as a Prestring
func (f *FldRead) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// FldReadPrestring is a Prestring storing the needed information to compactly encode a FldRead
type FldReadPrestring struct {
	FieldName string
}

func (f FldReadPrestring) String() string { _ = "STUB: not implemented"; return "" }

// ParamFldRead is used when a struct field value is determined to flow from the param of a function to a consumption
// site within the body of the function
type ParamFldRead struct {
	*TriggerIfNilable
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (f *ParamFldRead) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this ParamFldRead as a Prestring
func (f *ParamFldRead) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// ParamFldReadPrestring is a Prestring storing the needed information to compactly encode a ParamFldRead
type ParamFldReadPrestring struct {
	FieldName string
}

func (f ParamFldReadPrestring) String() string { _ = "STUB: not implemented"; return "" }

// FldReturn is used when a struct field value is determined to flow from a return value of a function
type FldReturn struct {
	*TriggerIfNilable
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (f *FldReturn) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

func (f FldReturn) String() string { _ = "STUB: not implemented"; return "" }

// Prestring returns this FldReturn as a Prestring
func (f *FldReturn) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// FldReturnPrestring is a Prestring storing the needed information to compactly encode a FldReturn
type FldReturnPrestring struct {
	RetNum    int
	FuncName  string
	FieldName string
}

func (f FldReturnPrestring) String() string { _ = "STUB: not implemented"; return "" }

// FuncReturn is used when a value is determined to flow from the return of a function. This
// consumer trigger can be used on top of two different sites: RetAnnotationKey &
// CallSiteRetAnnotationKey. RetAnnotationKey is the parameter site in the function declaration;
// CallSiteRetAnnotationKey is the argument site in the call expression. CallSiteRetAnnotationKey
// is specifically used for functions with contracts since we need to duplicate the sites for
// context sensitivity.
type FuncReturn struct {
	*TriggerIfNilable

	IsFromRichCheckEffectFunc bool
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (f *FuncReturn) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this FuncReturn as a Prestring
func (f *FuncReturn) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// FuncReturnPrestring is a Prestring storing the needed information to compactly encode a FuncReturn
type FuncReturnPrestring struct {
	RetNum   int
	FuncName string
	// Location is empty for a FuncReturn enclosing RetAnnotationKey. Location points to the
	// location of the result return at the call site for a FuncReturn enclosing CallSiteRetAnnotationKey.
	Location string
}

func (f FuncReturnPrestring) String() string { _ = "STUB: not implemented"; return "" }

// MethodReturn is used when a value is determined to flow from the return of a method
type MethodReturn struct {
	*TriggerIfNilable
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (m *MethodReturn) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this MethodReturn as a Prestring
func (m *MethodReturn) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// MethodReturnPrestring is a Prestring storing the needed information to compactly encode a MethodReturn
type MethodReturnPrestring struct {
	RetNum   int
	FuncName string
}

func (m MethodReturnPrestring) String() string { _ = "STUB: not implemented"; return "" }

// MethodResultReachesInterface is used when a result of a method is determined to flow into a result of an interface using inheritance
type MethodResultReachesInterface struct {
	*TriggerIfNilable
	AffiliationPair
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (m *MethodResultReachesInterface) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this MethodResultReachesInterface as a Prestring
func (m *MethodResultReachesInterface) Prestring() Prestring {
	_ = "STUB: not implemented"
	return *new(Prestring)
}

// MethodResultReachesInterfacePrestring is a Prestring storing the needed information to compactly encode a MethodResultReachesInterface
type MethodResultReachesInterfacePrestring struct {
	RetNum   int
	ImplName string
	IntName  string
}

func (m MethodResultReachesInterfacePrestring) String() string {
	_ = "STUB: not implemented"

	// InterfaceParamReachesImplementation is used when a param of a method is determined to flow into the param of an implementing method
	return ""
}

type InterfaceParamReachesImplementation struct {
	*TriggerIfNilable
	AffiliationPair
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (i *InterfaceParamReachesImplementation) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this InterfaceParamReachesImplementation as a Prestring
func (i *InterfaceParamReachesImplementation) Prestring() Prestring {
	_ = "STUB: not implemented"
	return *new(Prestring)
}

// InterfaceParamReachesImplementationPrestring is a Prestring storing the needed information to compactly encode a InterfaceParamReachesImplementation
type InterfaceParamReachesImplementationPrestring struct {
	ParamName string
	IntName   string
	ImplName  string
}

func (i InterfaceParamReachesImplementationPrestring) String() string {
	_ = "STUB: not implemented"

	// GlobalVarRead is when a value is determined to flow from a read to a global variable
	return ""
}

type GlobalVarRead struct {
	*TriggerIfNilable
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (g *GlobalVarRead) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this GlobalVarRead as a Prestring
func (g *GlobalVarRead) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// GlobalVarReadPrestring is a Prestring storing the needed information to compactly encode a GlobalVarRead
type GlobalVarReadPrestring struct {
	VarName string
}

func (g GlobalVarReadPrestring) String() string { _ = "STUB: not implemented"; return "" }

// MapRead is when a value is determined to flow from a map index expression
// These should always be instantiated with NeedsGuard = true
type MapRead struct {
	*TriggerIfDeepNilable
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (m *MapRead) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this MapRead as a Prestring
func (m *MapRead) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// MapReadPrestring is a Prestring storing the needed information to compactly encode a MapRead
type MapReadPrestring struct {
	TypeName string
}

func (m MapReadPrestring) String() string { _ = "STUB: not implemented"; return "" }

// ArrayRead is when a value is determined to flow from an array index expression
type ArrayRead struct {
	*TriggerIfDeepNilable
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (a *ArrayRead) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this ArrayRead as a Prestring
func (a *ArrayRead) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// ArrayReadPrestring is a Prestring storing the needed information to compactly encode a ArrayRead
type ArrayReadPrestring struct {
	TypeName string
}

func (a ArrayReadPrestring) String() string { _ = "STUB: not implemented"; return "" }

// SliceRead is when a value is determined to flow from a slice index expression
type SliceRead struct {
	*TriggerIfDeepNilable
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (s *SliceRead) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this SliceRead as a Prestring
func (s *SliceRead) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// SliceReadPrestring is a Prestring storing the needed information to compactly encode a SliceRead
type SliceReadPrestring struct {
	TypeName string
}

func (s SliceReadPrestring) String() string { _ = "STUB: not implemented"; return "" }

// PtrRead is when a value is determined to flow from a read to a pointer
type PtrRead struct {
	*TriggerIfDeepNilable
}

// Prestring returns this PtrRead as a Prestring
func (p *PtrRead) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (p *PtrRead) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// PtrReadPrestring is a Prestring storing the needed information to compactly encode a PtrRead
type PtrReadPrestring struct {
	TypeName string
}

func (p PtrReadPrestring) String() string { _ = "STUB: not implemented"; return "" }

// ChanRecv is when a value is determined to flow from a channel receive
type ChanRecv struct {
	*TriggerIfDeepNilable
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (c *ChanRecv) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this ChanRecv as a Prestring
func (c *ChanRecv) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// ChanRecvPrestring is a Prestring storing the needed information to compactly encode a ChanRecv
type ChanRecvPrestring struct {
	TypeName string
}

func (c ChanRecvPrestring) String() string { _ = "STUB: not implemented"; return "" }

// FuncParamDeep is used when a value is determined to flow deeply from a function parameter
type FuncParamDeep struct {
	*TriggerIfDeepNilable
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (f *FuncParamDeep) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this FuncParamDeep as a Prestring
func (f *FuncParamDeep) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// FuncParamDeepPrestring is a Prestring storing the needed information to compactly encode a FuncParamDeep
type FuncParamDeepPrestring struct {
	ParamName string
}

func (f FuncParamDeepPrestring) String() string { _ = "STUB: not implemented"; return "" }

// VariadicFuncParamDeep is used when a value is determined to flow deeply from a variadic function
// parameter, and thus be nilable iff the shallow Annotation on that parameter is nilable
type VariadicFuncParamDeep struct {
	*TriggerIfNilable
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (v *VariadicFuncParamDeep) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this VariadicFuncParamDeep as a Prestring
func (v *VariadicFuncParamDeep) Prestring() Prestring {
	_ = "STUB: not implemented"
	return *new(Prestring)
}

// VariadicFuncParamDeepPrestring is a Prestring storing the needed information to compactly encode a VariadicFuncParamDeep
type VariadicFuncParamDeepPrestring struct {
	ParamName string
}

func (v VariadicFuncParamDeepPrestring) String() string { _ = "STUB: not implemented"; return "" }

// FuncReturnDeep is used when a value is determined to flow from the deep Annotation of the return
// of a function
type FuncReturnDeep struct {
	*TriggerIfDeepNilable
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (f *FuncReturnDeep) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this FuncReturnDeep as a Prestring
func (f *FuncReturnDeep) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// FuncReturnDeepPrestring is a Prestring storing the needed information to compactly encode a FuncReturnDeep
type FuncReturnDeepPrestring struct {
	RetNum   int
	FuncName string
}

func (f FuncReturnDeepPrestring) String() string { _ = "STUB: not implemented"; return "" }

// FldReadDeep is used when a value is determined to flow from the deep Annotation of a field that is
// read and then indexed into - for example x.f[0]
type FldReadDeep struct {
	*TriggerIfDeepNilable
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (f *FldReadDeep) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this FldReadDeep as a Prestring
func (f *FldReadDeep) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// FldReadDeepPrestring is a Prestring storing the needed information to compactly encode a FldReadDeep
type FldReadDeepPrestring struct {
	FieldName string
}

func (f FldReadDeepPrestring) String() string { _ = "STUB: not implemented"; return "" }

// LocalVarReadDeep is when a value is determined to flow deeply from a local variable.
type LocalVarReadDeep struct {
	*TriggerIfDeepNilable
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (v *LocalVarReadDeep) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this LocalVarReadDeep as a Prestring
func (v LocalVarReadDeep) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// LocalVarReadDeepPrestring is a Prestring storing the needed information to compactly encode a LocalVarReadDeep
type LocalVarReadDeepPrestring struct {
	VarName string
}

func (v LocalVarReadDeepPrestring) String() string { _ = "STUB: not implemented"; return "" }

// GlobalVarReadDeep is when a value is determined to flow from the deep Annotation of a global variable
// that is read and indexed into
type GlobalVarReadDeep struct {
	*TriggerIfDeepNilable
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (g *GlobalVarReadDeep) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this GlobalVarReadDeep as a Prestring
func (g *GlobalVarReadDeep) Prestring() Prestring {
	_ = "STUB: not implemented"
	return *new(Prestring)
}

// GlobalVarReadDeepPrestring is a Prestring storing the needed information to compactly encode a GlobalVarReadDeep
type GlobalVarReadDeepPrestring struct {
	VarName string
}

func (g GlobalVarReadDeepPrestring) String() string { _ = "STUB: not implemented"; return "" }

// GuardMissing is when a value is determined to flow from a site that requires a guard,
// to a site that is not guarded by that guard.
//
// GuardMissing is never created during backpropagation, but on a call to RootAssertionNode.ProcessEntry
// that checks the guards on ever FullTrigger created, it is substituted for the producer in any
// FullTrigger whose producer has NeedsGuard = true and whose consumer has GuardMatched = false,
// guaranteeing that that producer triggers.
//
// For example, from a read to map without the `v, ok := m[k]` form, thus always resulting in nilable
// regardless of `m`'s deep nilability
type GuardMissing struct {
	*ProduceTriggerTautology
	OldAnnotation ProducingAnnotationTrigger
}

// equals returns true if the passed ProducingAnnotationTrigger is equal to this one
func (g *GuardMissing) equals(other ProducingAnnotationTrigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Prestring returns this GuardMissing as a Prestring
func (g *GuardMissing) Prestring() Prestring { _ = "STUB: not implemented"; return *new(Prestring) }

// GuardMissingPrestring is a Prestring storing the needed information to compactly encode a GuardMissing
type GuardMissingPrestring struct {
	OldPrestring Prestring
}

func (g GuardMissingPrestring) String() string { _ = "STUB: not implemented"; return "" }

// don't modify the ConsumeTrigger and ProduceTrigger objects after construction! Pointers
// to them are duplicated

// A ProduceTrigger represents a point at which a value is produced that may be nilable because of
// an Annotation (ProducingAnnotationTrigger). Will always be paired with a ConsumeTrigger.
// For semantics' sake, the Annotation field of a ProduceTrigger is all that matters - the Expr is
// included only to produce more informative error messages
type ProduceTrigger struct {
	Annotation ProducingAnnotationTrigger
	Expr       ast.Expr
}
