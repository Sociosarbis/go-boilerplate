package checker

import "math/rand"

type MajorityChecker struct {
	m   map[int][]int
	arr []int
}

func Constructor(arr []int) MajorityChecker {
	ret := MajorityChecker{
		m:   make(map[int][]int, 0),
		arr: arr,
	}

	for i, num := range arr {
		if _, ok := ret.m[num]; !ok {
			ret.m[num] = []int{i}
		} else {
			ret.m[num] = append(ret.m[num], i)
		}
	}

	return ret
}

func (this *MajorityChecker) Query(left int, right int, threshold int) int {
	for i := 0; i < 10; i++ {
		n := right - left + 1
		index := left + rand.Intn(n)
		num := this.arr[index]
		indices := this.m[num]
		l := 0
		r := len(indices) - 1
		var leftBound int
		var rightBound int
		for l <= r {
			mid := (l + r) >> 1
			if indices[mid] < left {
				l = mid + 1
			} else if indices[mid] > left {
				if mid > 0 && indices[mid-1] >= left {
					r = mid - 1
				} else {
					l = mid
					break
				}
			} else {
				l = mid
				break
			}
		}
		leftBound = l
		l = 0
		r = len(indices) - 1
		for l <= r {
			mid := (l + r) >> 1
			if indices[mid] > right {
				r = mid - 1
			} else if indices[mid] < right {
				if mid+1 < len(indices) && indices[mid+1] <= right {
					l = mid + 1
				} else {
					l = mid
					break
				}
			} else {
				l = mid
				break
			}
		}
		rightBound = l
		if leftBound <= rightBound {
			if rightBound-leftBound+1 >= threshold {
				return num
			}
		}
	}
	return -1
}
