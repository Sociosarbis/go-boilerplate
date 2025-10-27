package beams

func numberOfBeams(bank []string) int {
	var prevCount, count int
	n := len(bank[0])
	var out int
	for _, row := range bank {
		for i := 0; i < n; i++ {
			if row[i] == '1' {
				count++
			}
		}
		if count != 0 {
			out += prevCount * count
			prevCount, count = count, 0
		}
	}
	return out
}
