package value2

func maxValue(nums []int, k int) int {
	n := len(nums)
	dp := [2][][][128]bool{}
	dp[0], dp[1] = make([][][128]bool, k+1), make([][][128]bool, k+1)
	for i := 0; i <= k; i++ {
		dp[0][i], dp[1][i] = make([][128]bool, n), make([][128]bool, n)
	}
	for index, num := range nums {
		if index > 0 {
			copy(dp[0][1][index][:], dp[0][1][index-1][:])
			for i := k - 1; i > 0; i-- {
				copy(dp[0][i+1][index][:], dp[0][i+1][index-1][:])
				for mask, ok := range dp[0][i][index-1] {
					if ok {
						dp[0][i+1][index][mask|num] = true
					}
				}
			}
		}
		dp[0][1][index][num] = true
	}
	for index := n - 1; index >= 0; index-- {
		num := nums[index]
		if index+1 < n {
			copy(dp[1][1][index][:], dp[1][1][index+1][:])
			for i := k - 1; i > 0; i-- {
				copy(dp[1][i+1][index][:], dp[1][i+1][index+1][:])
				for mask, ok := range dp[1][i][index+1] {
					if ok {
						dp[1][i+1][index][mask|num] = true
					}
				}
			}
		}
		dp[1][1][index][num] = true
	}
	var ret int
	for i := 0; i+k < n; i++ {
		for m1, k1 := range dp[0][k][i] {
			if k1 {
				for m2, k2 := range dp[1][k][i+1] {
					if k2 {
						if m1^m2 > ret {
							ret = m1 ^ m2
						}
					}
				}
			}
		}
	}
	return ret
}
