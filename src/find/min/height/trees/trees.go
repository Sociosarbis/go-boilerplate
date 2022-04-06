package trees

func dfs(i int, neighbors *[][]int, visited *[]bool, temp int) int {
	ret := temp + 1
	(*visited)[i] = true
	for _, neighbor := range (*neighbors)[i] {
		if !(*visited)[neighbor] {
			tempRet := dfs(neighbor, neighbors, visited, temp+1)
			if tempRet > ret {
				ret = tempRet
			}
		}
	}
	(*visited)[i] = false
	return ret
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	neighbors := make([][]int, n)
	for _, edge := range edges {
		neighbors[edge[0]] = append(neighbors[edge[0]], edge[1])
		neighbors[edge[1]] = append(neighbors[edge[1]], edge[0])
	}
	visited := make([]bool, n)
	min := n
	ret := []int{}
	for i := range neighbors {
		temp := dfs(i, &neighbors, &visited, 0)
		if temp < min {
			min = temp
			ret = []int{i}
		} else if temp == min {
			ret = append(ret, i)
		}
	}
	return ret
}
