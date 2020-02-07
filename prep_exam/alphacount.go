package piscine

func AlphaCount(str string) int {
	l := 0
	for _, j := range str {
		if IsAlpha(j) {
			l++
		}

	}
	return l
}
