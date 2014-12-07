package wishful

type EitherT struct {
	m   Point
	Run Any
}

func NewEitherT(m Point) EitherT {
	return EitherT{
		m:   m,
		Run: Empty{},
	}
}

func (x EitherT) Of(v Any) Point {
	return EitherT{
		m:   x.m,
		Run: x.m.Of(NewRight(v)),
	}
}

func (x EitherT) From(v Any) EitherT {
	return EitherT{
		m:   x.m,
		Run: v,
	}
}

func (x EitherT) Fold(f, g Morphism) Any {
	return x.Run.(Monad).Chain(func(o Any) Monad {
		return x.m.Of(o.(Foldable).Fold(f, g)).(Monad)
	})
}

func (x EitherT) Ap(v Applicative) Applicative {
	mon := x.Chain(func(f Any) Monad {
		return v.(Functor).Map(func(x Any) Any {
			fun := NewFunction(f)
			res, _ := fun.Call(x)
			return res
		}).(Monad)
	})
	return mon.(Applicative)
}

func (x EitherT) Chain(f Step) Monad {
	mon := x.Run.(Monad)
	tra := EitherT{
		m: x.m,
		Run: mon.Chain(func(y Any) Monad {
			return y.(Foldable).Fold(
				func(v Any) Any {
					return x.m.Of(NewLeft(v))
				},
				func(v Any) Any {
					return f(v).(EitherT).Run
				},
			).(Monad)
		}),
	}
	return tra
}

func (x EitherT) Map(f Morphism) Functor {
	mon := x.Chain(func(y Any) Monad {
		app := NewEitherT(x.m).Of(f(y))
		return app.(Monad)
	})
	return mon.(Functor)
}

// Derived

func (x EitherT) Swap() Monad {
	return x.Fold(
		func(v Any) Any {
			return NewRight(v)
		},
		func(v Any) Any {
			return NewLeft(v)
		},
	).(Monad)
}

func (x EitherT) Bimap(f, g Morphism) Monad {
	return x.Fold(
		func(v Any) Any {
			return NewLeft(f(v))
		},
		func(v Any) Any {
			return NewRight(g(v))
		},
	).(Monad)
}
