package rectangles

import "math/rand"

type Solution struct {
	rects       [][]int
	prefix_sums []int
}

func Constructor(rects [][]int) Solution {
	prefix_sums := make([]int, len(rects)+1)

	prev := 0
	for i, rect := range rects {
		prev += (rect[2] - rect[0] + 1) * (rect[3] - rect[1] + 1)
		prefix_sums[i+1] = prev
	}

	return Solution{
		prefix_sums: prefix_sums,
		rects:       rects,
	}
}

func (this *Solution) Pick() []int {
	l := 0
	r := len(this.prefix_sums) - 1
	max := this.prefix_sums[r]
	i := rand.Intn(max)
	for l <= r {
		mid := (l + r) >> 1
		if this.prefix_sums[mid] <= i {
			if mid+1 < len(this.prefix_sums) && this.prefix_sums[mid+1] <= i {
				l = mid + 1
			} else {
				l = mid
				break
			}
		} else {
			r = mid - 1
		}
	}

	rest := i - this.prefix_sums[l]
	width := this.rects[l][2] - this.rects[l][0] + 1
	return []int{this.rects[l][0] + rest%width, this.rects[l][1] + rest/width}
}
