package submatrices

func countSubmatrices(grid [][]int, k int) int {
	n := len(grid[0])
	top := make([]int, n)
	sum := make([][2]int, n)
	var out int
	for i, row := range grid {
		top[0] += row[0]
		index := i & 1
		if top[0] <= k {
			out++
		}
		sum[0][index] = top[0]
		left := row[0]
		for j := 1; j < n; j++ {
			cell := row[j]
			left += cell
			temp := left + top[j] + sum[j-1][1-index]
			if temp <= k {
				out++
			}
			sum[j][index] = temp
			top[j] += cell
		}
	}
	return out
}
