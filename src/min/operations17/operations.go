package operations17

func minOperations(nums []int) int {
	var dir, left int
	n := len(nums)
	// 操作不能改变数字间的相对顺序
	// 所以需要在循环数组视角下是单调数组
	// 反转后左移等于右移
block:
	for i, num := range nums {
		if num == 0 {
			var prev = num
			for j := 1; j < n; j++ {
				index := (i + j) % n
				if nums[index] != prev+1 {
					continue block
				}
				prev = nums[index]
			}
			dir, left = 1, i
			break
		} else if num+1 == n {
			var prev = num
			for j := 1; j < n; j++ {
				index := (i + j) % n
				if nums[index] != prev-1 {
					continue block
				}
				prev = nums[index]
			}
			dir, left = 2, i
			break
		}
	}
	if dir == 0 {
		return -1
	}
	var out int
	if dir == 1 {
		out = left
		if n-left+2 < out {
			out = n - left + 2
		}
	} else {
		out = left + 1
		if n-left+1 < out {
			out = n - left + 1
		}
	}
	return out
}
