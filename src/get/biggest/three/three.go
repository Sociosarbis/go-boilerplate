package three

// 左上到右下
func getIndex1(i, j, n int) int {
	return i - j + n - 1
}

// 右上到左下
func getIndex2(i, j int) int {
	return i + j
}

func insert(out *[]int, value int) {
	if len(*out) == 0 {
		*out = append(*out, value)
		return
	}
	for i, num := range *out {
		if value > num {
			if i == 0 || (i > 0 && (*out)[i-1] != value) {
				*out = append(*out, 0)
				copy((*out)[i+1:], (*out)[i:])
				(*out)[i] = value
				if len(*out) > 3 {
					*out = (*out)[:3]
				}
			}
			return
		}
	}
	if len(*out) < 3 && (*out)[len(*out)-1] != value {
		*out = append(*out, value)
	}
}

func collect(prefix1, prefix2, grid [][]int, i, j, m, n int, out *[]int) {
	l := (m - 1 - i) / 2
	if n-j-1 < l {
		l = n - 1 - j
	}
	if j < l {
		l = j
	}
	insert(out, grid[i][j])
	for k := 1; k <= l; k++ {
		li, lj := i+k, j-k
		ri, rj := i+k, j+k
		bi, bj := i+2*k, j
		index1, index2 := getIndex1(i, j, n), getIndex2(i, j)
		sum := prefix2[index2][li+1] - prefix2[index2][i]
		sum += prefix1[index1][ri+1] - prefix1[index1][i]
		index1, index2 = getIndex1(bi, bj, n), getIndex2(bi, bj)
		sum += prefix2[index2][bi+1] - prefix2[index2][ri]
		sum += prefix1[index1][bi+1] - prefix1[index1][li]
		sum -= grid[i][j] + grid[li][lj] + grid[ri][rj] + grid[bi][bj]
		insert(out, sum)
	}
}

func getBiggestThree(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	l := m + n - 1
	prefix1 := make([][]int, l)
	prefix2 := make([][]int, l)
	for i := 0; i < l; i++ {
		prefix1[i] = make([]int, m+1)
		prefix2[i] = make([]int, m+1)
	}
	for i := n - 1; i >= 0; i-- {
		index := getIndex1(0, i, n)
		for j := 0; j < m && j+i < n; j++ {
			prefix1[index][j+1] = prefix1[index][j] + grid[j][j+i]
		}
	}
	for i := 1; i < m; i++ {
		index := getIndex1(i, 0, n)
		for j := 0; j < n && i+j < m; j++ {
			prefix1[index][i+j+1] = prefix1[index][i+j] + grid[i+j][j]
		}
	}
	for i := 0; i < n; i++ {
		index := getIndex2(0, i)
		for j := 0; j < m && i-j >= 0; j++ {
			prefix2[index][j+1] = prefix2[index][j] + grid[j][i-j]
		}
	}
	for i := 1; i < m; i++ {
		index := getIndex2(i, n-1)
		for j := 0; n-1-j >= 0 && i+j < m; j++ {
			prefix2[index][i+j+1] = prefix2[index][i+j] + grid[i+j][n-1-j]
		}
	}
	out := make([]int, 0, 4)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			collect(prefix1, prefix2, grid, i, j, m, n, &out)
		}
	}
	return out
}
