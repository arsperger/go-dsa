package bsearch

func Bsearch(arrs []int, x int) int {

	left := 0
	right := len(arrs) - 1

	for left <= right {
		midx := int((left + right) / 2)
		mid := arrs[midx]

		if x == mid {
			return midx
		}

		if x > mid {
			left = midx + 1
		}
		if x < mid {
			right = midx - 1
		}

	}

	return -1
}
