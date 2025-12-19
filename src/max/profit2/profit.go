package profit2

func dfs(graph [][]int, present, future []int, u, budget int) [2][]int {
	price := present[u-1]
	halfPrice := int(price / 2)
	sellPrice := future[u-1]
	subProfits := [2][]int{
		make([]int, budget+1),
		make([]int, budget+1),
	}
	dps := [2][]int{
		make([]int, budget+1),
		make([]int, budget+1),
	}
	for _, v := range graph[u] {
		childDps := dfs(graph, present, future, v, budget)
		for k := 0; k < 2; k++ {
			for i := budget; i >= 0; i-- {
				for j := 0; j <= i; j++ {
					if subProfits[k][i-j]+childDps[k][j] > subProfits[k][i] {
						subProfits[k][i] = subProfits[k][i-j] + childDps[k][j]
					}
				}
				if childDps[k][i] > subProfits[k][i] {
					subProfits[k][i] = childDps[k][i]
				}
			}
		}
	}
	for i := 1; i <= budget; i++ {
		for j := 0; j < 2; j++ {
			if subProfits[0][i] > dps[j][i] {
				dps[j][i] = subProfits[0][i]
			}
		}
	}
	end := budget - price
	profit := sellPrice - price
	for i := end; i >= 0; i-- {
		if subProfits[1][i]+profit > dps[0][i+price] {
			dps[0][i+price] = subProfits[1][i] + profit
		}
	}
	end = budget - halfPrice
	profit = sellPrice - halfPrice
	for i := end; i >= 0; i-- {
		if subProfits[1][i]+profit > dps[1][i+halfPrice] {
			dps[1][i+halfPrice] = subProfits[1][i] + profit
		}
	}
	return dps
}

func maxProfit(n int, present []int, future []int, hierarchy [][]int, budget int) int {
	graph := make([][]int, n+1)
	for _, edge := range hierarchy {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
	}
	dps := dfs(graph, present, future, 1, budget)
	var out int
	for j := budget; j >= 1; j-- {
		if dps[0][j] > out {
			out = dps[0][j]
		}
	}
	return out
}
