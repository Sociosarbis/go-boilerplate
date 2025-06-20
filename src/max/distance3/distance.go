package distance3

func maxDistance(s string, k int) int {
	var ret int
	counter := [4]int{}
	for _, c := range s {
		switch c {
		case 'N':
			counter[0]++
		case 'S':
			counter[1]++
		case 'E':
			counter[2]++
		case 'W':
			counter[3]++
		}
		counter2 := counter
		k2 := k
		var temp int
		if counter2[1] > counter2[0] {
			counter2[0], counter2[1] = counter2[1], counter2[0]
		}
		if counter2[1] >= k2 {
			temp = counter2[0] + k2*2 - counter2[1]
			k2 = 0
		} else {
			temp = counter2[0] + counter2[1]
			k2 -= counter2[1]
		}
		if counter2[3] > counter2[2] {
			counter2[2], counter2[3] = counter2[3], counter2[2]
		}
		if counter2[3] >= k2 {
			temp += counter2[2] + k2*2 - counter2[3]
		} else {
			temp += counter2[2] + counter2[3]
		}
		if temp > ret {
			ret = temp
		}
	}
	return ret
}
