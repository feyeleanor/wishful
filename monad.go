package wishful

//	http://en.wikipedia.org/wiki/Monad_(functional_programming)

type Step func(v Any) Monad

type Monad interface {
	Chain(Step) Monad
}

type MonadLaws struct {
	x Point
}

func NewMonadLaws(point Point) MonadLaws {
	return MonadLaws{
		x: point,
	}
}

func (o MonadLaws) LeftIdentity(run Morphism) (f, g Transform) {
	f = func(v int) Any {
		a := o.x.Of(v).(Monad)
		return run(a.Chain(func(x Any) Monad {
			return Apply(func(x Any) Any {
				return o.x.Of(x)
			})(x).(Monad)
		}))
	}
	g = func(v int) Any {
		return run(Apply(func(x Any) Any {
			return o.x.Of(x)
		})(v))
	}
	return
}

func (o MonadLaws) RightIdentity(run Morphism) (f, g Transform) {
	f = func(v int) Any {
		a := o.x.Of(v).(Monad)
		return run(a.Chain(func(x Any) Monad {
			return o.x.Of(x).(Monad)
		}))
	}
	g = func(v int) Any {
		return run(o.x.Of(v))
	}
	return
}

func (o MonadLaws) Associativity(run Morphism) (f, g Transform) {
	f = func(v int) Any {
		a := o.x.Of(v).(Monad)
		return run(a.Chain(func(x Any) Monad {
			return o.x.Of(x).(Monad)
		}).Chain(func(x Any) Monad {
			return o.x.Of(x).(Monad)
		}))
	}
	g = func(v int) Any {
		a := o.x.Of(v).(Monad)
		return run(a.Chain(func(x Any) Monad {
			b := o.x.Of(x).(Monad)
			return b.Chain(func(x Any) Monad {
				return o.x.Of(x).(Monad)
			})
		}))
	}
	return
}
