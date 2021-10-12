package fastpowering

func Fpowering(x, y int) int {
	if y == 0 {
		return 1
	}
	tmp := Fpowering(x, y/2)

	// even
	if y%2 == 0 {
		return tmp * tmp
	}

	// odd
	return tmp * tmp * x
}
