package reach

func canReach(arr []int, start int) bool {
	if arr[start] == 0 {
		return true
	}
	n := len(arr)
	visited := make([]bool, n)
	visited[start] = true
	multipliers := []int{-1, 1}
	bfs := make([]int, 0, n)
	bfs = append(bfs, start)
	l := 1
	for l != 0 {
		for i := 0; i < l; i++ {
			index := bfs[i]
			for _, multiplier := range multipliers {
				next := index + arr[index]*multiplier
				if next >= 0 && next < n && !visited[next] {
					if arr[next] == 0 {
						return true
					}
					visited[next] = true
					bfs = append(bfs, next)
				}
			}
		}
		bfs = bfs[l:]
		l = len(bfs)
	}
	return false
}
