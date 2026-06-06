package difference

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func leftRightDifference(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	for i := 1; i < n; i++ {
		left[i] = left[i-1] + nums[i-1]
	}
	var temp int
	left[n-1] = abs(temp - left[n-1])
	for i := n - 2; i >= 0; i-- {
		temp += nums[i+1]
		left[i] = abs(temp - left[i])
	}
	return left
}
