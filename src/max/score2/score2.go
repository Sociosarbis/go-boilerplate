package score2

func maxScore(nums []int, x int) int64 {
	dp := [2]int64{0, 0}
	visited := [2]bool{false, false}
	if nums[0]&1 == 0 {
		dp[0] = int64(nums[0])
		visited[0] = true
	} else {
		dp[1] = int64(nums[0])
		visited[1] = true
	}

	n := len(nums)

	for i := 1; i < n; i++ {
		var v0 int64
		if visited[0] {
			v0 = dp[0] + int64(nums[i])
		}
		var v1 int64
		if visited[1] {
			v1 = dp[1] + int64(nums[i])
		}
		index := nums[i] & 1
		if index == 0 {
			v1 -= int64(x)
		} else {
			v0 -= int64(x)
		}
		var temp int64
		if visited[0] && (!visited[1] || v0 > v1) {
			temp = v0
		} else {
			temp = v1
		}
		if temp > dp[index] || !visited[index] {
			dp[index] = temp
			visited[index] = true
		}
	}

	if dp[0] > dp[1] {
		return dp[0]
	}
	return dp[1]
}
