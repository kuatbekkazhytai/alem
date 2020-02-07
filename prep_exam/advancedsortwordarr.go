package piscine

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	arr := array
	arrLen := 0
	for range arr {
		arrLen++
	} 
	for i := 0; i < arrLen; i ++ {
		for j := i+1; j < arrLen; j++ {
			if f(arr[i], arr[j]) > 0 {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}