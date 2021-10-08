package lsearch

func Lsearch(arrs []int, x int) int {
	if arrs == nil {
		return -1
	}
	for m, n := range arrs {
		if x == n {
			return m
		}
	}
	return -1
}
