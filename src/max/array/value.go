package array

func maxArrayValue(nums []int) int64 {
	n := len(nums)
	ret := int64(nums[n-1])
	var temp int64
	for i := n - 1; i >= 0; i-- {
		if int64(nums[i]) <= temp {
			temp += int64(nums[i])
			if temp > ret {
				ret = temp
			}
		} else {
			temp = int64(nums[i])
		}
	}
	if temp > ret {
		return temp
	}
	return ret
}
