package keyed_set

import "testing"

type StructType struct {
	i int
	s string
}

type StructKeyProvider struct {
	i int
	s string
}

func (st *StructKeyProvider) Key() any {
	return st.s
}

func TestNew(t *testing.T) {
	keys := []string{
		"k0",
		"k1",
		"k2",
		"k3",
		"k4",
		"k4",
		"k3",
		"k2",
		"k1",
		"k0",
	}
	s1 := New[StructType](len(keys))
	for i, k := range keys {
		st := StructType{i, k}
		s1.Add(st)
	}
	if len(s1.elements) != len(keys) {
		t.Fatalf("len(s.elements) = %d: must be %d!", len(s1.elements), len(keys))
	}

	s2 := New[StructKeyProvider](len(keys))
	for i, k := range keys {
		st := StructKeyProvider{i, k}
		s2.Add(st)
	}
	if len(s2.elements) != len(keys)/2 {
		t.Fatalf("len(s.elements) = %d: must be %d!", len(s2.elements), len(keys)/2)
	}
}
