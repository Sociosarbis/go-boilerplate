package string

func findTheString(lcp [][]int) string {
	n := len(lcp)
	word := make([]byte, n)

	var current byte = 'a'

	for i := 0; i < n; i++ {
		if word[i] == 0 {
			if current > 'z' {
				return ""
			}
			word[i] = current
			j := i + 1
			for ; j < n; j++ {
				// 如果大于0，则表示首字母一定相同
				if lcp[i][j] > 0 {
					word[j] = word[i]
				}
			}
			current++
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if word[i] != word[j] {
				// 如果不同，则相同前缀不可能大于0
				if lcp[i][j] > 0 {
					return ""
				}
			} else {
				if i == n-1 || j == n-1 {
					// 如果相同前缀一定等于1
					if lcp[i][j] != 1 {
						return ""
					}
				} else {
					if lcp[i][j] != lcp[i+1][j+1]+1 {
						return ""
					}
				}
			}
		}
	}
	return string(word)
}
