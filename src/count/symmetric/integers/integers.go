package integers

func countSymmetricIntegers(low int, high int) int {
	var ret int
	end := 100
	if end > high+1 {
		end = high + 1
	}
	start := 11
	if start < low {
		start = low
	}
	for i := start; i < end; i++ {
		if i >= low {
			if i/10 == i%10 {
				ret++
			}
		}
	}
	start = 1001
	if start < low {
		start = low
	}
	if start <= high {
		end = 10000
		if end > high+1 {
			end = high + 1
		}
		for i := start; i < end; i++ {
			if (i/100)%10+i/1000 == i%10+(i/10)%10 {
				ret++
			}
		}
	}
	return ret
}
