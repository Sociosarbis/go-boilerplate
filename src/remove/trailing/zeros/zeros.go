package zeros

func removeTrailingZeros(num string) string {
	i := len(num)
	for i > 0 && num[i-1] == '0' {
		i--
	}
	return num[:i]
}
