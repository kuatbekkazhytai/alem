package piscine

func SortWordArr(array []string) {
	l := len(array)
	for i := 0; i <= l - 1; i ++ {
		for j :=0; j <= l-1; j ++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	
}