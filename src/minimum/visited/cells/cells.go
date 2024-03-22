package cells

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

func minimumVisitedCells(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	if m == 1 && n == 1 {
		return 1
	}
	hp := h.New([][2]int{{1, 0}}, func(a, b [2]int) bool {
		return a[0] < b[0]
	})
	hpList := make([]h.Heap[[2]int], n)
	for i := 0; i < n; i++ {
		hpList[i] = h.New([][2]int{}, func(a, b [2]int) bool {
			return a[0] < b[0]
		})
	}
	heap.Push(&hpList[0], [2]int{1, 0})
	for k := 1; k < n; k++ {
		for hp.Len() != 0 {
			top := heap.Pop(&hp).([2]int)
			dis, j := top[0], top[1]
			if grid[0][j]+j >= k {
				dis++
				if m == 1 && k+1 == n {
					return dis
				}
				if grid[0][k] > 0 {
					heap.Push(&hpList[k], [2]int{dis, 0})
					heap.Push(&hp, [2]int{dis, k})
				}
				if grid[0][j]+j > k {
					heap.Push(&hp, top)
				}
				break
			}
		}
	}

	for k := 1; k < m; k++ {
		hp := h.New([][2]int{}, func(a, b [2]int) bool {
			return a[0] < b[0]
		})
		for l := 0; l < n; l++ {
			var fromRow *[2]int
			for hp.Len() != 0 {
				top := heap.Pop(&hp).([2]int)
				j := top[1]
				if grid[k][j]+j >= l {
					fromRow = &top
					break
				}
			}
			var fromCol *[2]int
			for hpList[l].Len() != 0 {
				top := heap.Pop(&hpList[l]).([2]int)
				i := top[1]
				if grid[i][l]+i >= k {
					fromCol = &top
					break
				}
			}
			dis := -1
			if fromRow != nil {
				if fromCol != nil {
					if fromRow[0] <= fromCol[0] {
						dis = fromRow[0]
					} else {
						dis = fromCol[0]
					}
				} else {
					dis = fromRow[0]
				}
				dis++
			} else if fromCol != nil {
				dis = fromCol[0] + 1
			}
			if k+1 == m && l+1 == n {
				return dis
			}
			if dis != -1 {
				if grid[k][l] != 0 {
					heap.Push(&hp, [2]int{dis, l})
					heap.Push(&hpList[l], [2]int{dis, k})
				}
			}
			if fromRow != nil {
				if fromRow[1]+grid[k][fromRow[1]] > l {
					heap.Push(&hp, *fromRow)
				}
			}
			if fromCol != nil {
				if fromCol[1]+grid[fromCol[1]][l] > k {
					heap.Push(&hpList[l], *fromCol)
				}
			}
		}
	}
	return -1
}
