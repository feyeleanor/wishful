package wishful

type Foldable interface {
	Fold(f, g Transform) Any
}
