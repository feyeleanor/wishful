package wishful

//	http://en.wikipedia.org/wiki/Element_(category_theory)

type Point interface {
	Of(v Any) Point
}
