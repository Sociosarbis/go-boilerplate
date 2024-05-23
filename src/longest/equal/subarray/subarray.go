package subarray

type item struct {
	cost  int
	queue []int
}

func longestEqualSubarray(nums []int, k int) int {
	n := len(nums)
	dp := make([]item, n+1)
	var ret int
	for i, num := range nums {
		cost, queue := dp[num].cost, dp[num].queue
		remain := k - cost
		l := len(queue)
		var min int
		if l != 0 {
			min = i - queue[l-1] - 1
		}
		index := 0
		for index < l && min > remain {
			if index+1 == l {
				remain = k
			} else {
				remain += queue[index+1] - queue[index] - 1
			}
			index++
		}
		queue = queue[index:]
		l = len(queue)
		if min > remain {
			dp[num].cost = 0
		} else {
			dp[num].cost = k - remain + min
		}
		if ret < l+1 {
			ret = l + 1
		}
		dp[num].queue = append(queue, i)
	}
	return ret
}
