package swaps2

import "sort"

func minSwaps(grid [][]int) int {
	n := len(grid)
	rows := make([]int, n)
	for i, row := range grid {
		for j := n - 1; j >= 0; j-- {
			if row[j] == 1 {
				rows[i] = j
				break
			}
		}
	}
	rowsSorted := make([]int, n)
	copy(rowsSorted, rows)
	sort.Ints(rowsSorted)
	for i := 0; i < n; i++ {
		if rowsSorted[i] > i {
			return -1
		}
	}
	var out int
	// 需要从0开始，这样让目标元素，逐行交换到左边的时候，会使得后面的元素向右移
	// 使得后面的元素，更有可能满足条件
	for i := 0; i < n; i++ {
		if rows[i] > i {
			for j := i + 1; j < n; j++ {
				if rows[j] <= i {
					out += j - i
					value := rows[j]
					copy(rows[i+1:j+1], rows[i:j])
					rows[i] = value
					break
				}
			}
		}
	}
	return out
}
