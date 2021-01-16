package bricks

type group struct {
	p        int
	size     int
	isStable bool
}

// 先进行消除，反向遍历hits

func hitBricks(grid [][]int, hits [][]int) []int {
	ret := make([]int, len(hits))
	groupMarks := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		groupMarks[i] = make([]int, len(grid[i]))
	}
	groups := []group{{}}
	for _, h := range hits {
		x := h[0]
		y := h[1]
		if grid[x][y] == 1 {
			grid[x][y] = 2
		}
	}
	for i := len(hits) - 1; i >= 0; i-- {
		x := hits[i][0]
		y := hits[i][1]
		if grid[x][y] == 2 {
			count := 0
			dfs(&grid, x-1, y, 0, &groupMarks, &groups)
			dfs(&grid, x+1, y, 0, &groupMarks, &groups)
			dfs(&grid, x, y-1, 0, &groupMarks, &groups)
			dfs(&grid, x, y+1, 0, &groupMarks, &groups)
			dirs := [][]int{
				{-1, 0},
				{1, 0},
				{0, -1},
				{0, 1},
			}
			newGroupID := len(groups)
			groups = append(groups, group{
				p:        newGroupID,
				size:     1,
				isStable: x == 0,
			})
			for _, d := range dirs {
				xn := x + d[0]
				yn := y + d[1]
				if xn >= 0 && xn < len(grid) && yn >= 0 && yn < len(grid[0]) {
					index := getParent(&groups, groupMarks[xn][yn])
					if index != 0 && index != newGroupID {
						// 如果子组可以掉落则增加count
						if !groups[index].isStable {
							count += groups[index].size
						}
						if !groups[newGroupID].isStable {
							groups[newGroupID].isStable = groups[index].isStable
						}
						groupMarks[xn][yn] = newGroupID
						groups[index].p = newGroupID
						groups[newGroupID].size += groups[index].size
					}
				}
			}
			groupMarks[x][y] = newGroupID
			grid[x][y] = 1
			// 如果合并以后依然会掉落则count还是0
			if groups[newGroupID].isStable {
				ret[i] = count
			}
		}
	}
	return ret
}

func getParent(groups *[]group, id int) int {
	if id == 0 {
		return id
	}
	cur := (*groups)[id]
	for cur.p != id {
		id = cur.p
		cur = (*groups)[cur.p]
	}
	return id
}

func dfs(grid *[][]int, x int, y int, groupID int, groupMarks *[][]int, groups *[]group) {
	if x >= 0 && x < len(*grid) && y >= 0 && y < len((*grid)[0]) && (*grid)[x][y] == 1 {
		groupID := getParent(groups, groupID)
		curGroupID := getParent(groups, (*groupMarks)[x][y])
		if groupID != 0 && groupID == curGroupID {
			return
		}
		isTop := x == 0
		if curGroupID == 0 {
			if groupID == 0 {
				curGroupID = len(*groups)
				*groups = append(*groups, group{
					p:        curGroupID,
					size:     1,
					isStable: isTop,
				})
			} else {
				curGroupID = groupID
				(*groups)[curGroupID].size++
				if !(*groups)[curGroupID].isStable {
					(*groups)[curGroupID].isStable = isTop
				}
			}
		} else {
			if groupID != 0 {
				if curGroupID != groupID {
					(*groups)[curGroupID].size += (*groups)[groupID].size
					if !(*groups)[curGroupID].isStable {
						(*groups)[curGroupID].isStable = (*groups)[groupID].isStable
					}
					(*groups)[groupID].p = curGroupID
				}
			}
		}
		(*groupMarks)[x][y] = curGroupID
		dfs(grid, x-1, y, curGroupID, groupMarks, groups)
		dfs(grid, x, y-1, curGroupID, groupMarks, groups)
		dfs(grid, x+1, y, curGroupID, groupMarks, groups)
		dfs(grid, x, y+1, curGroupID, groupMarks, groups)
	}
}
