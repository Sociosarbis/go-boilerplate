package number

func maximumOddBinaryNumber(s string) string {
	bs := []byte(s)
	count := 0
	for _, b := range bs {
		if b == '1' {
			count++
		}
	}
	n := len(bs)
	bs[n-1] = '1'
	count--
	for i := 0; i+1 < n; i++ {
		if count > 0 {
			bs[i] = '1'
			count--
		} else {
			bs[i] = '0'
		}
	}
	return string(bs)
}
