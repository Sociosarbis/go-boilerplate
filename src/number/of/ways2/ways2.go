package ways2

const mask int64 = 1e9 + 7

func numberOfWays(corridor string) int {
	var out int64 = 1
	var temp int
	var count int
	for _, b := range corridor {
		if b == 'S' {
			if count < 2 {
				count++
			} else {
				out = out * int64(temp) % mask
				temp = 0
				count = 1
			}
		}
		if count == 2 {
			temp++
		}
	}
	if count != 2 {
		return 0
	}
	return int(out)
}
