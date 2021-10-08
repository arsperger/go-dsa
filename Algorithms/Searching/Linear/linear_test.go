package lsearch

import "testing"

func TestLseatch(t *testing.T) {

	data := []struct {
		arr []int
		x   int
		w   int
	}{
		{[]int{1, 23, 45, 2, 65, 22, 43}, 45, 2},
		{[]int{34, 55, 65, 23, 2, 43, 5}, 77, -1},
		{[]int{}, 77, -1},
	}

	for _, c := range data {
		got := Lsearch(c.arr, c.x)
		if got != c.w {
			t.Errorf("Error got %d, for searched %d", got, c.x)
		}
	}
}
