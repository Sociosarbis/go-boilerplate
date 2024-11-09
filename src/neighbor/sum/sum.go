package sum

type NeighborSum struct {
	grid         [][]int
	valueToIndex [][2]int
	n            int
}

func Constructor(grid [][]int) NeighborSum {
	n := len(grid)
	valueToIndex := make([][2]int, n*n)
	for i, row := range grid {
		for j, value := range row {
			valueToIndex[value][0], valueToIndex[value][1] = i, j
		}
	}
	return NeighborSum{
		grid,
		valueToIndex,
		len(grid),
	}
}

func (this *NeighborSum) AdjacentSum(value int) int {
	i, j := this.valueToIndex[value][0], this.valueToIndex[value][1]
	var ret int
	if i > 0 {
		ret += this.grid[i-1][j]
	}
	if j > 0 {
		ret += this.grid[i][j-1]
	}
	if i+1 < this.n {
		ret += this.grid[i+1][j]
	}
	if j+1 < this.n {
		ret += this.grid[i][j+1]
	}
	return ret
}

func (this *NeighborSum) DiagonalSum(value int) int {
	i, j := this.valueToIndex[value][0], this.valueToIndex[value][1]
	var ret int
	if i > 0 {
		if j > 0 {
			ret += this.grid[i-1][j-1]
		}
		if j+1 < this.n {
			ret += this.grid[i-1][j+1]
		}
	}
	if i+1 < this.n {
		if j > 0 {
			ret += this.grid[i+1][j-1]
		}
		if j+1 < this.n {
			ret += this.grid[i+1][j+1]
		}
	}
	return ret
}
