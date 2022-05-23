package tree

import "sort"

var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func cutOffTree(forest [][]int) int {
	treeList := []int{}
	for _, row := range forest {
		for _, cell := range row {
			if cell > 1 {
				treeList = append(treeList, cell)
			}
		}
	}

	sort.Ints(treeList)

	start := []int{0, 0}

	ret := 0
	for _, tree := range treeList {
		temp := 0
		if forest[start[0]][start[1]] != tree {
			bfs := [][]int{start}
			visited := make([][]bool, len(forest))
			for i := range visited {
				visited[i] = make([]bool, len(forest[0]))
			}
			visited[start[0]][start[1]] = true
		Search:
			for len(bfs) != 0 {
				n := len(bfs)
				for i := 0; i < n; i++ {
					for j := 0; j < 4; j++ {
						nextI := bfs[i][0] + dirs[j][0]
						nextJ := bfs[i][1] + dirs[j][1]
						if nextI >= 0 && nextI < len(forest) && nextJ >= 0 && nextJ < len(forest[0]) && forest[nextI][nextJ] != 0 && !visited[nextI][nextJ] {
							if forest[nextI][nextJ] == tree {
								temp++
								start[0] = nextI
								start[1] = nextJ
								break Search
							}
							visited[nextI][nextJ] = true
							bfs = append(bfs, []int{nextI, nextJ})
						}
					}
				}
				temp++
				bfs = bfs[n:]
			}
			if len(bfs) == 0 {
				return -1
			}
		}
		ret += temp
	}
	return ret
}
