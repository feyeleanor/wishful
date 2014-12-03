package useful

import (
	. "./wishful"
)

type Either interface {
	Of(Any) Point
	Ap(Applicative) Applicative
	Chain(func(Any) Monad) Monad
	Concat(Semigroup) Semigroup
	Map(Transform) Functor
	Bimap(f, g Transform) Monad
	Fold(f, g Transform) Any
	Swap() Monad
	Sequence(Point) Any
	Traverse(Transform, Point) Functor
}

type Left struct {
	x Any
}

type Right struct {
	x Any
}

func NewLeft(x Any) Left {
	return Left{x: x}
}

func NewRight(x Any) Right {
	return Right{x: x}
}

func (x Left) Of(v Any) Point {
	return NewRight(v)
}

func (x Right) Of(v Any) Point {
	return NewRight(v)
}

func (x Left) Ap(v Applicative) Applicative {
	return x
}

func (x Right) Ap(v Applicative) Applicative {
	return fromMonadToApplicativeAp(x, v)
}

func (x Left) Chain(f func(v Any) Monad) Monad {
	return x
}

func (x Right) Chain(f func(v Any) Monad) Monad {
	return f(x.x)
}

func (x Left) Map(f Transform) Functor {
	return x
}

func (x Right) Map(f Transform) Functor {
	res := x.Chain(func(v Any) Monad {
		return NewRight(f(v))
	})
	return res.(Functor)
}

func (x Left) Concat(y Semigroup) Semigroup {
	return x
}

func (x Right) Concat(y Semigroup) Semigroup {
	return concat(x, y)
}

// Derived
func (x Left) Swap() Monad {
	return NewRight(x.x)
}

func (x Right) Swap() Monad {
	return NewLeft(x.x)
}

func (x Left) Bimap(f, g Transform) Monad {
	return NewLeft(f(x.x))
}

func (x Right) Bimap(f, g Transform) Monad {
	return NewRight(g(x.x))
}

func (x Left) Fold(f, g Transform) Any {
	return f(x.x)
}

func (x Right) Fold(f, g Transform) Any {
	return g(x.x)
}

func (x Left) Sequence(p Point) Any {
	return x.Traverse(Identity, p)
}

func (x Right) Sequence(p Point) Any {
	return x.Traverse(Identity, p)
}

func (x Left) Traverse(f Transform, p Point) Functor {
	return p.Of(NewLeft(x.x)).(Functor)
}

func (x Right) Traverse(f Transform, p Point) Functor {
	return f(x.x).(Functor).Map(func(a Any) Any {
		return NewRight(a)
	})
}
