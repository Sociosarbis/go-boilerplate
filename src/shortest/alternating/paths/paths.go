package paths

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	ret := make([]int, n)
	for i := 1; i < n; i++ {
		ret[i] = -1
	}
	redNext := make([][]byte, n)
	blueNext := make([][]byte, n)
	for _, edge := range redEdges {
		redNext[edge[0]] = append(redNext[edge[0]], byte(edge[1]))
	}
	for _, edge := range blueEdges {
		blueNext[edge[0]] = append(blueNext[edge[0]], byte(edge[1]))
	}
	redVisited := map[byte]struct{}{
		0: {},
	}
	blueVisited := map[byte]struct{}{
		0: {},
	}
	bfs := [][]byte{
		{0, 0},
		{1, 0},
	}
	i := 0
	for len(bfs) != 0 {
		n := len(bfs)
		for j := 0; j < n; j++ {
			item := bfs[j]
			if ret[item[1]] == -1 || ret[item[1]] > i {
				ret[item[1]] = i
			}
			if item[0] == 0 {
				for _, next := range blueNext[item[1]] {
					if _, ok := blueVisited[next]; !ok {
						bfs = append(bfs, []byte{1, next})
						blueVisited[next] = struct{}{}
					}
				}
			} else {
				for _, next := range redNext[item[1]] {
					if _, ok := redVisited[next]; !ok {
						bfs = append(bfs, []byte{0, next})
						redVisited[next] = struct{}{}
					}
				}
			}
		}
		bfs = bfs[n:]
		i++
	}
	return ret
}
