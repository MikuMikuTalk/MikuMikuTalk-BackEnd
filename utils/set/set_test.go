package set

import (
	"testing"
)

func TestUnion(t *testing.T) {
	slice1 := NewSet(1, 2, 3)
	slice2 := NewSet(3, 4, 5)
	// want := NewSet(1, 2, 3, 4, 5)
	got := Union(slice1, slice2)
	got2 := InterSet(slice1, slice2)
	got3 := Difference(slice1, slice2)
	t.Log(got)
	t.Log(got2)
	t.Log(got3)

	slice3 := NewSet(1, 2, 3, 4)
	slice4 := NewSet(1, 2, 3, 4)
	t.Log(slice3.Equal(slice4))

	slice4.Add(1, 2, 3, 4, 5, 6, 7, 8)
	t.Log(slice4)
}
