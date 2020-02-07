package piscine

func Atoi(s string) int {
	var valid bool
	var neg bool
	num := 0
	if s[0] == '-' {
		if s[1] == '-' {
			valid = false
		} else {
			valid = true
			neg = true
		}
	} else if s[0] == '+' {
		if s[1] == '+' {
			valid = false
		} else {
			valid = true

		}
	} else {
		valid = true
	}

	if valid {
		for _, j := range s {

			num = num * 10
			if j == '0' {
				num = num + 0
			}
			if j == '1' {
				num = num + 1
			}
			if j == '2' {
				num = num + 2
			}
			if j == '3' {
				num = num + 3
			}
			if j == '4' {
				num = num + 4
			}
			if j == '5' {
				num = num + 5
			}
			if j == '6' {
				num = num + 6
			}
			if j == '7' {
				num = num + 7
			}
			if j == '8' {
				num = num + 8
			}
			if j == '9' {
				num = num + 9
			}
			if j == 32 {
				return 0
			}

		}
	} else {
		return 0
	}
	if neg {
		num = -1 * num
	}

	return num
}
