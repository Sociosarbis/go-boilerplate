package distance

func getMinDistance(nums []int, target int, start int) int {
	n := len(nums)
	out := n
	for i := 0; start+i < n; i++ {
		if nums[start+i] == target {
			out = i
			break
		}
	}
	if out < start+1 {
		n = out
	} else {
		n = start + 1
	}
	for i := 1; i < n; i++ {
		if nums[start-i] == target {
			return i
		}
	}
	return out
}
