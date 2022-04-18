package order

func dfs(ret *[]int, num int, n int, i *int) {
	for j := 0; j < 10; j++ {
		temp := num + j
		if temp > n || *i >= len(*ret) {
			break
		}
		(*ret)[*i] = temp
		*i++
		if temp*10 <= n {
			dfs(ret, temp*10, n, i)
		}
	}
}

func lexicalOrder(n int) []int {
	ret := make([]int, n)
	i := 0
	dfs(&ret, 1, n, &i)
	return ret
}
