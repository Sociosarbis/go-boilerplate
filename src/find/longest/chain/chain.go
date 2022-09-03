package chain

import (
	"sort"
)

type Pairs [][]int

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this Pairs) Less(i, j int) bool {
	if this[i][1] == this[j][1] {
		return this[i][0] > this[j][0]
	} else {
		return this[i][1] < this[j][1]
	}
}

func findLongestChain(pairs [][]int) int {
	sort.Sort(Pairs(pairs))
	ret := 1
	prev := pairs[0][1]
	for i := 1; i < len(pairs); i++ {
		if pairs[i][0] > prev {
			prev = pairs[i][1]
			ret++
		}
	}
	return ret
}
