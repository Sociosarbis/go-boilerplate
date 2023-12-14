package stamp

func canPlace(prefixSum [][]int, i, j, stampHeight, stampWidth int) bool {
	if i+1 < stampHeight || j+1 < stampWidth {
		return false
	}
	res := prefixSum[i][j]
	if i >= stampHeight {
		res -= prefixSum[i-stampHeight][j]
	}
	if j >= stampWidth {
		res -= prefixSum[i][j-stampWidth]
	}

	if i >= stampHeight && j >= stampWidth {
		res += prefixSum[i-stampHeight][j-stampWidth]
	}
	return res == 0
}

func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	m := len(grid)
	n := len(grid[0])

	prefixSum := make([][]int, m)

	for i := range prefixSum {
		prefixSum[i] = make([]int, n)
	}
	// 判断当前格能否作为邮票的右下角进行粘贴
	for i, row := range grid {
		for j, cell := range row {
			topLeft := 0
			left := 0
			top := 0
			if i > 0 {
				if j > 0 {
					topLeft = prefixSum[i-1][j-1]
				}
				top = prefixSum[i-1][j]
			}
			if j > 0 {
				left = prefixSum[i][j-1]
			}
			prefixSum[i][j] = left + top - topLeft + cell
		}
	}

	for i, row := range grid {
		for j, cell := range row {
			if cell == 0 && canPlace(prefixSum, i, j, stampHeight, stampWidth) {
				l := j - stampWidth + 1
				t := i - stampHeight + 1
				if canPlace(prefixSum, i, j-1, stampHeight, stampWidth) {
					l = j
				}
				if canPlace(prefixSum, i-1, j, stampHeight, stampWidth) {
					t = i
				}
				if canPlace(prefixSum, i-1, j-1, stampHeight, stampWidth) {
					for l <= j {
						grid[i][l] = 1
						l++
					}
					for t <= i {
						grid[t][j] = 1
						t++
					}
				} else {
					for t <= i {
						ll := l
						for ll <= j {
							grid[t][ll] = 1
							ll++
						}
						t++
					}
				}
			}
		}
	}
	for _, row := range grid {
		for _, cell := range row {
			if cell == 0 {
				return false
			}
		}
	}
	return true
}
