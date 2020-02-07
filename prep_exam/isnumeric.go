package piscine

func IsNumeric(str string) bool {

	for _, r := range str {
		if r >= '0' && r <= '9' {
			return true
		}
	}

	return false
}
