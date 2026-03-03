package bit

func dfs(n int, k int, isOdd bool) byte {
	mid := 1 << (n - 1)
	if k == mid {
		if n > 1 {
			if isOdd {
				return '0'
			}
			return '1'
		}
		if isOdd {
			return '1'
		}
		return '0'
	} else if k < mid {
		return dfs(n-1, k, isOdd)
	} else {
		return dfs(n-1, 2*mid-k, !isOdd)
	}
}

func findKthBit(n int, k int) byte {
	return dfs(n, k, false)
}
