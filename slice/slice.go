package slice

import "github.com/alchemy/gontainer/maps"

type void struct{}

var null = void{}

func FindNonUnique[T comparable](slice []T) []T {
	seen := make(map[T]int)
	for _, elem := range slice {
		if _, ok := seen[elem]; !ok {
			seen[elem]++
		}
	}
	return maps.Keys(maps.Filter(seen, func(k T, v int) bool { return v > 1 }))
}

func FindUnique[T comparable](slice []T) []T {
	seen := make(map[T]int)
	for _, elem := range slice {
		if _, ok := seen[elem]; !ok {
			seen[elem]++
		}
	}
	return maps.Keys(maps.Filter(seen, func(k T, v int) bool { return v == 1 }))
}

func RemoveDuplicate[T comparable](slice []T) []T {
	seen := make(map[T]void)
	list := []T{}
	for _, elem := range slice {
		if _, ok := seen[elem]; !ok {
			seen[elem] = null
			list = append(list, elem)
		}
	}
	return list
}

func Index[T comparable](slice []T, elem T) int {
	for i := range slice {
		if slice[i] == elem {
			return i
		}
	}
	return -1
}

func In[T comparable](elem T, slice []T) bool {
	return Index(slice, elem) >= 0
}

func Remove[T comparable](slice []T, elements ...T) []T {
	for i := range slice {
		if In(slice[i], elements) {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func Map[T, U any](s []T, t func(i int, v T) U) []U {
	ns := make([]U, len(s))
	for i, v := range s {
		ns[i] = t(i, v)
	}
	return ns
}

func Filter[T any](slice []T, f func(i int, v T) bool) []T {
	ns := make([]T, 0)
	for i, v := range slice {
		if f(i, v) {
			ns = append(ns, v)
		}
	}
	return ns
}
