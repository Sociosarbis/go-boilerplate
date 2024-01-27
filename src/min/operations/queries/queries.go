package queries

import (
	"math"

	"github.com/boilerplate/src/utils"
)

func minOperationsQueries(n int, edges [][]int, queries [][]int) []int {
	freq := make([][]int, n)
	temp := make([]int, 27)
	g := make([]map[int]int, n)
	freq[0] = make([]int, 27)
	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		if g[u] == nil {
			g[u] = make(map[int]int, 1)
		}
		g[u][v] = edge[2]
		if g[v] == nil {
			g[v] = make(map[int]int, 1)
		}
		g[v][u] = edge[2]
	}
	dfs1(0, 0, g, freq, temp)

	l := int(math.Ceil(math.Log2(float64(n))))
	up := make([][]int, n)
	for i := range up {
		up[i] = make([]int, l+1)
	}
	tin := make([]int, n)
	tout := make([]int, n)
	timer := 0
	dfs2(0, 0, g, up, &timer, tin, tout)

	ret := make([]int, len(queries))
	for i, query := range queries {
		if query[0] == 0 {
			ret[i] = sum(freq[query[1]]) - utils.Max(freq[query[1]]...)
		} else if query[1] == 0 {
			ret[i] = sum(freq[query[0]]) - utils.Max(freq[query[0]]...)
		} else {
			ancestor := lca(query[0], query[1], up, tin, tout)
			n := sum(freq[query[0]]) + sum(freq[query[1]]) - sum(freq[ancestor])*2
			for j := 1; j <= 26; j++ {
				temp := freq[query[0]][j] + freq[query[1]][j] - freq[ancestor][j]*2
				if temp > ret[i] {
					ret[i] = temp
				}
			}
			ret[i] = n - ret[i]
		}
	}
	return ret
}

func dfs1(u int, p int, g []map[int]int, freq [][]int, temp []int) {
	for v, w := range g[u] {
		if v != p {
			temp[w]++
			freq[v] = make([]int, 27)
			copy(freq[v], temp)
			dfs1(v, u, g, freq, temp)
			temp[w]--
		}
	}
}

// binary lifting
func dfs2(v int, p int, g []map[int]int, up [][]int, timer *int, tin []int, tout []int) {
	*timer++
	tin[v] = *timer
	up[v][0] = p
	l := len(up[v])
	for i := 1; i < l; i++ {
		up[v][i] = up[up[v][i-1]][i-1]
	}
	for u := range g[v] {
		if u != p {
			dfs2(u, v, g, up, timer, tin, tout)
		}
	}
	*timer++
	tout[v] = *timer
}

func isAncestor(u int, v int, tin []int, tout []int) bool {
	return tin[u] <= tin[v] && tout[u] >= tout[v]
}

func lca(u int, v int, up [][]int, tin []int, tout []int) int {
	if isAncestor(u, v, tin, tout) {
		return u
	} else if isAncestor(v, u, tin, tout) {
		return v
	}
	for i := len(up[u]) - 1; i >= 0; i-- {
		if !isAncestor(up[u][i], v, tin, tout) {
			u = up[u][i]
		}
	}
	return up[u][0]
}

func sum(nums []int) int {
	ret := 0
	for _, num := range nums {
		ret += num
	}
	return ret
}
