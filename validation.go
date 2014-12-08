package wishful

type Validation interface {
	Of(Any) Point
	Ap(Applicative) Applicative
	Monad
	Concat(Semigroup) Semigroup
	Map(Morphism) Functor
	Foldable
	Bimap(Morphism, Morphism) Monad
}

type failure struct {
	x Any
}

func Failure(x Any) failure {
	return failure{
		x: x,
	}
}

func (x failure) Of(v Any) Point {
	return Success(v)
}

func (x failure) Ap(v Applicative) Applicative {
	return v.(Foldable).Fold(
		func(y Any) Any {
			return Failure(concatAnyvals(x.x)(y))
		},
		func(y Any) Any {
			return Failure(x.x)
		},
	).(Applicative)
}

func (x failure) Chain(f Step) Monad {
	return x
}

func (x failure) Map(f Morphism) Functor {
	return x.Bimap(Identity, f).(Functor)
}

func (x failure) Concat(y Semigroup) Semigroup {
	a := y.(Validation)
	b := a.Bimap(
		concatAnyvals(x.x),
		Identity,
	)
	return b.(Semigroup)
}

// Derived

func (x failure) Fold(f, g Morphism) Any {
	return f(x.x)
}

func (x failure) Bimap(f, g Morphism) Monad {
	return Failure(f(x.x))
}

type success struct {
	x Any
}

func Success(x Any) success {
	return success{
		x: x,
	}
}

func (x success) Of(v Any) Point {
	return Success(v)
}

func (x success) Ap(v Applicative) Applicative {
	return v.(Functor).Map(func(g Any) Any {
		fun := NewFunction(x.x)
		res, _ := fun.Call(g)
		return res
	}).(Applicative)
}

func (x success) Chain(f Step) Monad {
	return f(x.x)
}

func (x success) Map(f Morphism) Functor {
	return x.Bimap(Identity, f).(Functor)
}

func (x success) Concat(y Semigroup) Semigroup {
	a := y.(Functor)
	b := a.Map(concatAnyvals(x.x))
	return b.(Semigroup)
}

// Derived

func (x success) Fold(f, g Morphism) Any {
	return g(x.x)
}

func (x success) Bimap(f, g Morphism) Monad {
	return Success(g(x.x))
}
