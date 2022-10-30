package permutation

func dfs(s string, ret []string, index *int, temp []byte, i int) {
	if i < len(s) {
		temp[i] = s[i]
		dfs(s, ret, index, temp, i+1)
		if s[i] >= 97 && s[i] <= 122 {
			temp[i] = s[i] - 32
			dfs(s, ret, index, temp, i+1)
		} else if s[i] >= 65 && s[i] <= 90 {
			temp[i] = s[i] + 32
			dfs(s, ret, index, temp, i+1)
		}
	} else {
		ret[*index] = string(temp)
		*index++
	}
}

func letterCasePermutation(s string) []string {
	n := 0
	for _, c := range s {
		if !(c >= 48 && c <= 57) {
			n++
		}
	}
	if n == 0 {
		return []string{s}
	}
	index := 0
	ret := make([]string, 1<<n)
	dfs(s, ret, &index, make([]byte, len(s)), 0)
	return ret
}
