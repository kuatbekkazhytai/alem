 package piscine

func AtoiBase(s string, base string) int {
	sLen := StrLen(s)
	baseLen := StrLen(base)
	if baseLen < 2 {
		return 0
	} 
	for _, j := range base {
		if j == '+' || j == '-' {
			return 0
		}
	}
	for i := 0; i < baseLen; i ++ {
		for j := i+1; j < baseLen; j++ {
			if   base[i] == base[j] {
				return 0
			}
		}
	}
	pow := RecursivePower(baseLen, sLen-1)
		//return pow
	return AtoiBaseRec(s, base, pow, sLen, baseLen)
}

func AtoiBaseRec(s, base string, pow, sLen, baseLen int) int {
	res := 0
	for i := 0; i < baseLen; i ++ {
		for j :=0; j <sLen; j ++ {
			if s[j] == base[i] {
				res=res +  pow*i
				pow = pow/baseLen
				
			} 
			
			
		}
	}
	return res
}

func StrLen(s string) int {
	l:=0
	for range s {
		l++
	}
	return l
} 