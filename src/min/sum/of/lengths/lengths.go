package lengths

func minSumOfLengths(arr []int, target int) int {
	var count, out int
	var sum int
	r := -1
	n := len(arr)
	queue := make([]int, 0, n)
	dp := make([]int, n)
	for i, num := range arr {
		if sum+num > target {
			break
		}
		r = i
		sum += num
		if sum == target {
			dp[i] = r + 1
			queue = append(queue, i)
			break
		}
	}
	var l int
	for i := 1; i < n; i++ {
		if r >= i-1 {
			sum -= arr[i-1]
		}
		if i >= r {
			sum = 0
			r = i - 1
		}
		for idx := l; idx < len(queue); idx++ {
			end := queue[idx]
			if end < i {
				l = idx + 1
				if count == 0 || dp[end] < count {
					count = dp[end]
				}
			} else {
				break
			}
		}
		for j := r + 1; j < n; j++ {
			if sum+arr[j] > target {
				break
			}
			r = j
			sum += arr[j]
			if sum == target {
				value := r - i + 1
				if count != 0 && (out == 0 || value+count < out) {
					out = value + count
				}
				dp[r] = value
				queue = append(queue, r)
				break
			}
		}
	}
	if out == 0 {
		return -1
	}
	return out
}
