package room

type ExamRoom struct {
	n        int
	occupied []int
	pq       [][]int
}

func Constructor(n int) ExamRoom {
	return ExamRoom{
		n:        n,
		occupied: []int{},
		pq:       [][]int{{0, n - 1}},
	}
}

func (this *ExamRoom) maxSpace(space []int) int {
	if space[1] == this.n-1 || space[0] == 0 {
		return space[1] - space[0]
	}
	return (space[1] - space[0]) / 2
}

func (this *ExamRoom) enqueue(space []int) {
	v := this.maxSpace(space)
	l := 0
	r := len(this.pq) - 1
	for l <= r {
		mid := (l + r) / 2
		temp := this.maxSpace(this.pq[mid])
		if temp > v || (temp == v && this.pq[mid][0] < space[0]) {
			if mid > 0 {
				r = mid - 1
			} else {
				l = mid
				break
			}
		} else if temp < v || this.pq[mid][0] > space[0] {
			l = mid + 1
		} else {
			l = mid
			break
		}
	}
	if l >= len(this.pq) {
		this.pq = append(this.pq, space)
	} else {
		this.pq = append(this.pq, []int{})
		copy(this.pq[l+1:], this.pq[l:])
		this.pq[l] = space
	}
}

func (this *ExamRoom) dequeue(space []int) {
	v := this.maxSpace(space)
	l := 0
	r := len(this.pq) - 1
	for l <= r {
		mid := (l + r) / 2
		temp := this.maxSpace(this.pq[mid])
		if temp > v || (temp == v && this.pq[mid][0] < space[0]) {
			r = mid - 1
		} else if temp < v || this.pq[mid][0] > space[0] {
			l = mid + 1
		} else {
			l = mid
			break
		}
	}
	copy(this.pq[l:], this.pq[l+1:])
	this.pq = this.pq[:len(this.pq)-1]
}

func (this *ExamRoom) Seat() int {
	top := this.pq[len(this.pq)-1]
	this.pq = this.pq[:len(this.pq)-1]
	var seat int
	if len(this.occupied) != 0 {
		if top[0] != 0 {
			if top[1] == this.n-1 {
				seat = top[1]
			} else {
				seat = (top[0] + top[1]) / 2
			}
		}
	}
	l := 0
	r := len(this.occupied) - 1
	for l <= r {
		mid := (l + r) / 2
		if this.occupied[mid] > seat {
			if mid > 0 {
				r = mid - 1
			} else {
				l = mid
				break
			}
		} else {
			l = mid + 1
		}
	}
	if l >= len(this.occupied) {
		if len(this.occupied) != 0 {
			if seat != this.occupied[len(this.occupied)-1]+1 {
				space := []int{this.occupied[len(this.occupied)-1] + 1, seat - 1}
				this.enqueue(space)
			}
		}
		if seat != this.n-1 {
			space := []int{seat + 1, this.n - 1}
			this.enqueue(space)
		}
		this.occupied = append(this.occupied, seat)
	} else {
		var space []int
		if l != 0 {
			if this.occupied[l-1]+1 != seat {
				space = []int{this.occupied[l-1] + 1, seat - 1}
				this.enqueue(space)
			}
		}
		if seat+1 != this.occupied[l] {
			space = []int{seat + 1, this.occupied[l] - 1}
			this.enqueue(space)
		}
		this.occupied = append(this.occupied, 0)
		copy(this.occupied[l+1:], this.occupied[l:])
		this.occupied[l] = seat
	}
	return seat
}

func (this *ExamRoom) Leave(p int) {
	l := 0
	r := len(this.occupied) - 1
	var left int
	right := this.n - 1
	for l <= r {
		mid := (l + r) / 2
		if this.occupied[mid] > p {
			r = mid - 1
		} else if this.occupied[mid] < p {
			l = mid + 1
		} else {
			l = mid
			break
		}
	}
	if l+1 == len(this.occupied) {
		if p != this.n-1 {
			this.dequeue([]int{p + 1, this.n - 1})
		}
	} else {
		right = this.occupied[l+1] - 1
		if p+1 != this.occupied[l+1] {
			this.dequeue([]int{p + 1, this.occupied[l+1] - 1})
		}
	}
	if l == 0 {
		if p != 0 {
			this.dequeue([]int{0, p - 1})
		}
	} else {
		left = this.occupied[l-1] + 1
		if this.occupied[l-1]+1 != p {
			this.dequeue([]int{this.occupied[l-1] + 1, p - 1})
		}
	}
	this.enqueue([]int{left, right})
	copy(this.occupied[l:], this.occupied[l+1:])
	this.occupied = this.occupied[:len(this.occupied)-1]
}
