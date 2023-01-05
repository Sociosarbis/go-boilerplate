package pairs

type node struct {
	l *node
	r *node
	c int
}

func (self *node) count(i int, num int, bound int) int {
	if i < 0 {
		return self.c
	}
	nd := (num >> i) & 1
	bd := (bound >> i) & 1
	ret := 0
	if bd == 1 {
		if nd == 1 {
			if self.l != nil {
				ret += self.l.count(i-1, num, bound)
			}
			if self.r != nil {
				ret += self.r.c
			}
			return ret
		} else {
			if self.r != nil {
				ret += self.r.count(i-1, num, bound)
			}
			if self.l != nil {
				ret += self.l.c
			}
		}
	} else {
		if nd == 0 {
			if self.l != nil {
				ret += self.l.count(i-1, num, bound)
			}
		} else {
			if self.r != nil {
				ret += self.r.count(i-1, num, bound)
			}
		}
	}
	return ret
}

func countPairs(nums []int, low int, high int) int {
	root := &node{}
	ret := 0
	for _, num := range nums {
		ret += root.count(14, num, high) - root.count(14, num, low-1)
		n := root
		for i := 14; i >= 0; i-- {
			d := (num >> i) & 1
			if d == 0 {
				if n.l == nil {
					n.l = &node{}
				}
				n.l.c++
				n = n.l
			} else {
				if n.r == nil {
					n.r = &node{}
				}
				n.r.c++
				n = n.r
			}
		}
	}
	return ret
}
