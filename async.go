package wishful

var (
	EitherPromise EitherT = NewEitherT(Promise{})
)

func Async(f Contract) func(x Any) EitherT {
	return func(x Any) EitherT {
		return EitherPromise.From(
			NewPromise(
				func(resolve Transform) Any {
					fun := NewFunction(f)
					res, _ := fun.Call(x)
					return res.(Promise).Fork(func(x Any) Any {
						return resolve(NewRight(x))
					})
				},
			),
		)
	}
}
