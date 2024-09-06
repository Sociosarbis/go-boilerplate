package length

func maximumLength(nums []int, k int) int {
	n := len(nums)
	dp := make(map[int][]int, n)
	ret := 1
	for _, num := range nums {
		if arr, ok := dp[num]; ok {
			for i := k; i >= 0; i-- {
				if arr[i] != 0 {
					arr[i]++
					if arr[i] > ret {
						ret = arr[i]
					}
				}
			}
		} else {
			dp[num] = make([]int, k+1)
			dp[num][0] = 1
		}

		for prev, arr := range dp {
			if prev != num {
				for i := 0; i < k; i++ {
					if arr[i] != 0 {
						if dp[num][i+1] < arr[i]+1 {
							dp[num][i+1] = arr[i] + 1
							if arr[i]+1 > ret {
								ret = arr[i] + 1
							}
						}
					}
				}
			}
		}
	}
	return ret
}
