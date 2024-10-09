package difference

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type empty struct{}

func minimumDifference(nums []int, k int) int {
	dp := [2]map[int]empty{}
	ret := abs(k - nums[0])
	for i, num := range nums {
		index := i % 2
		m := make(map[int]empty, 30)
		m[num] = empty{}
		if abs(k-num) < ret {
			ret = abs(k - num)
		}
		for mask := range dp[1-index] {
			v := mask | num
			if abs(k-v) < ret {
				ret = abs(k - v)
			}
			if _, ok := m[v]; !ok {
				m[v] = empty{}
			}
		}
		dp[index] = m
	}
	return ret
}
