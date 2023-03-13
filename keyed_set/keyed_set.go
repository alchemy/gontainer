package keyed_set

type KeyProvider interface {
	Key() any
}

type KeyedSet[T any] struct {
	elements map[any]T
}

func New[T any](size int) KeyedSet[T] {
	s := KeyedSet[T]{}
	s.elements = make(map[any]T, size)
	return s
}

func key[T any](elem *T) any {
	var i any = elem
	k, ok := i.(KeyProvider)
	if ok {
		return k.Key()
	}
	return i
}

func (s KeyedSet[T]) Add(elem T) {
	s.elements[key(&elem)] = elem
}

func (s KeyedSet[T]) Remove(elem T) {
	delete(s.elements, key(&elem))
}

func (s KeyedSet[T]) Contains(elem T) bool {
	_, ok := s.elements[key(&elem)]
	return ok
}

func (s KeyedSet[T]) Size() int {
	return len(s.elements)
}

func (s KeyedSet[T]) IsSubsetOf(t KeyedSet[T]) bool {
	for _, elem := range s.elements {
		if !t.Contains(elem) {
			return false
		}
	}
	return true
}

func (s KeyedSet[T]) Elements() []T {
	elements := make([]T, len(s.elements))
	for _, elem := range s.elements {
		elements = append(elements, elem)
	}
	return elements
}

func Union[T any](s KeyedSet[T], t KeyedSet[T]) KeyedSet[T] {
	u := KeyedSet[T]{}
	for _, e := range s.elements {
		u.Add(e)
	}
	for _, e := range t.elements {
		u.Add(e)
	}
	return u
}

func Intersection[T any](s KeyedSet[T], t KeyedSet[T]) KeyedSet[T] {
	i := KeyedSet[T]{}
	if s.Size() <= t.Size() {
		for _, e := range s.elements {
			if t.Contains(e) {
				i.Add(e)
			}
		}
	} else {
		for _, e := range t.elements {
			if s.Contains(e) {
				i.Add(e)
			}
		}
	}
	return i
}

func Difference[T any](s KeyedSet[T], t KeyedSet[T]) KeyedSet[T] {
	d := KeyedSet[T]{}
	for _, e := range s.elements {
		if !t.Contains(e) {
			d.Add(e)
		}
	}
	return d
}
