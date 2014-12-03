package useful

import (
	. "github.com/feyeleanor/wishful"
)

type Validation interface {
	Of(v Any) Point
	Ap(v Applicative) Applicative
	Monad
	Concat(y Semigroup) Semigroup
	Map(f Transform) Functor
	Fold(f, g Transform) Any
	Bimap(f, g Transform) Monad
}

type Failure struct {
	x Any
}

type Success struct {
	x Any
}

func NewFailure(x Any) Failure {
	return Failure{
		x: x,
	}
}

func NewSuccess(x Any) Success {
	return Success{
		x: x,
	}
}

func (x Failure) Of(v Any) Point {
	return NewSuccess(v)
}

func (x Success) Of(v Any) Point {
	return NewSuccess(v)
}

func (x Failure) Ap(v Applicative) Applicative {
	return v.(Validation).Fold(
		func(y Any) Any {
			return NewFailure(concatAnyvals(x.x)(y))
		},
		func(y Any) Any {
			return NewFailure(x.x)
		},
	).(Applicative)
}

func (x Success) Ap(v Applicative) Applicative {
	return v.(Functor).Map(func(g Any) Any {
		fun := NewFunction(x.x)
		res, _ := fun.Call(g)
		return res
	}).(Applicative)
}

func (x Failure) Chain(f Step) Monad {
	return x
}

func (x Success) Chain(f Step) Monad {
	return f(x.x)
}

func (x Failure) Map(f Transform) Functor {
	return x.Bimap(Identity, f).(Functor)
}

func (x Success) Map(f Transform) Functor {
	return x.Bimap(Identity, f).(Functor)
}

func (x Failure) Concat(y Semigroup) Semigroup {
	a := y.(Validation)
	b := a.Bimap(
		concatAnyvals(x.x),
		Identity,
	)
	return b.(Semigroup)
}

func (x Success) Concat(y Semigroup) Semigroup {
	a := y.(Functor)
	b := a.Map(concatAnyvals(x.x))
	return b.(Semigroup)
}

// Derived

func (x Failure) Fold(f, g Transform) Any {
	return f(x.x)
}

func (x Success) Fold(f, g Transform) Any {
	return g(x.x)
}

func (x Failure) Bimap(f, g Transform) Monad {
	return NewFailure(f(x.x))
}

func (x Success) Bimap(f, g Transform) Monad {
	return NewSuccess(g(x.x))
}
