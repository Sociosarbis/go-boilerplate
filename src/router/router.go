package router

import (
	"sort"
)

type ringBuffer[T any] struct {
	inner []T
	cap   int
	count int
	start int
	end   int
}

func newRingBuffer[T any](cap int) *ringBuffer[T] {
	return &ringBuffer[T]{
		inner: make([]T, 0, cap),
		cap:   cap,
		count: 0,
		start: 0,
		end:   -1,
	}
}

func (this *ringBuffer[T]) push(item T) *T {
	if len(this.inner) < this.cap {
		this.inner = append(this.inner, item)
		this.end++
		this.count++
		return nil
	} else {
		if this.count >= this.cap {
			head := this.inner[this.start]
			this.inner[this.start] = item
			this.end = this.start
			this.start++
			if this.start >= this.cap {
				this.start = 0
			}
			return &head
		} else {
			this.end++
			if this.end >= this.cap {
				this.end = 0
			}
			this.inner[this.end] = item
			this.count++
			return nil
		}
	}
}

func (this *ringBuffer[T]) lpop() *T {
	if this.count == 0 {
		return nil
	}
	item := this.inner[this.start]
	this.start++
	if this.start >= this.cap {
		this.start = 0
	}
	this.count--
	return &item
}

func (this *ringBuffer[T]) top() *T {
	if this.count == 0 {
		return nil
	}
	return &this.inner[this.end]
}

type Router struct {
	buffer      *ringBuffer[[3]int]
	included    map[int]map[int]bool
	destination map[int][]int
}

func Constructor(memoryLimit int) Router {
	return Router{
		buffer:      newRingBuffer[[3]int](memoryLimit),
		included:    map[int]map[int]bool{},
		destination: map[int][]int{},
	}
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	top := this.buffer.top()
	item := [3]int{source, destination, timestamp}
	if top == nil || top[2] != timestamp {
		this.included = map[int]map[int]bool{
			source: {
				destination: true,
			},
		}
	} else {
		if m, ok := this.included[source]; ok {
			if _, ok := m[destination]; ok {
				return false
			}
		}
	}
	removed := this.buffer.push(item)
	if removed != nil {
		if a, ok := this.destination[removed[1]]; ok {
			a = a[1:]
			this.destination[removed[1]] = a
		}
		if removed[2] == top[2] {
			delete(this.included[removed[0]], removed[1])
		}
	}
	if a, ok := this.destination[item[1]]; ok {
		a = append(a, item[2])
		this.destination[item[1]] = a
	} else {
		this.destination[item[1]] = []int{item[2]}
	}
	if top != nil && top[2] == timestamp {
		if m, ok := this.included[source]; ok {
			m[destination] = true
		} else {
			this.included[source] = map[int]bool{
				destination: true,
			}
		}
	}
	return true
}

func (this *Router) ForwardPacket() []int {
	removed := this.buffer.lpop()
	if removed != nil {
		top := this.buffer.top()
		if top == nil || top[2] == removed[2] {
			if m, ok := this.included[removed[0]]; ok {
				delete(m, removed[1])
			}
		}
		if a, ok := this.destination[removed[1]]; ok {
			a = a[1:]
			this.destination[removed[1]] = a
		}
		return []int{removed[0], removed[1], removed[2]}
	}
	return nil
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	if a, ok := this.destination[destination]; ok {
		left := sort.Search(len(a), func(i int) bool {
			return a[i] >= startTime
		})
		if left != len(a) {
			right := sort.Search(len(a), func(i int) bool {
				return a[i] > endTime
			})
			return right - left
		}
	}
	return 0
}
