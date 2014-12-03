package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
	"testing"
	"testing/quick"
)

// Applicative Laws

func Test_Id_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(Id{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Id_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(Id{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Id_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(Id{}).Homomorphism(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Id_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(Id{}).Interchange(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

func Test_Id_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(Id{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Id_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(Id{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

func Test_Id_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(Id{}).LeftIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Id_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(Id{}).RightIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Id_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(Id{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Semigroup Laws

func Test_Id_SemigroupLaws_Associativity(t *testing.T) {
	f, g := NewSemigroupLaws(Id{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Comonad Laws

func Test_Id_ComonadLaws_Identity(t *testing.T) {
	f, g := NewComonadLaws(Id{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Id_ComonadLaws_Composition(t *testing.T) {
	f, g := NewComonadLaws(Id{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Id_ComonadLaws_Associativity(t *testing.T) {
	f, g := NewComonadLaws(Id{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
