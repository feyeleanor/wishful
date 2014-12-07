package wishful

type Contract func(Any) Promise

type Promise struct {
	Fork func(resolve Morphism) Any
}

func NewPromise(f func(resolve Morphism) Any) Promise {
	return Promise{Fork: f}
}

func (x Promise) Of(v Any) Point {
	return Promise{func(resolve Morphism) Any {
		return resolve(v)
	}}
}

func (x Promise) Ap(v Applicative) Applicative {
	return Promise{func(resolve Morphism) Any {
		return x.Fork(func(f Any) Any {
			fun := v.(Functor)
			pro := fun.Map(func(x Any) Any {
				fun := NewFunction(f)
				res, _ := fun.Call(x)
				return res
			})
			return pro.(Promise).Fork(resolve)
		})
	}}
}

func (x Promise) Chain(f Step) Monad {
	return Promise{func(resolve Morphism) Any {
		return x.Fork(func(a Any) Any {
			p := f(a).(Promise)
			return p.Fork(resolve)
		})
	}}
}

func (x Promise) Map(f Morphism) Functor {
	return Promise{func(resolve Morphism) Any {
		return x.Fork(func(a Any) Any {
			return resolve(f(a))
		})
	}}
}

func (x Promise) Extract() Any {
	return x.Fork(Identity)
}

func (x Promise) Extend(f func(p Comonad) Any) Comonad {
	return x.Map(func(y Any) Any {
		fun := NewFunction(f)
		res, _ := fun.Call(x.Of(y))
		return res
	}).(Comonad)
}
