package deque

type MyCircularDeque struct {
	nums []int
	c    int
	k    int
	l    int
	r    int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		nums: make([]int, k),
		c:    0,
		k:    k,
		l:    -1,
		r:    -1,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsEmpty() {
		this.nums[0] = value
		this.l = 0
		this.r = 0
	} else {
		if this.IsFull() {
			return false
		}
		index := this.l - 1
		if index < 0 {
			index += this.k
		}
		this.nums[index] = value
		this.l = index
	}
	this.c++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsEmpty() {
		this.nums[0] = value
		this.l = 0
		this.r = 0
	} else {
		if this.IsFull() {
			return false
		}
		index := this.r + 1
		if index >= this.k {
			index %= this.k
		}
		this.nums[index] = value
		this.r = index
	}
	this.c++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.l = (this.l + 1) % this.k
	this.c--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	index := this.r - 1
	if index < 0 {
		index += this.k
	}
	this.r = index
	this.c--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.nums[this.l]
	}
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.nums[this.r]
	}
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.c == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.c >= this.k
}
