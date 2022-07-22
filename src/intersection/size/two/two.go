package two

import (
	"sort"
)

type Intervals [][]int

func (this Intervals) Len() int {
	return len(this)
}

func (this Intervals) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this Intervals) Less(i, j int) bool {
	if this[i][0] < this[j][0] {
		return true
	} else if this[i][0] == this[j][0] {
		return this[i][1] > this[j][1]
	}
	return false
}

func intersectionSizeTwo(intervals [][]int) int {
	ret := 0
	sort.Sort(Intervals(intervals))
	vals := make([]int, len(intervals))
	// vals表示重叠次数
	for i := len(intervals) - 1; i >= 0; i-- {
		val := intervals[i][0]
		for j := vals[i]; j < 2; j++ {
			ret++
			for k := i - 1; k >= 0; k-- {
				if intervals[k][1] < val {
					break
				}
				vals[k]++
			}
			val++
		}
	}

	return ret
}
