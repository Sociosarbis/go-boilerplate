package grid

func findPeakGrid(mat [][]int) []int {
	m := len(mat)
	n := len(mat[0])
	if m > n {
		l := 0
		r := m - 1
		for l <= r {
			mid := (l + r) >> 1
			midMax := getRowMax(mat, mid)
			leftMax := getRowMax(mat, mid-1)
			rightMax := getRowMax(mat, mid+1)
			if midMax > leftMax {
				if midMax > rightMax {
					l = mid
					break
				} else {
					l = mid + 1
				}
			} else {
				if leftMax > rightMax {
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
		}
		return []int{l, getRowMaxIndex(mat, l)}
	} else {
		l := 0
		r := n - 1
		for l <= r {
			mid := (l + r) >> 1
			midMax := getColMax(mat, mid)
			leftMax := getColMax(mat, mid-1)
			rightMax := getColMax(mat, mid+1)
			if midMax > leftMax {
				if midMax > rightMax {
					l = mid
					break
				} else {
					l = mid + 1
				}
			} else {
				if leftMax > rightMax {
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
		}
		return []int{getColMaxIndex(mat, l), l}
	}
}

func getRowMax(mat [][]int, i int) int {
	ret := 0
	if i < 0 || i >= len(mat) {
		return ret
	}
	for _, num := range mat[i] {
		if num > ret {
			ret = num
		}
	}
	return ret
}

func getRowMaxIndex(mat [][]int, i int) int {
	ret := 0
	for j := 1; j < len(mat[i]); j++ {
		if mat[i][j] > mat[i][ret] {
			ret = j
		}
	}
	return ret
}

func getColMax(mat [][]int, i int) int {
	ret := 0
	if i < 0 || i >= len(mat[0]) {
		return ret
	}
	for j := 0; j < len(mat); j++ {
		if mat[j][i] > ret {
			ret = mat[j][i]
		}
	}
	return ret
}

func getColMaxIndex(mat [][]int, i int) int {
	ret := 0
	for j := 1; j < len(mat); j++ {
		if mat[j][i] > mat[ret][i] {
			ret = j
		}
	}
	return ret
}
