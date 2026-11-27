package items

func maximumSaleItems(items [][]int, budget int) int {
	minPrice := 1500
	for _, item := range items {
		if item[1] < minPrice {
			minPrice = item[1]
		}
	}
	n := len(items)
	firstCopies := make([]int, n)
	for i, item := range items {
		factor := item[0]
		count := 1
		for j, other := range items {
			if i != j && other[0]%factor == 0 {
				count++
			}
		}
		firstCopies[i] = count
	}
	dp := make([]int, budget)
	if items[0][1] <= budget {
		dp[budget-items[0][1]] = firstCopies[0]
	}
	for i := 1; i < n; i++ {
		price := items[i][1]
		for j := price; j < budget; j++ {
			if dp[j] != 0 {
				if dp[j-price] < dp[j]+firstCopies[i] {
					dp[j-price] = dp[j] + firstCopies[i]
				}
			}
		}
		if price <= budget && dp[budget-price] < firstCopies[i] {
			dp[budget-price] = firstCopies[i]
		}
	}
	var out int
	for remain, count := range dp {
		if count != 0 {
			if remain/minPrice+count > out {
				out = remain/minPrice + count
			}
		}
	}
	return out
}
