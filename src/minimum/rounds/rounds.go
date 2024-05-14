package rounds

func minimumRounds(tasks []int) int {
	counter := make(map[int]int, len(tasks))

	for _, t := range tasks {
		if c, ok := counter[t]; ok {
			counter[t] = c + 1
		} else {
			counter[t] = 1
		}
	}

	var ret int

	for _, v := range counter {
		if v == 1 {
			return -1
		}
		temp := v % 3
		ret += v / 3
		if temp != 0 {
			ret++
		}
	}
	return ret
}
