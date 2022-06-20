package range_module

type RangeModule struct {
	ranges [][]int
}

func Constructor() RangeModule {
	return RangeModule{
		ranges: [][]int{},
	}
}

func (this *RangeModule) findPivot(target int) int {
	l := 0
	r := len(this.ranges) - 1
	for l <= r {
		mid := (l + r) >> 1
		rng := this.ranges[mid]
		if rng[0] > target {
			if mid > 0 {
				r = mid - 1
			} else {
				l = mid
				break
			}
		} else if rng[1] <= target {
			l = mid + 1
		} else {
			l = mid
			break
		}
	}
	return l
}

func (this *RangeModule) AddRange(left int, right int) {
	i := this.findPivot(left)
	for ; i < len(this.ranges); i++ {
		rng := this.ranges[i]
		if rng[0] >= right {
			this.ranges = append(this.ranges, nil)
			copy(this.ranges[i+1:], this.ranges[i:])
			this.ranges[i] = []int{left, right}
			left = right
			break
		} else if rng[1] > left {
			if rng[0] < left {
				l := rng[0]
				r := left
				this.ranges = append(this.ranges, nil)
				copy(this.ranges[i+1:], this.ranges[i:])
				this.ranges[i] = []int{l, r}
				this.ranges[i+1][0] = r
			} else if rng[0] > left {
				l := left
				r := rng[0]
				this.ranges = append(this.ranges, nil)
				copy(this.ranges[i+1:], this.ranges[i:])
				this.ranges[i] = []int{l, r}
				left = rng[0]
			} else {
				if right <= rng[1] {
					left = right
					break
				} else {
					left = rng[1]
				}
			}
		}
	}

	if i >= len(this.ranges) && left != right {
		this.ranges = append(this.ranges, []int{left, right})
	}
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	i := this.findPivot(left)

	for ; i < len(this.ranges); i++ {
		rng := this.ranges[i]
		if rng[0] > left {
			return false
		} else {
			if left < rng[1] {
				if right <= rng[1] {
					return true
				} else {
					left = rng[1]
				}
			}
		}
	}
	return left == right
}

func (this *RangeModule) RemoveRange(left int, right int) {
	i := this.findPivot(left)
	for ; i < len(this.ranges); i++ {
		rng := this.ranges[i]
		if rng[0] >= right {
			break
		} else {
			if rng[1] > left {
				if rng[0] < left {
					this.ranges = append(this.ranges, nil)
					copy(this.ranges[i+2:], this.ranges[i+1:])
					this.ranges[i+1] = []int{left, rng[1]}
					rng[1] = left
				} else if rng[1] <= right {
					copy(this.ranges[i:], this.ranges[i+1:])
					this.ranges = this.ranges[0 : len(this.ranges)-1]
					i--
				} else {
					this.ranges[i][0] = right
				}
			}
		}
	}
}
