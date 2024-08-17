package sum

const max int = 1e7
const mask int = (1 << 30) - 1

type pair struct {
	val int
	num int
}

func minimumValueSum(nums []int, andValues []int) int {
	n := len(nums)
	m := len(andValues)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = max
		}
	}
	temp := mask
	for i := 0; i < n; i++ {
		temp &= nums[i]
		if temp >= andValues[0] {
			if temp == andValues[0] {
				dp[0][i] = nums[i]
			}
		} else {
			break
		}
	}
	for i := 1; i < m; i++ {
		q := []pair{}
		for j := i; j < n; j++ {
			if dp[i-1][j-1] < max {
				q = append(q, pair{dp[i-1][j-1], nums[j]})
			}
			// 与前面的可能AND
			for k := range q {
				q[k].num &= nums[j]
			}
			// 如果已经小于目标值，则移出队列
			for len(q) != 0 && q[0].num < andValues[i] {
				q = q[1:]
			}
			if len(q) > 0 {
				var start int
				// 维护最小值
				for k, it := range q {
					if it.num == q[start].num {
						if q[start].val > it.val {
							q[start].val = it.val
						}
					} else {
						start++
						q[start] = q[k]
					}
				}
				q = q[:start+1]
				// 队列中num是从小到大排的，而且在上两步中已经去掉了小于andValue的值
				// 以及确保相同num的值，只保留最小那个
				if q[0].num == andValues[i] {
					dp[i][j] = q[0].val + nums[j]
				}
			}
		}
	}
	if dp[m-1][n-1] >= max {
		return -1
	}
	return dp[m-1][n-1]
}
