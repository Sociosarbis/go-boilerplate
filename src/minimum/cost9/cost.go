package cost9

func nonNegative(num int) int {
	if num < 0 {
		return 0
	}
	return num
}

func count(cost1, cost2, costBoth, need1, need2, i int) int64 {
	return int64(costBoth)*int64(i) + int64(nonNegative(need1-i))*int64(cost1) + int64(nonNegative(need2-i))*int64(cost2)
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func minimumCost(cost1 int, cost2 int, costBoth int, need1 int, need2 int) int64 {
	return min(min(min(1e15, count(cost1, cost2, costBoth, need1, need2, 0)), count(cost1, cost2, costBoth, need1, need2, need1)), count(cost1, cost2, costBoth, need1, need2, need2))
}
