package combinations

var ROOK_DIRS = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
var QUEEN_DIRS = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {1, 1}, {1, -1}, {-1, 1}}
var BISHOP_DIRS = [][2]int{{-1, -1}, {1, 1}, {1, -1}, {-1, 1}}

func countCombinations(pieces []string, positions [][]int) int {
	visited := make([][8]bool, 64)
	var ret int
	dfs(visited, positions, pieces, 0, &ret)
	return ret
}

func dfs(visited [][8]bool, positions [][]int, pieces []string, i int, out *int) {
	if i >= len(pieces) {
		*out++
		return
	}
	var x, y int
	piece := pieces[i]
	var dirs [][2]int
	switch piece {
	case "rook":
		dirs = ROOK_DIRS
	case "queen":
		dirs = QUEEN_DIRS
	default:
		dirs = BISHOP_DIRS
	}
	position := positions[i]
	for ii, dir := range dirs {
		var start int
		if ii != 0 {
			start = 1
		}
	loop:
		for j := start; j < 8; j++ {
			var index int
			x = position[0] + dir[0]*j - 1
			y = position[1] + dir[1]*j - 1
			if x < 0 || x >= 8 || y < 0 || y >= 8 {
				break
			}
			for k := start; k <= j; k++ {
				x = position[0] + dir[0]*k - 1
				y = position[1] + dir[1]*k - 1
				index = x*8 + y
				if visited[index][k] {
					break loop
				}
			}
			for k := j + 1; k < 8; k++ {
				if visited[index][k] {
					continue loop
				}
			}
			for k := start; k <= j; k++ {
				x = position[0] + dir[0]*k - 1
				y = position[1] + dir[1]*k - 1
				index = x*8 + y
				visited[index][k] = true
			}
			for k := j + 1; k < 8; k++ {
				visited[index][k] = true
			}
			dfs(visited, positions, pieces, i+1, out)
			for k := start; k <= j; k++ {
				x = position[0] + dir[0]*k - 1
				y = position[1] + dir[1]*k - 1
				index = x*8 + y
				visited[index][k] = false
			}
			for k := j + 1; k < 8; k++ {
				visited[index][k] = false
			}
		}
	}
}
