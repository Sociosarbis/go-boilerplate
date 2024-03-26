package graph

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

type Graph struct {
	n int
	m map[int]map[int]int
}

func Constructor(n int, edges [][]int) Graph {
	m := make(map[int]map[int]int, n)
	for i := 0; i < n; i++ {
		m[i] = map[int]int{}
	}
	for _, edge := range edges {
		a, b, c := edge[0], edge[1], edge[2]
		m[a][b] = c
	}
	return Graph{
		n,
		m,
	}
}

func (this *Graph) AddEdge(edge []int) {
	a, b, c := edge[0], edge[1], edge[2]
	this.m[a][b] = c
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	visited := make([]int, this.n+1)
	options := h.New[[2]int]([][2]int{{0, node1}}, func(a, b [2]int) bool {
		return a[0] < b[0]
	})
	for options.Len() != 0 {
		top := heap.Pop(&options).([2]int)
		if top[1] == node2 {
			return top[0]
		}
		for next, c := range this.m[top[1]] {
			temp := top[0] + c
			if visited[next] > temp || (next != node1 && visited[next] == 0) {
				visited[next] = temp
				heap.Push(&options, [2]int{temp, next})
			}
		}
	}
	return -1
}
