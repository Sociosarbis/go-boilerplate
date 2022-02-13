package utils

func Min(num ...int) int {
	min := num[0]
	for _, n := range num[1:] {
		if n < min {
			min = n
		}
	}
	return min
}
