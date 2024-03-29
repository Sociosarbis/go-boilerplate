package sum

func minimumSum(nums []int) int {
	n := len(nums)
	ret := 150
	suffix := make([]int, n)
	min := nums[n-1]
	suffix[n-1] = min
	for i := n - 2; i >= 0; i-- {
		if nums[i] < min {
			suffix[i] = nums[i]
			min = nums[i]
		} else {
			suffix[i] = min
		}
	}
	min = nums[0]
	n--
	for i := 1; i < n; i++ {
		if nums[i] > min && nums[i] > suffix[i+1] {
			temp := nums[i] + min + suffix[i+1]
			if temp < ret {
				ret = temp
			}
		}

		if nums[i] < min {
			min = nums[i]
		}
	}
	if ret == 150 {
		return -1
	}
	return ret
}
