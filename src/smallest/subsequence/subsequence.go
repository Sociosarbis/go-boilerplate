package subsequence

func smallestSubsequence(s string) string {
	n := len(s)
	counter := [26]int{}
	for i := n - 1; i >= 0; i-- {
		counter[s[i]-'a']++
	}
	out := []byte{s[0]}
	used := [26]bool{}
	used[s[0]-'a'] = true
	counter[s[0]-'a']--
	for i := 1; i < n; i++ {
		counter[s[i]-'a']--
		if used[s[i]-'a'] {
			continue
		}
		j := len(out) - 1
		for ; j >= 0; j-- {
			if out[j] > s[i] && counter[out[j]-'a'] > 0 {
				used[out[j]-'a'] = false
			} else {
				break
			}
		}
		out = append(out[:j+1], s[i])
		used[s[i]-'a'] = true
	}
	return string(out)
}
