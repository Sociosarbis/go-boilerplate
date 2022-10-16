package partition

func dfs(i int, graph [][]int, groupMap []int) bool {
	group := groupMap[i-1]
	dislikeGroup := 3 - group
	for _, num := range graph[i-1] {
		if groupMap[num-1] != 0 {
			if groupMap[num-1] != dislikeGroup {
				return false
			}
		} else {
			groupMap[num-1] = dislikeGroup
			if !dfs(num, graph, groupMap) {
				return false
			}
		}
	}
	return true
}

func possibleBipartition(n int, dislikes [][]int) bool {
	graph := make([][]int, n)
	for _, dislike := range dislikes {
		graph[dislike[0]-1] = append(graph[dislike[0]-1], dislike[1])
		graph[dislike[1]-1] = append(graph[dislike[1]-1], dislike[0])
	}
	groupMap := make([]int, n)
	for i := 1; i <= n; i++ {
		if groupMap[i-1] == 0 {
			groupMap[i-1] = 1
			if !dfs(i, graph, groupMap) {
				return false
			}
		}
	}
	return true
}
