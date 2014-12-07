package wishful

type Either interface {
	Of(Any) Point
	Ap(Applicative) Applicative
	Chain(Step) Monad
	Concat(Semigroup) Semigroup
	Map(Morphism) Functor
	Bimap(f, g Morphism) Monad
	Fold(f, g Morphism) Any
	Swap() Monad
	Sequence(Point) Any
	Traverse(Morphism, Point) Functor
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

func (x Left) Chain(f Step) Monad {
	return x
}

func (x Right) Chain(f Step) Monad {
	return f(x.x)
}

func (x Left) Map(f Morphism) Functor {
	return x
}

func (x Right) Map(f Morphism) Functor {
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

func (x Left) Bimap(f, g Morphism) Monad {
	return NewLeft(f(x.x))
}

func (x Right) Bimap(f, g Morphism) Monad {
	return NewRight(g(x.x))
}

func (x Left) Fold(f, g Morphism) Any {
	return f(x.x)
}

func (x Right) Fold(f, g Morphism) Any {
	return g(x.x)
}

func (x Left) Sequence(p Point) Any {
	return x.Traverse(Identity, p)
}

func (x Right) Sequence(p Point) Any {
	return x.Traverse(Identity, p)
}

func (x Left) Traverse(f Morphism, p Point) Functor {
	return p.Of(NewLeft(x.x)).(Functor)
}

func (x Right) Traverse(f Morphism, p Point) Functor {
	return f(x.x).(Functor).Map(func(a Any) Any {
		return NewRight(a)
	})
}
