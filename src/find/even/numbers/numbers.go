package numbers

func findEvenNumbers(digits []int) []int {
	counter := [10]int{}
	for _, digit := range digits {
		counter[digit]++
	}
	ret := []int{}
	for i := 1; i < 10; i++ {
		if counter[i] != 0 {
			counter[i]--
			for j := 0; j < 10; j++ {
				if counter[j] != 0 {
					counter[j]--
					for k := 0; k < 10; k += 2 {
						if counter[k] != 0 {
							ret = append(ret, 100*i+10*j+k)
						}
					}
					counter[j]++
				}
			}
			counter[i]++
		}
	}
	return ret
}
