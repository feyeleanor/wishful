package wishful

type Comonad interface {
	Extend(f func(a Comonad) Any) Comonad
	Extract() Any
}

type ComonadLaws struct {
	x Point
}

func NewComonadLaws(point Point) ComonadLaws {
	return ComonadLaws{
		x: point,
	}
}

func (o ComonadLaws) Identity(run Morphism) (f, g Transform) {
	f = func(v int) Any {
		a := o.x.Of(v).(Comonad)
		return run(a.Extend(func(x Comonad) Any {
			return a.Extract()
		}))
	}
	g = func(v int) Any {
		return run(o.x.Of(v))
	}
	return
}

func (o ComonadLaws) Composition(run Morphism) (f, g Transform) {
	extract := func(y Comonad) Any {
		return y.Extract()
	}
	f = func(v int) Any {
		a := o.x.Of(v).(Comonad)
		b := Compose(func(x Any) Any {
			return a.Extract()
		})(func(x Any) Any {
			return a.Extend(extract)
		})

		return run(b(a))
	}
	g = func(v int) Any {
		a := o.x.Of(v).(Comonad)
		return run(extract(a))
	}
	return
}

func (o ComonadLaws) Associativity(run Morphism) (f, g Transform) {
	extract := func(y Comonad) Any {
		return y.Extract()
	}
	duplicate := func(x Comonad) Morphism {
		return func(y Any) Any {
			return x.Extend(extract)
		}
	}
	f = func(v int) Any {
		a := o.x.Of(v).(Comonad)
		b := Compose(duplicate(a))(duplicate(a))
		return run(b(a))
	}
	g = func(v int) Any {
		a := o.x.Of(v).(Functor)
		c := a.(Comonad)
		return run(a.Map(Compose(duplicate(c))(duplicate(c))).(Comonad).Extract())
	}
	return
}
