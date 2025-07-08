package value3

import (
	"sort"
)

func maxValue(events [][]int, k int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0] || (events[i][0] == events[j][0] && events[i][1] < events[j][1])
	})
	dp := make([][][2]int, k)
	dp[0] = append(dp[0], [2]int{
		events[0][2],
		events[0][1],
	})
	n := len(events)
	insert := func(j, nextValue, end int) {
		i2 := sort.Search(len(dp[j]), func(i int) bool {
			return dp[j][i][1] >= end
		})

		if i2 == len(dp[j]) {
			if i2 == 0 {
				dp[j] = append(dp[j], [2]int{nextValue, end})
			} else {
				top := dp[j][i2-1]
				if top[0] < nextValue {
					dp[j] = append(dp[j], [2]int{nextValue, end})
				}
			}
		} else {
			if i2 > 0 {
				if dp[j][i2-1][0] >= nextValue {
					return
				}
			}
			if dp[j][i2][0] < nextValue {
				dp[j][i2][0], dp[j][i2][1] = nextValue, end
			} else if dp[j][i2][1] > end {
				dp[j] = append(dp[j], [2]int{})
				copy(dp[j][i2+1:], dp[j][i2:])
				dp[j][i2] = [2]int{nextValue, end}
			}

			i3 := i2 + 1
			nextValue = dp[j][i2][0]
			for ; i3 < len(dp[j]); i3++ {
				if dp[j][i3][0] > nextValue {
					break
				}
			}
			if i3 != i2+1 {

				copy(dp[j][i2+1:], dp[j][i3:])
				dp[j] = dp[j][:len(dp[j])-(i3-i2-1)]
			}
		}
	}
	out := dp[0][0][0]
	for i := 1; i < n; i++ {
		start, end, value := events[i][0], events[i][1], events[i][2]
		from := k - 1
		if from >= i {
			from = i
		}
		for j := from; j > 0; j-- {
			i1 := sort.Search(len(dp[j-1]), func(i int) bool {
				return dp[j-1][i][1] >= start
			})
			if i1 != 0 {
				i1--
				nextValue := value + dp[j-1][i1][0]
				insert(j, nextValue, end)
				if nextValue > out {
					out = nextValue
				}
			}
		}
		insert(0, value, end)
		if value > out {
			out = value
		}
	}
	return out
}
