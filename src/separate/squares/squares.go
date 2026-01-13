package squares

import "sort"

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func separateSquares(squares [][]int) float64 {
	var sum float64
	for _, square := range squares {
		sum += float64(square[2]) * float64(square[2])
	}
	var l float64
	r := 2 * 1e9
	torlerance := 1 / 1e5
	step := 1 / 1e6
	sort.Slice(squares, func(i, j int) bool {
		return squares[i][1]+squares[i][2] < squares[j][1]+squares[j][2]
	})
	n := len(squares)
	out := r
	for l <= r {
		mid := (r + l) / 2
		var midSum float64
		for i := n - 1; i >= 0; i-- {
			square := squares[i]
			temp := float64(square[1] + square[2])
			if temp <= mid {
				break
			}
			if float64(square[1]) >= mid {
				midSum += float64(square[2]) * float64(square[2])
			} else {
				midSum += (float64(square[1]+square[2]) - mid) * float64(square[2])
			}
		}
		other := sum - midSum
		if r-l <= torlerance {
			r = mid - step
			out = mid
		} else {
			if midSum <= other {
				r = mid - step
			} else {
				l = mid + step
			}
		}
	}
	return out
}
