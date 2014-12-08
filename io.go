package wishful

type IO struct {
	UnsafePerform Value
}

func NewIO(unsafe Value) IO {
	return IO{
		UnsafePerform: unsafe,
	}
}

func (x IO) Of(v Any) Point {
	return NewIO(func() interface{} {
		return v
	})
}

func (x IO) Ap(v Applicative) Applicative {
	res := x.Chain(func(f Any) Monad {
		fun := v.(Functor)
		res := fun.Map(func(x Any) Any {
			fun := NewFunction(f)
			res, _ := fun.Call(x)
			return res
		})
		return res.(Monad)
	})
	return res.(Applicative)
}

func (x IO) Chain(f Step) Monad {
	return NewIO(func() interface{} {
		io := f(x.UnsafePerform()).(IO)
		return io.UnsafePerform()
	})
}

func (x IO) Map(f Morphism) Functor {
	res := x.Chain(func(x Any) Monad {
		return IO{func() interface{} {
			return f(x)
		}}
	})
	return res.(Functor)
}
