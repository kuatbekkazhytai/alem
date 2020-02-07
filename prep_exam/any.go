package piscine

func Any(f func(string) bool, arr []string) bool {
	for _, j := range arr {
		if f(j) {
			return true
		}
	}
	return false
}
