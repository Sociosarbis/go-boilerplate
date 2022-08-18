package freq

const maxLen = 100001

func maxEqualFreq(nums []int) int {
	counter1 := make([]int, maxLen)
	counter2 := make([]int, maxLen)
	max := 0
	min := maxLen
	ret := 0
	for i, num := range nums {
		prevCount := counter1[num]
		counter1[num]++
		counter2[counter1[num]]++
		counter2[prevCount]--
		if max < counter1[num] {
			max = counter1[num]
		}

		if min == prevCount {
			if counter2[min] == 0 {
				min = prevCount + 1
			}
		} else if min > counter1[num] {
			min = counter1[num]
		}

		if max == min {
			if max == 1 || max == i+1 {
				ret = i + 1
			}
		} else {
			if (counter2[max] == 1 && max == min+1) || (counter2[min] == 1 && max*counter2[max] == i) {
				ret = i + 1
			}
		}
	}
	return ret
}
