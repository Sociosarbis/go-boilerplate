package averages

func largestSumOfAverages(nums []int, k int) float64 {
	prefixSum := make([]float64, len(nums))
	prefixSum[0] = float64(nums[0])
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + float64(nums[i])
	}
	if k == 1 {
		return prefixSum[len(nums)-1] / float64(len(nums))
	}
	if k == len(nums) {
		return prefixSum[k-1]
	}
	dp := make([][]float64, len(nums))
	for i := range dp {
		dp[i] = make([]float64, k)
		r1 := k - 2
		if i+1 == len(dp) {
			r1 = k - 1
		}
		if r1 > i {
			r1 = i
		}
		dp[i][0] = prefixSum[i] / float64(i+1)
		for j := 0; j < r1; j++ {
			for k := j; k < i; k++ {
				v := dp[k][j] + (prefixSum[i]-prefixSum[k])/(float64(i)-float64(k))
				if v > dp[i][j+1] {
					dp[i][j+1] = v
				}
			}
		}
	}
	return dp[len(nums)-1][k-1]
}
