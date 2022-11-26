package nodes

type item struct {
	i int
	c int
}

func findItem(items []item, i, c int) int {
	l := 0
	r := len(items) - 1
	for l <= r {
		mid := (l + r) / 2
		if items[mid].c < c || (items[mid].c == c && items[mid].i < i) {
			l = mid + 1
		} else {
			if mid > 0 {
				if items[mid-1].c > c || (items[mid-1].c == c && items[mid].i > i) {
					r = mid - 1
					continue
				}
			}
			l = mid
			break
		}
	}
	return l
}

func reachableNodes(edges [][]int, maxMoves int, n int) int {
	g := make([]map[int]int, n)
	for i := range g {
		g[i] = map[int]int{}
	}
	for _, e := range edges {
		g[e[0]][e[1]] = e[2]
		g[e[1]][e[0]] = e[2]
	}
	visited := make([]int, n)
	for i := 0; i < n; i++ {
		visited[i] = -1
	}
	bfs := []item{{
		0,
		maxMoves,
	}}
	visited[0] = maxMoves
	for len(bfs) != 0 {
		back := bfs[len(bfs)-1]
		bfs = bfs[:len(bfs)-1]
		i, c := back.i, back.c
		for next, cost := range g[i] {
			nextMoves := c - cost - 1
			if nextMoves >= 0 {
				if nextMoves > visited[next] {
					if nextMoves > 0 {
						if visited[next] != -1 {
							index := findItem(bfs, next, visited[next])
							if index < len(bfs) && bfs[index].i == i {
								copy(bfs[index:], bfs[index+1:])
								bfs = bfs[:len(bfs)-1]
							}
						}
						index := findItem(bfs, next, nextMoves)
						if index < len(bfs) {
							bfs = append(bfs, item{})
							copy(bfs[index+1:], bfs[index:])
							bfs[index] = item{
								next,
								nextMoves,
							}
						} else {
							bfs = append(bfs, item{next, nextMoves})
						}
					}
					visited[next] = nextMoves
				}
			}
		}
	}
	ret := 0
	for i, v := range visited {
		if v >= 0 {
			ret++
			for j, c := range g[i] {
				if v+visited[j] >= c {
					if i < j {
						ret += c
					}
				} else {
					ret += v
				}
			}
		}
	}
	return ret
}
