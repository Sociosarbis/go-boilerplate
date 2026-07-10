package queries2

import (
	"sort"
)

func getDistance(jumpTable map[int][]int, nums []int, u, v, maxDiff int) int {
	if u == v {
		return 0
	}
	if nums[v]-nums[u] <= maxDiff {
		return 1
	}
	if arr, ok := jumpTable[v]; ok {
		for i := len(arr) - 1; i >= 0; i-- {
			if arr[i] >= u {
				res := getDistance(jumpTable, nums, u, arr[i], maxDiff)
				if res != -1 {
					return (1 << i) + res
				}
				break
			}
		}
	}
	return -1
}

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		temp[i] = i
	}
	sort.Slice(temp, func(i, j int) bool {
		return nums[temp[i]] < nums[temp[j]]
	})
	oldIndexToNewIndex := make([]int, n)
	for i, index := range temp {
		oldIndexToNewIndex[index] = i
	}
	for i := 0; i < n; i++ {
		temp[i] = nums[temp[i]]
	}
	nums = temp
	jumpTable := make(map[int][]int, n)
	for i := 1; i < n; i++ {
		index := sort.Search(i, func(j int) bool {
			return nums[i]-nums[j] <= maxDiff
		})
		if index != i {
			var levels []int
			for j := 0; true; j++ {
				levels = append(levels, index)
				if targets, ok := jumpTable[index]; ok {
					if j < len(targets) {
						index = targets[j]
						continue
					}
				}
				break
			}
			jumpTable[i] = levels
		}
	}
	out := make([]int, len(queries))
	for i, query := range queries {
		u, v := oldIndexToNewIndex[query[0]], oldIndexToNewIndex[query[1]]
		if u > v {
			u, v = v, u
		}
		out[i] = getDistance(jumpTable, nums, u, v, maxDiff)
	}
	return out
}
