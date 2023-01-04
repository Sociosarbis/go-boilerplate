package value

func maxValue(n int, index int, maxSum int) int {
	l := 1
	r := maxSum - n + 1

	for l <= r {
		mid := (l + r) / 2
		leftCount := index
		left := mid - leftCount
		if left <= 0 {
			left = 1
			leftCount = mid - 1
		}
		rightCount := n - index - 1
		right := mid - rightCount
		if right <= 0 {
			right = 1
			rightCount = mid - 1
		}
		temp := leftCount*(mid-1+left)/2 + rightCount*(mid-1+right)/2 + mid + index - leftCount + n - index - 1 - rightCount
		if temp > maxSum {
			r = mid - 1
		} else if temp < maxSum {
			nextL := mid + 1
			leftCount = index
			left = nextL - leftCount
			if left <= 0 {
				left = 1
				leftCount = nextL - 1
			}
			rightCount = n - index - 1
			right = nextL - rightCount
			if right <= 0 {
				right = 1
				rightCount = nextL - 1
			}
			temp = leftCount*(nextL-1+left)/2 + rightCount*(nextL-1+right)/2 + nextL + index - leftCount + n - index - 1 - rightCount
			if temp <= maxSum {
				l = nextL
			} else {
				l = mid
				break
			}
		} else {
			l = mid
			break
		}
	}
	return l
}
