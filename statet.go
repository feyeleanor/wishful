package wishful

type StateT struct {
	m   Point
	Run func(x Any) Point
}

func NewStateT(m Point) StateT {
	return StateT{
		m: m,
		Run: func(x Any) Point {
			return nil
		},
	}
}

func (x StateT) Lift(m Functor) StateT {
	return StateT{
		m: x.m,
		Run: func(b Any) Point {
			return m.Map(func(c Any) Any {
				return Tuple2{c, b}
			}).(Point)
		},
	}
}

func (x StateT) Of(a Any) Point {
	return StateT{
		m: x.m,
		Run: func(b Any) Point {
			return x.m.Of(Tuple2{a, b})
		},
	}
}

func (x StateT) Chain(f Step) Monad {
	return StateT{
		m: x.m,
		Run: func(b Any) Point {
			result := x.Run(b)
			return result.(Monad).Chain(func(t Any) Monad {
				tup := t.(Tuple2)
				fun := NewFunction(f)
				res, _ := fun.Call(tup[0])
				return res.(StateT).Run(tup[1]).(Monad)
			}).(Point)
		},
	}
}

func (x StateT) Get() StateT {
	return StateT{
		m: x.m,
		Run: func(b Any) Point {
			return x.m.Of(Tuple2{b, b})
		},
	}
}

func (x StateT) Modify(f Transform) StateT {
	return StateT{
		m: x.m,
		Run: func(b Any) Point {
			fun := NewFunction(f)
			res, _ := fun.Call(b)
			return x.m.Of(Tuple2{Empty{}, res})
		},
	}
}

func (x StateT) Put(v Any) StateT {
	return x.Modify(func(a Any) Any {
		return v
	})
}

func (x StateT) EvalState(s Any) Any {
	return x.Run(s).(Functor).Map(func(t Any) Any {
		return t.(Tuple2)[0]
	})
}

func (x StateT) ExecState(s Any) Any {
	return x.Run(s).(Functor).Map(func(t Any) Any {
		return t.(Tuple2)[1]
	})
}

func (x StateT) Map(f Transform) Functor {
	return x.Chain(func(a Any) Monad {
		fun := NewFunction(f)
		res, _ := fun.Call(a)
		return x.Of(res).(Monad)
	}).(Functor)
}

func (x StateT) Ap(a Applicative) Applicative {
	return x.Chain(func(f Any) Monad {
		return a.(Functor).Map(func(b Any) Any {
			fun := NewFunction(f)
			res, _ := fun.Call(b)
			return res
		}).(Monad)
	}).(Applicative)
}
