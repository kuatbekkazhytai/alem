package piscine

func IsAlpha(a rune) bool {
	if (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') {
		return true
	}
	return false
}
