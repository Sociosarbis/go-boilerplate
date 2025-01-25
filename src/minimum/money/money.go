package money

import "sort"

func minimumMoney(transactions [][]int) int64 {
	sort.Slice(transactions, func(i, j int) bool {
		a, b := transactions[i][1]-transactions[i][0], transactions[j][1]-transactions[j][0]
		if a < 0 {
			if b < 0 {
				return transactions[i][1] < transactions[j][1]
			} else {
				return true
			}
		} else {
			if b < 0 {
				return false
			} else {
				return transactions[i][0] > transactions[j][0]
			}
		}
	})

	var ret int64
	var temp int
	for _, tx := range transactions {
		if temp < tx[0] {
			ret += int64(tx[0] - temp)
			temp = 0
		} else {
			temp -= tx[0]
		}
		temp += tx[1]
	}
	return ret
}
