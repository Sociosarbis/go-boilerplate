package product

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDotProduct(nums1 []int, nums2 []int) int {
	n := len(nums1)
	dp := make([]int, len(nums2))
	a := nums1[0]
	for i, b := range nums2 {
		dp[i] = a * b
	}
	for i := 1; i < n; i++ {
		var mx, nextMax int
		a := nums1[i]
		for i, b := range nums2 {
			nextMax = max(mx, dp[i])
			dp[i] = max(a*b+mx, dp[i])
			mx = nextMax
		}
	}
	out := dp[0]
	n = len(nums2)
	for i := 1; i < n; i++ {
		out = max(dp[i], out)
	}
	return out
}
