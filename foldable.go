package wishful

type Foldable interface {
	Fold(Morphism, Morphism) Any
}
