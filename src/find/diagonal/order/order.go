package order

func findDiagonalOrder(mat [][]int) []int {
	m := len(mat)
	n := len(mat[0])
	l := m + n - 1
	ret := make([]int, m*n)
	count := 0
	for i := 0; i < l; i++ {
		if i&1 == 0 {
			j := m - 1
			k := i - m + 1
			if k < 0 {
				j += k
				k = 0
			}
			for ii := 0; j-ii >= 0; ii++ {
				if k+ii < n {
					ret[count] = mat[j-ii][k+ii]
					count++
				}
			}
		} else {
			j := 0
			k := i
			if k >= n {
				j += k - n + 1
				k = n - 1
			}
			for ii := 0; j+ii < m; ii++ {
				if k-ii >= 0 {
					ret[count] = mat[j+ii][k-ii]
					count++
				}
			}
		}
	}
	return ret
}
