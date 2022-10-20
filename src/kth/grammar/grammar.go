package grammar

func kthGrammar(n int, k int) int {
	rows := make([]int, n)
	k--
	for i := 0; i < n; i++ {
		rows[n-i-1] = k
		k /= 2
	}

	ret := 0
	for i := 1; i < n; i++ {
		if rows[i] == rows[i-1]*2+1 {
			ret = 1 - ret
		}
	}
	return ret
}
