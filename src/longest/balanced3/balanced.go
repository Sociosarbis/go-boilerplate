package balanced3

func isBalanced(a, b []int) bool {
	var diff int
	for i := 0; i < 26; i++ {
		if (b[i] != 0 || a[i] != 0) && b[i] != a[i] {
			diff = b[i] - a[i]
			break
		}
	}
	for i := 0; i < 26; i++ {
		if (b[i] != 0 || a[i] != 0) && b[i] != a[i] {
			if b[i]-a[i] != diff {
				return false
			}
		}
	}
	return true
}

func longestBalanced(s string) int {
	n := len(s)
	prefixSum := make([][]int, n+1)
	counter := make([]int, 26)
	prefixSum[0] = make([]int, 26)
	var out int
	for i, c := range s {
		counter[c-'a']++
		copied := make([]int, 26)
		copy(copied, counter)
		prefixSum[i+1] = copied
	}
	for i := 1; i <= n; i++ {
		for j := 0; j+out < i; j++ {
			if isBalanced(prefixSum[j], prefixSum[i]) {
				out = i - j
				break
			}
		}
	}
	return out
}
