package partitions

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func countPartitions(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	var temp, out int
	n := len(nums)
	for i := 1; i < n; i++ {
		temp += nums[i-1]
		if diff(temp, sum-temp)&1 == 0 {
			out++
		}
	}
	return out
}
