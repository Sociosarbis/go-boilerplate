package items2

import (
	"sort"
)

func maximumSaleItems(items [][]int, budget int) int {
	counter := make([]int, 1e5+1)
	for _, item := range items {
		counter[item[0]]++
	}
	gains := make([]int, 1e5+1)
	for i := 1; i <= 1e5; i++ {
		if counter[i] == 0 {
			continue
		}
		for j := 1; j*j <= i; j++ {
			if i%j == 0 {
				a, b := j, i/j
				if counter[a] != 0 {
					if a == i {
						gains[a] += counter[i] - 1
					} else {
						gains[a] += counter[i]
					}
				}
				if b != a {
					if counter[b] != 0 {
						if b == i {
							gains[b] += counter[i] - 1
						} else {
							gains[b] += counter[i]
						}
					}
				}
			}
		}
	}
	var minPrice int = 1e9
	for _, item := range items {
		if item[1] < minPrice {
			minPrice = item[1]
		}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i][1] < items[j][1]
	})
	var out int
	for _, item := range items {
		factor, price := item[0], item[1]
		if gains[factor] == 0 {
			continue
		}
		if price > minPrice*2 {
			break
		}
		count := budget / price
		if count < gains[factor] {
			out += 2 * count
			budget -= price * count
			break
		}
		budget -= price * gains[factor]
		out += 2 * gains[factor]
	}
	return out + budget/minPrice
}
