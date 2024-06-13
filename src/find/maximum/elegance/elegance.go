package elegance

import "sort"

type empty struct{}

func findMaximumElegance(items [][]int, k int) int64 {
	// 先取profit最大项，然后逐个尝试替换重复项
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] > items[j][0]
	})

	visited := make(map[int]empty, k)

	duplicates := make([][]int, 0, k/2)
	var ret int64
	var total int64
	for i := 0; i < k; i++ {
		item := items[i]
		if _, ok := visited[item[1]]; !ok {
			visited[item[1]] = empty{}
		} else {
			duplicates = append(duplicates, item)
		}
		total += int64(item[0])
	}
	nu := len(visited)
	ret = total + int64(nu)*int64(nu)
	n := len(items)
	nd := len(duplicates)
	for i := k; i < n && nd != 0; i++ {
		if _, ok := visited[items[i][1]]; !ok {
			total += int64(items[i][0]) - int64(duplicates[nd-1][0])
			nd--
			duplicates = duplicates[:nd]
			visited[items[i][1]] = empty{}
			nu++
			temp := total + int64(nu)*int64(nu)
			if temp > ret {
				ret = temp
			}
		}
	}
	return ret
}
