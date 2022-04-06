package trees

// 等同于求树的中心节点
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	neighbors := make([][]int, n)
	for _, edge := range edges {
		neighbors[edge[0]] = append(neighbors[edge[0]], edge[1])
		neighbors[edge[1]] = append(neighbors[edge[1]], edge[0])
	}

	degrees := make([]int, n)

	queue := []int{}
	count := 0
	for i, neighbor := range neighbors {
		degrees[i] = len(neighbor)
		if degrees[i] == 1 {
			queue = append(queue, i)
			degrees[i] = 0
		}
	}
	count += len(queue)

	for count < n {
		nextQueue := []int{}
		for _, i := range queue {
			for _, neighbor := range neighbors[i] {
				if degrees[neighbor] > 0 {
					degrees[neighbor]--
					if degrees[neighbor] == 1 {
						nextQueue = append(nextQueue, neighbor)
					}
				}
			}
		}
		queue = nextQueue
		count += len(queue)
	}

	return queue
}
