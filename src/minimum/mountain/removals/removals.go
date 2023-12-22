package removals

func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	queue := [][2]int{}

	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		l := 0
		r := len(queue) - 1
		for l <= r {
			mid := (l + r) >> 1
			if queue[mid][0] < num {
				l = mid + 1
			} else {
				if mid > 0 && queue[mid-1][0] >= num {
					r = mid - 1
				} else {
					l = mid
					break
				}
			}
		}
		nextValue := 1
		if l > 0 {
			nextValue = queue[l-1][1] + 1
		}
		queue = append(queue, [2]int{})
		copy(queue[l+1:], queue[l:])
		queue[l] = [2]int{num, nextValue}
		l++
		j := l
		for j < len(queue) {
			if queue[j][1] > nextValue {
				break
			}
			j++
		}
		if l != j {
			copy(queue[l:], queue[j:])
			queue = queue[:len(queue)-(j-l)]
		}
		dp[i] = nextValue
	}

	ret := 0
	queue = queue[:0]
	for i := 0; i+1 < n; i++ {
		num := nums[i]
		l := 0
		r := len(queue) - 1
		for l <= r {
			mid := (l + r) >> 1
			if queue[mid][0] < num {
				l = mid + 1
			} else {
				if mid > 0 && queue[mid-1][0] >= num {
					r = mid - 1
				} else {
					l = mid
					break
				}
			}
		}
		nextValue := 1
		if l > 0 {
			nextValue = queue[l-1][1] + 1
		}
		queue = append(queue, [2]int{})
		copy(queue[l+1:], queue[l:])
		queue[l] = [2]int{num, nextValue}
		l++
		j := l
		for j < len(queue) {
			if queue[j][1] > nextValue {
				break
			}
			j++
		}
		if l != j {
			copy(queue[l:], queue[j:])
			queue = queue[:len(queue)-(j-l)]
		}
		if nextValue != 1 && dp[i] != 1 {
			temp := nextValue + dp[i] - 1
			if temp > ret {
				ret = temp
			}
		}
	}
	return n - ret
}
