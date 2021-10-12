package primefactors

import "math"

func PrimeFactors(n int) []int {

	f := []int{}

	for n%2 == 0 {
		f = append(f, 2)
		n = n / 2
	}

	i := 3
	maxN := int(math.Sqrt(float64(n)))

	for i <= maxN {

		for n%i == 0 {
			f = append(f, i)
			n = n / i
			maxN = int(math.Sqrt(float64(n)))
		}

		i = i + 2
	}

	if n > 1 {
		f = append(f, n)
	}

	return f

}
