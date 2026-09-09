package commas2

func countCommas(n int64) int64 {
	var min int64 = 1e3
	if n < min {
		return 0
	}
	var count int = 3
	var out int64
	for n >= min*10 {
		out += (min * 9) * int64(count/3)
		min *= 10
		count++
	}
	return out + (n-min+1)*int64(count/3)
}
