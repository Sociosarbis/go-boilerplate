package string

func maximumBinaryString(binary string) string {
	bytes := []byte(binary)
	previousZeroIndex := -1

	for i, b := range bytes {
		if b == '0' {
			if previousZeroIndex == -1 {
				previousZeroIndex = i
			} else {
				if previousZeroIndex+1 == i {
					bytes[previousZeroIndex] = '1'
				} else {
					bytes[previousZeroIndex], bytes[previousZeroIndex+1], bytes[i] = '1', '0', '1'
				}
				previousZeroIndex++
			}
		}
	}
	return string(bytes)
}
