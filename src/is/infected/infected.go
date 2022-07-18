package infected

var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

type group struct {
	enclose_cost int
	i            int
	j            int
	frontier     []int
}

func containVirus(isInfected [][]int) int {
	turn := 1
	ret := 0
	for {
		groups := []group{}
		for i := range isInfected {
			for j := range isInfected[i] {
				if isInfected[i][j] == turn {
					stack := [][]int{{i, j}}
					isInfected[i][j] = turn + 1
					group := group{
						enclose_cost: 0,
						i:            i,
						j:            j,
						frontier:     []int{},
					}
					for len(stack) != 0 {
						n := len(stack)
						for k := 0; k < n; k++ {
							x, y := stack[k][0], stack[k][1]
							for _, dir := range dirs {
								nextX, nextY := x+dir[0], y+dir[1]
								if nextX >= 0 && nextX < len(isInfected) && nextY >= 0 && nextY < len(isInfected[0]) {
									if isInfected[nextX][nextY] == turn {
										stack = append(stack, []int{nextX, nextY})
										isInfected[nextX][nextY] = turn + 1
									} else if isInfected[nextX][nextY] <= 0 {
										if isInfected[nextX][nextY] == 0 {
											group.frontier = append(group.frontier, nextX, nextY)
											isInfected[nextX][nextY] = -1
										}
										group.enclose_cost++
									}
								}
							}

						}
						stack = stack[n:]
					}
					if len(group.frontier) != 0 {
						for k := 0; k < len(group.frontier); k += 2 {
							x, y := group.frontier[k], group.frontier[k+1]
							isInfected[x][y] = 0
						}
						groups = append(groups, group)
					}
				}
			}
		}

		if len(groups) == 1 {
			ret += groups[0].enclose_cost
			break
		}
		max := len(groups[0].frontier)
		index := 0
		for i := 1; i < len(groups); i++ {
			if len(groups[i].frontier) > max {
				max = len(groups[i].frontier)
				index = i
			}
		}
		turn++
		ret += groups[index].enclose_cost
		stack := [][]int{{groups[index].i, groups[index].j}}
		for len(stack) != 0 {
			n := len(stack)
			for k := 0; k < n; k++ {
				x, y := stack[k][0], stack[k][1]
				for _, dir := range dirs {
					nextX, nextY := x+dir[0], y+dir[1]
					if nextX >= 0 && nextX < len(isInfected) && nextY >= 0 && nextY < len(isInfected[0]) {
						if isInfected[nextX][nextY] == turn {
							stack = append(stack, []int{nextX, nextY})
							isInfected[nextX][nextY] = 1
						}
					}
				}

			}
			stack = stack[n:]
		}
		copy(groups[index:], groups[index+1:])
		groups = groups[:len(groups)-1]
		for _, group := range groups {
			for i := 0; i < len(group.frontier); i += 2 {
				isInfected[group.frontier[i]][group.frontier[i+1]] = turn
			}
		}
	}
	return ret
}
