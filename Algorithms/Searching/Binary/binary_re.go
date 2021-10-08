package bsearch_re

func bsearch(arrs []int, x int, left int, right int) int {

	midx := int((left + right) / 2)
	mid := arrs[midx]

	if x == mid {
		return midx
	}

	if left == right {
		return -1
	}

	if x > mid {
		return bsearch(arrs, x, midx+1, right)
	}

	return bsearch(arrs, x, left, midx-1)
}
