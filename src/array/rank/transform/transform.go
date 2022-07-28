package transform

import (
	"sort"
)

func arrayRankTransform(arr []int) []int {
	tempArr := make([]int, len(arr))
	copy(tempArr, arr)
	sort.Ints(tempArr)
	temp := 1
	prev := tempArr[0]
	m := make(map[int]int, len(arr))
	m[prev] = temp
	for i := 1; i < len(tempArr); i++ {
		if tempArr[i] != prev {
			temp++
			prev = tempArr[i]
			m[prev] = temp
		}
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = m[arr[i]]
	}
	return arr
}
