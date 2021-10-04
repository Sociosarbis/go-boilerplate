package formatting

func licenseKeyFormatting(s string, k int) string {
	count := 0
	var c byte
	chars := make([]byte, 0, len(s))
	for i := len(s) - 1; i >= 0; i -= 1 {
		c = s[i]
		if c != '-' {
			if count == k {
				chars = append(chars, '-')
				count = 0
			}
			if c >= 97 {
				chars = append(chars, c-32)
			} else {
				chars = append(chars, c)
			}
			count += 1
		}
	}
	i, j := 0, len(chars)-1
	for i < j {
		chars[i], chars[j] = chars[j], chars[i]
		i += 1
		j -= 1
	}
	return string(chars)
}
