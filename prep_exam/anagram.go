package piscine

//import "fmt"

func IsAnagram(str1, str2 string) bool {
	n1 := 0
	n2 := 0
	l1 := 0
	l2 := 0
	var arr1 []rune
	var arr2 []rune

	for _, j := range str1 {
		arr1 = append(arr1, j)
		l1++
		if j == 32 {
			n1++
		}
	}
	for _, j := range str2 {
		arr2 = append(arr2, j)
		l2++
		if j == 32 {
			n2++
		}
	}
	for i := 0; i < l1; i++ {
		for j := i + 1; j < l1; j++ {
			if arr1[i] < arr1[j] {
				arr1[i], arr1[j] = arr1[j], arr1[i]
			}
		}
	}
	for i := 0; i < l2; i++ {
		for j := i + 1; j < l2; j++ {
			if arr2[i] < arr2[j] {
				arr2[i], arr2[j] = arr2[j], arr2[i]
			}
		}
	}
	arr1 = arr1[:l1-n1]
	arr2 = arr2[:l2-n2]
	str1 = string(arr1)
	str2 = string(arr2)
	if str1 == str2 {
		return true
	}
	return false

}
