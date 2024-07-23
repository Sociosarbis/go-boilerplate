package power

import "sort"

const mod int = 1e9 + 7

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sumOfPowers(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	// dp[i][p][v]表示以i结尾且长度为p的子序列中
	// 最小diff大于等于v的数量
	dp := make([][]map[int]int, n)
	var ret int
	for i := 1; i < n; i++ {
		dp[i] = make([]map[int]int, k+1)
		end := min(k, i+1)
		for j := 2; j <= end; j++ {
			dp[i][j] = map[int]int{}
		}
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			if cnt, ok := dp[i][2][diff]; ok {
				dp[i][2][diff] = cnt + 1
			} else {
				dp[i][2][diff] = 1
			}
		}
		for j := 1; j < i; j++ {
			diff := nums[i] - nums[j]
			end := min(k, j+2)
			for p := 3; p <= end; p++ {
				// 遍历以j结尾，长度小1的各种可能
				for v, cnt := range dp[j][p-1] {
					v1 := min(diff, v)
					dp[i][p][v1] = (dp[i][p][v1] + cnt) % mod
				}
			}
		}

		for v, cnt := range dp[i][k] {
			ret = (ret + v*cnt%mod) % mod
		}
	}
	return ret
}
