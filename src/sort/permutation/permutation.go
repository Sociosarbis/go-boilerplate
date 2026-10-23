package permutation

func sortPermutation(nums []int) int {
	out := 131071
	for i, num := range nums {
		if num != i {
			out &= num
		}
	}
	if out == 131071 {
		return 0
	}
	return out
}
