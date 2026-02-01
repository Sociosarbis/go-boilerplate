package cost6

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumCost(nums []int) int {
	out := 150
	n := len(nums)
	right := nums[n-1]
	for i := n - 2; i > 0; i-- {
		out = min(out, nums[i]+right+nums[0])
		right = min(right, nums[i])
	}
	return out
}
