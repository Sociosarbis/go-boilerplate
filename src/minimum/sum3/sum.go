package sum3

type rect struct {
	l int
	r int
	t int
	b int
}

func (this *rect) isValid() bool {
	return this.l <= this.r
}

func (this *rect) size() int {
	return (this.r - this.l + 1) * (this.b - this.t + 1)
}

func getRect(grid [][]int, l, r, t, b int) rect {
	x1, y1, x2, y2 := r, b, l, t
	for i := t; i <= b; i++ {
		for j := l; j <= r; j++ {
			if grid[i][j] == 1 {
				if i < y1 {
					y1 = i
				}
				if i > y2 {
					y2 = i
				}
				if j < x1 {
					x1 = j
				}
				if j > x2 {
					x2 = j
				}
			}
		}
	}
	return rect{
		l: x1,
		r: x2,
		t: y1,
		b: y2,
	}
}

func hsplit(grid [][]int, l, r, t, b, s int) (rect, rect) {
	r1 := getRect(grid, l, s, t, b)
	if !r1.isValid() {
		return r1, rect{}
	}
	r2 := getRect(grid, s+1, r, t, b)
	return r1, r2
}

func vsplit(grid [][]int, l, r, t, b, s int) (rect, rect) {
	r1 := getRect(grid, l, r, t, s)
	if !r1.isValid() {
		return r1, rect{}
	}
	r2 := getRect(grid, l, r, s+1, b)
	return r1, r2
}

func tryHsplit(grid [][]int, r rect) int {
	var out int
	for j := r.l; j < r.r; j++ {
		r1, r2 := hsplit(grid, r.l, r.r, r.t, r.b, j)
		if r1.isValid() && r2.isValid() {
			temp := r1.size() + r2.size()
			if out == 0 || temp < out {
				out = temp
			}
		}
	}
	return out
}

func tryVsplit(grid [][]int, r rect) int {
	var out int
	for j := r.t; j < r.b; j++ {
		r1, r2 := vsplit(grid, r.l, r.r, r.t, r.b, j)
		if r1.isValid() && r2.isValid() {
			temp := r1.size() + r2.size()
			if out == 0 || temp < out {
				out = temp
			}
		}
	}
	return out
}

func minimumSum(grid [][]int) int {
	r := getRect(grid, 0, len(grid[0])-1, 0, len(grid)-1)
	out := r.size()
	for i := r.l; i < r.r; i++ {
		r1, r2 := hsplit(grid, r.l, r.r, r.t, r.b, i)
		if r1.isValid() && r2.isValid() {
			s1 := r1.size()
			s2 := tryHsplit(grid, r2)
			if s2 != 0 {
				if s1+s2 < out {
					out = s1 + s2
				}
			}
			s2 = tryVsplit(grid, r2)
			if s2 != 0 {
				if s1+s2 < out {
					out = s1 + s2
				}
			}
			s1 = r2.size()
			s2 = tryVsplit(grid, r1)
			if s2 != 0 {
				if s1+s2 < out {
					out = s1 + s2
				}
			}
		}
	}

	for i := r.t; i < r.b; i++ {
		r1, r2 := vsplit(grid, r.l, r.r, r.t, r.b, i)
		if r1.isValid() && r2.isValid() {
			s1 := r1.size()
			s2 := tryVsplit(grid, r2)
			if s2 != 0 {
				if s1+s2 < out {
					out = s1 + s2
				}
			}
			s2 = tryHsplit(grid, r2)
			if s2 != 0 {
				if s1+s2 < out {
					out = s1 + s2
				}
			}
			s1 = r2.size()
			s2 = tryHsplit(grid, r1)
			if s2 != 0 {
				if s1+s2 < out {
					out = s1 + s2
				}
			}
		}
	}
	return out
}
