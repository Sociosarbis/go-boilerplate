package square

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func largestMagicSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	p1 := make([][]int, m)
	for i, row := range grid {
		p1[i] = make([]int, n+1)
		for j, cell := range row {
			p1[i][j+1] = p1[i][j] + cell
		}
	}
	p2 := make([][]int, n)
	for i := 0; i < n; i++ {
		p2[i] = make([]int, m+1)
		for j := 0; j < m; j++ {
			p2[i][j+1] = p2[i][j] + grid[j][i]
		}
	}
	k := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			w, h := n-j, m-i
			l := min(w, h)
		loop:
			for w := l; w > k; w-- {
				sum := p1[i][j+w] - p1[i][j]
				for d := 1; d < w; d++ {
					if p1[i+d][j+w]-p1[i+d][j] != sum {
						continue loop
					}
				}
				for d := 0; d < w; d++ {
					if p2[j+d][i+w]-p2[j+d][i] != sum {
						continue loop
					}
				}
				var temp int
				for d := 0; d < w; d++ {
					temp += grid[i+d][j+d]
				}
				if temp != sum {
					continue
				}
				temp = 0
				for d := 0; d < w; d++ {
					temp += grid[i+d][j+w-1-d]
				}
				if temp != sum {
					continue
				}
				k = w
				break
			}
		}
	}
	return k
}
