package operations19

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

type doubleHeap struct {
	lhp  h.Heap[[2]int]
	rhp  h.Heap[[2]int]
	sum  int64
	lsum int64
	lc   int
	rc   int
}

func newDoubleHeap() doubleHeap {
	return doubleHeap{
		lhp: h.New[[2]int]([][2]int{}, func(a, b [2]int) bool {
			return a[1] > b[1] || (a[1] == b[1] && a[0] < b[0])
		}),
		rhp: h.New[[2]int]([][2]int{}, func(a, b [2]int) bool {
			return a[1] < b[1] || (a[1] == b[1] && a[0] < b[0])
		}),
	}
}

func (self *doubleHeap) add(index, num int) {
	self.sum += int64(num)
	if self.rc <= self.lc {
		if self.lhp.Len() > 0 && self.lhp.Top()[1] > num {
			ltop := heap.Pop(&self.lhp).([2]int)
			heap.Push(&self.rhp, ltop)
			heap.Push(&self.lhp, [2]int{index, num})
			self.lsum += int64(num) - int64(ltop[1])
		} else {
			heap.Push(&self.rhp, [2]int{index, num})
		}
		self.rc++
	} else {
		if self.rhp.Len() > 0 && self.rhp.Top()[1] < num {
			rtop := heap.Pop(&self.rhp).([2]int)
			heap.Push(&self.lhp, rtop)
			heap.Push(&self.rhp, [2]int{index, num})
			self.lsum += int64(rtop[1])
		} else {
			heap.Push(&self.lhp, [2]int{index, num})
			self.lsum += int64(num)
		}
		self.lc++
	}
}

func (self *doubleHeap) remove(index, num int) {
	self.sum -= int64(num)
	if self.lhp.Len() > 0 && self.lhp.Top()[0] == index {
		heap.Pop(&self.lhp)
		self.lsum -= int64(num)
		self.lc--
	} else if self.rhp.Len() > 0 && self.rhp.Top()[0] == index {
		heap.Pop(&self.rhp)
		self.rc--
	} else {
		if self.lhp.Len() > 0 && num <= self.lhp.Top()[1] {
			self.lc--
			self.lsum -= int64(num)
		} else {
			self.rc--
		}
	}
	for self.lhp.Len() > 0 && self.lhp.Top()[0] <= index {
		heap.Pop(&self.lhp)
	}
	for self.rhp.Len() > 0 && self.rhp.Top()[0] <= index {
		heap.Pop(&self.rhp)
	}
	for self.lc+1 < self.rc {
		top := heap.Pop(&self.rhp).([2]int)
		heap.Push(&self.lhp, top)
		self.lc++
		self.lsum += int64(top[1])
		self.rc--
	}
	for self.rc+1 < self.lc {
		top := heap.Pop(&self.lhp).([2]int)
		heap.Push(&self.rhp, top)
		self.rc++
		self.lc--
		self.lsum -= int64(top[1])
	}
}

func (self *doubleHeap) median() int {
	if self.lc > self.rc {
		return self.lhp.Top()[1]
	} else if self.rc > self.lc {
		return self.rhp.Top()[1]
	} else {
		l, r := self.lhp.Top()[1], self.rhp.Top()[1]
		if (l+r)%2 == 0 {
			return (l + r) / 2
		}
		return l
	}
}

func (self *doubleHeap) cost() int64 {
	m := self.median()
	return self.sum + int64(m)*int64(self.lc-self.rc) - self.lsum*2
}

func min(a, b int64) int64 {
	if a == -1 {
		return b
	}
	if b == -1 {
		return a
	}
	if a < b {
		return a
	}
	return b
}

func reset(s []int64) {
	n := len(s)
	s[0] = 0
	for i := 1; i < n; i++ {
		s[i] = -1
	}
}

func minOperations(nums []int, x int, k int) int64 {
	dh := newDoubleHeap()
	n := len(nums)
	dp := make([][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int64, k+1)
		reset(dp[i])
	}
	best := make([]int64, k+1)
	for i := 0; i < x; i++ {
		num := nums[i]
		dh.add(i, num)
	}
	dp[x-1][1] = dh.cost()
	for i := x; i < n; i++ {
		for j := 1; j <= k; j++ {
			best[j] = min(best[j], dp[i-k][j])
		}
		dh.add(i, nums[i])
		dh.remove(i-x, nums[i-x])
		cost := dh.cost()
		for j := 1; j <= k; j++ {
			if best[j-1] >= 0 {
				dp[i][j] = best[j-1] + cost
			}
		}
	}
	var out int64 = -1
	for i := x*k - 1; i < n; i++ {
		out = min(out, dp[i][k])
	}
	return out
}
