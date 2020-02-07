package piscine

func Nauuo(plus, minus, rand int) string {
		if plus == minus && rand == 0 {
			return string('0')
		} else if (plus - minus <= rand&& plus - minus >=0) || (minus - plus <=rand&& minus - plus >= 0 )  {
		return string('?')
	}  else if plus - minus > rand  {
		return string('+')
	} else if minus - plus > rand {
		return string('-')
	}
	return string('a')
}