package scheduling

import "sort"

type item struct {
	i int
	t int
}

type items []item

func (this items) Len() int {
	return len(this)
}

func (this items) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this items) Less(i, j int) bool {
	if this[i].t < this[j].t {
		return true
	} else if this[i].t == this[j].t {
		return this[i].i < this[j].i
	}
	return false
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	dp := [][]int{{0, 0}}
	sortItems := make(items, len(startTime))
	for i, t := range startTime {
		sortItems[i].i = i
		sortItems[i].t = t
	}
	sort.Sort(sortItems)

	temp := make([]int, len(startTime))
	for i, item := range sortItems {
		temp[i] = startTime[item.i]
	}
	copy(startTime, temp)
	for i, item := range sortItems {
		temp[i] = endTime[item.i]
	}
	copy(endTime, temp)
	for i, item := range sortItems {
		temp[i] = profit[item.i]
	}
	copy(profit, temp)

	for i, p := range profit {
		l := 0
		r := len(dp) - 1
		s := startTime[i]
		for l <= r {
			mid := (l + r) / 2
			if s < dp[mid][0] {
				r = mid - 1
			} else {
				if mid+1 < len(dp) && dp[mid+1][0] <= s {
					l = mid + 1
				} else {
					l = mid
					break
				}
			}
		}
		item := []int{endTime[i], p + dp[l][1]}
		l = 0
		r = len(dp) - 1
		for l <= r {
			mid := (l + r) / 2
			if item[0] < dp[mid][0] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		if l > 0 {
			if item[1] < dp[l-1][1] {
				continue
			}
		}
		if l >= len(dp) {
			dp = append(dp, item)
		} else {
			j := l
			for ; j < len(dp); j++ {
				if dp[j][1] > item[1] {
					break
				}
			}
			if j > l {
				if j != l+1 {
					copy(dp[l+1:], dp[j:])
					dp = dp[:len(dp)-j+l+1]
				}
			} else {
				dp = append(dp, make([]int, 0))
				copy(dp[l+1:], dp[l:])
			}
			dp[l] = item
		}
	}
	return dp[len(dp)-1][1]
}
