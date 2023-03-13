package pair

type Pair[T any, U any] struct {
	First  T
	Second U
}

func New[T any, U any](first T, second U) *Pair[T, U] {
	return &Pair[T, U]{first, second}
}

func Make[T any, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{first, second}
}

func Zip[T any, U any](s1 []T, s2 []U) []Pair[T, U] {
	n := len(s1)
	if n > len(s2) {
		n = len(s2)
	}
	zipped := make([]Pair[T, U], n)
	for i := 0; i < n; i++ {
		zipped[i].First, zipped[i].Second = s1[i], s2[i]
	}
	return zipped
}
