package square

func dfs(matchsticks *[]int, i int, temp int, lengthPerSide int, mask int, dp *map[int]bool) bool {
	if res, ok := (*dp)[mask]; ok {
		return res
	}
	if temp == lengthPerSide {
		nextI := i + 1
		if nextI == 4 {
			return true
		}
		return dfs(matchsticks, nextI, 0, lengthPerSide, mask, dp)
	} else if temp > lengthPerSide {
		(*dp)[mask] = false
		return false
	}

	for k := 0; k < len(*matchsticks); k++ {
		if mask&(1<<k) == 0 {
			res := dfs(matchsticks, i, temp+(*matchsticks)[k], lengthPerSide, mask|(1<<k), dp)
			if res {
				return true
			}
		}
	}
	(*dp)[mask] = false
	return false
}

func makesquare(matchsticks []int) bool {
	sum := 0
	for _, s := range matchsticks {
		sum += s
	}
	if sum%4 != 0 {
		return false
	}
	dp := make(map[int]bool, 0)
	return dfs(&matchsticks, 0, 0, sum/4, 0, &dp)
}
