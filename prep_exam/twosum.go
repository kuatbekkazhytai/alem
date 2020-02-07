package piscine

func TwoSum(nums []int, target int) []int {
	arr := []int{0, 0}

	for i, j := range nums {
		for k, l := range nums {
			if j+l == target {
				arr[0] = i
				arr[1] = k
			}
		}
	}
	return arr
}
