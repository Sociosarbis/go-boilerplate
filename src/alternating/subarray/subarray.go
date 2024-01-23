package subarray

func alternatingSubarray(nums []int) int {
	n := len(nums)
	temp := 1
	ret := 1
	for i := 1; i < n; i++ {
		if temp == 1 {
			if nums[i]-nums[i-1] == 1 {
				temp++
				continue
			}
		} else if nums[i] == nums[i-2] {
			temp++
			continue
		}
		if temp > ret {
			ret = temp
		}
		if nums[i]-nums[i-1] == 1 {
			temp = 2
		} else {
			temp = 1
		}

	}
	if temp > ret {
		ret = temp
	}
	if ret == 1 {
		return -1
	}
	return ret
}
