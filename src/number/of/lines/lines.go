package lines

func numberOfLines(widths []int, s string) []int {
	temp := 0
	lines := 0

	for _, c := range s {
		if temp == 0 {
			lines++
		}

		if temp+widths[c-97] > 100 {
			temp = widths[c-97]
			lines++
		} else {
			temp += widths[c-97]
		}
	}
	return []int{lines, temp}
}
