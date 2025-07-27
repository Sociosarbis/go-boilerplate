package valley

func countHillValley(nums []int) int {
	s := 0
	prev := nums[0]
	n := len(nums)
	var out int
	for i := 1; i < n; i++ {
		if nums[i] > prev {
			if s == 2 {
				out++
			}
			s = 1
			prev = nums[i]
		} else if nums[i] < prev {
			if s == 1 {
				out++
			}
			s = 2
			prev = nums[i]
		}
	}
	return out
}
