package number2

func findKthNumber(m int, n int, k int) int {
	left := 0
	right := m * n
	for left < right {
		middle := (left + right) >> 1
		count := 0
		i := m
		j := 1
		for i >= 1 && j <= n {
			t := j
			if middle > n*i {
				j = n + 1
			} else {
				j = middle/i + 1
			}
			count += (j - t) * i
			i = middle / j
		}
		if count < k {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return left
}
