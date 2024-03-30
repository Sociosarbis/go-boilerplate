package coins

import "sort"

func minimumAddedCoins(coins []int, target int) int {
	sort.Ints(coins)
	n := len(coins)
	i := 0
	// temp表示之前的coin能够组成[1,temp]的数值
	temp := 0
	ret := 0
	for temp < target {
		if i < n {
			// 如果下个coin不能接上，则要新增
			if coins[i] <= temp+1 {
				temp += coins[i]
				i++
				continue
			} else {
				ret++
			}
		} else {
			ret++
		}
		temp += temp + 1
	}
	return ret
}
