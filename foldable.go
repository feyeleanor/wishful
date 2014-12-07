package wishful

type Foldable interface {
	Fold(f, g Morphism) Any
}
