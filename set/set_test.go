package set

import (
	"log"
	"testing"
)

type Elem struct {
	i int
	s string
}

func TestSet(t *testing.T) {
	var s Set[Elem]
	//s2 := New[Elem](4)
	s.Add(Elem{1, "alpha"})
	log.Println(s.Elements())
}
