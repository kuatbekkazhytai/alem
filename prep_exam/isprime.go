package piscine

func IsPrime(nb int) bool {
	if nb < 3 {
		return false
	}
	if nb % 2 == 0 {
		return false
	}
	//var i 
	for i := 3; i*i <= nb; i=i + 2 {
		if nb%i == 0 {
			return false
		}
	} 
	return true
}