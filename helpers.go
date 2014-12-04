package wishful

type Thunk func() Any
type Transform func(Any) Any

func Inc(a Any) (r Any) {
	switch obj := a.(type) {
	case int:
		r = obj + 1
	case float32:
		r = obj + 1
	case float64:
		r = obj + 1
	default:
		r = a
	}
	return
}

func Dec(a Any) (r Any) {
	switch obj := a.(type) {
	case int:
		r = obj - 1
	case float32:
		r = obj - 1
	case float64:
		r = obj - 1
	default:
		r = a
	}
	return
}
