package state2

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

const mod int64 = 1e9 + 7

type item struct {
	v int
	i int
}

func getFinalState(nums []int, k int, multiplier int) []int {
	if multiplier == 1 {
		return nums
	}
	n := len(nums)
	var max int
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	arr := make([]item, n)
	for i, num := range nums {
		arr[i].v, arr[i].i = num, i
	}

	hp := h.New(arr, func(a, b item) bool {
		if a.v < b.v {
			return true
		} else if a.v == b.v {
			return a.i < b.i
		}
		return false
	})
	for i := 0; i < k; i++ {
		top := heap.Pop(&hp).(item)
		v := int64(top.v) * int64(multiplier)
		if v <= int64(max) {
			nums[top.i] = int(v)
			top.v = nums[top.i]
			heap.Push(&hp, top)
		} else {
			k -= i
			remain := k % n
			avg := k / n
			indices := make([]int, 0, n)
			indices = append(indices, top.i)
			for hp.Len() > 0 {
				top = heap.Pop(&hp).(item)
				indices = append(indices, top.i)
			}
			a := divideMul(int64(multiplier), avg)
			b := a * int64(multiplier) % mod
			for _, i := range indices {
				var m int64
				if remain > 0 {
					m = b
					remain--
				} else {
					m = a
				}
				nums[i] = int(int64(nums[i]) * m % mod)
			}
			break
		}
	}
	return nums
}

func divideMul(multiplier int64, k int) int64 {
	if k == 0 {
		return 1
	}
	if k == 1 {
		return multiplier
	}
	v := divideMul(multiplier, k/2)
	if k%2 == 0 {
		return v * v % mod
	}
	return v * (v * multiplier % mod) % mod
}
