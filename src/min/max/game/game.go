package game

func minMaxGame(nums []int) int {
	n := len(nums)
	for n > 1 {
		n /= 2
		for i := 0; i < n; i++ {
			if i&1 == 1 {
				if nums[i*2] > nums[i*2+1] {
					nums[i] = nums[i*2]
				} else {
					nums[i] = nums[i*2+1]
				}
			} else {
				if nums[i*2] < nums[i*2+1] {
					nums[i] = nums[i*2]
				} else {
					nums[i] = nums[i*2+1]
				}
			}
		}
	}
	return nums[0]
}
