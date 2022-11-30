package stack

import "container/heap"

type item struct {
	i int
	v int
	c []int
}

type freqStack struct {
	i int
	h []*item
	m map[int]*item
}

type FreqStack struct {
	s freqStack
}

func (s freqStack) Len() int { return len(s.h) }

func (s freqStack) Less(i, j int) bool {
	c1 := s.h[i].c
	c2 := s.h[j].c
	return len(c1) > len(c2) || (len(c1) == len(c2) && c1[len(c1)-1] > c2[len(c2)-1])
}

func (s freqStack) Swap(i, j int) {
	s.h[i].i, s.h[j].i = j, i
	s.h[i], s.h[j] = s.h[j], s.h[i]
}

func (s *freqStack) Push(it any) {
	s.h = append(s.h, it.(*item))
}

func (s *freqStack) Pop() any {
	ret := s.h[len(s.h)-1]
	ret.c = ret.c[:len(ret.c)-1]
	s.h = s.h[:len(s.h)-1]
	return ret
}

func Constructor() FreqStack {
	return FreqStack{
		freqStack{0,
			[]*item{},
			map[int]*item{},
		},
	}
}

func (this *FreqStack) Push(val int) {
	if it, ok := this.s.m[val]; ok {
		it.c = append(it.c, this.s.i)
		heap.Fix(&this.s, it.i)
	} else {
		it := &item{
			len(this.s.h),
			val,
			[]int{this.s.i},
		}
		this.s.m[val] = it
		heap.Push(&this.s, it)
	}
	this.s.i++
}

func (this *FreqStack) Pop() int {
	it := heap.Pop(&this.s).(*item)
	if len(it.c) == 0 {
		delete(this.s.m, it.v)
	} else {
		heap.Push(&this.s, it)
	}
	return it.v
}
