package distance6

import "sort"

func encode(point []int, side int) int64 {
	if point[0] == 0 {
		return int64(point[1])
	}
	if point[0] == side {
		return int64(side)*2 + int64(side-point[1])
	}
	if point[1] == 0 {
		return int64(side)*3 + int64(side-point[0])
	}
	return int64(side) + int64(point[0])
}

func check(nums []int64, side, d, k int) bool {
	n := len(nums)
	r := n - k
	for i := 0; i <= r; i++ {
		num := nums[i]
		var count int
		end := num + 4*int64(side) - int64(d)
		start := num
		si := i
		for count < k && si < n {
			j := sort.Search(n-si, func(j int) bool {
				return nums[si+j] >= start
			})
			if si+j == n || nums[si+j] > end {
				break
			}
			count++
			start = nums[si+j] + int64(d)
			si += j + 1
		}
		if count == k {
			return true
		}
	}
	return false
}

func maxDistance(side int, points [][]int, k int) int {
	n := len(points)
	nums := make([]int64, 0, n)
	for _, point := range points {
		nums = append(nums, encode(point, side))
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	l, r := 1, side
	var out int
	for l <= r {
		mid := (l + r) >> 1
		if check(nums, side, mid, k) {
			l = mid + 1
			out = mid
		} else {
			r = mid - 1
		}
	}
	return out
}
