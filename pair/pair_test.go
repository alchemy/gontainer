package pair

import (
	"testing"
	"time"
)

func TestPair(t *testing.T) {
	now := time.Now()
	p1 := New("a", 1)
	p2 := New("b", now)

	if p1.First != "a" {
		t.Errorf("p1.First must be \"a\"")
	}
	if p1.Second != 1 {
		t.Errorf("p1.Second must be 1")
	}

	if p2.First != "b" {
		t.Errorf("p2.First must be \"b\"")
	}
	if p2.Second != now {
		t.Errorf("p2.Second must be %v", now)
	}
}

func TestZip(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []string{"a", "b", "c"}

	pairs := Zip(a, b)
	if len(pairs) != 3 {
		t.Errorf("len(pairs) must be 3")
	}
	for i, p := range pairs {
		if p.First != a[i] {
			t.Errorf("pairs[%d].First must be %v", i, a[i])
		}
		if p.Second != b[i] {
			t.Errorf("pairs[%d].Second must be %v", i, b[i])
		}
	}
}
