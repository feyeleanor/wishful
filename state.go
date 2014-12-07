package wishful

type State struct {
	Run func(a Any) (Any, Any)
}

func (x State) Of(y Any) Point {
	return State{func(z Any) (Any, Any) {
		return y, z
	}}
}

func (x State) Ap(v Applicative) Applicative {
	app := x.Chain(func(f Any) Monad {
		fun := v.(Functor)
		app := fun.Map(func(x Any) Any {
			fun := NewFunction(f)
			res, _ := fun.Call(x)
			return res
		})
		return app.(Monad)
	})
	return app.(Applicative)
}

func (x State) Chain(f Step) Monad {
	return State{func(s Any) (Any, Any) {
		a, b := x.Run(s)
		fun := NewFunction(f)
		res, _ := fun.Call(a)
		return res.(State).Run(b)
	}}
}

func (x State) Map(f Morphism) Functor {
	fun := x.Chain(func(y Any) Monad {
		return x.Of(f(y)).(Monad)
	})
	return fun.(Functor)
}

// Derived

func (x State) EvalState(y Any) Any {
	a, _ := x.Run(y)
	return a
}

func (x State) ExecState(y Any) Any {
	_, b := x.Run(y)
	return b
}

func (x State) Get() State {
	return State{func(z Any) (Any, Any) {
		return z, z
	}}
}

func (x State) Modify(f Morphism) State {
	return State{func(z Any) (Any, Any) {
		fun := NewFunction(f)
		res, _ := fun.Call(z)
		return nil, res
	}}
}

func (x State) Put(a Any, b Any) State {
	return State{func(z Any) (Any, Any) {
		return a, b
	}}
}
