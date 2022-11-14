package arrange

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func splitArraySameAverage(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum == 0 {
		return true
	}
	n := len(nums)
	m := gcd(sum, len(nums))
	sum /= m
	n /= m
	half := len(nums) / 2
	dp := make([]map[int]bool, half+1)
	for j := 0; j <= half; j++ {
		dp[j] = map[int]bool{}
	}
	dp[0][0] = true
	for _, num := range nums {
		for j := half; j >= 1; j-- {
			for k := range dp[j-1] {
				if _, ok := dp[j][k+num]; !ok {
					if j%n == 0 && sum*j/n == k+num {
						return true
					}
					dp[j][k+num] = true
				}
			}
		}
	}
	return false
}
