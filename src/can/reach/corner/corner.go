package corner

const left = 1
const right = 2
const bottom = 4
const top = 8

var canMasks = [4]int{left | right, left | bottom, right | top, top | bottom}

func power(x int) int {
	return x * x
}

func canMerge(c1, c2 []int, xCorner, yCorner int) bool {
	v := power(c1[0]-c2[0])+power(c1[1]-c2[1]) <= power(c1[2]+c2[2])
	if v {
		// 判断重叠部分是否在矩形内，不太理解
		return c1[0]*c2[2]+c2[0]*c1[2] < (c1[2]+c2[2])*xCorner && c1[1]*c2[2]+c2[1]*c1[2] < (c1[2]+c2[2])*yCorner
	}
	return v
}

func getMask(circle []int, xCorner, yCorner int) int {
	var mask = 0
	rSq := power(circle[2])
	t1 := power(circle[0])
	t2 := power(circle[0] - xCorner)
	t3 := power(circle[1])
	t4 := power(circle[1] - yCorner)
	if circle[1] >= yCorner {
		if t1+t4 <= rSq {
			mask |= left
		}
		if t2+t4 <= rSq {
			mask |= right
		}
	} else if circle[1] <= 0 {
		if t1+t3 <= rSq {
			mask |= left
		}
		if t2+t3 <= rSq {
			mask |= right
		}
	} else {
		if t1 <= rSq {
			mask |= left
		}
		if t2 <= rSq {
			mask |= right
		}
	}
	if circle[0] >= xCorner {
		if t3+t2 <= rSq {
			mask |= bottom
		}
		if t4+t2 <= rSq {
			mask |= top
		}
	} else if circle[0] <= 0 {
		if t3+t1 <= rSq {
			mask |= bottom
		}
		if t4+t1 <= rSq {
			mask |= top
		}
	} else {
		if t3 <= rSq {
			mask |= bottom
		}
		if t4 <= rSq {
			mask |= top
		}
	}
	return mask
}

func isBlocked(mask int) bool {
	for _, m := range canMasks {
		if mask&m == m {
			return true
		}
	}
	return false
}

type group struct {
	mask    int
	members []int
}

func canReachCorner(xCorner int, yCorner int, circles [][]int) bool {
	u := make([]int, len(circles))
	groups := []group{}
	for i, circle := range circles {
		mask := getMask(circle, xCorner, yCorner)
		if isBlocked(mask) {
			return false
		}
		for j := 0; j < i; j++ {
			if canMerge(circle, circles[j], xCorner, yCorner) {
				g1 := u[i]
				g2 := u[j]
				if g1 == 0 {
					u[i] = g2
					groups[g2-1].mask |= mask
					if isBlocked(groups[g2-1].mask) {
						return false
					}
					groups[g2-1].members = append(groups[g2-1].members, i)
				} else if g1 != g2 {
					groups[g1-1].mask |= groups[g2-1].mask
					if isBlocked(groups[g1-1].mask) {
						return false
					}
					for _, m := range groups[g2-1].members {
						u[m] = g1
					}
					members := make([]int, len(groups[g1-1].members)+len(groups[g2-1].members))
					copy(members, groups[g1-1].members)
					copy(members[len(groups[g1-1].members):], groups[g2-1].members)
					groups[g1-1].members = members
					groups[g2-1].mask = 0
					groups[g2-1].members = nil
				}
			}
		}
		if u[i] == 0 {
			groups = append(groups, group{
				mask,
				[]int{i},
			})
			u[i] = len(groups)
		}
	}
	return true
}
