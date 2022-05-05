package k

func numSubarrayProductLessThanK(nums []int, k int) int {
	ret := 0
	i := 0
	j := 0
	temp := 1
	for ; i < len(nums); i++ {
		for ; j < len(nums); j++ {
			temp *= nums[j]
			if temp >= k {
				temp /= nums[j]
				break
			}
		}
		ret += j - i
		if j != i {
			temp /= nums[i]
		}
		if j < i+1 {
			j = i + 1
		}
	}
	return ret
}
