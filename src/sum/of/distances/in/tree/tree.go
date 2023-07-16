package tree

func sumOfDistancesInTree(n int, edges [][]int) []int {
	ret := make([]int, n)
	count := make([]int, n)
	graph := make([][]int, n)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	// 先以第一个点为根，算出其他点到第一个点的总距离
	helper1(&graph, 0, -1, &ret, &count)
	for _, next := range graph[0] {
		// 深度遍历相邻的点，根据前后两个点的差值，算出其他点到当前点的总距离
		helper2(&graph, next, 0, &ret, &count)
	}
	return ret
}

func helper1(graph *[][]int, cur int, pre int, ret *[]int, count *[]int) {
	(*count)[cur] = 1
	for _, next := range (*graph)[cur] {
		if next != pre {
			helper1(graph, next, cur, ret, count)
			(*count)[cur] += (*count)[next]
			(*ret)[cur] += (*ret)[next] + (*count)[next]
		}
	}
}

func helper2(graph *[][]int, cur int, pre int, ret *[]int, count *[]int) {
	(*ret)[cur] = (*ret)[pre] + (*count)[0] - (*count)[cur]*2
	for _, next := range (*graph)[cur] {
		if next != pre {
			helper2(graph, next, cur, ret, count)
		}
	}
}
