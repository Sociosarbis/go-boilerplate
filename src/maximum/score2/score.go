package score2

func maximumScore(nums []int, k int) int {
	l, r := k, k
	n := len(nums)
	min := nums[k]
	var ret int
	var moveLeft bool
	for l > 0 || r+1 < n {
		if l > 0 {
			nextL := l - 1
			nextR := r + 1
			if nextR < n {
				if nums[nextL] >= nums[nextR] {
					moveLeft = true
				} else {
					moveLeft = false
				}
			} else {
				moveLeft = true
			}
		} else {
			moveLeft = false
		}

		var nextMin int
		if moveLeft {
			nextMin = nums[l-1]
		} else {
			nextMin = nums[r+1]
		}
		if nextMin < min {
			temp := (r - l + 1) * min
			if temp > ret {
				ret = temp
			}
			min = nextMin
		}
		if moveLeft {
			l--
		} else {
			r++
		}
	}
	temp := n * min
	if temp > ret {
		return temp
	}
	return ret
}
