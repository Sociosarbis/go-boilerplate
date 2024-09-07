package length2

func maximumLength(nums []int, k int) int {
	n := len(nums)
	dp := [2][]int{make([]int, n), make([]int, n)}
	var ret int
	for i := 0; i <= k; i++ {
		index := i % 2
		max := make(map[int]int, n)
		for j, num := range nums {
			var temp int
			if v, ok := max[num]; ok {
				temp = v
			}
			if j > 0 {
				if dp[1-index][j-1] > temp {
					temp = dp[1-index][j-1]
				}
				if temp+1 > dp[index][j-1] {
					dp[index][j] = temp + 1
				} else {
					dp[index][j] = dp[index][j-1]
				}
			} else {
				dp[index][j] = temp + 1
			}
			max[nums[j]] = temp + 1
		}
	}
	index := k % 2
	for i := 0; i < n; i++ {
		if dp[index][i] > ret {
			ret = dp[index][i]
		}
	}
	return ret
}
