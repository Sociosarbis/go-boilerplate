package array

func splitArray(nums []int, k int) int {
	n := len(nums)
	prefixSum := make([]int, n+1)
	left := 0
	for i, num := range nums {
		if num > left {
			left = num
		}
		prefixSum[i+1] = prefixSum[i] + num
	}
	right := prefixSum[n]

	for left <= right {
		mid := (left + right) / 2
		index := 0
		for i := 0; i < k; i++ {
			if index >= n {
				break
			}
			l := index
			r := n - 1
			for l <= r {
				m := (l + r) / 2
				if prefixSum[m+1]-prefixSum[index] > mid {
					r = m - 1
				} else {
					if m+1 < n && prefixSum[m+2]-prefixSum[index] <= mid {
						l = m + 1
					} else {
						l = m
						break
					}
				}
			}
			index = l + 1
		}
		if index < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
