package wishful

import (
	"testing"
	"testing/quick"
)

// Manual tests

func Test_Option_None_Empty(t *testing.T) {
	f := func(x int) Option {
		return None().Empty().(Option)
	}
	g := func(x int) Option {
		return some{}.Empty().(Option)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_Of(t *testing.T) {
	f := func(x int) Option {
		return None().Of(x).(Option)
	}
	g := func(x int) Option {
		return some{}.Of(x).(Option)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_Ap(t *testing.T) {
	f := func(x int) Option {
		return None().Ap(None()).(Option)
	}
	g := func(x int) Option {
		return None()
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_Chain(t *testing.T) {
	f := func(x int) Option {
		return None().Chain(func(v Any) Monad {
			return None()
		}).(Option)
	}
	g := func(x int) Option {
		return None()
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_Map(t *testing.T) {
	f := func(x int) Option {
		return None().Map(Identity).(Option)
	}
	g := func(x int) Option {
		return None()
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_Concat(t *testing.T) {
	f := func(x int) Option {
		return None().Concat(None()).(Option)
	}
	g := func(x int) Option {
		return None()
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_Some_GetOrElse(t *testing.T) {
	f := func(x int, y int) int {
		return x
	}
	g := func(x int, y int) int {
		return some{}.Of(x).(Option).GetOrElse(ConstantNoArgs(y)).(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_GetOrElse(t *testing.T) {
	f := func(x int, y int) int {
		return y
	}
	g := func(x int, y int) int {
		return None().GetOrElse(ConstantNoArgs(y)).(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_Some_OrElse(t *testing.T) {
	f := func(x int, y int) Option {
		return some{x}
	}
	g := func(x int, y int) Option {
		return some{x}.OrElse(some{y})
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_OrElse(t *testing.T) {
	f := func(x int, y int) Option {
		return some{y}
	}
	g := func(x int, y int) Option {
		return None().OrElse(some{y})
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Applicative Laws

// Some

func Test_Option_Some_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(some{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_Some_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(some{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_Some_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(some{}).Homomorphism(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_Some_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(some{}).Interchange(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// None

func Test_Option_None_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(None()).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(None()).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(None()).Homomorphism(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(None()).Interchange(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

// Some

func Test_Option_Some_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(some{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_Some_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(some{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// None

func Test_Option_None_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(None()).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(None()).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

// Some

func Test_Option_Some_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(some{}).LeftIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_Some_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(some{}).RightIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_Some_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(some{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// None

func Test_Option_None_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(None()).LeftIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(None()).RightIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(None()).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Semigroup Laws

// Some

func Test_Option_Some_SemigroupLaws_Associativity(t *testing.T) {
	f, g := NewSemigroupLaws(some{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// None

func Test_Option_None_SemigroupLaws_Associativity(t *testing.T) {
	f, g := NewSemigroupLaws(None()).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
