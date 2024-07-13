package array

func countBits(num int) int {
	var ret int
	for num != 0 {
		ret++
		num &= num - 1
	}
	return ret
}

func canSortArray(nums []int) bool {
	n := len(nums)
	prev := [2]int{0, 0}
	cur := [2]int{nums[0], nums[0]}
	prevBitCount := countBits(nums[0])
	for i := 1; i < n; i++ {
		bitCount := countBits(nums[i])
		if bitCount != prevBitCount {
			if cur[0] < prev[1] {
				return false
			}
			prev = cur
			cur = [2]int{nums[i], nums[i]}
			prevBitCount = bitCount
		} else {
			if nums[i] > cur[1] {
				cur[1] = nums[i]
			}

			if nums[i] < cur[0] {
				cur[0] = nums[i]
			}
		}
	}
	return cur[0] > prev[1]
}
