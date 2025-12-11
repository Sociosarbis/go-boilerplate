package buildings

import (
	"sort"
)

type empty struct{}

func visit(visited []bool, flags []byte, i, index int) {
	if visited[index] {
		flags[i]++
	}
	visited[index] = true

}

func countCoveredBuildings(n int, buildings [][]int) int {
	visited := make([]bool, n+1)
	l := len(buildings)
	arr := make([][3]int, l)
	for i, building := range buildings {
		arr[i][0], arr[i][1], arr[i][2] = building[0], building[1], i
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	flags := make([]byte, l)
	for _, building := range arr {
		visit(visited, flags, building[2], building[1])
	}
	visited = make([]bool, n+1)
	for i := l - 1; i >= 0; i-- {
		building := arr[i]
		visit(visited, flags, building[2], building[1])
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1]
	})
	visited = make([]bool, n+1)
	for _, building := range arr {
		visit(visited, flags, building[2], building[0])
	}
	visited = make([]bool, n+1)
	for i := l - 1; i >= 0; i-- {
		building := arr[i]
		visit(visited, flags, building[2], building[0])
	}
	var out int
	for _, flag := range flags {
		if flag == 4 {
			out++
		}
	}
	return out
}
