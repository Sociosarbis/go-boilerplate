package assignment

import (
	"sort"
)

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	n := len(difficulty)
	indices := make([]int, n)
	for i := 1; i < n; i++ {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		if difficulty[indices[i]] != difficulty[indices[j]] {
			return difficulty[indices[i]] < difficulty[indices[j]]
		}
		return profit[indices[i]] > profit[indices[j]]
	})

	for i := n - 1; i >= 0; i-- {
		ii := indices[i]
		n := len(indices)
		j := i
		for ; j+1 < n; j++ {
			jj := indices[j+1]
			if profit[jj] > profit[ii] {
				break
			}
		}
		if j != i {
			copy(indices[i+1:], indices[j+1:])
			indices = indices[:n-(j-i)]
		}
	}
	n = len(indices)

	var ret int
	for _, w := range worker {
		index := sort.Search(n, func(i int) bool {
			ii := indices[i]
			return difficulty[ii] > w
		})
		if index > 0 {
			ret += profit[indices[index-1]]
		}
	}
	return ret
}
