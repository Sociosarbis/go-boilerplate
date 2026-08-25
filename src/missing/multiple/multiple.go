package multiple

func missingMultiple(nums []int, k int) int {
	exists := [101]bool{}
	for _, num := range nums {
		exists[num] = true
	}
	i := k
	for ; i <= 100; i += k {
		if !exists[i] {
			return i
		}
	}
	return i
}
