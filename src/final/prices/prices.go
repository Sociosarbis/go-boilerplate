package prices

func finalPrices(prices []int) []int {
	stack := [][]int{{0, prices[0]}}
	ret := make([]int, len((prices)))
	for i := 1; i < len(prices); i++ {
		p := prices[i]
		for j := len(stack) - 1; j >= 0; j-- {
			if stack[j][1] >= p {
				ret[stack[j][0]] = stack[j][1] - p
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, []int{i, p})
	}
	for _, s := range stack {
		ret[s[0]] = s[1]
	}
	return ret
}
