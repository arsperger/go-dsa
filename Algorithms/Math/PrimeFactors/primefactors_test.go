package primefactors

import "testing"

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestPrimeFactors(t *testing.T) {

	data := []struct {
		n int
		w []int
	}{
		{126, []int{2, 3, 3, 7}}, {150, []int{2, 3, 5, 5}}, {127, []int{127}},
	}

	for _, idx := range data {
		b := PrimeFactors(idx.n)
		if Equal(idx.w, b) == false {
			t.Errorf("Error, want %-q got %-q", idx.w, b)
		}
	}

}
