package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Option interface {
}

type Some struct {
	x AnyVal
}

type None struct {
}

func NewSome(x AnyVal) Some {
	return Some{
		x: x,
	}
}

func NewNone() None {
	return None{}
}

func (x Some) Of(v AnyVal) Point {
	return NewSome(v)
}

func (x None) Of(v AnyVal) Point {
	return NewSome(v)
}

func (x Some) Empty() Monoid {
	return NewNone()
}

func (x None) Empty() Monoid {
	return NewNone()
}

func (x Some) Ap(v Applicative) Applicative {
	return fromMonadToApplicativeAp(x, v)
}

func (x None) Ap(v Applicative) Applicative {
	return x
}

func (x Some) Chain(f func(v AnyVal) Monad) Monad {
	return f(x.x)
}

func (x None) Chain(f func(v AnyVal) Monad) Monad {
	return x
}

func (x Some) Map(f func(v AnyVal) AnyVal) Functor {
	res := x.Chain(func(v AnyVal) Monad {
		return NewSome(f(v))
	})
	return res.(Functor)
}

func (x None) Map(f func(v AnyVal) AnyVal) Functor {
	return x
}

func (x Some) Concat(y Semigroup) Semigroup {
	return concat(x, y)
}

func (x None) Concat(y Semigroup) Semigroup {
	return x
}