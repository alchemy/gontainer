package maps

import "github.com/alchemy/gontainer/pair"

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0
	for k, _ := range m {
		keys[i] = k
		i++
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}

func Elements[K comparable, V any](m map[K]V) []pair.Pair[K, V] {
	elements := make([]pair.Pair[K, V], len(m))
	i := 0
	for k, v := range m {
		elements[i].First, elements[i].Second = k, v
		i++
	}
	return elements
}

func Map[K1, K2 comparable, V1, V2 any](m map[K1]V1, t func(k K1, v V1) (K2, V2)) map[K2]V2 {
	nm := make(map[K2]V2)
	for k, v := range m {
		nk, nv := t(k, v)
		nm[nk] = nv
	}
	return nm
}

func Filter[K comparable, V any](m map[K]V, f func(k K, v V) bool) map[K]V {
	nm := make(map[K]V)
	for k, v := range m {
		if f(k, v) {
			nm[k] = v
		}
	}
	return nm
}
