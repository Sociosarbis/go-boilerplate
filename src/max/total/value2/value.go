package value2

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

type item struct {
	value int
	i     int
	j     int
}

func log(n int) int {
	var count int
	for n > 1 {
		n >>= 1
		count++
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getValue(mxst, mist [][]int, i, j int) int {
	l := log(j - i + 1)
	mx := max(mxst[i][l], mxst[j-(1<<l)+1][l])
	mi := min(mist[i][l], mist[j-(1<<l)+1][l])
	return mx - mi
}

func maxTotalValue(nums []int, k int) int64 {
	n := len(nums)
	m := log(n)
	mxst := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		mxst[i] = make([]int, m+1)
		mxst[i][0] = nums[i]
		for j := 0; i+(1<<(j+1))-1 < n; j++ {
			mxst[i][j+1] = max(mxst[i][j], mxst[i+(1<<j)][j])
		}
	}
	mist := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		mist[i] = make([]int, m+1)
		mist[i][0] = nums[i]
		for j := 0; i+(1<<(j+1))-1 < n; j++ {
			mist[i][j+1] = min(mist[i][j], mist[i+(1<<j)][j])
		}
	}
	var out int64
	arr := make([]item, n)
	for i := 0; i < n; i++ {
		arr[i].i, arr[i].value, arr[i].j = i, getValue(mxst, mist, i, n-1), n-1
	}
	hp := h.New[item](arr, func(a, b item) bool {
		return a.value > b.value
	})
	for k > 0 {
		top := heap.Pop(&hp).(item)
		out += int64(top.value)
		if top.j > top.i {
			top.j--
			top.value = getValue(mxst, mist, top.i, top.j)
			heap.Push(&hp, top)
		}
		k--
	}
	return out
}
