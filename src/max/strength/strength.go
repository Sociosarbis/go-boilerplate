package strength

import "sort"

func maxStrength(nums []int) int64 {
	sort.Ints(nums)
	n := len(nums)
	i := n - 1
	var ret int64 = int64(nums[i])
	var left int64
	var right int64
	for i >= 0 && nums[i] > 0 {
		if right == 0 {
			right = int64(nums[i])
		} else {
			right *= int64(nums[i])
		}
		i--
	}
	i = 0
	for i < len(nums) && nums[i] < 0 {
		if left == 0 {
			left = int64(nums[i])
		} else {
			left *= int64(nums[i])
		}
		i++
	}

	if right > 0 {
		if left > 0 {
			return left * right
		} else if left == 0 {
			return right
		} else {
			return (left / int64(nums[i-1])) * right
		}
	} else {
		if left > 0 {
			return left
		} else {
			if i > 1 {
				return left / int64(nums[i-1])
			} else {
				return ret
			}
		}
	}
}
