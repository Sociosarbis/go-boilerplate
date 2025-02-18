package query

import "fmt"

type RangeFreqQuery struct {
	numToIndices [10001][]int
}

func Constructor(arr []int) RangeFreqQuery {
	numToIndices := [10001][]int{}
	for i, num := range arr {
		numToIndices[num] = append(numToIndices[num], i)
	}
	return RangeFreqQuery{
		numToIndices,
	}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	indices := this.numToIndices[value]
	n := len(indices)
	fmt.Println(left, right, value, n)
	if n == 0 {
		return 0
	}
	var l int
	r := n - 1
	for l <= r {
		mid := (l + r) >> 1
		if indices[mid] >= left {
			if mid > 0 && indices[mid-1] >= left {
				r = mid - 1
			} else {
				l = mid
				break
			}
		} else {
			l = mid + 1
		}
	}
	r = n - 1
	start := l
	for l <= r {
		mid := (l + r) >> 1
		if indices[mid] <= right {
			if mid+1 <= r && indices[mid+1] <= right {
				l = mid + 1
			} else {
				l = mid
				break
			}
		} else {
			r = mid - 1
		}
	}
	if l > r {
		l = r
	}
	if start <= l {
		return l - start + 1
	}
	return 0
}
