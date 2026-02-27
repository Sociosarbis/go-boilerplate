package operations15

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minOperations(s string, k int) int {
	n := len(s)
	options := [2][]int{
		make([]int, 0, n/2+1),
		make([]int, 0, n/2+1),
	}
	var m int
	for _, c := range s {
		if c == '0' {
			m++
		}
	}
	if m == 0 {
		return 0
	}
	for i := 0; i <= n; i++ {
		options[i%2] = append(options[i%2], i)
	}
	oi := m % 2
	index := m / 2
	copy(options[oi][index:], options[oi][index+1:])
	options[oi] = options[oi][:len(options[oi])-1]
	dist := make([]int, n+1)
	queue := make([]int, 0, n+1)
	queue = append(queue, m)
	var i int
	for i < len(queue) {
		z := queue[i]
		i++
		c1 := max(0, k-(n-z))
		c2 := min(k, z)
		// 不用考虑0在字符串的位置
		lnode := z + k - 2*c2
		rnode := z + k - 2*c1
		oi := lnode % 2
		on := len(options[oi])
		index := sort.Search(len(options[oi]), func(i int) bool {
			return options[oi][i] >= lnode
		})
		j := index
		for ; j < on && options[oi][j] <= rnode; j++ {
			dist[options[oi][j]] = dist[z] + 1
			queue = append(queue, options[oi][j])
		}
		copy(options[oi][index:], options[oi][j:])
		options[oi] = options[oi][:index+(on-j)]
	}
	if dist[0] == 0 {
		return -1
	}
	return dist[0]
}
