package win

import "sort"

type item struct {
	start int
	count int
}

func maximizeWin(prizePositions []int, k int) int {
	var prev int
	items := []item{}
	n := len(prizePositions)
	var j int
	for i, p := range prizePositions {
		if p != prev {
			prev = p
			if i+1 > j {
				j = i + 1
			}
			for ; j < n; j++ {
				if prizePositions[j]-p <= k {
					continue
				} else {
					break
				}
			}
			count := j - i
			items = append(items, item{
				start: p,
				count: count,
			})
		}
	}
	n = len(items)
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		i1 := indices[i]
		i2 := indices[j]
		return items[i1].count < items[i2].count
	})

	var ret int

	for _, item := range items {
		for len(indices) > 0 && items[indices[len(indices)-1]].start <= item.start+k {
			indices = indices[:len(indices)-1]
		}
		var temp int
		if len(indices) > 0 {
			temp = items[indices[len(indices)-1]].count
		}
		if temp+item.count > ret {
			ret = temp + item.count
		}
	}
	return ret
}
