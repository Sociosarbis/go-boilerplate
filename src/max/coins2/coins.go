package coins2

import "sort"

func maxCoins(piles []int) int {
	sort.Ints(piles)
	n := len(piles)
	var ret int
	j := n / 3
	for i := n - 2; i >= j; i -= 2 {
		ret += piles[i]
	}
	return ret
}
