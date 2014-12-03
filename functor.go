package wishful

type Functor interface {
	Map(f Transform) Functor
}

type FunctorLaws struct {
	x Point
}

func NewFunctorLaws(point Point) FunctorLaws {
	return FunctorLaws{
		x: point,
	}
}

func (o FunctorLaws) Identity(run Transform) (func(v int) Any, func(v int) Any) {
	f := func(v int) Any {
		a := o.x.Of(v).(Functor)
		return run(a.Map(Identity))
	}
	g := func(v int) Any {
		return run(o.x.Of(v))
	}
	return f, g
}

func (o FunctorLaws) Composition(run Transform) (func(v int) Any, func(v int) Any) {
	f := func(v int) Any {
		a := o.x.Of(v).(Functor)
		return run(a.Map(Compose(Identity)(Identity)))
	}
	g := func(v int) Any {
		a := o.x.Of(v).(Functor)
		return run(a.Map(Identity).Map(Identity))
	}
	return f, g
}
