package profit3

func maxProfit(prices []int, strategy []int, k int) int64 {
	n := len(prices)
	prefixSum1 := make([]int64, n+1)
	prefixSum2 := make([]int64, n+1)
	for i, price := range prices {
		prefixSum1[i+1] = prefixSum1[i] + int64(price*strategy[i])
		prefixSum2[i+1] = prefixSum2[i] + int64(price)
	}
	out := prefixSum1[n]
	for i := n - k; i >= 0; i-- {
		var temp int64
		if i > 0 {
			temp = prefixSum1[i]
		}
		if i+k < n {
			temp += prefixSum1[n] - prefixSum1[i+k]
		}
		temp += prefixSum2[i+k] - prefixSum2[i+k/2]
		if temp > out {
			out = temp
		}
	}
	return out
}
