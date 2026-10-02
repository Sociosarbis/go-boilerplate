package parenthesis

func dfs(s []byte, out []string, n, st int) []string {
	if n == 0 && st == 0 {
		return append(out, string(s))
	}
	if n > 0 {
		s = append(s, '(')
		out = dfs(s, out, n-1, st+1)
		s = s[:len(s)-1]
		if st > 0 {
			s = append(s, ')')
			out = dfs(s, out, n, st-1)
		}
	} else if st > 0 {
		s = append(s, ')')
		out = dfs(s, out, n, st-1)
	}
	return out
}

func generateParenthesis(n int) []string {
	return dfs(nil, nil, n, 0)
}
