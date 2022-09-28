package number

func getKthMagicNumber(k int) int {
	if k == 1 {
		return 1
	}
	nums := make([]int, k)
	i := 1
	nums[0] = 1
	dp := make([]int, 3)
	for i < len(nums) {
		n1, n2, n3 := 3*nums[dp[0]], 5*nums[dp[1]], 7*nums[dp[2]]
		if n1 <= n2 {
			if n1 <= n3 {
				nums[i] = n1
			} else {
				nums[i] = n3
			}
		} else {
			if n2 <= n3 {
				nums[i] = n2
			} else {
				nums[i] = n3
			}
		}
		if i+1 < len(nums) {
			for 3*nums[dp[0]] <= nums[i] {
				dp[0]++
			}
			for 5*nums[dp[1]] <= nums[i] {
				dp[1]++
			}
			for 7*nums[dp[2]] <= nums[i] {
				dp[2]++
			}
		}
		i++
	}
	return nums[len(nums)-1]
}
