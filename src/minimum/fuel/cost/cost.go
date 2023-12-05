package cost

type resType struct {
	full_car  int
	empty_car int
	cost      int64
}

func minimumFuelCost(roads [][]int, seats int) int64 {
	tree := make([][]int, len(roads)+1)
	visited := make([]bool, len(roads)+1)
	for _, road := range roads {
		tree[road[0]] = append(tree[road[0]], road[1])
		tree[road[1]] = append(tree[road[1]], road[0])
	}
	visited[0] = true
	return dfs(tree, 0, seats, visited).cost
}

func dfs(tree [][]int, i int, seats int, visited []bool) resType {
	ret := resType{}
	for _, child := range tree[i] {
		if !visited[child] {
			visited[child] = true
			res := dfs(tree, child, seats, visited)
			ret.full_car += res.full_car
			ret.cost += int64(res.full_car) + res.cost
			if res.empty_car != 0 {
				ret.empty_car += res.empty_car
				ret.cost += 1
				if ret.empty_car >= seats {
					ret.empty_car -= seats
					ret.full_car += 1
				}
			}
		}
	}
	ret.empty_car += 1
	if ret.empty_car >= seats {
		ret.empty_car -= seats
		ret.full_car += 1
	}
	return ret
}
