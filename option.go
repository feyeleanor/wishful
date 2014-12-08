package wishful

type Option interface {
	Of(Any) Point
	Empty() Monoid
	Ap(Applicative) Applicative
	Chain(Step) Monad
	Concat(Semigroup) Semigroup
	Foldable
	Map(Morphism) Functor
	GetOrElse(Value) Any
	OrElse(Option) Option
}

type none struct {
}

func None() none {
	return none{}
}

func (x none) Of(v Any) Point {
	return Some(v)
}

func (x none) Empty() Monoid {
	return None()
}

func (x none) Ap(v Applicative) Applicative {
	return x
}

func (x none) Chain(f Step) Monad {
	return x
}

func (x none) Fold(f, g Morphism) Any {
	return g(nil)
}

func (x none) Map(f Morphism) Functor {
	return x
}

func (x none) Concat(y Semigroup) Semigroup {
	return x
}

// Derived

func (x none) GetOrElse(f Value) Any {
	return f()
}

func (x none) OrElse(y Option) Option {
	return y
}

type some struct {
	x interface{}
}

func Some(x Any) some {
	return some{
		x: x,
	}
}

func (x some) Of(v Any) Point {
	return Some(v)
}

func (x some) Empty() Monoid {
	return None()
}

func (x some) Ap(v Applicative) Applicative {
	return fromMonadToApplicativeAp(x, v)
}

func (x some) Chain(f Step) Monad {
	return f(x.x)
}

func (x some) Fold(f, g Morphism) Any {
	return f(x.x)
}

func (x some) Map(f Morphism) Functor {
	res := x.Chain(func(v Any) Monad {
		return Some(f(v))
	})
	return res.(Functor)
}

func (x some) Concat(y Semigroup) Semigroup {
	return concat(x, y)
}

// Derived

func (x some) GetOrElse(f Value) Any {
	return x.x
}

func (x some) OrElse(y Option) Option {
	return some{}.Of(x.x).(Option)
}
