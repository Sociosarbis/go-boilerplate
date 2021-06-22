package permutation

func permutation(s string) []string {
	chars := [26]int{}
	for _, c := range s {
		chars[c-97] += 1
	}
	ret := []string{}
	dfs(&ret, "", &chars, len(s))
	return ret
}

func dfs(ret *[]string, s string, chars *[26]int, remain int) {
	if remain == 0 {
		*ret = append(*ret, s)
		return
	}
	for i := 0; i < 26; i += 1 {
		if (*chars)[i] != 0 {
			(*chars)[i] -= 1
			dfs(ret, s+string(rune(i+97)), chars, remain-1)
			(*chars)[i] += 1
		}
	}
}
