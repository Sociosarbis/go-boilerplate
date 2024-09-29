package buy

func timeRequiredToBuy(tickets []int, k int) int {
	var ret int
	for i := 0; i < len(tickets); i++ {
		v := tickets[i] - 1
		ret++
		if v == 0 {
			if i == k {
				break
			}
		} else {
			if i == k {
				k = len(tickets)
			}
			tickets = append(tickets, v)
		}
	}
	return ret
}
