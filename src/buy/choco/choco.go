package choco

func buyChoco(prices []int, money int) int {
	i1 := 0
	for i, num := range prices {
		if num < prices[i1] {
			i1 = i
		}
	}
	if prices[i1] >= money {
		return money
	}
	var i2 int
	if i1 == 0 {
		i2 = 1
	}
	for i, num := range prices {
		if i != i1 && num < prices[i2] {
			i2 = i
		}
	}
	ret := money - (prices[i1] + prices[i2])
	if ret < 0 {
		return money
	}
	return ret
}
