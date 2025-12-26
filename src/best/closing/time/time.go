package time

func bestClosingTime(customers string) int {
	var total int
	for _, customer := range customers {
		if customer == 'Y' {
			total++
		}
	}
	n := len(customers)

	temp := n
	out := n
	var count int
	for i, customer := range customers {
		if i+total-count*2 < temp {
			out = i
			temp = i + total - count*2
		}
		if customer == 'Y' {
			count++
		}
	}
	if n-total < temp {
		return n
	}
	return out
}
