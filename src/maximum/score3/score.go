package score3

import "sort"

func maxmiumScore(cards []int, cnt int) int {
	sort.Ints(cards)
	n := len(cards)

	var ret int

	for i := n - cnt; i < n; i++ {
		ret += cards[i]
	}
	if ret%2 == 0 {
		return ret
	}
	diff := -ret
	for i := n - cnt; i < n; i++ {
		if cards[i]%2 == 0 {
			for j := n - cnt - 1; j >= 0; j-- {
				if cards[j]%2 == 1 {
					diff = cards[j] - cards[i]
					break
				}
			}
			break
		}
	}
	for i := n - cnt; i < n; i++ {
		if cards[i]%2 == 1 {
			for j := n - cnt - 1; j >= 0; j-- {
				if cards[j]%2 == 0 {
					if cards[j]-cards[i] > diff {
						return ret + cards[j] - cards[i]
					}
					break
				}
			}
			break
		}
	}
	return ret + diff
}
