package integer

func largestGoodInteger(num string) string {
	var max byte
	n := len(num)
	end := n - 2
	for i := 0; i < end; i++ {
		if num[i] == num[i+1] && num[i] == num[i+2] && num[i] > max {
			max = num[i]
		}
	}
	if max == 0 {
		return ""
	}
	return string([]byte{max, max, max})
}
