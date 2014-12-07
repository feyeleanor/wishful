package wishful

type Endo struct {
	Fork Morphism
}

func NewEndo(x Morphism) Endo {
	return Endo{
		Fork: x,
	}
}

func (x Endo) Of(v Any) Point {
	return NewEndo(func(x Any) Any {
		return v
	})
}

func (x Endo) Empty() Monoid {
	return NewEndo(func(v Any) Any {
		return v
	})
}

func (x Endo) Concat(y Semigroup) Semigroup {
	return NewEndo(func(v Any) Any {
		a := y.(Endo)
		return x.Fork(a.Fork(v))
	})
}

func (x Endo) Map(f Morphism) Functor {
	return NewEndo(func(v Any) Any {
		return f(x.Fork(v))
	})
}
