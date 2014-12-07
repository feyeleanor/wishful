package wishful

type Option interface {
	Of(v Any) Point
	Empty() Monoid
	Ap(v Applicative) Applicative
	Chain(f func(v Any) Monad) Monad
	Concat(y Semigroup) Semigroup
	Fold(f Morphism, g Thunk) Any
	Map(f Morphism) Functor
	GetOrElse(f Thunk) Any
	OrElse(y Option) Option
}

type Some struct {
	x Any
}

type None struct {
}

func NewSome(x Any) Some {
	return Some{
		x: x,
	}
}

func NewNone() None {
	return None{}
}

func (x Some) Of(v Any) Point {
	return NewSome(v)
}

func (x None) Of(v Any) Point {
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

func (x Some) Chain(f Step) Monad {
	return f(x.x)
}

func (x None) Chain(f Step) Monad {
	return x
}

func (x Some) Fold(f Morphism, g Thunk) Any {
	return f(x.x)
}

func (x None) Fold(f Morphism, g Thunk) Any {
	return g()
}

func (x Some) Map(f Morphism) Functor {
	res := x.Chain(func(v Any) Monad {
		return NewSome(f(v))
	})
	return res.(Functor)
}

func (x None) Map(f Morphism) Functor {
	return x
}

func (x Some) Concat(y Semigroup) Semigroup {
	return concat(x, y)
}

func (x None) Concat(y Semigroup) Semigroup {
	return x
}

// Derived

func (x Some) GetOrElse(f Thunk) Any {
	return x.x
}

func (x None) GetOrElse(f Thunk) Any {
	return f()
}

func (x Some) OrElse(y Option) Option {
	return Some{}.Of(x.x).(Option)
}

func (x None) OrElse(y Option) Option {
	return y
}
