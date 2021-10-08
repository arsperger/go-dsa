package gcd

import "testing"

func TestGetGCD(t *testing.T) {

	data := []struct {
		a int
		b int
		w int
	}{
		{0, 0, 0}, {2, 0, 2}, {0, 2, 2}, {12, 4, 4}, {252, 105, 21}, {-462, -1071, -21},
	}

	for _, d := range data {
		if got := GetGCD(d.a, d.b); got != d.w {
			t.Errorf("failed %d, got: %d, want: %d", d.a, got, d.w)
		}
	}

}
