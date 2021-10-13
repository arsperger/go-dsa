package seive

import "math"

func FindPrimes(num int) []int {

	compose := make(map[int]bool, num)
	primes := []int{}

	for i := 4; i <= num; i += 2 {
		compose[i] = true
	}

	nprime := 3
	stop := int(math.Sqrt(float64(num)))

	for nprime <= stop {

		for i := nprime * 2; i <= num; i += nprime {
			compose[i] = true
		}

		nprime = nprime + 2
		if compose[nprime] {
			nprime = nprime + 2
		}
	}

	for i := 2; i <= num; i += 1 {
		if !compose[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

// n * log(log(n))
