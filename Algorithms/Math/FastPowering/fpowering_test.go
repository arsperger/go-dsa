package fastpowering

import "testing"

func TestFpowering(t *testing.T) {
	data := []struct {
		x int
		y int
		w int
	}{
		{2, 2, 4}, {3, 4, 81},
	}

	for _, i := range data {
		got := Fpowering(i.x, i.y)
		if got != i.w {
			t.Errorf("%d ^ %d is not equal to %d", i.x, i.y, i.w)
		}
	}
}
