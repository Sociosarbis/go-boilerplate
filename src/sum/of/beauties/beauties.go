package beauties

const max int = 1e5

func sumOfBeauties(nums []int) int {
	n := len(nums)
	suffix := make([]int, n+1)
	suffix[n] = max
	for i := n - 1; i >= 0; i-- {
		if nums[i] < suffix[i+1] {
			suffix[i] = nums[i]
		} else {
			suffix[i] = suffix[i+1]
		}
	}
	n--
	var ret int
	prefix := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > prefix && nums[i] < suffix[i+1] {
			ret += 2
		} else if nums[i] > nums[i-1] && nums[i] < nums[i+1] {
			ret++
		}
		if nums[i] > prefix {
			prefix = nums[i]
		}
	}
	return ret
}
