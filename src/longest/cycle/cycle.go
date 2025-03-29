package cycle

func longestCycle(edges []int) int {
	n := len(edges)
	inDeg := make([]int, n)
	for _, edge := range edges {
		if edge != -1 {
			inDeg[edge]++
		}
	}
	visited := make([]bool, n)
	queue := make([]int, 0, n)
	for i, deg := range inDeg {
		if deg == 0 {
			visited[i] = true
			queue = append(queue, i)
		}
	}
	l := len(queue)
	var ret int
	for l != 0 {
		for i := 0; i < l; i++ {
			next := edges[queue[i]]
			if next != -1 {
				if inDeg[next] > 0 {
					inDeg[next]--
					if inDeg[next] == 0 {
						visited[next] = true
						queue = append(queue, next)
					}
				}
			}
		}
		queue = queue[l:]
		l = len(queue)
	}
	for i := 0; i < n; i++ {
		if !visited[i] {
			var temp int
			next := edges[i]
			for next != -1 {
				temp++
				if next == i {
					if temp > ret {
						ret = temp
					}
					break
				}
				visited[next] = true
				next = edges[next]
			}
		}
	}
	if ret == 0 {
		return -1
	}
	return ret
}
