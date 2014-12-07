package wishful

// http://en.wikipedia.org/wiki/Functor

type Functor interface {
	Map(f Morphism) Functor
}

type FunctorLaws struct {
	x Point
}

func NewFunctorLaws(point Point) FunctorLaws {
	return FunctorLaws{
		x: point,
	}
}

func (o FunctorLaws) Identity(run Morphism) (f, g Transform) {
	f = func(v int) Any {
		a := o.x.Of(v).(Functor)
		return run(a.Map(Identity))
	}
	g = func(v int) Any {
		return run(o.x.Of(v))
	}
	return
}

func (o FunctorLaws) Composition(run Morphism) (f, g Transform) {
	f = func(v int) Any {
		a := o.x.Of(v).(Functor)
		return run(a.Map(Compose(Identity)(Identity)))
	}
	g = func(v int) Any {
		a := o.x.Of(v).(Functor)
		return run(a.Map(Identity).Map(Identity))
	}
	return
}
