package distance

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func maxAdjacentDistance(nums []int) int {
	n := len(nums)
	out := abs(nums[0] - nums[n-1])
	for i := 1; i < n; i++ {
		temp := abs(nums[i] - nums[i-1])
		if temp > out {
			out = temp
		}
	}
	return out
}
