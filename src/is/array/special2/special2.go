package special2

import "sort"

func isArraySpecial(nums []int, queries [][]int) []bool {
	isOdd := nums[0]%2 == 1
	groups := [][2]int{{0, 0}}
	n := len(nums)
	for i := 1; i < n; i++ {
		if isOdd != (nums[i]%2 == 1) {
			isOdd = !isOdd
			groups[len(groups)-1][1] = i
		} else {
			groups = append(groups, [2]int{i, i})
		}
	}
	ret := make([]bool, len(queries))
	for i, query := range queries {
		from, to := query[0], query[1]
		index := sort.Search(len(groups), func(i int) bool {
			return groups[i][1] >= from
		})
		if groups[index][1] >= to {
			ret[i] = true
		} else {
			ret[i] = false
		}
	}
	return ret
}
