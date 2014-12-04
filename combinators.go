package wishful

// A combinator
func Apply(f Transform) Transform {
	return func(x Any) Any {
		return f(x)
	}
}

// B combinator
func Compose(f Transform) func(Transform) Transform {
	return func(g Transform) Transform {
		return func(a Any) Any {
			return f(g(a))
		}
	}
}

// K combinator
func Constant(a Any) Transform {
	return func(b Any) Any {
		return a
	}
}
func ConstantNoArgs(a Any) Thunk {
	return func() Any {
		return a
	}
}

// I combinator
func Identity(x Any) Any {
	return x
}

// T combinator
func Thrush(x Any) func(f Transform) Any {
	return func(f Transform) Any {
		return f(x)
	}
}
