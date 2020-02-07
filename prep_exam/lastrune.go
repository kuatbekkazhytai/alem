package piscine

func LastRune(s string) rune {
	l := 0
	runes := []rune(s)
	for range s {
		l++
	}
	return runes[l-1]
}