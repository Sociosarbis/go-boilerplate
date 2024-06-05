package array

import "sort"

type tree []int

func (t tree) add(num, i, start, end int) {
	if start == end {
		t[i]++
	} else {
		mid := (start + end) / 2
		if num >= start && num <= mid {
			t.add(num, i*2+1, start, mid)
		} else {
			t.add(num, i*2+2, mid+1, end)
		}
		t[i] = t[i*2+1] + t[i*2+2]
	}
}

func (t tree) getGreaterCount(num, i, start, end int) int {
	if start > num {
		return t[i]
	}
	if end <= num || start == end {
		return 0
	}
	mid := (start + end) / 2
	return t.getGreaterCount(num, i*2+1, start, mid) + t.getGreaterCount(num, i*2+2, mid+1, end)
}

func getN(num int) int {
	var n int
	var delta int
	if n^(n-1) == 0 {
		delta = 0
	} else {
		delta = 1
	}
	for num != 0 {
		num >>= 1
		n++
	}
	return (1 << (n + delta)) - 1
}

func resultArray(nums []int) []int {
	n := len(nums)
	arr1 := []int{nums[0]}
	arr2 := []int{nums[1]}

	sorted := make([]int, n)
	copy(sorted, nums)
	sort.Ints(sorted)
	var index int
	m := make(map[int]int, n)

	for i, num := range sorted {
		if i == 0 || num != sorted[i-1] {
			m[num] = index
			index++
		}
	}

	N := getN(index)

	leftRoot := tree(make([]int, N))
	rightRoot := tree(make([]int, N))

	leftRoot.add(m[nums[0]], 0, 0, index-1)
	rightRoot.add(m[nums[1]], 0, 0, index-1)

	for i := 2; i < n; i++ {
		num := nums[i]
		v := m[num]
		a := leftRoot.getGreaterCount(v, 0, 0, index-1)
		b := rightRoot.getGreaterCount(v, 0, 0, index-1)
		var isLeft bool
		if a > b {
			isLeft = true
		} else if a == b {
			if len(arr1) <= len(arr2) {
				isLeft = true
			}
		}
		if isLeft {
			arr1 = append(arr1, num)
			leftRoot.add(v, 0, 0, index-1)
		} else {
			arr2 = append(arr2, num)
			rightRoot.add(v, 0, 0, index-1)
		}
	}
	ret := make([]int, n)
	copy(ret[:len(arr1)], arr1)
	copy(ret[len(arr1):], arr2)
	return ret
}
