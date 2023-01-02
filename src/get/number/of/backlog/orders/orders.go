package orders

func getNumberOfBacklogOrders(orders [][]int) int {
	buyOrders := [][]int{}
	sellOrders := [][]int{}
	var m int = 1e9 + 7
	ret := 0
	for _, order := range orders {
		if order[2] == 0 {
			for i := len(sellOrders) - 1; i >= 0; i-- {
				if sellOrders[i][0] <= order[0] {
					if sellOrders[i][1] >= order[1] {
						sellOrders[i][1] -= order[1]
						order[1] = 0
					} else {
						order[1] -= sellOrders[i][1]
						sellOrders[i][1] = 0
					}
					if sellOrders[i][1] == 0 {
						sellOrders = sellOrders[:i]
					}
					if order[1] == 0 {
						break
					}
				}
			}
			if order[1] != 0 {
				l := 0
				r := len(buyOrders) - 1
				for l <= r {
					mid := (l + r) / 2
					if buyOrders[mid][0] > order[0] {
						if mid > 0 {
							r = mid - 1
						} else {
							l = mid
							break
						}
					} else if buyOrders[mid][0] < order[0] {
						l = mid + 1
					} else {
						l = mid
						break
					}
				}
				if l >= len(buyOrders) {
					buyOrders = append(buyOrders, order)
				} else {
					buyOrders = append(buyOrders, []int{})
					copy(buyOrders[l+1:], buyOrders[l:])
					buyOrders[l] = order
				}
			}
		} else {
			for i := len(buyOrders) - 1; i >= 0; i-- {
				if buyOrders[i][0] >= order[0] {
					if buyOrders[i][1] >= order[1] {
						buyOrders[i][1] -= order[1]
						order[1] = 0
					} else {
						order[1] -= buyOrders[i][1]
						buyOrders[i][1] = 0
					}
					if buyOrders[i][1] == 0 {
						buyOrders = buyOrders[:i]
					}
					if order[1] == 0 {
						break
					}
				} else {
					break
				}
			}
			if order[1] != 0 {
				l := 0
				r := len(sellOrders) - 1
				for l <= r {
					mid := (l + r) / 2
					if sellOrders[mid][0] < order[0] {
						if mid > 0 {
							r = mid - 1
						} else {
							l = mid
							break
						}
					} else if sellOrders[mid][0] > order[0] {
						l = mid + 1
					} else {
						l = mid
						break
					}
				}
				if l >= len(sellOrders) {
					sellOrders = append(sellOrders, order)
				} else {
					sellOrders = append(sellOrders, []int{})
					copy(sellOrders[l+1:], sellOrders[l:])
					sellOrders[l] = order
				}
			}
		}
	}
	for _, order := range buyOrders {
		ret += order[1]
		if ret > m {
			ret %= m
		}
	}
	for _, order := range sellOrders {
		ret += order[1]
		if ret > m {
			ret %= m
		}
	}
	return ret
}
