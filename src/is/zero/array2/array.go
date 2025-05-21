package array2

func minZeroArray(nums []int, queries [][]int) int {
	var i int
	n := len(queries)
	m := len(nums)
	dp := make([]int, m)
	var temp int
	for j, num := range nums {
		temp += dp[j]
		for num > temp && i < n {
			if queries[i][1] >= j {
				if queries[i][0] > j {
					dp[queries[i][0]] += queries[i][2]
				} else {
					temp += queries[i][2]
				}
				if queries[i][1]+1 < m {
					dp[queries[i][1]+1] -= queries[i][2]
				}
			}
			i++
		}
		if num > temp {
			if i == n {
				return -1
			}
		}
	}
	return i
}
