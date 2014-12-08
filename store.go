package wishful

type Store struct {
	Set Morphism
	Get Value
}

func NewStore(set Morphism, get Value) Store {
	return Store{
		set,
		get,
	}
}

func (x Store) Map(f Morphism) Functor {
	return x.Extend(func(x Store) Any {
		return f(x.Extract())
	})
}

// Derived

func (x Store) Extend(f func(x Store) Any) Store {
	return Store{
		func(y Any) Any {
			fun := NewFunction(f)
			res, _ := fun.Call(Store{
				x.Set,
				func() interface{} {
					return y
				},
			})
			return res
		},
		x.Get,
	}
}

func (x Store) Extract() Any {
	return x.Set(x.Get())
}
