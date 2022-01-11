package possible

var DIRS = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func getVisited(visited *map[int]map[int]int, i, j int) int {
	if _, ok := (*visited)[i]; !ok {
		return 0
	}
	if _, ok := (*visited)[i][j]; !ok {
		return 0
	}
	return (*visited)[i][j]
}

func setVisited(visited *map[int]map[int]int, i, j, v int) {
	if _, ok := (*visited)[i]; !ok {
		(*visited)[i] = map[int]int{}
	}
	(*visited)[i][j] = v
}

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	temp := source
	if source[0]+source[1] > target[0]+target[1] {
		source = target
		target = temp
	}
	visited := map[int]map[int]int{}
	for _, item := range blocked {
		setVisited(&visited, item[0], item[1], 1)
	}
	return dfs(&visited, source, target, source, 2) && dfs(&visited, target, source, target, 3)
}

func dfs(visited *map[int]map[int]int, source []int, target []int, origin []int, v int) bool {
	if source[0] < 0 || source[1] < 0 || source[0] >= 1000000 || source[1] >= 1000000 {
		return false
	}
	visitedFlag := getVisited(visited, source[0], source[1])
	if visitedFlag == 1 || visitedFlag == v {
		return false
	} else if visitedFlag != 0 {
		return true
	}
	setVisited(visited, source[0], source[1], v)
	if source[0] == target[0] && source[1] == target[1] || abs(source[0]-origin[0])+abs(source[1]-origin[1]) > 200 {
		return true
	}
	for _, dir := range DIRS {
		if dfs(visited, []int{source[0] + dir[0], source[1] + dir[1]}, target, origin, v) {
			return true
		}
	}
	return false
}
