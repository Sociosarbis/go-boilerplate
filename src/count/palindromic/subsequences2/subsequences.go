package subsequences2

import "strings"

func countPalindromicSubsequence(s string) int {
	var out int
	var i byte
	for i = 'a'; i <= 'z'; i++ {
		l := strings.IndexByte(s, i)
		r := strings.LastIndexByte(s, i)
		if l != -1 && l+2 <= r {
			visited := make([]bool, 26)
			for j := l + 1; j < r; j++ {
				if !visited[s[j]-'a'] {
					visited[s[j]-'a'] = true
					out++
				}
			}
		}
	}
	return out
}
