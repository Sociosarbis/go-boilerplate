package fancy

import "sort"

func qpow(n int64, k int) int64 {
	var out int64 = 1
	for k != 0 {
		if k&1 == 1 {
			out = (out * n) % mask
			k--
		} else {
			n = (n * n) % mask
			k >>= 1
		}
	}
	return out
}

const mask int64 = 1e9 + 7

type stack struct {
	end int
	mul int64
	add int64
}

type Fancy struct {
	stacks []stack
	nums   []int
}

func Constructor() Fancy {
	return Fancy{
		stacks: []stack{},
		nums:   []int{},
	}
}

func (this *Fancy) Append(val int) {
	this.nums = append(this.nums, val)
}

func (this *Fancy) AddAll(inc int) {
	if len(this.nums) == 0 {
		return
	}
	if len(this.stacks) == 0 {
		this.stacks = append(this.stacks, stack{
			end: len(this.nums) - 1,
			mul: 1,
			add: int64(inc),
		})
	} else {
		index := len(this.stacks) - 1
		top := this.stacks[index]
		top.add = (top.add + int64(inc)) % mask
		if top.end+1 == len(this.nums) {
			this.stacks[index] = top
		} else {
			top.end = len(this.nums) - 1
			this.stacks = append(this.stacks, top)
		}
	}
}

func (this *Fancy) MultAll(m int) {
	if len(this.nums) == 0 {
		return
	}
	if len(this.stacks) == 0 {
		this.stacks = append(this.stacks, stack{
			end: len(this.nums) - 1,
			mul: int64(m),
		})
	} else {
		index := len(this.stacks) - 1
		top := this.stacks[index]
		top.mul = (top.mul * int64(m)) % mask
		top.add = (top.add * int64(m)) % mask
		if top.end+1 == len(this.nums) {
			this.stacks[index] = top
		} else {
			top.end = len(this.nums) - 1
			this.stacks = append(this.stacks, top)
		}
	}
}

func (this *Fancy) GetIndex(idx int) int {
	if idx >= len(this.nums) {
		return -1
	}
	if len(this.stacks) == 0 {
		return this.nums[idx]
	}
	index := sort.Search(len(this.stacks), func(i int) bool {
		return this.stacks[i].end >= idx
	})
	var mul, add int64
	top := this.stacks[len(this.stacks)-1]
	if index == 0 {
		mul = top.mul
		add = top.add
	} else {
		mul = (top.mul * qpow(this.stacks[index-1].mul, int(mask-2))) % mask
		b := this.stacks[index-1].add * mul % mask
		if top.add < b {
			add = top.add + mask - b
		} else {
			add = top.add - b
		}
	}
	return int(((int64(this.nums[idx])*mul)%mask + add) % mask)
}
