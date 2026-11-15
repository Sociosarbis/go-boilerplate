package subarrays

const mod int = 1e9 + 7

func numOfSubarrays(arr []int) int {
	n := len(arr)
	dp := [2][2]int{}
	if arr[0]&1 == 0 {
		dp[0][0] = 1
	} else {
		dp[0][1] = 1
	}
	var out int
	for i := 1; i < n; i++ {
		index := i & 1
		prevIndex := 1 - index
		if arr[i]&1 == 0 {
			dp[index][1] = dp[prevIndex][1]
			dp[index][0] = dp[prevIndex][0] + 1
		} else {
			dp[index][1] = dp[prevIndex][0] + 1
			dp[index][0] = dp[prevIndex][1]
		}
		out = (out + dp[index][1]) % mod
	}
	return out
}
