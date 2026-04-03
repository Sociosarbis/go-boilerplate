package walls

import (
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func count(walls []int, start, end int) int {
	if start > end {
		return 0
	}
	n := len(walls)
	index := sort.Search(n, func(i int) bool {
		return walls[i] >= start
	})
	if index == n {
		return 0
	}
	return sort.Search(n-index, func(i int) bool {
		return walls[index+i] > end
	})
}

func maxWalls(robots []int, distance []int, walls []int) int {
	n := len(robots)
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		return robots[indices[i]] < robots[indices[j]]
	})
	sort.Ints(walls)
	rw := walls[len(walls)-1]
	dp := [2][2][2]int{}
	r, d := robots[indices[0]], distance[indices[0]]
	var next int
	if n > 1 {
		next = robots[indices[1]] - 1
	} else {
		next = rw
	}
	dp[0][0][0], dp[0][0][1] = count(walls, max(r-d, walls[0]), r), r+1
	dp[0][1][0], dp[0][1][1] = count(walls, r, min(r+d, next)), max(min(r+d+1, next+1), r+1)
	for i := 1; i < n; i++ {
		index := i & 1
		r, d := robots[indices[i]], distance[indices[i]]
		dp[index][0][0], dp[index][0][1] = dp[1-index][0][0]+count(walls, max(r-d, dp[1-index][0][1]), r), r+1
		temp1, temp2 := dp[1-index][1][0]+count(walls, max(r-d, dp[1-index][1][1]), r), r+1
		if temp1 > dp[index][0][0] {
			dp[index][0][0], dp[index][0][1] = temp1, temp2
		}
		var next int
		if i+1 < n {
			next = robots[indices[i+1]] - 1
		} else {
			next = rw
		}
		dp[index][1][0], dp[index][1][1] = max(dp[1-index][0][0], dp[1-index][1][0])+count(walls, r, min(r+d, next)), max(min(r+d+1, next+1), r+1)
	}
	index := (n - 1) & 1
	return max(dp[index][0][0], dp[index][1][0])
}
