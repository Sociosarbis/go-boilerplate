package insertions

func minInsertions(s string) int {
	var out, left, right int
	for _, c := range s {
		if c == '(' {
			if right == 1 {
				if left == 0 {
					out += 2
				} else {
					left--
					out++
				}
				right = 0
			}
			left++
		} else {
			right++
			if right == 2 {
				needLeft := right / 2
				if needLeft > left {
					out += needLeft - left
					left = 0
				} else {
					left -= needLeft
				}
				right = 0
			}
		}
	}
	if right == 1 {
		if left == 0 {
			out += 2
		} else {
			left--
			out++
		}
	}
	out += left * 2
	return out
}
