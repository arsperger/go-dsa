package interpolation

func ISearch(arr []int, x int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		midx := start + (x-arr[start])*(end-start)/(arr[end]-arr[start])
		mid := arr[midx]

		if x == mid {
			return midx
		}

		if x > arr[midx] {
			start = midx + 1
		}

		if x < mid {
			end = midx - 1
		}
	}
	return -1
}
