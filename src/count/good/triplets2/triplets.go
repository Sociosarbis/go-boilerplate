package triplets2

import (
	"github.com/boilerplate/src/tree/fenwick"
)

func goodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	numToIndex1, numToIndex2 := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		numToIndex1[nums1[i]], numToIndex2[nums2[i]] = i, i
	}
	suffix := make([]int, n)
	p1, p2 := fenwick.NewFenwickTree(n), fenwick.NewFenwickTree(n)
	for i := n - 1; i >= 0; i-- {
		num1, num2 := nums1[i], nums2[i]
		if numToIndex1[num1] <= numToIndex2[num1] {
			c := p1.Query(n - 2 - numToIndex2[num1])
			suffix[num1] = c
		}
		if num1 != num2 {
			if numToIndex2[num2] <= numToIndex1[num2] {
				c := p2.Query(n - 2 - numToIndex1[num2])
				suffix[num2] = c
			}
		}
		p1.Add(n-1-numToIndex2[num1], 1)
		p2.Add(n-1-numToIndex1[num2], 1)
	}
	p1, p2 = fenwick.NewFenwickTree(n), fenwick.NewFenwickTree(n)
	var ret int64
	for i := 0; i < n; i++ {
		num1, num2 := nums1[i], nums2[i]
		if numToIndex1[num1] >= numToIndex2[num1] {
			c := p1.Query(numToIndex2[num1] - 1)
			ret += int64(c) * int64(suffix[num1])
		}
		if num1 != num2 {
			if numToIndex2[num2] >= numToIndex1[num2] {
				c := p2.Query(numToIndex1[num2] - 1)
				ret += int64(c) * int64(suffix[num2])
			}
		}
		p1.Add(numToIndex2[num1], 1)
		p2.Add(numToIndex1[num2], 1)
	}
	return ret
}
