package apart

func kLengthApart(nums []int, k int) bool {
	last := -k - 1
	for i, num := range nums {
		if num == 1 {
			if i-last-1 < k {
				return false
			}
			last = i
		}
	}
	return true
}
