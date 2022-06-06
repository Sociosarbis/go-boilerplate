package three

type MyCalendarThree struct {
	intervals [][]int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		[][]int{},
	}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	l := 0
	r := len(this.intervals) - 1
	i := 0
	for l < r {
		mid := (l + r) >> 1
		if this.intervals[mid][1] <= start {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	i = l
	k := i
	replaceIntervals := [][]int{}
	ret := 1
	for start < end {
		if k < len(this.intervals) {
			if start >= this.intervals[k][1] {
				replaceIntervals = append(replaceIntervals, this.intervals[k])
				k++
			} else if start < this.intervals[k][0] {
				if end <= this.intervals[k][0] {
					replaceIntervals = append(replaceIntervals, []int{start, end, 1})
					break
				} else {
					replaceIntervals = append(replaceIntervals, []int{start, this.intervals[k][0], 1})
					start = this.intervals[k][0]
				}
			} else if start == this.intervals[k][0] {
				if end >= this.intervals[k][1] {
					this.intervals[k][2]++
					replaceIntervals = append(replaceIntervals, this.intervals[k])
					start = this.intervals[k][1]
					k++
				} else {
					replaceIntervals = append(replaceIntervals, []int{start, end, this.intervals[k][2] + 1})
					replaceIntervals = append(replaceIntervals, []int{end, this.intervals[k][1], this.intervals[k][2]})
					k++
					break
				}
			} else {
				if start >= this.intervals[k][1] {
					replaceIntervals = append(replaceIntervals, this.intervals[k])
				} else {
					replaceIntervals = append(replaceIntervals, []int{this.intervals[k][0], start, this.intervals[k][2]})
					if this.intervals[k][1] <= end {
						replaceIntervals = append(replaceIntervals, []int{start, this.intervals[k][1], this.intervals[k][2] + 1})
						start = this.intervals[k][1]
					} else {
						replaceIntervals = append(replaceIntervals, []int{start, end, this.intervals[k][2] + 1})
						replaceIntervals = append(replaceIntervals, []int{end, this.intervals[k][1], this.intervals[k][2]})
						start = this.intervals[k][1]
					}
				}
				k++
			}
		} else {
			replaceIntervals = append(replaceIntervals, []int{start, end, 1})
			break
		}
	}
	left := this.intervals[:i]
	var right [][]int
	if k < len(this.intervals) {
		right = make([][]int, len(this.intervals)-k)
		copy(right, this.intervals[k:])
	} else {
		right = make([][]int, 0)
	}
	this.intervals = append(append(left, replaceIntervals...), right...)
	for _, interval := range this.intervals {
		if interval[2] > ret {
			ret = interval[2]
		}
	}
	return ret
}
