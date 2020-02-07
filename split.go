package student

func Length(str string) int {
	num := 0
	for range str {
		num++
	}
	return num
}

func ArrayLength(str, charset string) int {
	num := 0
	for i := 0; i < Length(str); i++ {
		if i < Length(str)-Length(charset) {
			if str[i:i+Length(charset)] == charset {
				num += ArrayLength(str[i+Length(charset):], charset)
				break
			}
		}
	}
	return num + 1
}

func SplitStr(str string) []string {
	var ret []string
	for _, i := range str {
		ret = append(ret, string(i))
	}
	return ret
}

func RecursiveSplit(str, charset string, ret []string, m int) []string {
	word := ""
	for i := 0; i < Length(str); i++ {
		if i < Length(str)-Length(charset)+1 {
			if str[i] == charset[0] {
				if str[i:i+Length(charset)] == charset {
					RecursiveSplit(str[i+Length(charset):], charset, ret, m+1)
					break
				}
			}
		}
		word += string(str[i])
	}
	if Length(word) > 0 {
		ret[m] = word
	}
	return ret
}

func Split(str, charset string) []string {
	ret := make([]string, 0)
	if Length(charset) == 0 {
		return SplitStr(str)
	} else if Length(str) == 0 && Length(charset) == 0 {
		return ret
	} else if Length(str) == 0 {
		return ret
	}
	ret = make([]string, ArrayLength(str, charset))
	// fmt.Printf("len of ret: %v\n", len(ret))

	m := 0
	return RecursiveSplit(str, charset, ret, m)
}
