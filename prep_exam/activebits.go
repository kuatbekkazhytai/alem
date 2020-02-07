package piscine

func ActiveBits(n int) uint {
	var x uint
	for n > 0 {
		if n%2 == 1 {
			x++
		}
		n = n / 2
	}
	return x
}
