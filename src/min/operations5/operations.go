package operations

import "sort"

func minOperations(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		start := i
		for i < len(nums) && nums[i] == nums[i-1] {
			i++
		}
		if start != i {
			copy(nums[start:], nums[i:])
			nums = nums[:len(nums)-(i-start)]
			i = start
		}
	}
	r := 0
	ret := n
	n2 := len(nums)
	for i := 0; i < n2; i++ {
		if r < i {
			r = i
		}
		end := n2 - 1
		num := nums[i]
		for r <= end {
			mid := (r + end) / 2
			if nums[mid]-num < n {
				if mid+1 <= end && nums[mid+1]-num < n {
					r = mid + 1
				} else {
					r = mid
					break
				}
			} else {
				end = mid - 1
			}
		}
		temp := n - (r - i + 1)
		if temp < ret {
			ret = temp
		}
	}
	return ret
}
