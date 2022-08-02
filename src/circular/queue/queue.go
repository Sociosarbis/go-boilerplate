package queue

type MyCircularQueue struct {
	nums  []int
	start int
	end   int
	count int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		make([]int, k),
		0,
		0,
		0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsEmpty() {
		this.nums[0] = value
		this.start = 0
		this.end = 0
		this.count++
	} else {
		if this.IsFull() {
			return false
		}
		this.end = (this.end + 1) % len(this.nums)
		this.nums[this.end] = value
		this.count++
	}
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.start = (this.start + 1) % len(this.nums)
	this.count--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.nums[this.start]
	}
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.nums[this.end]
	}
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.count == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.count == len(this.nums)
}
