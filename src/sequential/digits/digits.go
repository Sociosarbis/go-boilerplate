package digits

func sequentialDigits(low int, high int) []int {
	if low == 1e9 {
		low--
	}
	if high == 1e9 {
		high--
	}
	var step, start, count int
	temp := low
	nextLevel := 1
	for temp != 0 {
		count++
		start = start*10 + count
		step = step*10 + 1
		nextLevel *= 10
		temp /= 10
	}
	temp = start
	var out []int
	for temp < low {
		if temp%10 != 9 && temp+step < nextLevel {
			temp += step
		} else {
			count++
			temp = start*10 + count
			step = step*10 + 1
			nextLevel *= 10
		}
	}
	for temp <= high && count != 10 {
		out = append(out, temp)
		if temp%10 != 9 && temp+step < nextLevel {
			temp += step
		} else {
			count++
			start = start*10 + count
			step = step*10 + 1
			nextLevel *= 10
			temp = start
		}
	}
	return out
}
