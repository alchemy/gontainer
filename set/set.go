package set

type void struct{}

var null = void{}

type Set[T comparable] struct {
	elements map[T]void
}

func New[T comparable](size int) Set[T] {
	return Set[T]{
		elements: make(map[T]void, size),
	}
}

func (s *Set[T]) Add(elements ...T) {
	if s.elements == nil {
		s.elements = make(map[T]void)
	}
	for _, e := range elements {
		s.elements[e] = null
	}
}

func (s *Set[T]) Remove(elem T) {
	delete(s.elements, elem)
}

func (s *Set[T]) Contains(elem T) bool {
	_, ok := s.elements[elem]
	return ok
}

func (s *Set[T]) Clear() {
	s.elements = make(map[T]void)
}

func (s *Set[T]) Size() int {
	return len(s.elements)
}

func (s *Set[T]) IsSubsetOf(t Set[T]) bool {
	for elem := range s.elements {
		if !t.Contains(elem) {
			return false
		}
	}
	return true
}

func (s *Set[T]) IsEqualTo(t Set[T]) bool {
	diff := Difference(*s, t)
	return diff.Size() == 0
}

func (s *Set[T]) Elements() []T {
	elements := make([]T, len(s.elements))
	i := 0
	for elem := range s.elements {
		elements[i] = elem
		i++
	}
	return elements
}

func Union[T comparable](s, t Set[T]) Set[T] {
	u := Set[T]{}
	for e := range s.elements {
		u.Add(e)
	}
	for e := range t.elements {
		u.Add(e)
	}
	return u
}

func Intersection[T comparable](s Set[T], t Set[T]) Set[T] {
	i := Set[T]{}
	if s.Size() <= t.Size() {
		for e := range s.elements {
			if t.Contains(e) {
				i.Add(e)
			}
		}
	} else {
		for e := range t.elements {
			if s.Contains(e) {
				i.Add(e)
			}
		}
	}
	return i
}

func Difference[T comparable](s Set[T], t Set[T]) Set[T] {
	d := Set[T]{}
	for e := range s.elements {
		if !t.Contains(e) {
			d.Add(e)
		}
	}
	return d
}

func Map[S comparable, T comparable](s Set[S], f func(S) T) Set[T] {
	m := Set[T]{}
	for elem := range s.elements {
		mapped := f(elem)
		m.Add(mapped)
	}
	return m
}

func Filter[T comparable](s Set[T], f func(T) bool) Set[T] {
	filtered := Set[T]{}
	for elem := range s.elements {
		if f(elem) {
			filtered.Add(elem)
		}
	}
	return filtered
}
