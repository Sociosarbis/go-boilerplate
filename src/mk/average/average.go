package average

import (
	"math"
)

type MKAverage struct {
	stream []int
	nums   []int
	m      int
	k      int
	p      float64
	v      float64
}

func Constructor(m int, k int) MKAverage {
	return MKAverage{
		stream: []int{},
		nums:   []int{},
		m:      m,
		k:      k,
		p:      float64(m - 2*k),
		v:      0,
	}
}

func (this *MKAverage) add(num int) {
	l := 0
	r := len(this.nums) - 1
	for l <= r {
		mid := (l + r) >> 1
		if this.nums[mid] > num {
			if mid > 0 && this.nums[mid] > num {
				r = mid - 1
			} else {
				l = mid
				break
			}
		} else if this.nums[mid] < num {
			l = mid + 1
		} else {
			l = mid
			break
		}
	}
	if l >= len(this.nums) {
		this.nums = append(this.nums, num)
	} else {
		this.nums = append(this.nums, 0)
		copy(this.nums[l+1:], this.nums[l:])
		this.nums[l] = num
	}
	if l >= this.k && l < len(this.nums)-this.k {
		this.v += float64(num) / this.p
	} else if l < this.k {
		if this.k < len(this.nums)-this.k {
			this.v += float64(this.nums[this.k]) / this.p
		}
	} else {
		if len(this.nums)-this.k-1 >= this.k {
			this.v += float64(this.nums[len(this.nums)-this.k-1]) / this.p
		}
	}
}

func (this *MKAverage) remove(num int) {
	l := 0
	r := len(this.nums) - 1
	for l <= r {
		mid := (l + r) >> 1
		if this.nums[mid] > num {
			r = mid - 1
		} else if this.nums[mid] < num {
			l = mid + 1
		} else {
			l = mid
			break
		}
	}
	if l >= this.k && l < this.m-this.k {
		this.v -= float64(num) / this.p
	} else if l >= this.m-this.k {
		this.v -= float64(this.nums[this.m-this.k-1]) / this.p
	} else {
		this.v -= float64(this.nums[this.k]) / this.p
	}
	copy(this.nums[l:], this.nums[l+1:])
	this.nums = this.nums[:len(this.nums)-1]
}

func (this *MKAverage) AddElement(num int) {
	this.stream = append(this.stream, num)
	if len(this.stream) > this.m {
		removed := this.stream[0]
		this.stream = this.stream[1:]
		this.remove(removed)
	}
	this.add(num)
}

func (this *MKAverage) CalculateMKAverage() int {
	if len(this.nums) < this.m {
		return -1
	}
	return int(math.Floor(this.v))
}
