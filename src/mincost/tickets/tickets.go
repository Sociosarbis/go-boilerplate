package tickets

var intervals = [3]int{1, 7, 30}

func mincostTickets(days []int, costs []int) int {
	n := len(days)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		d := days[i]
		for ii, interval := range intervals {
			minDay := d - interval + 1
			var prevValue int
			if days[0] >= minDay {
				prevValue = 0
			} else {
				j := i - 1
				for ; j >= 0 && days[j+1] >= minDay; j-- {
					if prevValue == 0 || dp[j] < prevValue {
						prevValue = dp[j]
					}
				}
			}
			if dp[i] == 0 || prevValue+costs[ii] < dp[i] {
				dp[i] = prevValue + costs[ii]
			}
		}
	}
	return dp[n-1]
}
