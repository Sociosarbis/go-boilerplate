package parentheses

func dfs(s string, i int) (int, int) {
	score := 0
	for i < len(s) {
		if s[i] == '(' {
			subScore, nextI := dfs(s, i+1)
			score += subScore
			i = nextI
		} else {
			if score == 0 {
				return 1, i + 1
			} else {
				return score * 2, i + 1
			}
		}
	}
	return score, i
}

func scoreOfParentheses(s string) int {
	ret, _ := dfs(s, 0)
	return ret
}
