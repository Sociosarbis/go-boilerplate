package paths

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func countPaths(n int, edges [][]int) int64 {
	g := make([][]int, n+1)

	for _, e := range edges {
		a := e[0]
		b := e[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	var ret int64
	dfs(g, 1, 0, &ret)
	return ret
}

func dfs(g [][]int, cur, prev int, ans *int64) [2]int64 {
	// 表示经过该节点的，无质数路径数量及单质数路径数量
	count := [2]int64{}
	prime := isPrime(cur)
	if prime {
		count[1]++
	} else {
		count[0]++
	}
	for _, next := range g[cur] {
		if next != prev {
			res := dfs(g, next, cur, ans)
			// 经过子节点的路径与经过父节点的路径的组合数
			*ans += count[0] * res[1]
			*ans += count[1] * res[0]
			// 经过子节点的路径可以跟父节点组成新的路径
			if prime {
				count[1] += res[0]
			} else {
				count[0] += res[0]
				count[1] += res[1]
			}
		}
	}
	return count
}
