package set

import "strconv"

func dfs(nums, target []int, i, j int) int {
	if nums[j] < target[i] {
		ret := 1
		for k := i - 1; k >= 0; k-- {
			ret *= len(nums)
		}
		return ret
	} else if nums[j] == target[i] {
		if i == 0 {
			return 1
		} else {
			ret := 0
			for j, num := range nums {
				if num > target[i-1] {
					break
				} else {
					ret += dfs(nums, target, i-1, j)
				}
			}
			return ret
		}
	}
	return 0
}

func atMostNGivenDigitSet(digits []string, n int) int {
	nums := make([]int, len(digits))
	for i, digit := range digits {
		num, _ := strconv.ParseInt(digit, 10, 32)
		nums[i] = int(num)
	}
	target := make([]int, 10)
	i := 0
	for n != 0 {
		target[i] = n % 10
		i++
		n /= 10
	}
	target = target[:i]
	ret := 0
	for j, num := range nums {
		if num > target[len(target)-1] {
			break
		} else {
			ret += dfs(nums, target, len(target)-1, j)
		}
	}
	base := 1
	for i := len(target) - 2; i >= 0; i-- {
		base *= len(nums)
		ret += base
	}
	return ret
}
