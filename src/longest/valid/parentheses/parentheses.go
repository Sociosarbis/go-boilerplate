package parenthesis

func longestValidParentheses(s string) int {
	var count, out int
	var left []int
	for i, c := range s {
		if c == '(' {
			left = append(left, i)
		} else {
			if len(left) > 0 {
				var length int
				if len(left) > 1 {
					length = i - left[len(left)-2]
				} else {
					length = i - left[len(left)-1] + 1
					count += length
					length = count
				}
				if length > out {
					out = length
				}
				left = left[:len(left)-1]
			} else if count > 0 {
				count = 0
			}
		}
	}
	return out
}
