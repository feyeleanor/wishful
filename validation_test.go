package wishful

import (
	"testing"
	"testing/quick"
)

type value struct {
	Val int
}

func (v value) Concat(a Semigroup) Semigroup {
	return value{v.Val + a.(value).Val}
}

// Manual tests

func Test_Validation_Failure_New(t *testing.T) {
	f := func(x int) failure {
		return Failure(x)
	}
	g := func(x int) failure {
		return failure{x}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_Of(t *testing.T) {
	f := func(x int) Validation {
		return success{}.Of(x).(Validation)
	}
	g := func(x int) Validation {
		return failure{}.Of(x).(Validation)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_Ap(t *testing.T) {
	f := func(x int) Validation {
		return failure{value{x}}.Ap(failure{value{1}}).(Validation)
	}
	g := func(x int) Validation {
		return failure{value{x + 1}}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_ApWithSuccess(t *testing.T) {
	f := func(x int) Validation {
		return failure{value{x}}.Ap(success{value{1}}).(Validation)
	}
	g := func(x int) Validation {
		return failure{value{x}}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_Chain(t *testing.T) {
	f := func(x int) Validation {
		return failure{}.Chain(func(v Any) Monad {
			return failure{}
		}).(Validation)
	}
	g := func(x int) Validation {
		return failure{}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_Map(t *testing.T) {
	f := func(x int) Validation {
		return failure{}.Map(Identity).(Validation)
	}
	g := func(x int) Validation {
		return failure{}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_Concat(t *testing.T) {
	f := func(x int) Validation {
		return failure{value{x}}.Concat(failure{value{1}}).(Validation)
	}
	g := func(x int) Validation {
		return failure{value{x + 1}}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_Bimap(t *testing.T) {
	f := func(x int) Validation {
		return failure{x}.Bimap(func(v Any) Any {
			return v.(int) + 1
		}, Identity).(Validation)
	}
	g := func(x int) Validation {
		return failure{x + 1}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Success_Bimap(t *testing.T) {
	f := func(x int) Validation {
		return success{x}.Bimap(Identity, func(v Any) Any {
			return v.(int) + 1
		}).(Validation)
	}
	g := func(x int) Validation {
		return success{x + 1}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Applicative Laws

// Failure

func Test_Validation_Failure_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(failure{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(failure{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(failure{}).Homomorphism(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(failure{}).Interchange(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Success

func Test_Validation_Success_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(success{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Success_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(success{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Success_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(success{}).Homomorphism(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Success_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(success{}).Interchange(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

// Failure

func Test_Validation_Failure_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(failure{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(failure{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Success

func Test_Validation_Success_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(success{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Success_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(success{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

// Failure

func Test_Validation_Failure_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(failure{}).LeftIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(failure{}).RightIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(failure{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Success

func Test_Validation_Success_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(success{}).LeftIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Success_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(success{}).RightIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Success_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(success{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Semigroup Laws

// Failure

func Test_Validation_Failure_SemigroupLaws_Associativity(t *testing.T) {
	f, g := NewSemigroupLaws(failure{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Success

func Test_Validation_Success_SemigroupLaws_Associativity(t *testing.T) {
	f, g := NewSemigroupLaws(success{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
