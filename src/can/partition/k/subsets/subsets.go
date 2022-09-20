package subsets

func dfs(nums []int, curSum, subsetSum, mask int, dp map[int]bool) bool {
	if mask == (1<<len(nums))-1 {
		return true
	}
	for i, num := range nums {
		if mask&(1<<i) == 0 {
			if num+curSum <= subsetSum {
				temp := num + curSum
				var nextCurSum int
				if temp == subsetSum {
					nextCurSum = 0
				} else {
					nextCurSum = temp
				}
				nextMask := mask | (1 << i)
				if _, ok := dp[nextMask]; !ok {
					dp[nextMask] = true
					if dfs(nums, nextCurSum, subsetSum, nextMask, dp) {
						return true
					}
				}
			}
		}
	}
	return false
}

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	dp := map[int]bool{}
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	subsetSum := sum / k
	res := dfs(nums, 0, subsetSum, 0, dp)
	return res
}
