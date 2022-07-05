package calendar

type MyCalendar struct {
	ranges [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{
		[][]int{},
	}
}

func insert(ranges *[][]int, i int, item []int) {
	if i >= len(*ranges) {
		*ranges = append(*ranges, item)
	}

	*ranges = append(*ranges, nil)
	copy((*ranges)[i+1:], (*ranges)[i:])
	(*ranges)[i] = item
}

func (this *MyCalendar) Book(start int, end int) bool {
	l := 0
	r := len(this.ranges) - 1
	for l <= r {
		mid := (l + r) >> 1
		if this.ranges[mid][1] <= start {
			l = mid + 1
		} else {
			if mid > 0 && this.ranges[mid-1][1] > start {
				r = mid - 1
			} else {
				l = mid
				break
			}
		}
	}

	ret := true

	if l >= len(this.ranges) {
		this.ranges = append(this.ranges, []int{start, end})
	} else {
		if end <= this.ranges[l][0] {
			if end == this.ranges[l][0] {
				this.ranges[l][0] = start
			} else {
				insert(&this.ranges, l, []int{start, end})
			}
		} else {
			ret = false
		}
	}
	if l > 0 && this.ranges[l-1][1] == this.ranges[l][0] {
		this.ranges[l-1][1] = this.ranges[l][1]
		copy(this.ranges[l:], this.ranges[l+1:])
		this.ranges = this.ranges[:len(this.ranges)-1]
	}
	return ret
}
